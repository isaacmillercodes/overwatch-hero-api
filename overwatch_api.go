package main

import (
  _ "github.com/lib/pq"
  "database/sql"
  // "fmt"
  "log"
  "net/http"
  "encoding/json"
  // "strconv"
)

type Hero struct {
	Id          string
	Name        string
	Description string
	Role        string
	Secondary   string
	Image       string
	Difficulty  int
}

var db *sql.DB

func init() {
  var err error
  db, err = sql.Open("postgres", "postgres://Isaac:password@localhost/overwatch_api?sslmode=disable")
  if err != nil {
    log.Fatal(err)
  }

  if err = db.Ping(); err != nil {
    log.Fatal(err)
  }
}

func main() {
  http.HandleFunc("/heroes", getAllHeroes)
  http.HandleFunc("/heroes/:id", getSingleHero)
  // http.HandleFunc("/heroes", createHero)
  http.ListenAndServe(":3000", nil)
}

func getAllHeroes(w http.ResponseWriter, r *http.Request) {
  if r.Method != "GET" {
    http.Error(w, http.StatusText(405), 405)
    return
  }

  rows, err := db.Query("SELECT * FROM heroes")
  if err != nil {
    http.Error(w, http.StatusText(500), 500)
    return
  }
  defer rows.Close()

  heroes := make([]*Hero, 0)
  for rows.Next() {
    hero := new(Hero)
    err := rows.Scan(&hero.Id, &hero.Name, &hero.Description, &hero.Role, &hero.Secondary, &hero.Image, &hero.Difficulty)
    if err != nil {
      http.Error(w, http.StatusText(500), 500)
      return
    }
    heroes = append(heroes, hero)
  }
  if err = rows.Err(); err != nil {
    http.Error(w, http.StatusText(500), 500)
    return
  } else {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    json.NewEncoder(w).Encode(heroes)
  }

}

func getSingleHero(w http.ResponseWriter, r *http.Request) {
  if r.Method != "GET" {
    http.Error(w, http.StatusText(405), 405)
    return
  }

  id := r.URL.Query().Get("id")
  if id == "" {
    http.Error(w, http.StatusText(400), 400)
    return
  }

  row := db.QueryRow("SELECT * FROM books WHERE id = $1", id)

  hero := new(Hero)
  err := row.Scan(&hero.Id, &hero.Name, &hero.Description, &hero.Role, &hero.Secondary, &hero.Image, &hero.Difficulty)
  if err == sql.ErrNoRows {
    http.NotFound(w, r)
    return
  } else if err != nil {
    http.Error(w, http.StatusText(500), 500)
    return
  }
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  json.NewEncoder(w).Encode(hero)
}
