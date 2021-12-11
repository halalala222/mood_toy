package apis

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/gin-gonic/gin"
	"log"
	"lwh.com/models"
	"lwh.com/setting"
	"net/http"
)

type XXG struct {
	Messages []struct {
		Type       string `json:"Type"`
		AnswerType string `json:"AnswerType"`
		Text       struct {
			Ext struct {
			} `json:"Ext"`
			ContentType          string `json:"ContentType"`
			UserDefinedChatTitle string `json:"UserDefinedChatTitle"`
			AnswerSource         string `json:"AnswerSource"`
			Content              string `json:"Content"`
			HitStatement         string `json:"HitStatement"`
		} `json:"Text"`
		Knowledge struct {
		} `json:"Knowledge"`
	} `json:"Messages"`
	RequestID string `json:"RequestId"`
	SessionID string `json:"SessionId"`
	MessageID string `json:"MessageId"`
}

func GetXXG(c *gin.Context) {
	var txt = models.Txt{}
	var xxg = XXG{}
	client, err := sdk.NewClientWithAccessKey("cn-shanghai", setting.Configone.XXG.AccessKeyID, setting.Configone.XXG.AccessKeySecret)
	/* use STS Token
	client, err := sdk.NewClientWithStsToken("cn-shanghai", "<your-access-key-id>", "<your-access-key-secret>", "<your-sts-token>")
	*/
	if err != nil {
		c.JSON(http.StatusBadGateway,gin.H{
			"message" : err.Error(),
		})
	}else {
		request := requests.NewCommonRequest()
		request.Method = "POST"
		request.Scheme = "https" // https | http
		request.Domain = "chatbot.cn-shanghai.aliyuncs.com"
		request.Version = "2017-10-11"
		request.ApiName = "Chat"
		request.QueryParams["Action"] = "Chat"
		request.QueryParams["InstanceId"] = "chatbot-cn-62Vf6QIZRe"
		request.QueryParams["Utterance"] = txt.Text
		request.QueryParams["Format"] = "JSON"
		response, err1 := client.ProcessCommonRequest(request)
		if err1 != nil {
			panic(err1)
		}
		resp := response.GetHttpContentString()
		fmt.Println(resp)
		if err2 := json.Unmarshal([]byte(resp),&xxg);err2 != nil{
			log.Panic(err2.Error())
		}else {
			c.JSON(http.StatusOK,gin.H{
				"content" : xxg.Messages[0].Text.Content,
			})
		}
	}

}

