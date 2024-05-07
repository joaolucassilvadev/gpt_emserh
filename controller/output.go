package controller

import (
	"github.com/gin-gonic/gin"
)

// Handler para servir o arquivo HTML
func ServeIndex(ctx *gin.Context) {
	// Define o caminho do arquivo HTML
	htmlFilePath := "/home/joao/test-gpt/gpt_emserh/index.html"

	// Serve o arquivo HTML usando ctx.File
	ctx.File(htmlFilePath)
}
