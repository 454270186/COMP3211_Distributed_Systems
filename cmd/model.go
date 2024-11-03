package main

import "time"

type InvokeResponse struct {
	Outputs     map[string]interface{}
	Logs        []string
	ReturnValue interface{}
}

type EnvironmentalData struct {
	ID               int
	SensorID         int       // 传感器 ID
	Temperature      int       // 温度（摄氏度）
	WindSpeed        int       // 风速（英里/小时）
	RelativeHumidity int       // 相对湿度（百分比）
	CO2Level         int       // 二氧化碳浓度（ppm）
	Timestamp        time.Time // 时间戳
}
