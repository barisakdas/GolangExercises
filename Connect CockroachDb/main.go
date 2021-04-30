package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {

	const (
		hostname     = "<your_host_name>"
		host_port    = "<your_port>"
		username     = ">your_username>"
		password     = "<your_pass>"
		databasename = "postgre"
		sslmode      = "require"
	)

	/*************** CREATE CONNECTİON *******************************************************************************************************/
	// Bağlantı esnasında db adı verilerek direkt db ye de bağlanılabilir. Stringe databasename=<database_adi> verilerek yapabiliriz.
	pg_con_string := fmt.Sprintf("port=%d host=%s user=%s "+"password=%s sslmode=%s", host_port, hostname, username, password, sslmode)

	db, err := sql.Open("postgres", pg_con_string)
	if err != nil {
		panic(err)
	}
	if err == nil {
		fmt.Println("Bağlantı başarılı")
	}
	defer db.Close()

	/*************** CREATE DATABASE ********************************************************************************************************/
	/*
		// Database oluşturma.
		if _, err2 := db.Exec("CREATE DATABASE IF NOT EXISTS bank"); err2 != nil {
			fmt.Println("DB oluşturulamadı.")
			log.Fatal(err2)
		}
	*/
	/*************** CREATE TABLE **********************************************************************************************************/
	/*
		// Table oluşturma. Burada hangi db ye yapacaksak onun ismini tablonun başına koymamız gerekiyor. '<db_name>.<table_name>' şeklinde olacak.
		if _, err2 := db.Exec("CREATE TABLE bank.tblusers(id INT PRIMARY KEY, name STRING, age INT, accessLevel STRING, department STRING, title STRING );"); err2 != nil {
			fmt.Println("Table oluşturulamadı.")
			log.Fatal(err2)
		}
	*/

	/*************** CREATE DATA ***********************************************************************************************************/
	/*
		// bank veri tabanı altındaki tblusers tablosuna kayıt ekleme.
		stmt, err := db.Prepare("INSERT INTO bank.tblusers(id,name,age,accessLevel,department,title ) VALUES($1,$2,$3,$4,$5,$6)")
		if err != nil {
			log.Fatal(err)
		}

		defer stmt.Close()

		res, err := stmt.Exec(4, "Tufan", 27, "Director", "System", "Developer")
		if err != nil {
			log.Fatal(err)
		}

		affect, err := res.RowsAffected()
		if err != nil {
			log.Fatal(err)
			fmt.Println("Error !!!")
		}

		fmt.Println(affect, "record added")
	*/
	/*************** GET DATA **************************************************************************************************************/

	// bank veri tabanı altındaki tblusers tablosundaki kayıtları getir.

	rows, err := db.Query("SELECT name FROM bank.tblusers;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		fmt.Println(name)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
