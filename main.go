package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

const dsn = "database.sqlite"

type Employee struct {
	ID   int
	Name string
	Dep  sql.NullString
}

func (e Employee) String() string {
	return fmt.Sprintf("ID:%v, Name: %v, Dep: %v", e.ID, e.Name, e.Dep.String)
}

func main() {
	db, err := sql.Open("sqlite3", dsn)
	// here you should configure the db: maopenconn, maxidleconns, etc...
	if err != nil {
		log.Fatalf("Error while opening the db: %v\n", err)
		return
	}
	// always test if the db is available
	err = db.Ping()
	if err != nil {
		log.Fatalf("DB is not available: %v\n", err)
		return
	} else {
		log.Println("Connected [OK]")
	}

	//empl, errGet := GetEmpl(db, 1)
	//if errGet != nil {
	//	log.Fatalf("Error geting emp with id %v: %v\n", 1, errGet)
	//	return
	//}

	empls, errGet := Get(db)
	if errGet != nil {
		log.Fatal(errGet)
		return
	}

	for i := range empls {
		fmt.Println(empls[i])
	}

}

func GetEmpl(db *sql.DB, id int) (Employee, error) {
	stmt, err := db.Prepare("select * from employees where employee_id= ?;")
	if err != nil {
		log.Fatalf("failed to prepare the stmt: %v\n", err)
		return Employee{}, err
	}
	row := stmt.QueryRow(1)
	empl := Employee{}

	err = row.Scan(&empl.ID, &empl.Name, &empl.Dep)
	if err != nil {
		log.Fatalf("failed to scan the row to struct empl: %v\n", err)
		return Employee{}, err
	}
	return empl, nil
}

func Get(db *sql.DB) ([]Employee, error) {
	rows, err := db.Query("select * from employees;")
	if err != nil {
		log.Fatalf("failed query all empls: %v\n", err)
		return []Employee{}, err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Fatalf("error while closing rows: %v\n", err)
			return
		}
	}()

	var empls []Employee

	for rows.Next() {
		e := Employee{}
		err := rows.Scan(&e.ID, &e.Name, &e.Dep)
		if err != nil {
			log.Fatalf("error while sanning rows: %v\n", err)
			return nil, err
		}
		empls = append(empls, e)
	}

	// Checks rows.Err to determine whether the previous loop stopped because of an error
	if err := rows.Err(); err != nil {
		log.Fatalf("error while preparing the next row in the loop: %v\n", err)
		return []Employee{}, err
	}

	return empls, nil
}
