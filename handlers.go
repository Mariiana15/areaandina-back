package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func HandleRoot(write_ http.ResponseWriter, req *http.Request) {
	write_.WriteHeader(http.StatusOK)
}

func GetZonesSummary(write_ http.ResponseWriter, req *http.Request) {

	var gen Generic
	gen.Id = "La sabana"
	gen.Temperature = 38
	var zoneArray []ZoneGeneric

	var zones []DataForestZone
	getZones(&zones)
	var t float64
	for i := len(zones) - 1; i >= 0; i-- {
		var zone ZoneGeneric
		zone.Id = zones[i].Zone
		zone.Temperature = zones[i].Temp
		zoneArray = append(zoneArray, zone)
		t = t + float64(zone.Temperature)

	}
	gen.Temperature = 0
	if len(zones) > 0 {
		z := float64(len(zones))
		gen.Temperature = float64(t) / float64(z)
	}
	gen.Zones = zoneArray
	byteData, _ := json.Marshal(gen)
	write_.Write(byteData)
	write_.WriteHeader(http.StatusOK)
}

func getZoneDataHumedad(zone *Zone, summary []DataSummary) []DataSummary {

	df, _ := getDataZone(zone.Id, "humedad")
	var data DataSummary
	data.D0 = 100
	data.D1 = df[0].Summary
	data.D2 = 0
	data.D3 = 0
	data.D4 = 0
	data.D5 = 0
	if len(df) > 1 {
		data.D2 = df[1].Summary
	}
	if len(df) > 2 {
		data.D3 = df[2].Summary
	}
	if len(df) > 3 {
		data.D4 = df[3].Summary
	}
	if len(df) > 4 {
		data.D5 = df[4].Summary
	}
	data.Desc = "Los datos de la cantidad de H2O por cada 100 m es."
	data.Id = "1"
	data.Image = "https://images.unsplash.com/photo-1619529079006-873f44d61e1e?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=500&q=80"
	data.Title = "Humedad ambiente"
	data.Und = "% H2O"
	data.Alert = "_"
	var m map[int32][]int32
	err := json.Unmarshal([]byte(fmt.Sprintln(df[0].Data)), &m)
	if err != nil {
		fmt.Println("{\"error\": \"%v\"}", err)
		return summary
	}
	zone.Data_H = m
	summary = append(summary, data)
	return summary

}

func getZoneDataHumedadT(zone *Zone, summary []DataSummary) []DataSummary {

	df, _ := getDataZone(zone.Id, "humedad suelo")
	var data DataSummary
	data.D0 = 100
	data.D1 = df[0].Summary
	data.D2 = 0
	data.D3 = 0
	data.D4 = 0
	data.D5 = 0
	if len(df) > 1 {
		data.D2 = df[1].Summary
	}
	if len(df) > 2 {
		data.D3 = df[2].Summary
	}
	if len(df) > 3 {
		data.D4 = df[3].Summary
	}
	if len(df) > 4 {
		data.D5 = df[4].Summary
	}
	data.Desc = "La humedad media en el suelo en un radio de 50m es."
	data.Id = "2"
	data.Image = "https://images.unsplash.com/photo-1589742377572-425d09d89b7a?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=5000&q=80"
	data.Title = "Humedad en el suelo"
	data.Und = "% H2O"
	data.Alert = "_"
	var m map[int32][]int32
	err := json.Unmarshal([]byte(fmt.Sprintln(df[0].Data)), &m)
	if err != nil {
		fmt.Println("{\"error\": \"%v\"}", err)
		return summary
	}

	zone.Data_HT = m
	summary = append(summary, data)
	return summary

}

func getZoneDataTemperature(zone *Zone, summary []DataSummary) []DataSummary {

	df, _ := getDataZone(zone.Id, "temperatura")
	var data DataSummary
	data.D0 = 75
	data.D1 = df[0].Summary
	data.D2 = 0
	data.D3 = 0
	data.D4 = 0
	data.D5 = 0
	if len(df) > 1 {
		data.D2 = df[1].Summary
	}
	if len(df) > 2 {
		data.D3 = df[2].Summary
	}
	if len(df) > 3 {
		data.D4 = df[3].Summary
	}
	if len(df) > 4 {
		data.D5 = df[4].Summary
	}
	data.Desc = "La humedad media en el suelo en un radio de 50m es."
	data.Id = "3"
	data.Image = "https://images.unsplash.com/photo-1572339152651-c3f1995ab1f2?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=500&q=80"
	data.Title = "Temperatura Ambiente"
	data.Und = "(°C)"
	data.Alert = "_"
	var m map[int32][]int32
	err := json.Unmarshal([]byte(fmt.Sprintln(df[0].Data)), &m)
	if err != nil {
		fmt.Println("{\"error\": \"%v\"}", err)
		return summary
	}

	zone.Data_T = m

	summary = append(summary, data)
	return summary
}

func getZoneDataCarbono(zone *Zone, summary []DataSummary) []DataSummary {

	df, _ := getDataZone(zone.Id, "carbono")
	var data DataSummary
	data.D0 = 2000
	data.D1 = df[0].Summary
	data.D2 = 0
	data.D3 = 0
	data.D4 = 0
	data.D5 = 0
	if len(df) > 1 {
		data.D2 = df[1].Summary
	}
	if len(df) > 2 {
		data.D3 = df[2].Summary
	}
	if len(df) > 3 {
		data.D4 = df[3].Summary
	}
	if len(df) > 4 {
		data.D5 = df[4].Summary
	}
	data.Desc = "La cantidad de particulas de carbono en el aire son."
	data.Id = "4"
	data.Image = "https://images.unsplash.com/photo-1531418580067-04bc72baa73c?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=500&q=80"
	data.Title = "Particulas CO2"
	data.Und = "Ud/M"
	data.Alert = "_"
	var m map[int32][]int32
	err := json.Unmarshal([]byte(fmt.Sprintln(df[0].Data)), &m)
	if err != nil {
		fmt.Println("{\"error\": \"%v\"}", err)
		return summary
	}

	zone.Data_C = m
	summary = append(summary, data)
	return summary

}

func HandleRootZone(write_ http.ResponseWriter, req *http.Request) {

	var zone Zone
	var data []DataSummary
	zone.Id = req.Header.Get("id")
	zone.Date = time.Now().Unix()
	data = getZoneDataHumedad(&zone, data)
	data = getZoneDataHumedadT(&zone, data)
	data = getZoneDataTemperature(&zone, data)
	data = getZoneDataCarbono(&zone, data)
	zone.Summary = data
	byteData, _ := json.Marshal(zone)
	write_.Write(byteData)
	write_.WriteHeader(http.StatusOK)
}

func ZoneDataRequest(write_ http.ResponseWriter, req *http.Request) {
	var zd ZoneInData
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&zd)
	write_.Header().Set("Content-Type", "application/json")
	if err != nil {
		write_.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(write_, "{\"error\": \"%v\"}", err)
		return
	}
	err = insertZone(zd)
	if err != nil {
		write_.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(write_, "{\"error\": \"%v\"}", msgDatabase)
		return
	}

	var m ResponseZone
	m.Message = "Operation make successfully"
	write_.WriteHeader(http.StatusOK)
	byteData, _ := json.Marshal(m)
	write_.Write(byteData)
}
