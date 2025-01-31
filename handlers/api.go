package handlers

import (
    "net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/Tonyrealzy/Basic-Information-Public-API/config"
)

func GetBasicInfo(c *gin.Context) {
	response:= map[string]string {
		"email": config.Email,
		"current_timestamp": time.Now().UTC().Format(time.RFC3339),
		"github_url": config.GithubURL,
	}
	c.JSON(http.StatusOK, response)
}