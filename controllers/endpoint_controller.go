package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nepile/api-monitoring/database"
	"github.com/nepile/api-monitoring/models"
)

type AddEndpointRequest struct {
	URL            string `json:"url" binding:"required,url"`
	ExpectedStatus int    `json:"expected_status"`
	CheckInterval  int    `json:"check_interval"`
}

func AddEndpoint(c *gin.Context) {
	userID := c.GetString("user_id")
	uid, _ := uuid.Parse(userID)

	var body AddEndpointRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ep := models.Endpoint{
		UserID:         uid,
		URL:            body.URL,
		ExpectedStatus: body.ExpectedStatus,
		CheckInterval:  body.CheckInterval,
	}
	if ep.CheckInterval == 0 {
		ep.CheckInterval = 60
	}

	if err := database.DB.Create(&ep).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"endpoint": ep})
}

func ListEndpoints(c *gin.Context) {
	userID := c.GetString("user_id")
	var eps []models.Endpoint
	database.DB.Where("user_id = ?", userID).Find(&eps)
	c.JSON(http.StatusOK, gin.H{"endpoints": eps})
}

func GetEndpointLogs(c *gin.Context) {
	id := c.Param("id")
	var logs []models.CheckLog
	database.DB.Where("endpoint_id = ?", id).Order("checked_at desc").Limit(50).Find(&logs)
	c.JSON(http.StatusOK, gin.H{"logs": logs})
}
