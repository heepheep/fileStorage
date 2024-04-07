package server

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	err := router.Run(":8081")
	if err != nil {
		return
	}
}
