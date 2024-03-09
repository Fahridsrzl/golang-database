package main

import (
	"database/sql"
	"fmt"
	"golang-database/entity"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "yourpassword"
	dbname   = "enigmacamp"
)

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

func main() {
	// student := entity.Student{Id: 9, Name: "marco", Email: "marco@mail.com", Address: "kuta", BirthDate: time.Date(2003, 03, 23, 0, 0, 0, 0, time.Local), Gender: "M"}

	// addStudent(student)
	// updateStudent(student)
	// deleteStudent("9")

	// students := getAllStudent()
	// for _, student := range students {
	// 	fmt.Println(student.Id, student.Name, student.Email, student.Address, student.BirthDate, student.Gender)
	// }

	// fmt.Println(getStudentById(7))
	students := searchBy("me")
	for _, student := range students {
		fmt.Println(student.Id, student.Name, student.Address, student.BirthDate, student.Gender)
	}

}

func addStudent(student entity.Student) {
	db := connectDb()
	defer db.Close()
	var err error

	sqlStatement := "INSERT INTO mst_student (id, name, email, address, birth_date, gender) VALUES ($1, $2, $3 ,$4, $5, $6);"

	_, err = db.Exec(sqlStatement, student.Id, student.Name, student.Email, student.Address, student.BirthDate, student.Gender)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Insert Success")
	}

}

func updateStudent(student entity.Student) {
	db := connectDb()
	defer db.Close()
	var err error

	sqlStatement := "UPDATE mst_student SET name = $2, email = $3, address = $4, birth_date = $5, gender = $6 WHERE id = $1;"

	_, err = db.Exec(sqlStatement, student.Id, student.Name, student.Email, student.Address, student.BirthDate, student.Gender)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Update Success")
	}
}

func deleteStudent(id string) {
	db := connectDb()
	defer db.Close()
	var err error

	sqlStatement := "DELETE FROM mst_student WHERE id = $1;"

	_, err = db.Exec(sqlStatement, id)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Delete Success")
	}
}

func getAllStudent() []entity.Student {
	db := connectDb()
	defer db.Close()

	sqlStatement := "SELECT * FROM mst_student;"

	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	students := scanStudent(rows)

	return students
}

func scanStudent(rows *sql.Rows) []entity.Student {
	students := []entity.Student{}
	var err error

	for rows.Next() {
		student := entity.Student{}
		err := rows.Scan(&student.Id, &student.Name, &student.Email, &student.Address, &student.BirthDate, &student.Gender)
		if err != nil {
			panic(err)
		}

		students = append(students, student)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return students
}

func getStudentById(id int) entity.Student {
	db := connectDb()
	defer db.Close()
	var err error

	sqlStatement := "SELECT * FROM mst_student WHERE id = $1"

	student := entity.Student{}
	err = db.QueryRow(sqlStatement, id).Scan(&student.Id, &student.Name, &student.Email, &student.Address, &student.BirthDate, &student.Gender)
	if err != nil {
		panic(err)
	}

	return student
}

func searchBy(name string) []entity.Student {
	db := connectDb()
	defer db.Close()

	sqlStatement := "SELECT * FROM mst_student WHERE name LIKE $1"

	rows, err := db.Query(sqlStatement, "%"+name+"%")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	students := scanStudent(rows)

	return students
}

func connectDb() *sql.DB {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Succesfully Connected!")
	}
	return db
}
