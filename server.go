package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/miguelito5625/gymBackend/routes"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	routes.SystemRoutes(r)
	r.Run()
}
