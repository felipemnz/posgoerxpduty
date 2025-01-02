package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Estrutura para armazenar os dados retornados pela API
type Address struct {
	Cep        string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Bairro     string `json:"bairro"`
	Localidade string `json:"localidade"`
	Uf         string `json:"uf"`
}

// Função para fazer a requisição a uma API específica
func fetchFromAPI(ctx context.Context, url string, resultChan chan<- Address, sourceChan chan<- string) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Println("Erro ao criar requisição:", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Erro ao fazer requisição:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("Erro: Status code inválido:", resp.StatusCode)
		return
	}

	var address Address
	err = json.NewDecoder(resp.Body).Decode(&address)
	if err != nil {
		log.Println("Erro ao decodificar resposta:", err)
		return
	}

	resultChan <- address
	sourceChan <- url
}

func main() {
	cep := "01153000"
	api1 := "https://brasilapi.com.br/api/cep/v1/" + cep
	api2 := "http://viacep.com.br/ws/" + cep + "/json/"

	// Contexto com timeout de 1 segundo
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// Canais para comunicação entre goroutines
	resultChan := make(chan Address, 1)
	sourceChan := make(chan string, 1)

	// Inicia as requisições simultaneamente
	go fetchFromAPI(ctx, api1, resultChan, sourceChan)
	go fetchFromAPI(ctx, api2, resultChan, sourceChan)

	select {
	case address := <-resultChan:
		source := <-sourceChan
		fmt.Println("Resultado recebido da API:", source)
		fmt.Printf("Endereço: %+v\n", address)
	case <-ctx.Done():
		fmt.Println("Erro: Timeout ao buscar as informações do CEP.")
	}
}