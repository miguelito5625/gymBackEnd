package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// MainRoutes define las rutas principales de la aplicación
func MainRoutes(r *gin.RouterGroup) {
	r.GET("/", func(c *gin.Context) {
		fmt.Println("Hola mundo!!!")
		c.String(200, "Hola, mundo!")
	})
}
