package gorrm

import (
	"fmt"
	"time"
)

type WafLog struct {
	ID         uint      `gorm:"primaryKey;autoIncrement"`
	AttackerIP string    `gorm:"size:45;not null"`
	Request    string    `gorm:"type:mediumtext;not null"`
	Createtime time.Time `gorm:"type:datetime;autoCreateTime"`
	FullUrl    string    `gorm:"column:full_url;type:text"`
	AttackType string    `gorm:"column:attack_type;size:60;not null"`
}

func InsertWafLog(AttackerIp, Request, FullUrl, AttackType string) {

	log := WafLog{
		AttackerIP: AttackerIp,
		Request:    Request,
		FullUrl:    FullUrl,
		AttackType: AttackType,
	}

	err := DB.Table("waf_log").Create(&log).Error
	if err != nil {
		fmt.Println("插入表错误 model.go")
	}
}
