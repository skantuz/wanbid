package models

import (
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/skantuz/godotenv"
)

type Paginator struct {
	TotalRows  int         `json:"total_rows"`
	TotalPages int         `json:"total_pages"`
	Data       interface{} `json:"data"`
	Offset     int         `json:"offset"`
	Limit      int         `json:"limit"`
	Page       int         `json:"tpage"`
	PrevPage   int         `json:"prev_page"`
	NextPage   int         `json:"next_page"`
}

var db *gorm.DB

func init() {
	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}
	//carga configuracion base de datos .env
	dbUser := os.Getenv("db_user")
	dbPass := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")
	dbType := os.Getenv("db_type")
	dbUri := ""

	//seleccion de base de datos
	switch dbType {
	case "psql":
		dbUri = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", dbHost, dbPort, dbUser, dbName, dbPass) //cadena de texto de conexion a base de datos Postgres
		break
	case "mysql":
		if dbHost == "localhost" {
			dbUri = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbName) //cadena de texto de conexion a base de datos Mysql
		} else {
			dbUri = fmt.Sprintf("%s:%s@%s:%s/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName) //cadena de texto de conexion a base de datos Mysql
		}
		break
	case "mssql":
		dbUri = fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", dbUser, dbPass, dbHost, dbPort, dbName) //cadena de texto de conexion a base de datos Microsoft SQL
		break
	case "sqlite3":
		dbUri = fmt.Sprintf("%s", dbName) //cadena de texto de conexion a base de datos Sqlite3
		break
	}

	conn, err := gorm.Open(dbType, dbUri)
	if err != nil {
		fmt.Print(err)
	}

	conn.Debug().AutoMigrate()

}
func GetDB() *gorm.DB {
	return db
}

func Paging(result interface{}, page, rowsPage int, orderBy ...string) *Paginator {
	paginator := &Paginator{}
	//verificamos si hay limite de registros por pagina si no traemos el default
	if rowsPage == 0 {
		dbpagination, _ := strconv.Atoi(os.Getenv("db_pagination"))
		rowsPage = dbpagination
	}
	//verificamos primera pag
	if page < 1 {
		page = 1
	}
	dbdebug, _ := strconv.ParseBool(os.Getenv("db_debug"))
	if dbdebug {
		db.Debug()
	}

	if len(orderBy) > 0 {
		for _, o := range orderBy {
			db.Order(o)
		}
	}
	done := make(chan bool, 1)
	var count int
	var offset int
	go countRecords(db, result, done, &count)
	<-done
	paginator.TotalRows = count
	paginator.Data = result
	paginator.Page = page
	paginator.Offset = offset
	paginator.Limit = rowsPage
	paginator.TotalPages = int(math.Ceil(float64(count) / float64(rowsPage)))

	if page > 1 {
		paginator.PrevPage = page - 1
	} else {
		paginator.PrevPage = page
	}

	if page == paginator.TotalPages {
		paginator.NextPage = page
	} else {
		paginator.NextPage = page + 1
	}
	return paginator
}

func countRecords(db *gorm.DB, anyType interface{}, done chan bool, count *int) {
	db.Model(anyType).Count(count)
	done <- true
}
