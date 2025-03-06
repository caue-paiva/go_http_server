package datastructs

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
)

type Message struct { //struct para uma mensagem do TCP-IP, com os dados transmitidos e a struct da conexão que passou os dados
	Request RequestInfo
	Conn    net.Conn
}

type RequestInfo struct {
	FirstLine RequestLine
	Headers   RequestHeaderLines
	Body      string
}

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
	ContenType       ContentType
	ContentLen       int
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
	var headerFields = map[string]string{}

	for _, line := range headers { //criar mapa dos headers
		splitLines := strings.Split(line, ":")
		fieldKey := strings.ToLower(strings.TrimSpace(splitLines[0]))
		fieldVal := strings.ToLower(strings.TrimSpace(splitLines[1]))
		headerFields[fieldKey] = fieldVal
	}

	var connectionIsPersistant bool
	connectAction, ConnExists := headerFields["connection"]

	if !ConnExists { //conexão não especificada, usar default
		connectionIsPersistant = defaultConnectionIsPersistant(httpVersion)
	} else if connectAction == "close" { //explicitada  para fechar conexão
		connectionIsPersistant = false
	} else { //explicitada  para manter conexão
		connectionIsPersistant = true
	}

	len, lenExists := headerFields["content-length"] //parsing no content len
	var contentLenght int
	if !lenExists {
		contentLenght = 0
	} else {
		var err error
		contentLenght, err = strconv.Atoi(len)
		if err != nil {
			contentLenght = 0
		}
	}

	return RequestHeaderLines{
		Host:             headerFields["host"],
		ConnectionPersis: connectionIsPersistant,
		UserAgent:        headerFields["user-agent"],
		AcceptLanguage:   headerFields["accept"],
		ContenType:       ContentType(headerFields["content-type"]),
		ContentLen:       contentLenght,
	}, err
}

func ParseMetadata(lines []string, lineBreakIndex int) (RequestLine, RequestHeaderLines, error) {
	request, parseErr := ParseRequestLine(lines[0])

	var headers RequestHeaderLines
	var headersErr error
	if lineBreakIndex != -1 {
		headers, headersErr = ParseRequestHeaders(lines[1:lineBreakIndex], request.HttpVersion)
	} else {
		headers, headersErr = ParseRequestHeaders(lines[1:], request.HttpVersion)
	}

	if parseErr != nil {
		return request, headers, fmt.Errorf("erro no parsing da linha de request: %v", parseErr.Error())
	}
	if headersErr != nil {
		return request, headers, fmt.Errorf("erro no parsing dos headers: %v", headersErr.Error())
	}

	return request, headers, nil
}
