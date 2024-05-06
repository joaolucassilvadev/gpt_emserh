package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/sashabaranov/go-openai"
	"github.com/spf13/viper"
)

func main() {
	// Define o arquivo de configuração
	viper.SetConfigFile(".env")

	// Lê as configurações do arquivo
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Erro ao ler arquivo de configuração: %v", err)
	}

	// Obtém a chave da API do arquivo de configuração
	apikey := viper.GetString("API_KEY")

	// Verifica se a chave da API está vazia
	if apikey == "" {
		log.Fatal("Chave da API não encontrada")
	}

	// Cria um cliente OpenAI com a chave da API
	client := openai.NewClient(apikey)

	// Define o caminho para o arquivo de entrada
	const inputFile = "./input_with_code.txt"

	// Lê o conteúdo do arquivo de entrada
	filebytes, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("Falha ao ler arquivo: %v", err)
	}

	// Concatena o conteúdo do arquivo com prefixo e sufixo
	msgPrefix := "Explique sobre o texto\n"
	msgSuffix := "\n```"

	contexto := msgPrefix + string(filebytes) + msgSuffix

	// Cria um contexto para a execução das operações
	ctx := context.Background()

	// Inicia a interação com o usuário
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Pergunte algo sobre a Emserh:")
	for {
		fmt.Print("> ")
		// Lê a pergunta do usuário
		pergunta, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Erro ao ler pergunta: %v", err)
		}
		// Remove espaços em branco
		pergunta = strings.TrimSpace(pergunta)

		// Se o usuário digitar "sair", encerre o programa
		if strings.EqualFold(pergunta, "sair") {
			fmt.Println("Encerrando o programa...")
			break
		}

		// Cria uma solicitação de chat com o modelo
		req := openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo, // Pode ajustar o modelo se necessário
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: pergunta,
				},
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: contexto, // Passa o contexto do código
				},
			},
			MaxTokens:   1000,
			Temperature: 0.9,
		}

		// Faz a solicitação de chat ao modelo
		resp, err := client.CreateChatCompletion(ctx, req)
		if err != nil {
			log.Fatalf("Erro ao obter resposta do modelo: %v", err)
		}

		// Exibe a resposta ao usuário
		fmt.Println("Resposta:")
		fmt.Println(resp.Choices[0].Message.Content)
	}

	fmt.Println("Programa encerrado.")
}
