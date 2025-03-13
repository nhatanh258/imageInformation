package models

//"time"

type Image struct {
	ID         int64  `json:"id"` // ID ảnh (Primary Key)
	URL        string `json:"url"`
	Path1      string `json:"path1" binding:"required"`
	Path2      string `json:"path2" binding:"required"`
	Text1      string `json:"text1" binding:"required"` // Văn bản liên quan đến ảnh
	Text2      string `json:"text2" binding:"required"`
	Width      int32  `json:"width" binding:"required"` //
	UploadedAt string `json:"uploaded_at"`              // Xử lý NULL timestamp
	UpdatedAt  string `json:"updated_at"`               // Thời gian cập nhật gần nhất
}
