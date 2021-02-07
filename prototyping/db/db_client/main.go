package main

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Define default settings for the database defined via `docker-compose.yml`
const (
	host     = "localhost" // Should be changed to "db" if this go service runs in same docker container
	port     = 5432
	user     = "db_user"
	password = "secret"
	dbName   = "demo_db"
)

// getConnStr prepares and returns the connection string to the default database
func getConnStr() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)
}

// checkConn tests if the default database is accessible
func checkConn() {
	connStr := getConnStr()

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Printf("\nSuccessfully connected to the database!\n")
}

// Create tables and prepare some mock data
var schema = `
CREATE TABLE IF NOT EXISTS country (
	id integer,
	name text,
	description text
);

DELETE FROM country;

INSERT INTO country(id, name, description) VALUES
	(1, 'Greece', 'Great country'),
	(2, 'Poland', 'Yet another great country');
`

// Country binds to the `country` table in the db
type Country struct {
	ID          int
	Name        string
	Description string
}

// GetAllCountries returns all countries stored in the `country` table
func GetAllCountries(db *sqlx.DB) ([]Country, error) {
	countries := []Country{}
	err := db.Select(&countries, `SELECT * FROM country ORDER BY id ASC`)
	if err != nil {
		return countries, err
	}
	return countries, nil
}

func main() {
	// We are going to try to write and read data from the default database.

	// Quick test if db is accessible.
	fmt.Println("\nStart. Trying to test the connection to the database...")
	checkConn()

	fmt.Println("\nTrying to connect to the database...")
	db, err := sqlx.Connect("postgres", getConnStr())
	if err != nil {
		panic(err)
	}

	// Create tables in the database, add data.
	fmt.Println("\nTrying to deploy the schema...")
	db.MustExec(schema)

	// Fetch all rows.
	fmt.Println("\nTrying to fetch existing data from the database...")
	countries := []Country{}
	err = db.Select(&countries, `SELECT * FROM country ORDER BY id ASC`)
	if err != nil {
		fmt.Println(err)
	}
	for _, c := range countries {
		fmt.Printf("%#v\n", c)
	}

	// Fetch specific row.
	fmt.Println("\nTrying to fetch the data for a specific country (Greece)...")
	greece := Country{}
	err = db.Get(&greece, `SELECT * FROM country WHERE name=$1`, "Greece")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", greece)

	// Insert a new row.
	fmt.Println("\nTrying to create a new country and save the data in the database (Bolivia)...")
	bolivia := Country{ID: 3, Name: "Bolivia", Description: "This is a cool country, too"}
	newCountries := []Country{bolivia}
	_, err = db.NamedQuery(`INSERT INTO country (id, name, description)
	VALUES (:id, :name, :description)`, newCountries)
	if err != nil {
		fmt.Println(err)
	}

	// Fetch all rows, again, to see if the insert succeeded.
	fmt.Println("\nTrying to fetch current data from the database...")
	allCountries, err := GetAllCountries(db)
	if err != nil {
		fmt.Println(err)
	}
	for _, c := range allCountries {
		fmt.Printf("%#v\n", c)
	}

	// No errors at this point? SUCCESS!
	fmt.Println("\nSuccess! All tests completed successfully!")
}
