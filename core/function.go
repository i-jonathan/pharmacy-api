package core

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db          *gorm.DB
	redisClient *redis.Client
)

func init() {
	// Initialize postgres database
	var (
		name = os.Getenv("POSTGRES_DB")
		pass = os.Getenv("POSTGRES_PASSWORD")
		user = os.Getenv("POSTGRES_USER")
		host = os.Getenv("database_host")
		port = os.Getenv("database_port")
		ssl  = os.Getenv("database_ssl")
	)

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		host, port, user, name, pass, ssl)
	log.Println(dsn)
	var err error

	// Open connection to database
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println(err)
	}

	// Initialize Connection to redis DB
	redisDB, _ := strconv.Atoi(os.Getenv("redis_database"))
	redisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("redis_address"),
		Password: os.Getenv("redis_pass"),
		DB:       redisDB,
	})
}

func RegisterModel(model interface{}) {
	log.Println("Say cheese")
	err := db.AutoMigrate(model)

	if err != nil {
		log.Println(err)
	}
}

func GetDB() *gorm.DB {
	return db
}

func GetRedisDB() *redis.Client {
	return redisClient
}

func Paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(r.URL.Query().Get("page"))
		if page == 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(r.URL.Query().Get("page_size"))
		switch {
		case pageSize > 50:
			pageSize = 50
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func ResponseData(count int, r *http.Request) (int, bool, bool) {
	prev, next := false, false
	var page, pageSize int
	if count != 0 {
		var err error
		if r.URL.Query().Get("page") == "" {
			page = 1
		} else {
			page, err = strconv.Atoi(r.URL.Query().Get("page"))
			if err != nil {
				log.Println(err)
			}
		}
		if r.URL.Query().Get("page_size") == "" {
			pageSize = 10
		} else {
			pageSize, err = strconv.Atoi(r.URL.Query().Get("page_size"))
			if pageSize > 50 {
				pageSize = 50
			}
			if err != nil {
				log.Println(err)
			}
		}

		if (page * pageSize) < count {
			next = true
		}

		if page > 1 {
			prev = true
		}
	}
	return page, prev, next
}
