package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	uuid "github.com/satori/go.uuid"
)

type dbManager interface {
	insert() error
	get(id string) error
	delete() error
}

func newConnect() *sql.DB {
	db, err := sql.Open("mysql", "dbforest1:dbforest1@tcp(35.202.243.156:3306)/dbforest1")
	if err != nil {
		fmt.Println(err.Error())
	}
	return db
}

func insertZone(z ZoneInData) error {
	var db = newConnect()

	atCreate := time.Now().Unix()
	id := uuid.NewV4().String()
	jsonStr, err := json.Marshal(z.Data)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
	_, err = db.Query(fmt.Sprintf("INSERT INTO Data_Forest VALUES ( '%s' ,'%s','%s','%s','%s', '%s' , '%s' );", id, string(jsonStr), fmt.Sprint(atCreate), z.TypeSensor, z.Zone, z.Summary, "active"))
	if err != nil {
		return err
	}

	defer db.Close()
	return err
}

func getZones(zone *[]DataForestZone) error {

	var db = newConnect()
	var z DataForestZone
	response, err := db.Query("Select Zone, AVG(Summary)  from Data_Forest where `Type`in ('humedad', 'temperatura', 'humedad suelo') group BY Zone  ORDER BY Date DESC  limit 6;")
	if err != nil {
		fmt.Println(err)
		return err
	}
	for response.Next() {
		response.Scan(&z.Zone, &z.Temp)
		*zone = append(*zone, z)
	}
	defer db.Close()
	return err
}

func getDataZone(zone string, types string) ([]DataForest, error) {

	var db = newConnect()
	var data []DataForest
	var d DataForest
	response, err := db.Query(fmt.Sprintf("Select  Data, Date, Type, Summary from Data_Forest where Zone = '%s' and Type = '%s' order by Date desc limit 5;", zone, types))
	if err != nil {
		return data, err
	}
	for response.Next() {

		response.Scan(&d.Data, &d.Date, &d.Type, &d.Summary)
		data = append(data, d)
	}
	defer db.Close()
	return data, err
}

func (auth *Auth) get(id string) error {
	var db = newConnect()
	response, err := db.Query(fmt.Sprintf("SELECT user,user_pass FROM users WHERE user = '%s'", id))
	if err != nil {
		fmt.Println(err)
		return err
	}
	for response.Next() {
		response.Scan(&auth.User, &auth.Pass)
	}
	defer db.Close()
	return err
}
