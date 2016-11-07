package main

import (
  "database/sql"
  "github.com/gin-gonic/gin"
  _ "github.com/lib/pq"
  "gopkg.in/gorp.v1"
  "log"
  // "strconv"
)

type Hero struct {
	Id          string `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
	Role        string `db:"role" json:"role"`
	Type        string `db:"type" json:"type"`
	Image       string `db:"image" json:"image"`
	Difficulty  int    `db:"difficulty" json:"difficulty"`
}

var dbmap = initDb()

func initDb() *gorp.DbMap {
  db, err := sql.Open("postgres", "postgres://Isaac:password@localhost/overwatch_api?sslmode=disable")
  checkErr(err, "sql.Open failed")
  dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

  return dbmap
}

func checkErr(err error, msg string) {
  if err != nil {
    log.Fatalln(msg, err)
  }
}

func main() {
  r := gin.Default()
  v1 := r.Group("api/v1")
  {
  v1.GET("/heroes", GetHeroes)
  v1.GET("/heroes/:id", GetHero)
  v1.POST("/heroes", PostHero)
  v1.PUT("/heroes/:id", UpdateHero)
  v1.DELETE("/heroes/:id", DeleteHero)
  }
  r.Run(":8080")
}

func GetHeroes(c *gin.Context) {
  var heroes []Hero
  _, err := dbmap.Select(&heroes, "SELECT * FROM heroes")
  if err == nil {
    c.JSON(200, heroes)
  } else {
    c.JSON(404, gin.H{"error": "no heroes found"})
  }
  // curl -i http://localhost:8080/api/v1/heroes
}

func GetHero(c *gin.Context) {
  id := c.Params.ByName("id")
  var hero Hero
  err := dbmap.SelectOne(&hero, "SELECT * FROM heroes WHERE id=$1", id)
  if err == nil {
    content := &Hero {
    Id: id,
    Name: hero.Name,
    Description: hero.Description,
    Role: hero.Role,
    Type: hero.Type,
    Image: hero.Image,
    Difficulty: hero.Difficulty,
    }
    c.JSON(200, content)
  } else {
  c.JSON(404, gin.H{"error": "hero not found"})
  }
// curl -i http://localhost:8080/api/v1/users/1
}

func PostHero(c *gin.Context) {
 // The futur code…
}
func UpdateHero(c *gin.Context) {
 // The futur code…
}
func DeleteHero(c *gin.Context) {
 // The futur code…
}
