package routes

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/miguelito5625/gymBackend/middlewares"
)

// AuthRoutes define las rutas de autenticación
func AuthRoutes(r *gin.RouterGroup) {
	auth := r.Group("/auth")
	{
		auth.POST("/login", middlewares.PrintCredentialsMiddleware, func(c *gin.Context) {
			// Obtener la hora actual
			horaSesion := time.Now().Format(time.RFC3339)
			// Preparar el objeto a devolver en la respuesta
			response := gin.H{
				"mensaje":     "Sesión iniciada",
				"hora_sesion": horaSesion,
			}
			// Devolver el objeto en formato JSON
			c.JSON(200, response)
		})
	}
}
