package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	hostname     = "<your_host_name>"
	host_port    = "<your_port>"
	username     = ">your_username>"
	password     = "<your_pass>"
	databasename = "<your_db_name>"
	sslmode      = "require"
)

func main() {

	/*********** Connection *************************************************************/
	// Veri tabanını vererek bağlan.

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		hostname, host_port, username, password, databasename)

	// Veri tabanı belirtmeden bağlan.
	/*
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s sslmode=disable",
			host, port, user, password)
	*/
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	if err == nil {
		fmt.Println("Bağlantı başarılı!")
	}
	defer db.Close()

	/*********** /Connection ************************************************************/

	/*********** Create Database ********************************************************/

	if _, err2 := db.Exec("CREATE KEYSPACE schduled"); err2 != nil {
		fmt.Println("DB oluşturulamadı.")
		log.Fatal(err2)
	}

	/*********** /Create Database *******************************************************/

	/*********** Create Table ***********************************************************/
	/*
		var createStmt = `CREATE TABLE weather (
				coord_longitude text,
				coord_latitude text,
				weather_id text,
				weather_main text,
				weather_description text,
				weather_icon text,
				base text,
				main_temp text,
				main_feels_like text,
				main_temp_min text,
				main_temp_max text,
				main_pressure text,
				main_humidity text,
				visibility text,
				wind_speed text,
				wind_deg text,
				clouds_all text,
				dt text,
				sys_type text,
				sys_id text,
				sys_country text,
				sys_sunrise text,
				sys_sunset text,
				timezone text,
				id text,
				name text,
				cod text
				)`
		if _, err := db.Exec(createStmt); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Created table weather")
	*/
	/*
		var createStmt = `CREATE TABLE scheduledFunctions (
			ID int,
			FunctionName text,
			SchuledTime text,
			Description text,
			CreatedTime timestamp,
			CreatedBy text) PRIMARY KEY ID`
		if _, err := db.Exec(createStmt); err != nil {
			log.Fatal(err)
		}
	*/
	/*********** /Create Table **********************************************************/

	/*********** İnsert Data ************************************************************/
	/*
		var insertStmt string = "INSERT INTO public.openweather(coord_longitude, coord_latitude, weather_id, weather_main, weather_description, weather_icon, base, main_temp, main_feels_like, main_temp_min, main_temp_max, main_pressure, main_humidity, visibility, wind_speed,wind_deg ,clouds_all, dt, sys_type, sys_id, sys_country, sys_sunrise, sys_sunset, timezone, id, name, cod) VALUES('" + string(obj.Coord.Lon) + "','" + lat + "','" + weatherid + "','" + weatherid + "','" + obj.Weather[0].Main + "','" + obj.Weather[0].Description + "','" + obj.Weather[0].Icon + "','" + obj.Base + "','" + temp + "','" + feelsLike + "','" + tempMin + "','" + tempMax + "','" + string(obj.Main.Pressure) + "','" + string(obj.Main.Humidity) + "','" + string(obj.Visibility) + "','" + speed + "','" + degree + "','" + string(obj.Clouds.All) + "','" + string(obj.Dt) + "','" + string(obj.Sys.Type) + "','" + string(obj.Sys.ID) + "','" + obj.Sys.Country + "','" + string(obj.Sys.Sunrise) + "','" + string(obj.Sys.Sunset) + "','" + string(obj.Timezone) + "','" + string(obj.ID) + "','" + obj.Name + "','" + string(obj.Cod) + "','" + ");"
		if _, err := db.Exec(insertStmt); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Inserted data: %s\n", insertStmt)
	*/
	/*********** /İnsert Data ************************************************************/

	/*********** Read Data ***************************************************************/
	/*
		var coord_longitude string
		var coord_latitude string
		var weather_id string
		//rows, err := db.Query(`SELECT coord_longitude, coord_latitude, weather_id, weather_main, weather_description, weather_icon, base, main_temp, main_feels_like, main_temp_min, main_temp_max, main_pressure, main_humidity, visibility, wind_speed, clouds_all, dt, sys_type, sys_id, sys_countrye, sys_sunrise, sys_sunset, timezone, id, name, cod FROM weather;`)
		rows, err := db.Query(`SELECT coord_longitude, coord_latitude, weather_id FROM weather;`)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		fmt.Printf("Query for id=1 returned: ")
		for rows.Next() {
			err := rows.Scan(&coord_longitude, &coord_latitude, &weather_id)
			if err != nil {
				log.Fatal("hata1: ", err)
			}
			fmt.Printf("Row[%s, %s, %s]\n", coord_longitude, coord_latitude, weather_id)
		}
		err = rows.Err()
		if err != nil {
			log.Fatal("hata2: ", err)
		}
	*/
	/*
		var name string
		var id int
		rows, err := db.Query(`SELECT ID, FunctionName FROM scheduledFunctions;`)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&id, &name)
			if err != nil {
				log.Fatal("hata1: ", err)
			}
			fmt.Printf("Row[%s, %s, %s]\n", id, name)
		}
		err = rows.Err()
		if err != nil {
			log.Fatal("hata2: ", err)
		}
	*/
	/*********** /Read Data **************************************************************/
}
