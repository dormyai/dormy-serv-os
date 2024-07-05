package router

import (
	"dormy/app/frontend/controller/common"
	"dormy/app/frontend/controller/property"
	"dormy/app/frontend/controller/user"
	"dormy/interceptor"

	"github.com/gin-gonic/gin"
)

func Load(r *gin.RouterGroup) {
	userGroup := r.Group("/user", interceptor.LoggerMiddleware(), interceptor.FrontendSignMiddleware(), interceptor.FrontendAuthMiddleware())

	userGroup.POST("/sign_msg", user.GetSignMsg)
	userGroup.POST("/login", user.Login)
	userGroup.POST("/login_by_addr", user.LoginByAddr)
	userGroup.GET("/login_user_info", user.LoginUserInfo)

	propertyGroup := r.Group("/property", interceptor.LoggerMiddleware(), interceptor.FrontendSignMiddleware(), interceptor.FrontendAuthMiddleware())

	propertyGroup.GET("/list", property.PropertyList)
	propertyGroup.GET("/detail", property.PropertyDetail)
	propertyGroup.GET("/userPropertList", property.UserPropertList)
	propertyGroup.GET("/calculateUserProperty", property.CalculateUserProperty)
	propertyGroup.GET("/rentals", property.PropertyRentalList)

	commonGroup := r.Group("/common", interceptor.LoggerMiddleware(), interceptor.FrontendSignMiddleware(), interceptor.FrontendAuthMiddleware())
	commonGroup.POST("/mail_subscribe", common.Subscribe)
	commonGroup.POST("/get_rate", common.GetRate)
}
