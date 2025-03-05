package services

import (
	"errors"
	"fmt"
	. "learningGo/datastructures"
	"os"
	"path/filepath"
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

// lida com operações GET, le arquivos dos diretórios locais
func getHandler(finalPath string) ([]byte, error) {
	bytes, err := os.ReadFile(finalPath)

	if err != nil {
		return []byte{}, errors.New("falha ao responder a request get")
	}

	return bytes, nil
}

func putHandler(finalPath string, content string) error {
	return os.WriteFile(finalPath, []byte(content), os.FileMode(os.O_RDWR))
}

// Faz o routing da request e tenta realizar a operação ditada pelo verbo HTTP, retornando uma string com o contéudo
func RouteRequest(requestInfo RequestLine, requestBody string) ([]byte, error) {
	var route string = requestInfo.EndPoint
	var method HttpMethod = requestInfo.Method
	finalPath, _ := filepath.Abs("content/" + route)

	var err error
	switch method {
	case GET:
		data, err := getHandler(finalPath)
		return data, err
	case PUT:
		err = putHandler(finalPath, requestBody)
		return []byte{}, err
	default:
		return []byte{}, fmt.Errorf("método http da request: %s, ainda não foi implementado", method)
	}

}
