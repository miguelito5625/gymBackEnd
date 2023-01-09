package routes

import "github.com/gin-gonic/gin"

// SystemRoutes define todas las rutas del sistema
func SystemRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		AuthRoutes(api)
		MainRoutes(api)
	}
}
