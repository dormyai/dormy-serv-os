package model

const TableNamePropertyInfoMedia = "property_info_media"

// PropertyInfoMedium mapped from table <property_info_media>
type PropertyInfoMedia struct {
	ID             int64  `gorm:"column:id;primaryKey" json:"id"`
	Title          string `gorm:"column:title" json:"title"`
	URL            string `gorm:"column:url" json:"url"`
	Type           int32  `gorm:"column:type;comment:1.图片2.视频" json:"type"` // 1.图片2.视频
	PropertyInfoID int64  `gorm:"column:property_info_id" json:"property_info_id"`
	Status         int32  `gorm:"column:status" json:"status"`
	Del            int32  `gorm:"column:del" json:"del"`
	CreateAt       string `gorm:"column:create_at" json:"create_at"`
}

// TableName PropertyInfoMedium's table name
func (*PropertyInfoMedia) TableName() string {
	return TableNamePropertyInfoMedia
}
