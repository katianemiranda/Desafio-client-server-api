package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		panic(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	io.Copy(os.Stdout, res.Body)

	file, err := os.Create("cotacao.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// colocar valor correto para o dolar aqui
	_, err = file.WriteString(fmt.Sprintf("Dolar: {%v}", string(body)))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao escrever arquivo: %v\n", err)
	}
}
