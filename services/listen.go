package services

import (
	"fmt"
	"io"
	"net"
)

const hostAdress = "127.0.0.1:8080"
const DataBufferSize = 10000

func StartServer() (net.Listener, error) {
	listener, err := net.Listen("tcp", hostAdress)

	if err != nil {
		return listener, fmt.Errorf("falha ao estabelecer um listener para o endereço %s", hostAdress)
	}

	return listener, nil
}

// espera um cliente se conectar ao server e lida com ele, função para ser usada numa goroutine
func HandleClient(listener net.Listener, channel chan []byte) {

	conn, err := listener.Accept()
	if err != nil {
		panic(fmt.Sprintf("Falha crítica ao aceitar conexão: %v", err))
	}

	data, readErr := io.ReadAll(conn)
	if readErr != nil {
		panic(fmt.Sprintf("Erro ao ler dado da conexão: %v", readErr))
	}

	channel <- data
}
