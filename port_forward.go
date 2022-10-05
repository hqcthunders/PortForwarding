package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

var remoteServerHost string

// var localHP string

func main() {
	banner("materials/go.txt")
	menu()
}

func banner(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func menu() {
	var option int
	banner("materials/menu.txt")
	fmt.Print("Input your option: ")
	fmt.Scan(&option)

	switch option {
	case 1:
		fmt.Println("Port forwarding configuration")
	case 2:
		fmt.Println("Start")
		fmt.Println()

		var remoteH, remoteP, localH, localP string

		fmt.Print("Remote host: ")
		fmt.Scan(&remoteH)
		fmt.Print("Remote port: ")
		fmt.Scan(&remoteP)
		fmt.Print("Local host: ")
		fmt.Scan(&localH)
		fmt.Print("Local port: ")
		fmt.Scan(&localP)

		localHP := localH + ":" + localP
		remoteServerHost = remoteH + ":" + remoteP
		Launch(localHP)
	default:
		os.Exit(0)
	}
}

func Launch(localHP string) {
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
