package datastructs

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

type RequestLine struct {
	Method      HttpMethod
	EndPoint    string
	HttpVersion float64
}

type RequestHeaderLines struct {
	Host             string
	ConnectionPersis bool
	UserAgent        string
	AcceptLanguage   string
}

func ParseRequestLine(requestLine string) (RequestLine, error) {
	var err error
	strs := strings.Split(requestLine, " ")
	versionStr := strings.Split(strs[2], "/")[1] //[HTTP, 1.1] -> "1.1"
	version, err := strconv.ParseFloat(versionStr, 64)

	if err != nil {
		return RequestLine{}, errors.New("falha ao achar a versão do http")
	}

	return RequestLine{
		Method:      HttpMethod(strs[0]),
		EndPoint:    strs[1],
		HttpVersion: version,
	}, nil
}

func defaultConnectionIsPersistant(httpVersion float64) bool {
	switch httpVersion {

	case 1.1:
		return true

	case 1.0:
		return false

	default:
		return false
	}
}

func ParseRequestHeaders(headers []string, httpVersion float64) (RequestHeaderLines, error) {
	var err error
	var numFields int = reflect.TypeOf(RequestHeaderLines{}).NumField()
	var headerFields = map[string]string{}

	for _, line := range headers { //criar mapa dos headers
		splitLines := strings.Split(line, ":")
		fieldKey := strings.ToLower(strings.TrimSpace(splitLines[0]))
		fieldVal := strings.ToLower(strings.TrimSpace(splitLines[1]))
		headerFields[fieldKey] = fieldVal
	}

	if len(headers) == numFields || len(headers) == (numFields-1) { //O campo conexão é opcional

		var connectionIsPersistant bool
		connectAction, exists := headerFields["Connection"]

		if !exists { //conexão não especificada, usar default
			connectionIsPersistant = defaultConnectionIsPersistant(httpVersion)
		} else if connectAction == "close" { //explicitada  para fechar conexão
			connectionIsPersistant = false
		} else { //explicitada  para manter conexão
			connectionIsPersistant = true
		}

		return RequestHeaderLines{
			Host:             headerFields["host"],
			ConnectionPersis: connectionIsPersistant,
			UserAgent:        headerFields["user-agent"],
			AcceptLanguage:   headerFields["accept"],
		}, err
	} else {
		return RequestHeaderLines{}, errors.New("não foi possível dar parsing nas linhas de header, o número de linhas não é compatível com o número de campos")
	}

}
