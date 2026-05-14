package web

import (
	gorrm "RatWAF/database"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetData(c *gin.Context) {
	waf_logs := []gorrm.WafLog{}

	err := gorrm.DB.Table("waf_log").Order("id desc").Select("id", "attacker_ip", "full_url", "createtime", "attack_type").Find(&waf_logs).Error
	if err != nil {
		c.JSON(500, gin.H{
			"err": err.Error(),
		})
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"log": waf_logs,
	})
	// c.JSON(http.StatusOK, gin.H{
	// 	"log": waf_logs,
	// })
}

func GetDetail(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	waf_log := gorrm.WafLog{}
	err = gorrm.DB.Table("waf_log").Order("id desc").Where("id=?", id).First(&waf_log).Error
	fmt.Println("waf_log", waf_log)
	if err != nil {
		fmt.Println("errr", err)
		c.JSON(404, gin.H{"log": err})
		return
	}
	c.HTML(http.StatusOK, "detail.html", gin.H{
		"log": waf_log,
	})

	// c.JSON(http.StatusOK, gin.H{
	// 	"log": waf_log,
	// })
}
