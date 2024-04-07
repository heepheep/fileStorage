package server

import "github.com/gin-gonic/gin"

func Main() {
	router := gin.Default()

	api := router.Group("/api")
	{
		file := api.Group("/file")
		{
			file.GET("/all")    // in progress
			file.GET("/:id")    // in progress
			file.POST("/")      // in progress
			file.PUT("/:id")    // in progress
			file.DELETE("/:id") // in progress
		}
		auth := api.Group("/auth")
		{
			auth.POST("/registration") // in progress
			auth.PUT("/change-pass")   // in progress
			auth.PUT("/change-email")  // in progress
			auth.POST("/login")        // in progress
			auth.POST("/logout")       // in progress
		}
	}
	err := router.Run(":8081")
	if err != nil {
		return
	}

}
