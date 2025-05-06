package manager

import (
	"context"
	"database/sql"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/qlixes/vocagame/entity"
)

type Manager interface {
	Update(key string, value entity.Park)
}

type manager struct {
	db *sql.DB
}

var Mgr = newManager()
var ctx = context.Background()

func newManager() *manager {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := rdb.Ping(ctx).Result()

	if err != nil {
		log.Fatalf("Could not connect redis : %v \n", err)
	}

	log.Printf("Connected redis : %v \n", pong)

	return &manager{
		cache: rdb,
	}
}

func (m *manager) Update(key string, value string) {
	err := m.cache.Set(ctx, key, value, 0)

	if err != nil {
		log.Fatalf("Failed update key : %s \n", err)
	}
}
