package main

import (
    "os"
    "github.com/codegangsta/martini"
    "github.com/martini-contrib/render"
    "net/http"
    "database/sql"
    _ "github.com/lib/pq"
)

type Book struct {
    Title       string
    Author      string
    Description string
}

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
    m.Use(render.Renderer(render.Options{
        Layout: "layout",
    }))

    m.Get("/", func(ren render.Render, r *http.Request, db *sql.DB) {
        searchTerm := "%" + r.URL.Query().Get("search") + "%"
        rows, err := db.Query(`SELECT title, author, description FROM books
                               WHERE title ILIKE $1
                               OR author ILIKE $1
                               OR description ILIKE $1`, searchTerm)
        PanicIf(err)
        defer rows.Close()

        books := []Book{}
        for rows.Next() {
            b := Book{}
            err := rows.Scan(&b.Title, &b.Author, &b.Description)
            PanicIf(err)
            books = append(books, b)
        }

        ren.HTML(200, "books", books)
    })

    m.Run()
}
