package model

const TableNamePropertyChainInfo = "property_chain_info"

// PropertyChainInfo mapped from table <property_chain_info>
type PropertyChainInfo struct {
	ID                   int64   `gorm:"column:id;primaryKey" json:"id"`
	PropertyInfoID       int64   `gorm:"column:property_info_id" json:"property_info_id"`
	ChainID              int64   `gorm:"column:chain_id" json:"chain_id"`
	SoltID               int64   `gorm:"column:solt_id" json:"solt_id"`
	TotalQuantity        int64   `gorm:"column:total_quantity;comment:总份数" json:"total_quantity"` // 总份数
	TokenAmount          int64   `gorm:"column:token_amount;comment:总份数" json:"token_amount"`     // 总份数
	TokenPrice           float64 `gorm:"column:token_price" json:"token_price"`
	MinPurchase          int64   `gorm:"column:min_purchase" json:"min_purchase"`
	MaxPurchase          int64   `gorm:"column:max_purchase" json:"max_purchase"`
	MinIncrement         int64   `gorm:"column:min_increment" json:"min_increment"`
	PropertyContractAddr string  `gorm:"column:property_contract_addr" json:"property_contract_addr"`
}

// TableName PropertyChainInfo's table name
func (*PropertyChainInfo) TableName() string {
	return TableNamePropertyChainInfo
}
