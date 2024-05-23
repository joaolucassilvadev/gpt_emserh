package service

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/sashabaranov/go-openai"
	"github.com/spf13/viper"
)

func Gpt(b string) string {
	var output string

	// Lê a chave da API do arquivo de configuração
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Erro ao ler arquivo de configuração: %v", err)
	}

	apikey := viper.GetString("API_KEY")
	if apikey == "" {
		log.Fatal("Chave da API não encontrada")
	}

	// Cria um cliente OpenAI com a chave da API
	client := openai.NewClient(apikey)

	// Lê o conteúdo do arquivo de entrada
	inputFile := "./input_with_code.txt"
	filebytes, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("Falha ao ler arquivo: %v", err)
	}

	// Define o contexto para a solicitação
	contexto := fmt.Sprintf("Explique sobre o texto:\n\n%s\n\n", string(filebytes))

	// Cria um contexto de execução
	ctx := context.Background()

	// Cria a solicitação de chat com o modelo
	req := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo0125,
		Messages: []openai.ChatCompletionMessage{
			{Role: openai.ChatMessageRoleSystem, Content: contexto},
			{Role: openai.ChatMessageRoleUser, Content: b},
		},
		MaxTokens:   50,
		Temperature: 0.5,
	}

	// Envia a solicitação de chat ao modelo
	resp, err := client.CreateChatCompletion(ctx, req)
	if err != nil {
		log.Fatalf("Erro ao obter resposta do modelo: %v", err)
		return "Erro ao obter resposta"
	}

	// Exibe a resposta ao usuário
	output = resp.Choices[0].Message.Content
	// fmt.Printf("Resposta: %s\n", output)

	return output
}
