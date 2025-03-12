package routers

import (
	"github.com/gin-gonic/gin"
	
)

func SetupRoutes(server *gin.Engine) {
	imagesGroup := server.Group("/images")
	         // Tạo middleware để kiểm tra quyền truy cập
	// xu ly anh router
	imagesGroup.POST("/create", CreateImage)
	imagesGroup.GET("/getAllImage", GetAllImages)
	imagesGroup.GET("/:id", GetImage)
	imagesGroup.PUT("/update/:id", UpdateImage)
	imagesGroup.DELETE("/delete/:id", DeleteImage)
	imagesGroup.GET("/text/:id", GetText)         
	imagesGroup.GET("/fulInform/:id", GetFullImageInfo)   // Lấy tất cả thông tin ảnh


}
