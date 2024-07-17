package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"net/http"
	"one-api/common"
	"strings"
)

var StopFinishReason = "stop"

// 定义一个新的错误类型
type CustomError struct {
	StatusCode int
	Message    string
	Type       string
	Code       string
}

func RelayErrorHandler(resp *http.Response) *CustomError {
	if resp == nil {
		return &CustomError{
			StatusCode: 500,
			Message:    "请稍后再试试，如长时间不可用，请通知管理员q2411728460 进行处理",
			Type:       "model_unavailable",
			Code:       "model_unavailable",
		}
	}

	customError := &CustomError{
		StatusCode: resp.StatusCode,
		Message:    "请稍后再试试，如长时间不可用，请通知管理员q2411728460 进行处理",
		Type:       "model_unavailable",
		Code:       "model_unavailable",
	}

	// 读取上游响应体（可选，如果需要记录上游错误信息）
	_, err := io.ReadAll(resp.Body)
	if err != nil {
		// 不记录错误
	}
	err = resp.Body.Close()
	if err != nil {
		// 不记录错误
	}

	return customError
}

func GetFullRequestURL(baseURL string, requestURL string, channelType int) string {
	fullRequestURL := fmt.Sprintf("%s%s", baseURL, requestURL)

	if strings.HasPrefix(baseURL, "https://gateway.ai.cloudflare.com") {
		switch channelType {
		case common.ChannelTypeOpenAI:
			fullRequestURL = fmt.Sprintf("%s%s", baseURL, strings.TrimPrefix(requestURL, "/v1"))
		case common.ChannelTypeAzure:
			fullRequestURL = fmt.Sprintf("%s%s", baseURL, strings.TrimPrefix(requestURL, "/openai/deployments"))
		}
	}
	return fullRequestURL
}

func GetAPIVersion(c *gin.Context) string {
	query := c.Request.URL.Query()
	apiVersion := query.Get("api-version")
	if apiVersion == "" {
		apiVersion = c.GetString("api_version")
	}
	return apiVersion
}
