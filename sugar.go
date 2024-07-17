package main

import (
	"context"
	"fmt"
	"os"

	"log/slog"

	"github.com/blinkbean/dingtalk"
	"github.com/golang-module/carbon"
	"github.com/mittwald/goharbor-client/v5/apiv2"
)

func service() {
	checkSync()
}

func EnvGet(envName string, defaultValue string) (data string) {
	data = os.Getenv(envName)
	if len(data) == 0 {
		data = defaultValue
		return
	}
	return
}

func checkSync() {
	harborClient := NewHarborClient()
	ctx := context.Background()

	bot := dingtalk.InitDingTalkWithSecret(config.DingRobot.AccessToken, config.DingRobot.Secret)
	resp, err := harborClient.ListReplicationExecutions(ctx, nil, nil, nil)
	if err != nil {
		slog.Error("list replication executions failed ", "error", err)
	}
	beginDayTime := carbon.Now().SubDuration(config.BeforTime)
	for _, v := range resp {
		localtime := carbon.Parse(fmt.Sprint(v.StartTime)).ToString()

		if carbon.Parse(localtime).Gt(beginDayTime) {
			if v.Status == "Failed" {
				repoResp, _ := harborClient.GetReplicationPolicyByID(ctx, v.PolicyID)
				slog.Error("task 同步失败",
					"task_id", v.ID,
					"start_time", localtime,
					"policy_name", repoResp.Name,
				)

				text := fmt.Sprintf("[镜像同步swr失败]\n* taskID: %v \n* 时间: %v \n* 同步策略: %v ", v.ID, localtime, repoResp.Name)

				err := bot.SendTextMessage(text)
				if err != nil {
					slog.Error("Send msg to dingding failed", "error", err)
				}

			}
			slog.Info("task 同步成功",
				"task_id", v.ID,
				"start_time", localtime,
			)
		}
	}

}

func NewHarborClient() (harborClient *apiv2.RESTClient) {
	apiURL := fmt.Sprintf("%v/api", config.Harbor.URL)

	// opts := &config.Options{
	// 	PageSize: 50,
	// 	Page:     1,
	// 	Sort:     "",
	// 	Query:    "",
	// 	Timeout:  30 * time.Second,
	// }

	harborClient, err := apiv2.NewRESTClientForHost(apiURL, config.Harbor.Username, config.Harbor.Password, nil)
	if err != nil {
		slog.Error("cat not get harbor clinet", "error", err)

	}
	return
}
