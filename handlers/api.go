package handlers

import (
    "net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/Tonyrealzy/Basic-Information-Public-API/config"
)

// Define a struct to maintain key order
type BasicInfoResponse struct {
	Email            string `json:"email"`
	CurrentDatetime string `json:"current_datetime"`
	GithubURL        string `json:"github_url"`
}

func GetBasicInfo(c *gin.Context) {
	response := BasicInfoResponse{
		Email:            config.Email,
		CurrentDatetime: time.Now().UTC().Format(time.RFC3339),
		GithubURL:        config.GithubURL,
	}
	c.JSON(http.StatusOK, response)
}

// func GetBasicInfo(c *gin.Context) {
// 	response:= map[string]string {
// 		"email": config.Email,
// 		"current_datetime": time.Now().UTC().Format(time.RFC3339),
// 		"github_url": config.GithubURL,
// 	}
// 	c.JSON(http.StatusOK, response)
// }