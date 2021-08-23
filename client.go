package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

func main() {
	fmt.Print("enter a username:")
	var user string = scanner()
	conn, err := net.Dial("tcp", "localhost:9090")
	fmt.Println("connection established")
	var wg sync.WaitGroup
	if err != nil{
		panic("dial up error")
	}
	defer conn.Close()
	wg.Add(1)
	go readMessage(wg, conn)
	go writeMessage(conn, user)
	wg.Wait()
}

func scanner() string {
	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')
	line = strings.Replace(line, "\n", "", strings.LastIndex(line, "\n"))
	return line
}

func readMessage(wg sync.WaitGroup, conn net.Conn)  {
	defer wg.Done()
	for true {
		data := make([]byte, 1024)
		_, err := conn.Read(data)
		if err != nil {
			panic(err)
		}
		if string(data) == "end" {
			break
		}
		fmt.Println(string(data))
	}
}

func writeMessage(conn net.Conn, username string)  {
	for true {
		temp := scanner()
		message := username + ":" + temp
		conn.Write([]byte(message))
	}
}


