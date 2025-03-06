package services

import (
	"errors"
	"fmt"
	datastructs "learningGo/datastructures"
	"os"
	"path/filepath"
)

// lida com operações GET, le arquivos dos diretórios locais
func getHandler(finalPath string) ([]byte, error) {
	bytes, err := os.ReadFile(finalPath)

	if err != nil {
		return []byte{}, errors.New("falha ao responder a request get")
	}

	return bytes, nil
}

// caso o arquivo não tenha uma extensão, add ele de acordo com o tipo do conteudo
func addExtension(filePath string, contentType datastructs.ContentType) string {
	var extension string

	switch contentType {
	case datastructs.TextPlain:
		extension = ".txt"
	case datastructs.TextHTML:
		extension = ".html"
	case datastructs.TextCSS:
		extension = ".css"
	case datastructs.TextJavaScript:
		extension = ".js"
	case datastructs.ApplicationJSON:
		extension = ".json"
	case datastructs.ApplicationXML:
		extension = ".xml"
	case datastructs.ApplicationForm:
		extension = ".form"
	case datastructs.MultipartForm:
		extension = ".form"
	case datastructs.ImageJPEG:
		extension = ".jpg"
	case datastructs.ImagePNG:
		extension = ".png"
	case datastructs.ImageGIF:
		extension = ".gif"
	case datastructs.ImageSVG:
		extension = ".svg"
	default:
		extension = ".txt" // Default to text extension if content type is unknown
	}

	return filePath + extension
}

func putHandler(finalPath string, content string, contentType datastructs.ContentType) error {
	ext := filepath.Ext(finalPath)
	fmt.Println(ext)
	if ext == "" { //não tem extensao
		finalPath = addExtension(finalPath, contentType)
	}

	return os.WriteFile(finalPath, []byte(content), 0644)
}

// Faz o routing da request e tenta realizar a operação ditada pelo verbo HTTP, retornando uma string com o contéudo
func RouteRequest(requestInfo datastructs.RequestLine, requestBody string, contentType datastructs.ContentType) ([]byte, error) {
	var route string = requestInfo.EndPoint
	var method datastructs.HttpMethod = requestInfo.Method
	finalPath, _ := filepath.Abs("content/" + route)

	var err error
	switch method {
	case datastructs.GET:
		data, err := getHandler(finalPath)
		return data, err
	case datastructs.PUT:
		err = putHandler(finalPath, requestBody, contentType)
		return []byte{}, err
	default:
		return []byte{}, fmt.Errorf("método http da request: %s, ainda não foi implementado", method)
	}

}
