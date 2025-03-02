package main

import (
	"fmt"
	. "learningGo/datastructures"
	. "learningGo/services"
	"os"
)

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

func main() {
	var response string = mockServer("request.txt", "content/texto.txt")
	fmt.Println(response)
}
