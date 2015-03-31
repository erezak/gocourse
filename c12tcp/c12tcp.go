package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:4444")
	if err != nil {
		fmt.Printf("Couldn't create address - %v\n", err)
		return
	}
	ln, err := net.ListenTCP("tcp", addr)
	if err != nil {
		fmt.Printf("Couldn't open listener - %v\n", err)
		return
	}
	defer ln.Close()

	fmt.Println("Listening for incoming connections.")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("Couldn't accept connection - %v\n", err)
		}
		go func(incoming net.Conn) {
			defer conn.Close()
			bsReader := bufio.NewReader(incoming)
			bs := make([]byte, 500)

			for {
				readBytes, err := bsReader.Read(bs)
				if err != nil {
					if err == io.EOF {
						fmt.Println("Client connection closed.")
					} else {
						fmt.Printf("Connection error - %v\n", err)
					}
					return
				}
				// Check for Ctrl-D
				if readBytes == 1 && bs[0] == 4 {
					fmt.Println("\nClient connection closed.")
					return
				}
				fmt.Print(string(bs[:readBytes]))
			}
		}(conn)
	}
}
