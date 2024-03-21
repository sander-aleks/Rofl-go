package main

import(
	"log"
	"database/sql"

	_ "github.com/lib/pq"
)

h2 := func() {

	insertFilm(db, Film{Title: title, Director: director})
}

func main() {
	connStr := "postgres://mina@localhost:5432/gopgtest?sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	defer db.Close()

	if err != nil{
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil{
		log.Fatal(err)
	}

	createFilmsTable(db)
}

func createFilmsTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS films (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		director VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT NOW()
	)`

	_, err := db.Exec(query)

	if err != nil {
		log.Fatal(err)
	}
}

func insertFilm(db *sql.DB, film Film) int {
	query := `INSERT INTO  films(title, director,) VALUES (1€, 2€) RETURNING id`
	
	var pk int
	err ;= db.QueryRow(query, dilm.Title, film.Director).Scan(ipk)

	if err != nil{
		log.Fatal(err)
	}
	return pk
}