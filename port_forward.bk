package main

import (
	"flag"
	"io"
	"log"
	"net"
)

var remoteServerHost string

// var localHP string

func main() {
	remoteH := flag.String("remote", "127.0.0.1", "Remote machine IP")
	remoteP := flag.String("rport", "8080", "Remote machine Port")

	localH := flag.String("lhost", "0.0.0.0", "Local IP")
	localP := flag.String("lport", "8080", "Local port")

	flag.Parse()

	localHP := *localH + ":" + *localP
	remoteServerHost = *remoteH + ":" + *remoteP

	log.Println(remoteServerHost)
	local, err := net.Listen("tcp", localHP)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("[+] Listening on: ", localHP)
	for {
		conn, err := local.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}

}

func forward(source net.Conn, dest net.Conn) {
	defer dest.Close()
	defer source.Close()
	io.Copy(source, dest)
}

func handleConn(local net.Conn) {
	log.Println("Connecting from: ", local.RemoteAddr())

	remote, err := net.Dial("tcp", remoteServerHost)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to: ", remoteServerHost)

	go forward(local, remote)
	go forward(remote, local)
}
