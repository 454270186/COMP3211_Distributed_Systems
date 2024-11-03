package main

import "math"

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
