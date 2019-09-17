package main

import "fmt"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"

type database struct {
	db         *sql.DB
	dbName     string
	dbUser     string
	dbPassword string
}

type phone struct {
	id     int
	number string
}

// const (
// 	dbName     = "phone"
// 	dbUser     = "root"
// 	dbPassword = "k9eUnehq#KD72"
// )

func (database database) connectToDatabase(dbName string) (*sql.DB, error) {
	fmt.Println("Connecting to database...")
	dbSource := fmt.Sprintf("%s:%s@/%s", database.dbUser, database.dbPassword, dbName)
	var err error
	database.db, err = sql.Open("mysql", dbSource)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return database.db, err
}

func (database database) resetDatabase() error {
	var err error
	database.db, err = database.connectToDatabase("")
	if err != nil {
		panic(err)
	}
	_, err2 := database.db.Exec("DROP DATABASE IF EXISTS " + database.dbName)
	if err2 != nil {
		return err2
	}
	defer database.closeConnection()
	return database.createDatabase()
}

func (database database) createDatabase() error {
	_, err := database.db.Exec("CREATE DATABASE " + database.dbName)
	return err
}

func (database database) createPhoneTable() error {
	var err error
	statement := `
		CREATE TABLE IF NOT EXISTS phone_number (
			id INT PRIMARY KEY AUTO_INCREMENT,
			value VARCHAR(100)
		)`
	_, err = database.db.Exec(statement)
	return err
}

func (database database) insertPhone(phone string) (int, error) {
	statement := `INSERT INTO phone_number(value) VALUES(?)`
	result, err := database.db.Exec(statement, phone)
	if err != nil {
		return -1, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}
	return int(id), nil
}

func (database database) getAllPhones() ([]phone, error) {
	rows, err := database.db.Query("SELECT id, value FROM phone_number")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var ret []phone
	for rows.Next() {
		var p phone
		if err := rows.Scan(&p.id, &p.number); err != nil {
			return nil, err
		}
		ret = append(ret, p)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return ret, nil
}

func (database database) getPhone(id int) (string, error) {
	var number string
	row := database.db.QueryRow("SELECT * FROM phone_number WHERE id = ?", id)
	err := row.Scan(&id, &number)
	if err != nil {
		return "", err
	}
	return number, nil
}

func (database database) findPhone(value string) (*phone, error) {
	var p phone
	row := database.db.QueryRow("SELECT * FROM phone_number WHERE value = ?", value)
	err := row.Scan(&p.id, &p.number)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return &p, nil
}

func (database database) updatePhone(p phone) error {
	statement := `UPDATE phone_number SET value = ? WHERE id = ?`
	_, err := database.db.Exec(statement, p.number, p.id)
	return err
}

func (database database) deletePhone(p phone) error {
	statement := `DELETE FROM phone_number WHERE id = ?`
	_, err := database.db.Exec(statement, p.id)
	return err
}

func (database database) closeConnection() {
	fmt.Println("Closing connection...")
	database.db.Close()
}
