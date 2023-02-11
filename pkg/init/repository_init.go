package init

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

const (
	host   = "localhost"
	port   = 5432
	dbname = "product-marketplace"
)

var DBInstance *sqlx.DB

func InitDbConnection() {
	psqlInfo := fmt.Sprintf("host=%s port=%d dbname=%s sslmode=disable", host, port, dbname)

	dbInstance, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		log.Print(err)
	}

	DBInstance = dbInstance
}
