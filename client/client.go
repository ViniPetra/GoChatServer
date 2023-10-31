package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:3000")
	fmt.Println("Conectado!")
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})

	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("Concluído")
		done <- struct{}{}
	}()

	for {
		fmt.Print("Digite uma mensagem ou comando (/private <nick> <mensagem>, /exit, /nick <novo_apelido>): ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "/exit" {
			conn.Close()
			break
		} else if strings.HasPrefix(input, "/private") {
			parts := strings.Fields(input)
			if len(parts) < 3 {
				fmt.Println("Formato incorreto. Use /private <nick> <mensagem>")
				continue
			}
			targetNickname := parts[1]
			message := strings.Join(parts[2:], " ")
			fmt.Fprintf(conn, "/private %s %s\n", targetNickname, message)
		} else if strings.HasPrefix(input, "/nick") {
			newNickname := strings.TrimSpace(strings.TrimPrefix(input, "/nick"))
			fmt.Fprintf(conn, "/nick %s\n", newNickname)
		} else {
			fmt.Fprintf(conn, "%s\n", input)
		}
	}

	<-done
}
