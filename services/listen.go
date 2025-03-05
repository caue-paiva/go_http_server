package services

import (
	"fmt"
	"io"
	datastructs "learningGo/datastructures"
	"net"
	"time"
)

const hostAdress = "0.0.0.0:8080"
const DataBufferSize = 10000

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

		conn.SetReadDeadline(time.Now().Add(2 * time.Second)) //deadline de 2 segundos para ler tudo
		data, readErr := io.ReadAll(conn)

		if readErr != nil {
			if netErr, ok := readErr.(net.Error); ok && netErr.Timeout() {
				fmt.Println("Request parcial recebida depois do timeout")
			} else {
				fmt.Println("Erro lendo da conexão:", readErr)
				conn.Close()
				continue
			}
		}

		channel <- datastructs.Message{Data: data, Conn: conn}
	}
}
