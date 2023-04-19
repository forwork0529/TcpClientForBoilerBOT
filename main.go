package main

import (
	"bufio"
	"fmt"
	"github.com/tarm/serial"
	"io/ioutil"
	"net"
	"os"
	"time"
)

func main() {

	port, err := ioutil.ReadFile("port.txt")

	c := &serial.Config{Name : string(port) ,Baud : 9600}

	s, err := serial.OpenPort(c)
	if err != nil {
		fmt.Printf("connection to the local %v: ERROR; %v\n", string(port), err)
		time.Sleep(time.Second * 3 )
		os.Exit(1)
	}
	fmt.Println("connection to the local port COM16: OK")

	reader := bufio.NewReader(s)

	// Подключаемся к сокету

	for {
		conn, err := net.Dial("tcp", "labhome.space:30002")
		if err != nil{
			fmt.Printf("connection to the remote server: ERROR: %v\n", err)
			time.Sleep(time.Second * 5)
			continue
		}
		fmt.Println("Connection to the remote server: OK")
		for{
			str, err := reader.ReadString('.')
			if err != nil {
				fmt.Printf("reading from comPort ERROR: %v\n", err)
			}

			n, err := fmt.Fprintf(conn, str)
			if err != nil || n < 1{
				fmt.Printf("cant write to remote server: %v", err)
				time.Sleep(time.Second * 3)
				break
			}else{
				fmt.Printf("send to server: %v\n", str)
			}
			time.Sleep(time.Second * 3)
		}

	}
}
