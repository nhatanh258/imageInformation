package routers

import (
	"GivingData/models"
	"GivingData/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAllImages lấy tất cả hình ảnh từ database
func GetAllImages(c *gin.Context) {
	if err := services.GetAllImage(c); err != nil {
		return
	}
}


// GetImage lấy thông tin ảnh theo ID
func GetImage(c *gin.Context) {
	idInput, err := strconv.ParseInt(c.Param("id"), 10, 64) // chuyen doi id tu string sang int64
	if err != nil || idInput < 0 {                         // neu id nhap vao ko hop le hoac nho hon 0
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}
	image, err := services.GetImageByID(idInput)
	if err != nil { // neu co loi trong lay du lieu tu database
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}
	c.JSON(http.StatusOK, image)
}

// CreateImage tạo mới ảnh và lưu vào database
func CreateImage(c *gin.Context) {

	var img models.Image
	// lay doi tuong da quy dinh trong event struct
	err := c.ShouldBindJSON(&img) // nhan du lieu tu client dang json nhu quet fmt tu stdin
	if err != nil {               // neu co loi trong nhap content tu client
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = services.Save(&img)
	if err != nil { // neu co loi trong luu du lieu vao database
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save event to database"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Event created successfully"})
}

// UpdateImage cập nhật thông tin ảnh theo ID

func UpdateImage(c *gin.Context) {
	IdInput, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil || IdInput < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	var updated models.Image

	image, err := services.GetImageByID(IdInput)
	if err != nil { // neu co loi trong lay du lieu tu database
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	if IdInput != int64(image.ID) { // neu userId truyen vao khong phai userId cua event
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized to update this event"})
		return
	}
	// luu lai event duoc cap nhat
	err = c.ShouldBindJSON(&updated)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updated.ID = IdInput // do day la cap naht su kien len van giu nguyen  id

	err = services.Update(IdInput, &updated)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update event in database"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Event updated successfully", "event": updated})
}

//DeleteImage c

func DeleteImage(c *gin.Context) {
	IdInput, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil || IdInput <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	err = services.Delete(IdInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete event from database"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})

}

//get text

func GetText(c *gin.Context) {
	idInput, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil || idInput <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image ID"})
		return
	}

	text, err := services.GetText(idInput)
	if err != nil { // neu co loi trong lay du lieu tu database
		c.JSON(http.StatusNotFound, gin.H{"error": "image not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Text of image", "text": text})
}

//get all inforf

func GetFullImageInfo(c *gin.Context) {
	idInput, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil || idInput <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid image ID"})
		return
	}

	image, err := services.GetFullImageInfo(idInput)
	if err != nil { // neu co loi trong lay du lieu tu database
		c.JSON(http.StatusNotFound, gin.H{"error": "image not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Full info of image", "image": image})
}
