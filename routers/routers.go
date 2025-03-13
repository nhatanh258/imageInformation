package routers

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(server *gin.Engine) {

	// Cấu hình CORS
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // Cho phép truy cập từ frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	imagesGroup := server.Group("/images")
	// Tạo middleware để kiểm tra quyền truy cập
	// xu ly anh router
	imagesGroup.POST("/create", CreateImage)
	imagesGroup.GET("/getAllImage", GetAllImages)
	imagesGroup.GET("/getDoubleImage/:id", getDoubleImage)
	imagesGroup.GET("/:id", GetImage)
	imagesGroup.PUT("/update/:id", UpdateImage)
	imagesGroup.DELETE("/delete/:id", DeleteImage)
	imagesGroup.GET("/text/:id", GetText)
	imagesGroup.GET("/fulInform/:id", GetFullImageInfo) // Lấy tất cả thông tin ảnh

}
