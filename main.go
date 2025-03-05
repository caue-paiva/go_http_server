package main

import (
	"fmt"
	datastructs "learningGo/datastructures"
	"learningGo/services"
	"net"
	"strings"
)

const NUM_WORKERS = 5

func runServer() {
	listener, err := services.StartServer()

	if err != nil {
		panic(fmt.Errorf("falha ao começar o servidor: %v ", err))
	}

	channel := make(chan datastructs.Message)

	for range NUM_WORKERS { //cria os 5 workers
		go services.HandleClient(listener, channel)
	}

	var data []byte
	var conn net.Conn
	for { //loop inf

		var msg datastructs.Message = <-channel
		data = msg.Data
		conn = msg.Conn

		lines, lineBreak := services.DeserializeRequest(data)
		if len(lines) == 0 {
			continue
		}
		reqInfo, _, err := datastructs.ParseMetadata(lines, lineBreak) //metadados da request, a linha da request e os headers

		var reqBody string
		if reqInfo.Method == "PUT" || reqInfo.Method == "POST" {
			reqBody = strings.Join(lines[lineBreak+1:], "\n") //concat de um array de strings em uma só
		} else {
			reqBody = ""
		}

		if err != nil {
			fmt.Println(err.Error())
		}

		content, _ := services.RouteRequest(reqInfo, reqBody)
		var contentType datastructs.ContentType = services.GetResponseType(reqInfo.EndPoint)

		var response string = services.SucessResponse(reqInfo.HttpVersion, string(contentType), content)
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
