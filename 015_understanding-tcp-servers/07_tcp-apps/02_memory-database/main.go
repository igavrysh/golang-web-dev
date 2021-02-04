package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			panic(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	// instructions
	io.WriteString(conn, "\nIN-MEMORY DATABASE\n\n" +
		"USE:\n" +
		"SET key value\n" +
		"GET key\n" +
		"DEL key\n\n" +
		"EXAMPLE:\n" +
		"SET fav chocolate\n" +
		"GET fav\n\n\n")

	// read & write
	data := make(map[string]string)
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln)

		if len(fs) == 0 {
			fmt.Fprintf(conn, "EXPECTED COMMAND\n")
			continue
		}

		// logic
		switch fs[0] {
		case "GET":
			k := fs[1]
			v := data[k]
			fmt.Fprintf(conn, "%s\n", v)
		case "SET":
			if len(fs) != 3 {
				fmt.Fprintf(conn, "EXPECTED VALUE\n")
				continue
			}
			k := fs[1]
			v := fs[2]
			data[k] = v
		case "DEL":
			k := fs[1]
			delete(data, k)
		default:
			fmt.Fprintf(conn, "INVALID COMMAND\n")
		}
	}
}
