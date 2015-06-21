package main

import (
    "fmt"
    "os"
    "github.com/codegangsta/martini"
    "net/http"
    "database/sql"
    _ "github.com/lib/pq"
)

func SetupDB() *sql.DB {
    dbUrl := os.Getenv("DATABASE_URL")
    if len(dbUrl) == 0 {
        dbUrl = "postgres://postgres:postgres@127.0.0.1:5432/go-books?sslmode=disable"
    }

    db, err := sql.Open("postgres", dbUrl)
    PanicIf(err)
    return db
}

func PanicIf(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {
    m := martini.Classic()
    m.Map(SetupDB())

    m.Get("/", func(rw http.ResponseWriter, r *http.Request, db *sql.DB) {
        searchTerm := "%" + r.URL.Query().Get("search") + "%"
        rows, err := db.Query(`SELECT title, author, description FROM books
                               WHERE title ILIKE $1
                               OR author ILIKE $1
                               OR description ILIKE $1`, searchTerm)
        PanicIf(err)
        defer rows.Close()

        var title, author, description string
        for rows.Next() {
            err := rows.Scan(&title, &author, &description)
            PanicIf(err)
            fmt.Fprintf(rw, "Title: %s\nAuthor: %s\nDescription: %s\n\n",
                title, author, description)
        }
    })

    m.Run()
}
