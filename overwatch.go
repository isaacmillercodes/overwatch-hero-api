package main

import (
  "database/sql"
  "github.com/gin-gonic/gin"
  _ "github.com/lib/pq"
  "gopkg.in/gorp.v1"
  "log"
  // "fmt"
  // "strconv"
)

type Hero struct {
	Id          string `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
	Role        string `db:"role" json:"role"`
	Secondary       string `db:"secondary" json:"secondary"`
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
      Secondary: hero.Secondary,
      Image: hero.Image,
      Difficulty: hero.Difficulty,
    }
    c.JSON(200, content)
  } else {
  c.JSON(404, gin.H{"error": "hero not found"})
  }
// curl -i http://localhost:8080/api/v1/heroes/01fc2354-d257-4ccf-b863-98f76a185c1d
}

// func PostHero(c *gin.Context) {
//   var hero Hero
//   c.Bind(&hero)
//   fmt.Println(hero)
//
//   if hero.Name != "" && hero.Description != "" && hero.Role != "" && hero.Secondary != "" && hero.Image != "" && hero.Difficulty != 0 {
//     fmt.Println("not empty!");
//     content := &Hero {
//       Name: hero.Name,
//       Description: hero.Description,
//       Role: hero.Role,
//       Secondary: hero.Secondary,
//       Image: hero.Image,
//       Difficulty: hero.Difficulty,
//     }
//     fmt.Println(content)
//     dbmap.Insert(content)
//     // insert != nil {
//       // fmt.Println(insert, "hi");
//       // user_id, err := insert.LastInsertId()
//       // if err == nil {
//         // c.JSON(201, content)
//       // } else {
//       //   checkErr(err, "Insert failed")
//       //   fmt.Println("oh no, there were errors");
//       // }
//     // }
//   } else {
//     c.JSON(422, gin.H{"error": "fields can't be empty"})
//     fmt.Println("some fields were empty, apparently");
//   }
//   fmt.Println("finished the conditional");
// // curl -i -X POST -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Queen\" }" http://localhost:8080/api/v1/users
// }

// func PostHero(c *gin.Context) {
//   var hero Hero
//   c.Bind(&hero)
//   fmt.Println(hero)
//
//   if hero.Name != "" && hero.Description != "" && hero.Role != "" && hero.Secondary != "" && hero.Image != "" && hero.Difficulty != 0 {
//     fmt.Println("not empty!");
//     if insert, _ := dbmap.Exec(`INSERT INTO heroes (name, description, role, secondary, image, difficulty) VALUES ($1, $1, $1, $1, $1, $1)`, hero.Name, hero.Description, hero.Role, hero.Secondary, hero.Image, hero.Difficulty); insert != nil {
//       fmt.Println(insert, "hi");
//       // user_id, err := insert.LastInsertId()
//       // if err == nil {
//         content := &Hero {
//           Name: hero.Name,
//           Description: hero.Description,
//           Role: hero.Role,
//           Secondary: hero.Secondary,
//           Image: hero.Image,
//           Difficulty: hero.Difficulty,
//         }
//         c.JSON(201, content)
//       // } else {
//       //   checkErr(err, "Insert failed")
//       //   fmt.Println("oh no, there were errors");
//       // }
//     }
//   } else {
//     c.JSON(422, gin.H{"error": "fields can't be empty"})
//     fmt.Println("some fields were empty, apparently");
//   }
//   fmt.Println("finished the conditional");
// // curl -i -X POST -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Queen\" }" http://localhost:8080/api/v1/users
// }

func PostHero(c *gin.Context) {


  dbmap.Exec(`INSERT INTO heroes (name, description, role, secondary, image, difficulty) VALUES ('Isaac', 'Description', 'Role', 'Secondary', 'Image', 3)`);

  c.JSON(201, gin.H{"status": "success!"})
// curl -i -X POST -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Queen\" }" http://localhost:8080/api/v1/users
}

// func PostHero(c *gin.Context) {
//   var hero Hero
//   c.Bind(&hero)
//   fmt.Println(hero)
//
//   if hero.Name != "" && hero.Description != "" && hero.Role != "" && hero.Secondary != "" && hero.Image != "" && hero.Difficulty != 0 {
//     fmt.Println("not empty!");
//     insert := dbmap.Insert(hero.Name, hero.Description, hero.Role, hero.Secondary, hero.Image, hero.Difficulty);
//     if insert != nil {
//       fmt.Println(insert, "hi");
//       // user_id, err := insert.LastInsertId()
//       // if err == nil {
//         content := &Hero {
//           Name: hero.Name,
//           Description: hero.Description,
//           Role: hero.Role,
//           Secondary: hero.Secondary,
//           Image: hero.Image,
//           Difficulty: hero.Difficulty,
//         }
//         c.JSON(201, content)
//       // } else {
//       //   checkErr(err, "Insert failed")
//       //   fmt.Println("oh no, there were errors");
//       // }
//     }
//   } else {
//     c.JSON(422, gin.H{"error": "fields can't be empty"})
//     fmt.Println("some fields were empty, apparently");
//   }
//   fmt.Println("finished the conditional");
// // curl -i -X POST -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Queen\" }" http://localhost:8080/api/v1/users
// }

// func PostHero(c *gin.Context) {
//   var hero Hero
//   c.Bind(&hero)
//   fmt.Println(hero)
//
//   if hero.Name != "" && hero.Description != "" && hero.Role != "" && hero.Secondary != "" && hero.Image != "" && hero.Difficulty != 0 {
//     fmt.Println("not empty!");
//     insert := dbmap.Insert(hero.Name, hero.Description, hero.Role, hero.Secondary, hero.Image, hero.Difficulty);
//     if insert != nil {
//       fmt.Println(insert, "hi");
//       // user_id, err := insert.LastInsertId()
//       // if err == nil {
//         content := &Hero {
//           Name: hero.Name,
//           Description: hero.Description,
//           Role: hero.Role,
//           Secondary: hero.Secondary,
//           Image: hero.Image,
//           Difficulty: hero.Difficulty,
//         }
//         c.JSON(201, content)
//       // } else {
//       //   checkErr(err, "Insert failed")
//       //   fmt.Println("oh no, there were errors");
//       // }
//     }
//   } else {
//     c.JSON(422, gin.H{"error": "fields can't be empty"})
//     fmt.Println("some fields were empty, apparently");
//   }
//   fmt.Println("finished the conditional");
// // curl -i -X POST -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Queen\" }" http://localhost:8080/api/v1/users
// }

// func PostHero(c *gin.Context) {
//     var json Hero
//
//     c.Bind(&json) // This will infer what binder to use depending on the content-type header.
//     hero := createHero()
//     if hero.Name == json.Name {
//         content := gin.H{
//             "result": "Success",
//             "name": hero.Name,
//             "description": hero.Description,
//         }
//         c.JSON(201, content)
//     } else {
//         c.JSON(500, gin.H{"result": "An error occured"})
//     }
// }
//
// func createHero() Hero {
//     hero := Hero{
//         Name: "Isaac",
//         Description: "Description",
//         Role: "Role",
//         Secondary: "Secondary",
//         Image: "Image",
//         Difficulty: 2,
//     }
//
//     err := dbmap.Insert(&hero)
//     checkErr(err, "Insert failed")
//     return hero
// }

func UpdateHero(c *gin.Context) {
 // The futur code…
}
func DeleteHero(c *gin.Context) {
 // The futur code…
}
