package main

import "fmt"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"

type student struct {
	id string
	name string
	age int
	grade int
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/laravue")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func sqlQuery() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var age = 27
	rows, err := db.Query("Select id, name, age, grade from golang where age = ?", age)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []student

	for rows.Next() {
		var  each = student{}
		var err = rows.Scan(&each.id, &each.name, &each.age, &each.grade)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range result {
		fmt.Println(each.name)
	}

}

func sqlQueryRow() {
	var db, err = connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var result = student{}
	var id = 4
	err = db.QueryRow("select name from golang where id = ?", id).Scan(&result.name)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("name: %s\n", result.name)

}

func sqlExec() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("insert into golang values(?,?,?,?)", nil,"karno", 30, 1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("insert success!")


}

func main() {
	//sqlQuery()
	//sqlQueryRow()
	sqlExec()
}