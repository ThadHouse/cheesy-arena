package forwarder

import (
	"io"
	"log"
	"net"
)

func ForwardAllAccess() {
	go ForwardRadioAccess()
	go ForwardSwitchAccess()
}

func HandleForward(local net.Conn, addr string) {
	forwarded, err := net.Dial("tcp", addr)
	if err != nil {
		defer local.Close()
		log.Printf("Failed to connect to forwarder %s: %v", addr, err)
		return
	}

	done := make(chan struct{})

	go func() {
		defer local.Close()
		defer forwarded.Close()
		io.Copy(forwarded, local)
		done <- struct{}{}
	}()

	go func() {
		defer local.Close()
		defer forwarded.Close()
		io.Copy(local, forwarded)
		done <- struct{}{}
	}()

	<-done
	<-done
}

func ForwardRadioAccess() {
	listener, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Printf("Failed to start radio forwarder: %v", err)
		return
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Radio failed accept: %v", err)
			continue
		}
		go HandleForward(conn, "10.0.100.2:80")
	}
}

func ForwardSwitchAccess() {
	listener, err := net.Listen("tcp", ":8083")
	if err != nil {
		log.Printf("Failed to start telnet forwarder: %v", err)
		return
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Telnet failed accept: %v", err)
			continue
		}
		go HandleForward(conn, "10.0.100.3:23")
	}
}
