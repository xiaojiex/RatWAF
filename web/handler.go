package web

import (
	gorrm "RatWAF/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetData(c *gin.Context) {
	waf_logs := []gorrm.WafLog{}

	err := gorrm.DB.Table("waf_log").Select("attacker_ip", "full_url", "createtime").Find(&waf_logs).Error
	if err != nil {
		c.JSON(500, gin.H{
			"err": err.Error(),
		})
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"log": waf_logs,
	})
}

func GetDetail(c *gin.Context) {
	id := c.Param("id")
	waf_log := gorrm.WafLog{}
	err := gorrm.DB.Table("waf_log").First(&waf_log).Where("id=?", id).Error
	if err != nil {
		c.JSON(404, gin.H{"log": err})
	}
	c.HTML(http.StatusOK, "detail.html", gin.H{
		"log": waf_log,
	})
}
