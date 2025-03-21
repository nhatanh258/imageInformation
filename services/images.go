package services

import (
	"GivingData/db"
	"GivingData/models"
   
	"errors"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetAllImage(c *gin.Context) error {
	var images []models.Image
	rows, err := db.DB.Query("SELECT id, url, path1, path2, text1, text2, width, uploaded_at, updated_at FROM images")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "should not query image table"})
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var img models.Image

		if err := rows.Scan(&img.ID, &img.URL, &img.Path1, &img.Path2, &img.Text1, &img.Text2, &img.Width, &img.UploadedAt, &img.UpdatedAt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return err
		}
		images = append(images, img)
	}
	c.JSON(http.StatusOK, images)
	return nil
}

func GetImageByID(id int64) (*models.Image, error) {
	var img models.Image
	query := `SELECT id, url, path1, path2, text1, text2, width , uploaded_at, updated_at FROM images WHERE id = ?`
	row := db.DB.QueryRow(query, id)

	err := row.Scan(&img.ID, &img.URL, &img.Path1, &img.Path2, &img.Text1, &img.Text2, &img.Width, &img.UploadedAt, &img.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &img, nil
}

// CreateImage tạo một hình ảnh mới
func Save(e *models.Image) error {
	if e == nil {
		return errors.New("image is nil")
	}
	query := `INSERT INTO images ( url, path1, path2, text1, text2, width) 
		  VALUES ( ?, ?, ?, ?, ?,?)`

	stmt, err := db.DB.Prepare(query) //Chuẩn bị câu lệnh SQL để thực thi.
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.URL, e.Path1, e.Path2 , e.Text1, e.Text2, e.Width)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId() // Lấy ID của bản ghi vừa được thêm vào database.
	if err != nil {
		return err
	}
	e.ID = int64(id) // ✅ Cập nhật ID cho event
	return nil
}

// UpdateImage cập nhật thông tin ảnh
func Update(id int64, apdateImage *models.Image) error {
	query := `UPDATE images SET url=?, path1=?, path2=?, text1=?, text2=?, width=? WHERE id=?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(apdateImage.URL, apdateImage.Path1, apdateImage.Path2, apdateImage.Text1,  apdateImage.Text2, apdateImage.Width, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("event not found")
	}
	return nil
}

// DeleteImage xóa ảnh theo ID
func Delete(id int64) error {
	query := `DELETE FROM images WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("event not found")
	}
	return nil
}

// GetText lấy văn bản liên quan đến ảnh
func GetText(id int64) (string, string, error) {
	query := "SELECT text FROM images WHERE id = ?"
	row := db.DB.QueryRow(query, id)
	var imag models.Image
	err := row.Scan(&imag.Text1, &imag.Text2)
	if err != nil {
		return "", "", err
	}
	return imag.Text1, imag.Text2, nil
}

// GetFullImageInfo lấy tất cả thông tin của ảnh
func GetFullImageInfo(id int64) (*models.Image, error) {
	var img models.Image
	query := `SELECT id, url, path1, path2, text1 , text2, width , uploaded_at, updated_at FROM images WHERE id = ?`
	row := db.DB.QueryRow(query, id)
	err := row.Scan(&img.ID, &img.URL, &img.Path1, &img.Path2, &img.Text1, &img.Text2, &img.Width, &img.UploadedAt, &img.UpdatedAt)
	if err != nil {
		panic(err.Error())
	}
	return &img, nil
}
func GetDouble(id int64) (*models.Image,*models.Image, error){
		var img1 models.Image
		var img2 models.Image 
	query1 := `SELECT id, url, path1, path2, text1, text2, width , uploaded_at, updated_at FROM images WHERE id = ?`
	row := db.DB.QueryRow(query1, id)

	err := row.Scan(&img1.ID, &img1.URL, &img1.Path1, &img1.Path2, &img1.Text1, &img1.Text2, &img1.Width, &img1.UploadedAt, &img1.UpdatedAt)
	//
	id = id + 1
	query2 := `SELECT id, url, path1, path2, text1, text2, width , uploaded_at, updated_at FROM images WHERE id = ?`
	row = db.DB.QueryRow(query2, id)
	err = row.Scan(&img2.ID, &img2.URL, &img2.Path1, &img2.Path2, &img2.Text1, &img2.Text2, &img2.Width, &img2.UploadedAt, &img2.UpdatedAt)
	if err != nil {
		return nil, nil, err
	}// 
	return &img1, &img2, nil
}