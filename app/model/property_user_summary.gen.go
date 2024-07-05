package model

import (
	"time"
)

const TableNamePropertyUserSummary = "property_user_summary"

// PropertyUserSummary mapped from table <property_user_summary>
type PropertyUserSummary struct {
	ID                     int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	PropertyInfoID         int64     `gorm:"column:property_info_id" json:"property_info_id"`
	UserID                 int64     `gorm:"column:user_id" json:"user_id"`
	Address                string    `gorm:"column:address" json:"address"`
	Amount                 int64     `gorm:"column:amount" json:"amount"`
	StartTime              time.Time `gorm:"column:start_time;not null;default:CURRENT_TIMESTAMP;comment:用户账户余额汇总表，每有一条流水触发，需要新增流水和根据流水汇总余额信息，统计房租的时候要根据每个人不同时间段所持有的份额进行统计" json:"start_time"` // 用户账户余额汇总表，每有一条流水触发，需要新增流水和根据流水汇总余额信息，统计房租的时候要根据每个人不同时间段所持有的份额进行统计
	EndTime                time.Time `gorm:"column:end_time" json:"end_time"`
	Status                 int64     `gorm:"column:status;comment:状态1.有效 持有中" json:"status"` // 状态1.有效 持有中
	PropertyUserSummarycol string    `gorm:"column:property_user_summarycol" json:"property_user_summarycol"`
	TxHash                 string    `gorm:"column:tx_hash" json:"tx_hash"`
	EventSignHash          string    `gorm:"column:event_sign_hash" json:"event_sign_hash"`
	TokenID                int64     `gorm:"column:token_id" json:"token_id"`
}

// TableName PropertyUserSummary's table name
func (*PropertyUserSummary) TableName() string {
	return TableNamePropertyUserSummary
}
