package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func PrintCredentialsMiddleware(c *gin.Context) {
	// Crea una instancia de la estructura
	var credentials Credentials

	// Parsea el cuerpo de la solicitud y almacena los datos en la instancia de la estructura
	if err := c.ShouldBindJSON(&credentials); err != nil {
		fmt.Println("Error al parsear el cuerpo de la solicitud:", err)
		return
	}

	// Imprime los valores en pantalla
	fmt.Println("Credenciales enviadas:", credentials.Username, credentials.Password)

	// Ejecuta el siguiente middleware o handler
	c.Next()
}

// PrintCredentialsMiddleware imprime las credenciales enviadas
// func PrintCredentialsMiddleware(c *gin.Context) {
// 	username := c.PostForm("username")
// 	password := c.PostForm("password")
// 	fmt.Println("Credenciales enviadas:", username, password)
// 	c.Next()
// }
