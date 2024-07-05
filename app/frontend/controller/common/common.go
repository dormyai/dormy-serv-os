package common

import (
	"dormy/app/model"
	"dormy/config"
	"dormy/database"
	ml "dormy/middleware"
	"dormy/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
)

func Subscribe(c *gin.Context) {

	lang := c.GetHeader("I18n-Language")
	email := c.PostForm("email")

	if len(email) == 0 {
		c.JSON(http.StatusOK, ml.Fail(lang, "100006"))
		return
	}

	if !isEmailValid(email) {
		c.JSON(http.StatusOK, ml.Fail(lang, "100007"))
		return
	}

	id, _ := util.Sf.GenerateID()

	var mailSubscribe model.MailSubscribe

	mailSubscribe.ID = id
	mailSubscribe.Email = email
	mailSubscribe.CreateAt = time.Now()

	database.DB.Save(&mailSubscribe)

	c.JSON(http.StatusOK, ml.Succ(lang, nil))
}

// isEmailValid 检查邮箱格式是否正确
func isEmailValid(email string) bool {
	// 这是一个简单的正则表达式，用于验证大多数电子邮件地址的格式
	// 对于非常严格或特定格式的电子邮件验证，您可能需要一个更复杂的正则表达式
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}

type ExchangeRatesResponse struct {
	Rates map[string]float64 `json:"rates"`
}

type Rate struct {
	Type       string
	Value      float64
	UpdateTime time.Time
}

func GetRate(c *gin.Context) {

	lang := c.GetHeader("I18n-Language")

	rate, bo := util.CacheGet("COMMON_RATE_GBP")

	ml.Log.Info("rate:", rate, bo)
	if bo {
		c.JSON(http.StatusOK, ml.Succ(lang, map[string]interface{}{"rate": rate}))
	} else {
		// 构建请求 URL
		url := fmt.Sprintf("https://openexchangerates.org/api/latest.json?app_id=%s", config.Get().Common.OpenExchangeRateKey)

		// 发起 HTTP GET 请求
		response, err := http.Get(url)
		if err != nil {
			ml.Log.Info(err)
			c.JSON(http.StatusOK, ml.Fail(lang, "100010"))
			return
		}
		defer response.Body.Close()

		// 读取响应体
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			ml.Log.Info(err)
			c.JSON(http.StatusOK, ml.Fail(lang, "100010"))
			return
		}

		// 解析 JSON 响应
		var exchangeRates ExchangeRatesResponse
		if err := json.Unmarshal(body, &exchangeRates); err != nil {
			ml.Log.Info(err)
			c.JSON(http.StatusOK, ml.Fail(lang, "100010"))
			return
		}

		// 获取美元兑英镑的汇率并打印
		gbpRate, exists := exchangeRates.Rates["GBP"]
		if !exists {
			ml.Log.Info("GBP rate not found")
			c.JSON(http.StatusOK, ml.Fail(lang, "100010"))
			return
		}

		ml.Log.Infof("USD to GBP rate: %f\n", gbpRate)

		rate := Rate{}
		rate.Type = "GBP"
		rate.Value = gbpRate
		rate.UpdateTime = time.Now()
		util.CachePut("COMMON_RATE_GBP", rate, 1*time.Hour)
		c.JSON(http.StatusOK, ml.Succ(lang, map[string]interface{}{"rate": rate}))
	}
}
