package main

import (
	"log"
	"net"
	"strings"
)

var clients = make(map[net.Conn]string)
var messages = make(chan string)

func broadcast(message string, sender net.Conn) {
	for client := range clients {
		if client != sender {
			_, err := client.Write([]byte(message))
			if err != nil {
				log.Printf("Erro ao enviar mensagem para %s: %s", clients[client], err)
			}
		}
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	nickname := "Anonimo"

	clients[conn] = nickname
	log.Printf("%s entrou na sala.", nickname)

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			delete(clients, conn)
			log.Printf("%s saiu da sala.", nickname)
			broadcast(nickname+" saiu da sala.", conn)
			break
		}

		message := string(buffer[:n])
		message = strings.TrimSpace(message)

		if strings.HasPrefix(message, "/nick") {
			newNickname := strings.TrimSpace(strings.TrimPrefix(message, "/nick"))
			clients[conn] = newNickname
			broadcast(nickname+" mudou o apelido para "+newNickname, conn)
			nickname = newNickname
		} else if strings.HasPrefix(message, "/private") {
			parts := strings.Fields(message)
			if len(parts) >= 3 {
				targetNickname := parts[1]
				targetMessage := strings.Join(parts[2:], " ")
				for client := range clients {
					if clients[client] == targetNickname {
						_, err := client.Write([]byte(nickname + " sussurrou para você: " + targetMessage))
						if err != nil {
							log.Printf("Erro ao enviar mensagem para %s: %s", targetNickname, err)
						}
						break
					}
				}
			}
		} else {
			broadcast(nickname+": "+message, conn)
			log.Printf("%s: %s", nickname, message)
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	log.Println("Servidor de chat iniciado em localhost:3000")

	go func() {
		for message := range messages {
			log.Println(message)
		}
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleClient(conn)
	}
}
