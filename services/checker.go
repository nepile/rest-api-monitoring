package services

import (
	"io"
	"net/http"
	"time"

	"github.com/nepile/api-monitoring/database"
	"github.com/nepile/api-monitoring/models"
)

func CheckEndpoint(e models.Endpoint) (int, float64, error) {
	start := time.Now()
	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	res, err := client.Get(e.URL)
	duration := time.Since(start).Seconds() * 1000
	if err != nil {
		database.DB.Create(&models.CheckLog{
			EndpointID:   e.ID,
			StatusCode:   0,
			ResponseTime: duration,
			CheckedAt:    time.Now(),
		})
		return 0, duration, err
	}

	_, _ = io.Copy(io.Discard, res.Body)
	res.Body.Close()

	database.DB.Create(&models.CheckLog{
		EndpointID:   e.ID,
		StatusCode:   res.StatusCode,
		ResponseTime: duration,
		CheckedAt:    time.Now(),
	})

	return res.StatusCode, duration, nil
}
