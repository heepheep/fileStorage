package server

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/file/:id", func(c *gin.Context) {

	})

	err := router.Run(":8081")
	if err != nil {
		return
	}

}
