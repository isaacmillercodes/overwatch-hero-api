package main

import (
  "database/sql"
  "github.com/gin-gonic/gin"
  _ "github.com/lib/pq"
  "gopkg.in/gorp.v1"
  "log"
  "fmt"
  // "strconv"
)

type Hero struct {
	Id          string `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
	Role        string `db:"role" json:"role"`
	Secondary   string `db:"secondary" json:"secondary"`
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
      Secondary: hero.Secondary,
      Image: hero.Image,
      Difficulty: hero.Difficulty,
    }
    c.JSON(200, content)
  } else {
  c.JSON(404, gin.H{"error": "hero not found"})
  }
}

func PostHero(c *gin.Context) {

  var hero Hero
  c.Bind(&hero)

  if hero.Name != "" && hero.Description != "" && hero.Role != "" && hero.Secondary != "" && hero.Image != "" && hero.Difficulty > 0 && hero.Difficulty < 4 {

    if insert, _ := dbmap.Exec(`INSERT INTO heroes (name, description, role, secondary, image, difficulty) VALUES ($1, $2, $3, $4, $5, $6)`, hero.Name, hero.Description, hero.Role, hero.Secondary, hero.Image, hero.Difficulty); insert != nil {

      c.JSON(201, gin.H{"success": "New hero added!"})
    }

  } else {
    c.JSON(422, gin.H{"error": "fields were missing"})
  }
  //http POST http://localhost:8080/api/v1/heroes name="CoolCharacter" description="This character is super cool" role="Offense" secondary="Being Cool" image="image.jpg" difficulty:=3
}

func UpdateHero(c *gin.Context) {

  id := c.Params.ByName("id")
  var hero Hero
  err := dbmap.SelectOne(&hero, "SELECT * FROM heroes WHERE id=$1", id)
  if err == nil {
    var changed Hero
    c.Bind(&changed)

    if changed.Name == "" && changed.Description == "" && changed.Role == "" && changed.Secondary == "" && changed.Image == "" && changed.Difficulty == 0 {
      c.JSON(422, gin.H{"error": "all fields were empty"})
      fmt.Println("all fields were empty!")
    }

    if changed.Name != "" {
      hero.Name = changed.Name
    }
    if changed.Description != "" {
      hero.Description = changed.Description
    }
    if changed.Role != "" {
      hero.Role = changed.Role
    }
    if changed.Secondary != "" {
      hero.Secondary = changed.Secondary
    }
    if changed.Image != "" {
      hero.Image = changed.Image
    }
    if changed.Difficulty != 0 {
      hero.Difficulty = changed.Difficulty
    }

    if update, err2 := dbmap.Exec(`UPDATE heroes SET name=$1, description=$2, role=$3, secondary=$4, image=$5, difficulty=$6 WHERE id=$7`, hero.Name, hero.Description, hero.Role, hero.Secondary, hero.Image, hero.Difficulty, id); update != nil {
      c.JSON(200, hero)
    } else {
      checkErr(err2, "Updated failed")
    }

  } else {
    c.JSON(404, gin.H{"error": "hero not found"})
  }
}

func DeleteHero(c *gin.Context) {
 // The future code…
}
