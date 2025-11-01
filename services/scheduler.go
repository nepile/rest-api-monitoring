package services

import (
	"log"
	"time"

	"github.com/nepile/api-monitoring/config"
	"github.com/nepile/api-monitoring/database"
	"github.com/nepile/api-monitoring/models"
	"github.com/nepile/api-monitoring/utils"
)

func StartScheduler(cfg *config.Config) {
	log.Println("Scheduler started")
	go func() {
		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()
		for {
			<-ticker.C
			runChecks(cfg)
		}
	}()
}

func runChecks(cfg *config.Config) {
	var endpoints []models.Endpoint
	if err := database.DB.Find(&endpoints).Error; err != nil {
		log.Println("failed fetch endpoints:", err)
		return
	}

	for _, ep := range endpoints {
		go func(e models.Endpoint) {
			var last models.CheckLog
			if err := database.DB.Where("endpoint_id = ?", e.ID).Order("checked_at desc").First(&last).Error; err != nil {
				elapsed := time.Since(last.CheckedAt).Seconds()
				if int(elapsed) < e.CheckInterval {
					return
				}
			}

			status, rt, err := CheckEndpoint(e)
			if err != nil || status != e.ExpectedStatus {
				msg := ""
				if err != nil {
					msg = "<b>API DOWN</b>\n" +
						"URL: " + e.URL + "\n" +
						"Error: " + err.Error() + "\n" +
						"Response Time: " + utils.FormatMs(rt)
				} else {
					msg = "<b>Unexpected Status</b>\n" +
						"URL: " + e.URL + "\n" +
						"Got: " + utils.Itoa(status) + " Expected: " + utils.Itoa(e.ExpectedStatus) + "\n" +
						"Response Time: " + utils.FormatMs(rt)
				}
				if err := SendTelegramAlert(cfg, msg); err != nil {
					log.Println("telegram send error:", err)
				}
			}
		}(ep)
	}
}
