package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

type MetaData interface{}

type Auth struct {
	User string `json:"user"`
	Pass string `json:"user_pass"`
}

type Car struct {
	Id          string `json:"id"`
	Brand       string `json:"brand" validate:"required,alphanum"`
	Model       string `json:"model" validate:"required,alphanum"`
	Horse_power uint32 `json:"horse_power" validate:"required,gte=0,lte=10000"`
}

type Zone struct {
	Id      string            `json:"id"`
	Data_H  map[int32][]int32 `json:"data_h"`
	Data_HT map[int32][]int32 `json:"data_ht"`
	Data_T  map[int32][]int32 `json:"data_t"`
	Data_C  map[int32][]int32 `json:"data_c"`
	Summary []DataSummary     `json:"summary"`
	Date    int64             `json:"date"`
}

type ResponseZone struct {
	Message string `json:"message"`
}

type ZoneInData struct {
	Data       map[int32][]int32 `json:"data"`
	Summary    string            `json:"summary"`
	TypeSensor string            `json:"typeSensor"`
	Zone       string            `json:"zone"`
}

type DataZoneFormat struct {
	Data map[int32][]int32 `json:"data"`
}

type Generic struct {
	Id          string        `json:"id"`
	Temperature float64       `json:"temperature"`
	Zones       []ZoneGeneric `json:"zones"`
}

type ZoneGeneric struct {
	Id          string  `json:"id"`
	Temperature float64 `json:"temperature"`
}

type DataSummary struct {
	Id    string `json:"id"`
	Desc  string `json:"desc"`
	Image string `json:"image"`
	Title string `json:"title"`
	Alert string `json:"alert"`
	Und   string `json:"und"`
	D0    int32  `json:"d0"`
	D1    int32  `json:"d1"`
	D2    int32  `json:"d2"`
	D3    int32  `json:"d3"`
	D4    int32  `json:"d4"`
	D5    int32  `json:"d5"`
}

type DataForest struct {
	Date    int64  `json:"Date"`
	Data    string `json:"Data"`
	Type    string `json:"Type"`
	Summary int32  `json:"Summary"`
}

type DataForestZone struct {
	Zone string  `json:"zone"`
	Temp float64 `json:"temp"`
}

var validate *validator.Validate

func (car *Car) ToJson() ([]byte, error) {
	return json.Marshal(car)
}

func (car *Car) ValidateStructure() (string, error) {
	validate = validator.New()
	var msg string
	err := validate.Struct(car)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err.Error(), err
		}
		return fmt.Sprintf("The field %s is mal format", err.(validator.ValidationErrors)[0].StructField()), err
	}
	return msg, err
}
