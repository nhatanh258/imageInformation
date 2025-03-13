package main

import (
	"GivingData/db"
	"GivingData/routers"
	"encoding/json"
	"fmt"
	"os"
    
	"github.com/gin-gonic/gin"
)

// Cấu trúc dữ liệu cho file cấu hình
type Config struct {
	Port     int    `json:"port"`
	Database string `json:"database"`
	LogLevel string `json:"log_level"`
}

// Hàm load cấu hình từ file JSON
func LoadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// Chương trình chính
func main() {
	database, _ := db.InitDB()
	defer database.Close()
	fmt.Println("dang tien hanh lay anh ...")
	server := gin.Default() // tao mot instance  gin router
	routers.SetupRoutes(server)
	server.Run(":8080")
}
