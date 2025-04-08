package main

import (
	"fmt"
	"time"
)

func enviar(c chan string, message string) {
	c <- message
}

func imprimir(c chan string) {
	msg := <-c
	time.Sleep(1 * time.Second)
	fmt.Println(msg)
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	for range 20 {
		go enviar(c1, "ping")
		go enviar(c2, "pong")

		imprimir(c1)
		imprimir(c2)
	}
}
