package driver

import (
	"github.com/go-redis/redis"
)

// DB ...
type DB struct {
	SQL *redis.Client
}

// DBConn ...
var dbConn = &DB{}

// ConnectStorage ...
func ConnectStorage(addr, pass string, db int) (*DB, error) {

	dbx := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       db,
	})
	_, err := dbx.Ping().Result()

	dbConn.SQL = dbx
	return dbConn, err

}
