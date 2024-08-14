package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    // Crea una instancia del enrutador Gin
    r := gin.Default()
	

	 // Ruta con parámetros de la URL
	 r.GET("/usuarios/:id", func(c *gin.Context) {
        // Recibir parámetros de la URL
        userID := c.Param("id")
        c.String(200, "Usuario ID: %s", userID)
    

})
}

