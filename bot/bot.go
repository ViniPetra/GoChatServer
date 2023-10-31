package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strings"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:3000")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	fmt.Println("Conectado como BobinhoBot. Você pode usar os comandos: /oi, /piada, /hora")

	go func() {
		reader := bufio.NewReader(conn)
		for {
			message, err := reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
				break
			}

			message = strings.TrimSpace(message)

			if strings.HasPrefix(message, "/oi") {
				fmt.Fprintln(conn, "BobinhoBot diz: Oi!")
			} else if strings.HasPrefix(message, "/piada") {
				joke := getRandomJoke()
				fmt.Fprintln(conn, "BobinhoBot conta uma piada:", joke)
			} else if strings.HasPrefix(message, "/hora") {
				currentTime := getCurrentTime()
				fmt.Fprintln(conn, "BobinhoBot diz: A hora atual é", currentTime)
			} else if strings.HasPrefix(message, "/inverte") {
				reversedText := reverseText(strings.TrimPrefix(message, "/inverte"))
				fmt.Fprintln(conn, "BobinhoBot inverteu o texto: "+reversedText)
			}
		}
	}()
	select {}
}

func getCurrentTime() string {
	currentTime := time.Now()
	return currentTime.Format("15:04:05")
}

func getRandomJoke() string {
	jokes := []string{
		"Por que o computador foi ao médico? Porque estava com vírus!",
		"O que o zero disse para o oito? Estou adorando o seu cinto!",
		"O que o advogado do frango fez? Foi bicar a papeleta.",
		"Por que o esqueleto não brigou com ninguém? Porque não tinha estômago para isso.",
		"Qual é o cúmulo da burrice? Dois carecas brigando por um pente.",
		"Por que a matemática é como o amor? Uma ideia simples, mas pode ficar complicada.",
		"O que um dente falou para o outro? Eu já vou, você fica!",
		"Por que o livro de matemática ficou triste? Porque tinha muitos problemas.",
		"O que a impressora falou para o papel? Esse é o seu fim!",
		"O que o canibal vegetariano come? Cabeças de alface.",
		"Como se faz para amarrar um elefante? Basta usar uma corda grande e um nó bem apertado.",
		"Por que o pássaro não usou o computador? Porque ele já tinha twitter.",
		"O que a foca disse para o pinguim? Nada, porque foca não fala.",
		"O que o tomate foi fazer no banco? Foi tirar extrato.",
		"Como o esqueleto liga para os amigos? Com o táxi-ossos.",
		"O que o cinto falou para a calça? Você segura as pontas que eu entro por onde dá!",
		"O que um prédio disse para o outro? Nossa, como você é alto!",
		"O que um elevador disse para o outro? Estou contigo até o último degrau.",
		"O que o zero disse para o oito? Bonito cinto!",
		"Por que o livro de história estava chateado? Porque ele tinha muitos passados.",
	}

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(jokes))
	return jokes[randomIndex]
}

func reverseText(text string) string {
	runes := []rune(text)
	reversedText := make([]rune, len(runes))
	for i, j := 0, len(runes)-1; i < len(runes); i, j = i+1, j-1 {
		reversedText[i] = runes[j]
	}
	return string(reversedText)
}
