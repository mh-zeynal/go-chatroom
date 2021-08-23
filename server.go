package main

import (
	"fmt"
	"net"
)

func main() {
	connections := make(map[net.Conn]bool, 0)
	li, err := net.Listen("tcp", ":9090")
	fmt.Println("now server is listening...")
	if err != nil{
		panic("something is wrong with listener")
	}
	launchTheServer(li, connections)
}

func launchTheServer(li net.Listener, connections map[net.Conn]bool)  {
	for {
		conn, _ := li.Accept()
		connections[conn] = true
		go func(){
			defer conn.Close()
			for {
				data := make([]byte, 1024)
				length, err := conn.Read(data)
				fmt.Println(string(data))
				if err != nil{
					connections[conn] = false
					delete(connections, conn)
					fmt.Println("client disconnected")
					break
				}
				for k, v := range connections {
					if v == true {
						_, error := k.Write(data[:length])
						if error != nil {
							panic(error)
						}
					}
				}
			}
		}()
	}
}
