package config

import (
	"github.com/redis/go-redis/v9"
	"os"
)

type Configs struct {
	Jwt struct {
		Key string
		Exp int
	}
	RedisClient *redis.Client
}

func Load() *Configs {
	conf := &Configs{}
	conf.Jwt.Key = os.Getenv("JWT_KEY")
	conf.Jwt.Exp = 60 // expired dalam menit

	conf.RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // ubah jika perlu
		Password: "",               // set jika ada password redis
		DB:       0,
	})

	return conf
}
