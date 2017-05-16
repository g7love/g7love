package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"os"
	"fmt"
	"strconv"
)

const defaultMaxConnections = 20

var (
	connection *gorm.DB
)

func init() {
	connection = connect()
}

func GetDB() *gorm.DB {
	return connection
}

func connect() *gorm.DB {
	url := getDatabaseUrl()
	max := getMaxConnection()
	fmt.Println(url)
	conn, err := gorm.Open("mysql", "lzx:lzx@/socialnetworking?charset=utf8&loc=Local")
	if err != nil {
		panic(err)
	}

	conn.DB().SetMaxIdleConns(max / 5)
	conn.DB().SetMaxOpenConns(max)
	conn.SingularTable(true)
	return conn
}

func getDatabaseUrl() string {
	url := os.Getenv("DATABASE_URL")
	if url == "" {
		panic("Environment variable 'DATABASE_URL' not defined")
	}

	return url
}

func getMaxConnection() int {
	env := os.Getenv("DATABASE_MAX_CONNECTIONS")
	if env == "" {
		return defaultMaxConnections
	}

	max, _ := strconv.Atoi(env)
	return max
}
