package main

import (
	"fmt"
	. "learningGo/datastructures"
	. "learningGo/services"
	"net"
	"strings"
)

const NUM_WORKERS = 5

/*
func mockServer(requestPath string, endpointPath string) string {
	var lines, lineBreakIndex = GetRequestContents(requestPath)
	request, parseErr := ParseRequestLine(lines[0])

	var headers RequestHeaderLines
	var headersErr error
	if lineBreakIndex != -1 {
		headers, headersErr = ParseRequestHeaders(lines[1:lineBreakIndex], request.HttpVersion)
	} else {
		headers, headersErr = ParseRequestHeaders(lines[1:], request.HttpVersion)
	}

	if parseErr != nil {
		fmt.Println(parseErr.Error())
	}
	if headersErr != nil {
		fmt.Println(headersErr.Error())
	}

	byteArr, _ := os.ReadFile(endpointPath)
	var content = string(byteArr)
	var response string = SucessResponse(request.HttpVersion, headers.AcceptLanguage, content)

	return response
}
*/

func runServer() {
	listener, err := StartServer()

	if err != nil {
		panic(fmt.Errorf("falha ao começar o servidor: %v ", err))
	}

	channel := make(chan Message)

	for range NUM_WORKERS { //cria os 5 workers
		go HandleClient(listener, channel)
	}

	var data []byte
	var conn net.Conn
	for { //loop inf

		var msg Message = <-channel
		data = msg.Data
		conn = msg.Conn

		lines, lineBreak := DeserializeRequest(data)
		if len(lines) == 0 {
			continue
		}
		reqInfo, _, err := ParseMetadata(lines, lineBreak) //metadados da request, a linha da request e os headers

		var reqBody string
		if reqInfo.Method == "PUT" || reqInfo.Method == "POST" {
			reqBody = strings.Join(lines[lineBreak+1:], "\n") //concat de um array de strings em uma só
		} else {
			reqBody = ""
		}

		if err != nil {
			fmt.Println(err.Error())
		}

		content, _ := RouteRequest(reqInfo, reqBody)
		var contentType ContentType = GetResponseType(reqInfo.EndPoint)

		var response string = SucessResponse(reqInfo.HttpVersion, string(contentType), content)
		fmt.Println(response)

		_, WriteErr := conn.Write([]byte(response))
		if WriteErr != nil {
			fmt.Println("Falha ao escreve resposta:", WriteErr)
		}

	}

}

func main() {
	runServer()
}
