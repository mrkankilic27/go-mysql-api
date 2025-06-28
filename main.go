package main

import (
	"database/sql"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db := connectDb()
	defer db.Close()
	CreateUser(db)
	getAllUsers(db)
}

var (
	ID        int
	Username  string
	FirstName string
	LastName  string
	Password  string
	Email     string
	BirthDate string
	IsActive  bool
)

func CreateUser(db *sql.DB) {
	if db.Ping() == nil {
		// createDb := "`users`. (`ID` INT NOT NULL AUTO_INCREMENT , `Username` VARCHAR(50) NOT NULL , `Email` VARCHAR(50) NOT NULL , `Password` VARCHAR(50) NOT NULL , `FirstName` VARCHAR(50) NOT NULL , `LastName` VARCHAR(50) NOT NULL , `BirthDate` DATE NOT NULL , `IsActive` BOOLEAN NOT NULL , PRIMARY KEY (`ID`)) ENGINE = MyISAM;"
		// _, errCreateDB := db.Exec("CREATE TABLE IF NOT EXİSTS " + createDb)

		// if errCreateDB != nil {
		// 	log.Fatal(errCreateDB)
		// }

		response, errCreateUser := db.Exec("INSERT INTO `users` (`Username`, `Email`, `Password`, `FirstName`, `LastName`, `BirthDate`, `IsActive`) VALUES ('workbenchoff2', 'deneme2@deneme.com', 'deneme2', 'deneme2', 'denemeoğlu', '2004-02-10', '1')")
		if errCreateUser != nil {
			log.Fatal(errCreateUser)
		}

		rowCount, errAffectedRows := response.RowsAffected()
		if errAffectedRows != nil {
			log.Fatal(errAffectedRows)
		}
		log.Printf("Inserted %d rows", rowCount)
	}
}

func getAllUsers(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		rows.Scan(&ID, &Username, &Email, &Password, &FirstName, &LastName, &BirthDate, &IsActive)
		log.Printf("Bulunan satır içeriği : %q", strconv.Itoa(ID)+" "+Username+" "+Email+" "+Password+" "+FirstName+" "+LastName+" "+BirthDate+" "+strconv.FormatBool(IsActive))
	}
}

func selectUser(id int, db *sql.DB) {
	rows, errSelectUser := db.Query("SELECT * FROM users WHERE ID= ?", id)
	if errSelectUser != nil {
		log.Fatal(errSelectUser)
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&ID, &Username, &Email, &Password, &FirstName, &LastName, &BirthDate, &IsActive)
		log.Printf("Bulunan satır içeriği : %q", strconv.Itoa(ID)+" "+Username+" "+Email+" "+Password+" "+FirstName+" "+LastName+" "+BirthDate+" "+strconv.FormatBool(IsActive))
	}
}

func connectDb() *sql.DB {
	db, err := sql.Open("mysql", "root:toor@tcp(127.0.0.1:3306)/users")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
