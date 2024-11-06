package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		fmt.Println("health")
	})
	r.GET("/api/SimulatedDataTask1", func(c *gin.Context) {
		sensorCnts := []int{20, 50, 70, 100, 150, 300, 500}
		respTimes := make([]int64, len(sensorCnts))
		var m sync.Mutex
		var wg sync.WaitGroup

		for i, sensorCnt := range sensorCnts {
			wg.Add(1)
			go func(j int) {
				var sum int64
				for i := 0; i < 10; i++ {
					start := time.Now()
					if err := InsertRandomRecord(context.Background(), sensorCnt); err != nil {
						fmt.Printf("err while insert simulated data: %v\n", err)
						c.JSON(500, InvokeResponse{})
						return
					}
					sum += time.Since(start).Milliseconds()
				}

				m.Lock()
				respTimes[j] = sum / 10
				m.Unlock()

				wg.Done()
			}(i)
		}

		wg.Wait()

		if err := PlotPerformanceChart(sensorCnts, respTimes); err != nil {
			fmt.Printf("err while draw chart: %v\n", err)
			c.JSON(500, InvokeResponse{})
			return
		}

		c.JSON(200, InvokeResponse{})
	})
	r.POST("/SimulatedData", func(c *gin.Context) {
		err := InsertRandomRecord(context.Background(), 20)
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
