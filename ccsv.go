package main

import (
	"fmt"
	"encoding/csv"
	"os"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

type Usermove struct {
	utctime int `json:"utctime"`
	idfa string `json:"idfa"`
	geohash string `json:"geohash"`
	latitude float64 `json:"latitude"`
	longitude float64 `json:"longitude"`
	horizontal float64 `json:"horizontal"`
	userid int `json:"userid"`
	hourofday int `json:"hourofday"`
	lcltime string `json:"lcltime"`

}

func main() {
	fmt.Printf("hello world")

	f, err := os.Open("../../../../../reverse_geocoder/output.csv")
	if err != nil {
		fmt.Printf("error")
		panic(err.Error())
	}

	defer f.Close()

	csvr := csv.NewReader(f)

	for i :=0; i <= 10; i++{
		row, err := csvr.Read()
		if err != nil {
			fmt.Printf("error 1")
			panic(err.Error())
		}

		fmt.Printf(row[0]+" "+row[1]+" "+row[3]+"\n")
	}

	db, err := sql.Open("mysql", "root:welcome123@tcp(127.0.0.1:3306)/test")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	results, err := db.Query("SELECT utctime, idfa, geohash, latitude, longitude, horizontal, userid, hourofday, lcltime from outputv3")

	if err != nil {
		panic(err.Error())
	}

	count := 0
	for results.Next() {
		var usrmove Usermove
		count = count + 1
		if count == 10 {
			break
		}
		err = results.Scan(&usrmove.utctime, &usrmove.idfa, &usrmove.geohash, &usrmove.latitude, &usrmove.longitude, &usrmove.horizontal, &usrmove.userid, &usrmove.hourofday, &usrmove.lcltime)
		if err != nil {
			panic(err.Error())
		}

		fmt.Printf(strconv.Itoa(usrmove.utctime)+" "+strconv.FormatFloat(usrmove.latitude,'f',2,32)+" "+strconv.FormatFloat(usrmove.longitude,'f',2,32)+"\n")
	}
}