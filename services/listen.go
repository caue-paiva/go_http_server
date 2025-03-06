package services

import (
	"bufio"
	"fmt"
	"io"
	datastructs "learningGo/datastructures"
	"net"
	"strings"
)

const hostAdress = "0.0.0.0:8080"
const DataBufferSize = 10000

func readClientRequest(conn net.Conn) (datastructs.RequestInfo, error) {
	reader := bufio.NewReader(conn)                   //io reader
	firstLine, err := reader.ReadString('\n')         //leu primeira linha
	firstLine = strings.TrimSuffix(firstLine, "\r\n") //tira o \r
	if err != nil {
		return datastructs.RequestInfo{}, fmt.Errorf("falha ao ler a primeira linha da request")
	}

	reqLine, parseErr := datastructs.ParseRequestLine(firstLine) //struct da primeira linha
	if parseErr != nil {
		return datastructs.RequestInfo{}, fmt.Errorf("falha ao dar parsing primeira linha da request")
	}

	var headerLines = make([]string, 0, 50) //parsing nos headers
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return datastructs.RequestInfo{}, fmt.Errorf("falha ao ler a linhas dos Headers")
		}

		if line == "\r\n" {
			break //cabou os headers
		}
		line = strings.TrimSuffix(line, "\r\n") //tira o \r
		headerLines = append(headerLines, line)
	}

	headers, parseErr := datastructs.ParseRequestHeaders(headerLines, reqLine.HttpVersion)

	if parseErr != nil {
		return datastructs.RequestInfo{}, fmt.Errorf("falha ao dar parsing nas linhas dos Headers")
	}

	var body string //parsing no body
	if headers.ContentLen != 0 {
		buf := make([]byte, headers.ContentLen) //buffer do tamanho de content len
		_, err = io.ReadFull(reader, buf)       //le os bytes correspondentes ao tamanho do buffer
		if err != nil {
			return datastructs.RequestInfo{}, fmt.Errorf("falha ao ler body da request")
		}
		body = string(buf)
	} else {
		body = ""
	}

	return datastructs.RequestInfo{
		FirstLine: reqLine,
		Headers:   headers,
		Body:      body,
	}, nil

}

func StartServer() (net.Listener, error) {
	listener, err := net.Listen("tcp", hostAdress)

	if err != nil {
		return listener, fmt.Errorf("falha ao estabelecer um listener para o endereço %s", hostAdress)
	}

	return listener, nil
}

// espera um cliente se conectar ao server e lida com ele, função para ser usada numa goroutine
func HandleClient(listener net.Listener, channel chan datastructs.Message) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(fmt.Sprintf("Falha crítica ao aceitar conexão: %v", err))
		}

		requestInfo, readErr := readClientRequest(conn)
		if readErr != nil {
			fmt.Printf("Falha ao receber a request do endereço %v, erro: %v", conn.LocalAddr(), readErr)
		}

		channel <- datastructs.Message{Request: requestInfo, Conn: conn}
	}
}
