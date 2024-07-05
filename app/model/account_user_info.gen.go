package model

import (
	"time"
)

const TableNameAccountUserInfo = "account_user_info"

// AccountUserInfo mapped from table <account_user_info>
type AccountUserInfo struct {
	ID       int64      `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name     string     `gorm:"column:name" json:"name"`
	Address  string     `gorm:"column:address" json:"address"`
	CreateAt *time.Time `gorm:"column:create_at" json:"create_at"`
	Type     int32      `gorm:"column:type" json:"type"`
	IP       string     `gorm:"column:ip" json:"ip"`
	Kyc      int32      `gorm:"column:kyc" json:"kyc"`
}

// TableName AccountUserInfo's table name
func (*AccountUserInfo) TableName() string {
	return TableNameAccountUserInfo
}
