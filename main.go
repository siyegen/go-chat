package main

import (
	"bufio"
	"io"
	"log"
	"net"
)

type Server struct{}

func (s *Server) Serve(l net.Listener) error {

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go s.handle(conn)
	}
}

func (s *Server) handle(c net.Conn) {
	log.Println("New Connection")

	bufRead := bufio.NewReader(c)
	for {
		msg, _ := bufRead.ReadString('\n')
		io.WriteString(c, msg)
	}
}

func main() {
	l, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatal("unable to listen", err)
	}

	s := &Server{}

	log.Println("starting server")
	log.Fatal(s.Serve(l))
}
