package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		fmt.Println("health")
	})
	r.POST("/SimulatedData", func(c *gin.Context) {
		err := InsertRandomRecord(context.Background())
		if err != nil {
			fmt.Printf("err while insert simulated data: %v\n", err)
		}
	})
	r.POST("/Statistics", func(c *gin.Context) {
		logs := make([]string, 0)
		datas := GetRows(context.Background())

		output := AnalysisData(datas)
		for sensorId, sensorDataMap := range output {
			log := fmt.Sprintf("[Sensor %v]\n", sensorId)
			for _, dataType := range DataTypes {
				log += fmt.Sprintf("%v: Min(%v) Max(%v) Avg(%v)\n", dataType, sensorDataMap[dataType].Min, sensorDataMap[dataType].Max, sensorDataMap[dataType].Avg)
			}
			log += "\n"

			logs = append(logs, log)
		}

		c.JSON(200, InvokeResponse{
			Logs: logs,
		})
	})

	return r
}
