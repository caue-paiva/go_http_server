package services

import (
	"fmt"
	. "learningGo/datastructures"
	"path/filepath"
	"strings"
	"time"
)

func contentTypeFromExtension(fileExtension string) ContentType {
	switch fileExtension {
	case ".txt":
		return TextPlain
	case ".html":
		return TextHTML
	default:
		return TextPlain
	}
}

func GetResponseType(route string) ContentType {
	fileExten := filepath.Ext(route)
	return contentTypeFromExtension(fileExten)
}

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

func SucessResponse(httpVersion float64, responseType string, responseContent []byte) string {
	var response = ""

	response += (fmt.Sprintf("HTTP/%.1f", httpVersion) + " 200 OK\r\n")
	response += fmt.Sprintf("Date: %s\r\n", time.Now().UTC().Format(time.RFC1123))
	response += ("Content-Type: " + responseType + "\r\n")
	response += ("Content-Length: " + fmt.Sprintf("%d", len(responseContent)) + "\r\n")
	response += "\r\n" //line break

	//agora vem o conteúdo
	response += (string(responseContent))
	return response
}
