package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("template.txt"))
}

func main() {
	l, err := net.Listen("tcp", ":8080")
	defer l.Close()

	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := l.Accept()

		if err != nil {
			log.Fatalln(err)
			continue
		}

		// in := bufio.NewScanner(conn)
		// var str string

		// for in.Scan() {
		// 	str += in.Text()
		// 	if in.Text() == "" {
		// 		break
		// 	}
		// 	tpl.Execute(conn, str)
		// }

		body := "<!DOCTYPE html><header><title>Go Server</title></header><body><H1>My Body</H1></body></html>"

		io.WriteString(conn, "HTTP/1.1 200 OK\r\n")

		fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
		fmt.Fprint(conn, "Content-Type: text/html\r\n")

		io.WriteString(conn, "\r\n")

		io.WriteString(conn, body)

		// tpl.Execute(conn, str)

		conn.Close()
	}

}
