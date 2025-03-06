package main

import (
	"fmt"
	datastructs "learningGo/datastructures"
	"learningGo/services"
	"net"
)

const NUM_WORKERS = 5

func runServer() {
	listener, err := services.StartServer()

	if err != nil {
		panic(fmt.Errorf("falha ao começar o servidor: %v ", err))
	}

	channel := make(chan datastructs.Message)

	for range NUM_WORKERS { //cria os 5 workers
		go services.HandleClient(listener, channel)
	}

	var req datastructs.RequestInfo
	var conn net.Conn
	for { //loop inf

		var msg datastructs.Message = <-channel //pega request mais recent do canal
		req = msg.Request
		conn = msg.Conn

		content, _ := services.RouteRequest(req.FirstLine, req.Body) //routing na mensagem
		var contentType datastructs.ContentType = services.GetResponseType(req.FirstLine.EndPoint)

		var response string = services.SucessResponse(req.FirstLine.HttpVersion, contentType, content)

		_, WriteErr := conn.Write([]byte(response))
		if WriteErr != nil {
			fmt.Println("Falha ao escreve resposta:", WriteErr)
		}

		var addr net.Addr = conn.RemoteAddr()
		closeErr := conn.Close()
		if closeErr != nil {
			fmt.Println(closeErr.Error())
		}
		fmt.Printf("Request conectada do endereço %v for servida com sucesso. A rota foi %v e o tipo de contéudo servido foi %v \n", addr, req.FirstLine.EndPoint, contentType)

	}

}

func main() {
	runServer()
}
