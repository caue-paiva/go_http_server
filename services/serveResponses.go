package services

import (
	"fmt"
	datastructs "learningGo/datastructures"
	"path/filepath"
	"strings"
	"time"
)

func contentTypeFromExtension(fileExtension string) datastructs.ContentType {
	switch fileExtension {
	case ".txt":
		return datastructs.TextPlain
	case ".html":
		return datastructs.TextHTML
	default:
		return datastructs.TextPlain
	}
}

func GetResponseType(method datastructs.HttpMethod, route string) datastructs.ContentType {
	if method != datastructs.GET { //caso não seja request de get, retorna uma string
		return datastructs.TextPlain
	}
	fileExten := filepath.Ext(route)
	return contentTypeFromExtension(fileExten)
}

func ErrorResponse(errorCode datastructs.HttpError, httpVersion float64, responseType string) string {
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

func SucessResponse(httpVersion float64, responseType datastructs.ContentType, responseContent []byte) string {
	var response = ""

	response += (fmt.Sprintf("HTTP/%.1f", httpVersion) + " 200 OK\r\n")
	response += fmt.Sprintf("Date: %s\r\n", time.Now().UTC().Format(time.RFC1123))
	response += ("Content-Type: " + string(responseType) + "\r\n")
	response += ("Content-Length: " + fmt.Sprintf("%d", len(responseContent)) + "\r\n")
	response += "\r\n" //line break

	//agora vem o conteúdo
	response += (string(responseContent))
	return response
}
