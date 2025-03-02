package services

import (
	"fmt"
	"net"
)

const hostAdress = "127.0.0.1:8080"
const DataBufferSize = 10000

func StartServer() (net.Listener,error) {
	listener, err := net.Listen("tcp",hostAdress)

	if err != nil {
		return listener,fmt.Errorf("falha ao estabelecer um listener para o endereço %s",hostAdress)
	}

	return listener, nil
}

func HandleClient(listener net.Listener) {

	conn,err := listener.Accept()
	if err != nil {
		panic(fmt.Sprintf("Falha crítica ao aceitar conexão: %v", err))
	}

	var data = make([]byte,0,DataBufferSize)
	readBytes,err := conn.Read(data)

	if err != nil {
		panic(fmt.Sprintf("Erro ao ler dado da conexão: %v", err))
	} 
}