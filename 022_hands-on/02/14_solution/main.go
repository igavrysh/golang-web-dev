package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func serve(c net.Conn) {
	defer c.Close()
	scanner := bufio.NewScanner(c)
	i := 0
	var rMethod, rURL string

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			rMethod = strings.Fields(ln)[0]
			fmt.Println("METHOD:" + rMethod)
			rURL = strings.Fields(ln)[1]
			fmt.Println("PATH:" + rURL)
		}
		if ln == "" {
			fmt.Println("THIS IS THE END OF THE HTTP REQUEST HEADERS")
			break
		}
		i++
	}
	body := fmt.Sprintf("METHOD:%s\r\nPATH:%s\r\nCHECK OUT THE RESPONSE BODY PAYLOAD", rMethod, rURL)
	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(c, "Content-Type: text/plain; charset=utf-8\r\n")
	io.WriteString(c, "\r\n")
	io.WriteString(c, body)
}

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go serve(conn)
	}
}
