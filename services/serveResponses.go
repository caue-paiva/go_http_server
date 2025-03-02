package services

import (
	"fmt"
	. "learningGo/datastructures"
	"strings"
	"time"
)

func ErrorResponse(errorCode HttpError, httpVersion float64, responseType string) string {
	var response = ""
	var errorNum string = strings.Split(string(errorCode), " ")[0]
	var errorName string = strings.Split(string(errorCode), " ")[1]

	response += (fmt.Sprintf("HTTP/%.1f %s %s\n", httpVersion, errorNum, errorName))
	response += ("Content-Type: " + responseType + "\n")
	response += ("Content-Length: " + fmt.Sprintf("%d", len(errorName)) + "\n")
	response += "\n" //line break
	response += errorName

	return response
}

func SucessResponse(httpVersion float64, responseType string, responseContent string) string {
	var response = ""

	response += (fmt.Sprintf("HTTP/%.1f", httpVersion) + " 200 OK\n")
	response += fmt.Sprintf("Date: %s\n", time.Now().UTC().Format(time.RFC1123))
	response += ("Content-Type: " + responseType + "\n")
	response += ("Content-Length: " + fmt.Sprintf("%d", len(responseContent)) + "\n")
	response += "\n" //line break

	//agora vem o conteúdo
	response += string(responseContent)
	return response
}
