package models

import (
	//"time"
)

type Image struct {
	ID         int64     `json:"id"` // ID ảnh (Primary Key)
	URL        string    `json:"url"`
	Path       string    `json:"path" binding:"required"`
	Text       string    `json:"text" binding:"required"`  // Văn bản liên quan đến ảnh
	Width      int32     `json:"width" binding:"required"` //
	UploadedAt string `json:"uploaded_at"`              // Xử lý NULL timestamp
	UpdatedAt  string `json:"updated_at"`               // Thời gian cập nhật gần nhất
}
