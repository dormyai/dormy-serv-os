package property

import (
	"dormy/app/model"
	"dormy/database"
	ml "dormy/middleware"
	"dormy/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 根据状态和时间查询最新出售的房屋信息
func PropertyList(c *gin.Context) {
	lang := c.GetHeader("I18n-Language")

	ps := c.DefaultQuery("ps", "1")
	pn := c.DefaultQuery("pn", "10")

	page := util.NewPageWithStr(pn, ps)

	db := database.DB

	var propertyInfos []model.PropertyInfo
	db.Where(" property_status = 1 ").Offset(page.Offset()).Limit(int(page.PageSize)).Find(&propertyInfos)

	var resultList []map[string]interface{}
	for i := 0; i < len(propertyInfos); i++ {

		tmp := make(map[string]interface{})

		tmp["property_info"] = propertyInfos[i]

		//查询出合约信息
		var propertyChainInfo model.PropertyChainInfo
		db.Where(" property_info_id = ? ", propertyInfos[i].ID).First(&propertyChainInfo)

		tmp["property_chain_info"] = propertyChainInfo

		var propertyInfoMedia []model.PropertyInfoMedia
		db.Where(" property_info_id = ? ", propertyInfos[i].ID).Find(&propertyInfoMedia)

		tmp["property_info_medium"] = propertyInfoMedia
		//查询出媒体信息
		resultList = append(resultList, tmp) // 将 tmp 添加到 resultList
	}

	var totalRecords int64
	db.Model(&model.PropertyInfo{}).Where(" property_status = 1 ").Count(&totalRecords)

	page.Result = resultList
	page = page.SetTotal(totalRecords)

	c.JSON(http.StatusOK, ml.Succ(lang, page))
}

func PropertyDetail(c *gin.Context) {
	lang := c.GetHeader("I18n-Language")
	id := c.Query("id")

	if len(id) == 0 {
		c.JSON(http.StatusOK, ml.Fail(lang, "100009"))
		return
	}

	db := database.DB

	var propertyInfo model.PropertyInfo
	db.First(&propertyInfo, id)

	if propertyInfo == (model.PropertyInfo{}) {
		c.JSON(http.StatusOK, ml.Fail(lang, "100009"))
		return
	}

	//查询出合约信息
	var propertyChainInfo model.PropertyChainInfo
	db.Where(" property_info_id = ? ", propertyInfo.ID).First(&propertyChainInfo)

	var propertyInfoMedia []model.PropertyInfoMedia
	db.Where(" property_info_id = ? ", propertyInfo.ID).Find(&propertyInfoMedia)

	c.JSON(http.StatusOK, ml.Succ(lang, map[string]interface{}{"property_info": propertyInfo,
		"property_chain_info": propertyChainInfo, "property_info_medium": propertyInfoMedia}))
}

func BuyProperty(c *gin.Context) {

}

func BuyOrderList(c *gin.Context) {

}

func UserPropertList(c *gin.Context) {
	lang := c.GetHeader("I18n-Language")
	token := c.GetHeader("token")
	session, _ := util.ParseToken(token)

	ps := c.DefaultQuery("ps", "1")
	pn := c.DefaultQuery("pn", "10")
	page := util.NewPageWithStr(pn, ps)

	db := database.DB

	var propertyUserSummarys []model.PropertyUserSummary

	db.Where(" (user_id = ? or address = ?) and end_time = '0000-00-00 00:00:00' ", session.Id, session.Addr).Offset(page.Offset()).Limit(int(page.PageSize)).Find(&propertyUserSummarys)

	var resultList []map[string]interface{}

	for i := 0; i < len(propertyUserSummarys); i++ {

		tmp := make(map[string]interface{})

		var propertyInfo model.PropertyInfo
		db.First(&propertyInfo, propertyUserSummarys[i].PropertyInfoID)

		var propertyChainInfo model.PropertyChainInfo
		db.Where(" property_info_id = ? ", propertyUserSummarys[i].PropertyInfoID).First(&propertyChainInfo)

		tmp["property_info"] = propertyInfo
		tmp["property_user_summary"] = propertyUserSummarys[i]
		tmp["property_chain_info"] = propertyChainInfo

		var propertyInfoMedia []model.PropertyInfoMedia
		db.Where(" property_info_id = ? ", propertyInfo.ID).Find(&propertyInfoMedia)

		tmp["property_info_medium"] = propertyInfoMedia
		//查询出媒体信息
		resultList = append(resultList, tmp) // 将 tmp 添加到 resultList
	}

	var totalRecords int64
	db.Model(&model.PropertyUserSummary{}).Where(" (user_id = ? or address = ?) and end_time = '0000-00-00 00:00:00' ", session.Id, session.Addr).Count(&totalRecords)

	page.Result = resultList
	page = page.SetTotal(totalRecords)

	c.JSON(http.StatusOK, ml.Succ(lang, page))

}

func CalculateUserProperty(c *gin.Context) {
	lang := c.GetHeader("I18n-Language")
	token := c.GetHeader("token")
	session, _ := util.ParseToken(token)

	db := database.DB
	//计算出这个用户的持有的不同的资产及数量
	var propertyUserSummary []model.PropertyUserSummary
	db.Where(" (user_id = ? or address = ?) and end_time = '0000-00-00 00:00:00' ", session.Id, session.Addr).Find(&propertyUserSummary)

	totalInvestment := int64(0)
	totalAssetValue := int64(0)
	totalRentalYield := int64(0)
	actualRoi := int64(0)

	for i := 0; i < len(propertyUserSummary); i++ {
		pus := propertyUserSummary[i]

		var pci model.PropertyChainInfo
		db.Where(" property_info_id = ? ", pus.PropertyInfoID).First(&pci)

		if pci == (model.PropertyChainInfo{}) {
			continue
		}

		totalInvestment += (pus.Amount * int64(pci.TokenPrice))
		totalAssetValue += (pus.Amount * int64(pci.TokenPrice))

	}

	if totalInvestment != 0 {
		totalRentalYield = 10
		actualRoi = 8
	}

	c.JSON(http.StatusOK, ml.Succ(lang, map[string]interface{}{"total_investment": totalInvestment,
		"total_asset_value": totalAssetValue, "total_rental_yield": totalRentalYield, "actual_roi": actualRoi}))
}

func PropertyRentalList(c *gin.Context) {
	lang := c.GetHeader("I18n-Language")
	token := c.GetHeader("token")
	session, _ := util.ParseToken(token)

	propertInfoId := c.DefaultQuery("propert_info_id", "0")
	ps := c.DefaultQuery("ps", "1")
	pn := c.DefaultQuery("pn", "10")
	page := util.NewPageWithStr(pn, ps)

	db := database.DB

	var propertyRentals []model.PropertyRental

	db.Where(" propert_info_id = ? ", propertInfoId).Offset(page.Offset()).Limit(int(page.PageSize)).Order("id desc ").Find(&propertyRentals)

	var totalRecords int64
	db.Model(&model.PropertyUserSummary{}).Where(" (user_id = ? or address = ?) and end_time = '0000-00-00 00:00:00' ", session.Id, session.Addr).Count(&totalRecords)

	page.Result = propertyRentals
	page = page.SetTotal(totalRecords)

	c.JSON(http.StatusOK, ml.Succ(lang, page))

}
