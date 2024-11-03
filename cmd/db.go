package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "github.com/microsoft/go-mssqldb"
)

var (
	db *sql.DB
	server = "erfeiyu.database.windows.net"
	dbPort = 1433
	user = "erfei"
	password = "terryhimself88."
	database = "erfeiyu"

	dsn = ""
)

func init() {
	dsn = fmt.Sprintf(
		"server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, dbPort, database,
	)
	fmt.Printf("dsn: %v\n", dsn)

	var err error
	db, err = sql.Open("sqlserver", dsn)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println("db init successfully")
}

func InsertRandomRecord(ctx context.Context) error {
	stmt := `
		INSERT INTO EnvironmentalData
		(SensorID, Temperature, WindSpeed, RelativeHumidity, CO2Level, Timestamp)
		VALUES
		(@p1, @p2, @p3, @p4, @p5, @p6);
	`

	sensorId := 1 + rand.Intn(20)
	temperature := 8 + rand.Intn(8)
	windSpeed := 15 + rand.Intn(11)
	humidity := 40 + rand.Intn(31)
	co2 := 500 + rand.Intn(1001)

	_, err := db.ExecContext(
		ctx,
		stmt,
		sql.Named("p1", sensorId),
		sql.Named("p2", temperature),
		sql.Named("p3", windSpeed),
		sql.Named("p4", humidity),
		sql.Named("p5", co2),
		sql.Named("p6", time.Now()),
	)
	if err != nil {
		return err
	}

	fmt.Printf(
		"SensorID:(%v) Temperature(%v) WindSpeed(%v) RelativeHumidity(%v) CO2Level(%v)\n",
		sensorId, temperature, windSpeed, humidity, co2,
	)
	return nil
}

func GetRows(ctx context.Context) []*EnvironmentalData {
    query := `
        SELECT ID, SensorID, Temperature, WindSpeed, RelativeHumidity, CO2Level
        FROM EnvironmentalData
    `

    data := make([]*EnvironmentalData, 0)
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		log.Println(err)
		return nil
	}

	for rows.Next() {
		d := &EnvironmentalData{}

		err := rows.Scan(
			&d.ID,
			&d.SensorID,
			&d.Temperature,
			&d.WindSpeed,
			&d.RelativeHumidity,
			&d.CO2Level,
		)
		if err != nil {
			log.Printf("error while scan rows: %v\n", err)
		}

		data = append(data, d)
	}


    return data
}