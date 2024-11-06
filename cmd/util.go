package main

import (
	"fmt"
	"log"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

var (
	DataTypes = []string{"Temperature", "WindSpeed", "RelativeHumidity", "CO2Level"}
)

type AnalysisOutput struct {
	Min int
	Max int
	Avg float64
	Sum int
}

func AnalysisData(datas []*EnvironmentalData) map[int]map[string]*AnalysisOutput {
	output := make(map[int]map[string]*AnalysisOutput)

	sensorIdCntMap := make(map[int]int)

	for _, data := range datas {
		if output[data.SensorID] == nil {
			output[data.SensorID] = make(map[string]*AnalysisOutput)
		}

		sensorDataMap := output[data.SensorID]
		for _, dataType := range DataTypes {
			if sensorDataMap[dataType] == nil {
				sensorDataMap[dataType] = &AnalysisOutput{
					Min: math.MaxInt,
					Max: math.MinInt,
				}
			}
		}

		tempOutput := sensorDataMap["Temperature"]
		windOutput := sensorDataMap["WindSpeed"]
		humidityOutput := sensorDataMap["RelativeHumidity"]
		co2Output := sensorDataMap["CO2Level"]

		tempOutput.Min = min(tempOutput.Min, data.Temperature)
		tempOutput.Max = max(tempOutput.Max, data.Temperature)
		tempOutput.Sum += data.Temperature

		windOutput.Min = min(windOutput.Min, data.WindSpeed)
		windOutput.Max = max(windOutput.Max, data.WindSpeed)
		windOutput.Sum += data.WindSpeed

		humidityOutput.Min = min(humidityOutput.Min, data.RelativeHumidity)
		humidityOutput.Max = max(humidityOutput.Max, data.RelativeHumidity)
		humidityOutput.Sum += data.RelativeHumidity

		co2Output.Min = min(co2Output.Min, data.CO2Level)
		co2Output.Max = max(co2Output.Max, data.CO2Level)
		co2Output.Sum += data.CO2Level

		sensorIdCntMap[data.SensorID]++
	}

	// avg
	for sensorId, sensorDataMap := range output {
		totalCnt := sensorIdCntMap[sensorId]

		for _, dataType := range DataTypes {
			sensorDataMap[dataType].Avg = float64(sensorDataMap[dataType].Sum) / float64(totalCnt)
		}
	}

	return output
}

func PlotPerformanceChart(sensorCounts []int, responseTimes []int64) error {
    // 检查传感器数量和响应时间的长度是否一致
    if len(sensorCounts) != len(responseTimes) {
        return fmt.Errorf("sensorCounts and responseTimes slices must have the same length")
    }

    // 创建一个绘图对象
    p := plot.New()
    p.Title.Text = "Insert Performance for Different Sensor Counts"
    p.X.Label.Text = "Number of Sensors"
    p.Y.Label.Text = "Response Time (ms)"

    // 将数据转换为 plotter.XYs 结构
    points := make(plotter.XYs, len(sensorCounts))
    for i := range sensorCounts {
        points[i].X = float64(sensorCounts[i])     // 将传感器数量转换为 float64
        points[i].Y = float64(responseTimes[i])    // 将响应时间从 int64 转换为 float64
    }

    // 创建折线图并添加到绘图对象中
    line, err := plotter.NewLine(points)
    if err != nil {
        return fmt.Errorf("failed to create line plot: %v", err)
    }
    line.LineStyle.Width = vg.Points(2)
    line.LineStyle.Dashes = []vg.Length{vg.Points(5), vg.Points(5)} // 设置虚线样式
    p.Add(line)

    // 添加数据点
    scatter, err := plotter.NewScatter(points)
    if err != nil {
        return fmt.Errorf("failed to create scatter plot: %v", err)
    }
    scatter.GlyphStyle.Radius = vg.Points(3)
    p.Add(scatter)

    // 保存图表为 PNG 文件
    if err := p.Save(8*vg.Inch, 4*vg.Inch, "performance_chart.png"); err != nil {
        return fmt.Errorf("failed to save plot: %v", err)
    }

    log.Println("Chart saved as performance_chart.png")
    return nil
}