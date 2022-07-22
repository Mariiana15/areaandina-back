package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
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
	db, err := sql.Open("mysql", "dbForest:dbForest@tcp(localhost:3306)/dbForest")
	if err != nil {
		fmt.Println(err.Error())
	}
	return db
}
func insertDB(db dbManager) error {
	err := db.insert()
	return err
}

func getDB(db dbManager, id string) error {
	err := db.get(id)
	return err
}

func deleteDB(db dbManager) error {
	err := db.delete()
	return err
}

func (car *Car) insert() error {
	var db = newConnect()
	car.Id = car.Brand[1:3] + strconv.Itoa(rand.Intn(1000)) + car.Model[1:3]
	_, err := db.Query(fmt.Sprintf("INSERT INTO cars VALUES ( '%s' ,'%s','%s',%d );", car.Id, car.Brand, car.Model, car.Horse_power))
	if err != nil {
		_, err = db.Query(fmt.Sprintf("UPDATE cars SET brand = '%s' , model = '%s', horse_power = %d  WHERE  id = '%s';", car.Brand, car.Model, car.Horse_power, car.Id))
		if err != nil {
			return err
		}
	}
	defer db.Close()
	return err
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

func (car *Car) get(id string) error {
	var db = newConnect()
	err := db.QueryRow("SELECT id,brand,model,horse_power FROM cars WHERE id = ?", id).Scan(&car.Id, &car.Brand, &car.Model, &car.Horse_power)
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

func (car *Car) delete() error {
	var db = newConnect()
	_, err := db.Query(fmt.Sprintf("DELETE FROM cars WHERE id = '%s'", car.Id))
	if err != nil {
		return err
	}
	defer db.Close()
	return err
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