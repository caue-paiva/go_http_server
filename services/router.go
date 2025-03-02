package services

import (
	"errors"
	"fmt"
	. "learningGo/datastructures"
	"os"
)

// le diretório com os arquivos estáticos a serem servidos e retorna um mapa com o nome do arquivo e seu conteúdo
/*TODO
func ReadContentDir() (map[string]string, error) {

	//dirContents, err := os.ReadDir(contentFolder)
	if err != nil {
		return map[string]string{}, errors.New("falha ao ler diretório de conteúdo")
	}

	// for _, content := range dirContents {

	// }

}
*/

func getHandler(finalPath string) (string, error) {
	bytes, err := os.ReadFile(finalPath)

	if err != nil {
		return "", errors.New("falha ao responder a request get")
	}

	return string(bytes), nil
}

func putHandler(finalPath string, content string) error {
	return os.WriteFile(finalPath, []byte(content), os.FileMode(os.O_RDWR))
}

func RouteRequest(requestInfo RequestLine, requestBody string, contentType ContentType) (string, error) {
	var route string = requestInfo.EndPoint
	var method HttpMethod = requestInfo.Method
	var finalPath string = ".." + route

	var err error
	switch method {
	case GET:
		return getHandler(finalPath)
	case PUT:
		err = putHandler(finalPath, requestBody)
		return "", err
	default:
		return "", fmt.Errorf("método http da request: %s, ainda não foi implementado", method)
	}

}
