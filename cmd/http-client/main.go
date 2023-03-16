package main

import (
	"fmt"
	"net"
	"os"

	"github.com/Goose97/tiny-http-server/pkg/client"
	"github.com/Goose97/tiny-http-server/pkg/common"
)

func main() {
	request, err := client.Parse()

	if err != nil {
		fmt.Printf("Error when parse request %+v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Request %+v\n", request)
	conn, err := net.Dial("tcp", request.Url.Host)

	if err != nil {
		panic("Can not connect to tcp server")
	}

	defer conn.Close()
	execute(request, conn)
}

func execute(request common.ClientRequest, conn net.Conn) {
	requestLine := []byte(fmt.Sprintf("GET %v HTTP/1.0\r\n", request.Url.Path))
	conn.Write(requestLine)

	var headers []byte

	for _, h := range request.Headers {
		headers = append(headers, []byte(fmt.Sprintf("%v: %v\r\n", h.Name, h.Value))...)
	}

	conn.Write(headers)
	conn.Write([]byte("\r\n"))
}
