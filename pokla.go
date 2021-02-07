package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Poklo struct {
	Id            int
	Uid           sql.NullString
	Plan_uid      sql.NullString
	Floor         int
	Rooms         int
	Price         float64
	Area          sql.NullString
	Views         sql.NullString
	Level         sql.NullString
	Name          sql.NullString
	Bath          sql.NullString
	Fireplace     sql.NullString
	Deal          sql.NullString
	Baths         sql.NullString
	Status        sql.NullString
	Pdf_id        sql.NullString
	Pdf_filename  sql.NullString
	Plan_id       sql.NullString
	Plan_filename sql.NullString
}

const pokloSql = `
SELECT
	a.id,
	a.uid,
	plan_uid,
	floor,
	rooms,
	price,
	area,
	views,
	level,
	name,
	detached_bath AS bath,
	fireplace,
	deal_uid AS deal,
	baths,
	status,
	pdf_id,
	pdf_filename,
	plans.plan_id,
	plans.plan_filename
FROM
	apts a
	INNER JOIN plans ON a.plan_uid = plans.uid
WHERE
	a.status <> 'outdated'
ORDER BY
	price ASC
LIMIT
	50
OFFSET
	250
`

func main() {
	r := gin.Default()

	db, err := sqlx.Connect("postgres", "host=localhost user=postgres dbname=poklo password=postgres port=6432 sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	r.GET("/", func(c *gin.Context) {

		// this Pings the database trying to connect
		// use sqlx.Open() for sql.Open() semantics

		poklos := []Poklo{}

		err = db.Select(&poklos, pokloSql)

		if err != nil {
			fmt.Println(err)
			return
		}

		c.JSON(200, poklos)
	})
	r.Run("0.0.0.0:8082")
}
