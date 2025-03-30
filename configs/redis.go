package configs

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)
var (
	client *redis.Client
	oncer   sync.Once
)

func RedisConnect() {
	oncer.Do(func() {
		client = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			ClientName: "",
			Password: "",
			DB:       0,
		})

		ctx := context.Background()
		_,err := client.Ping(ctx).Result()
		if err != nil {
			log.Fatal("Gagal Connect redis")
		}
		
		
		log.Println("Berhasil connect redis")
	})
}

func KeepToRedis(id,token, refresh string) {
	// rdb := RedisConnect()
	ctx := context.Background()
	key := "JwtToken: "+id
	refreshkey := "RefreshToken: "+id

	if err := client.Set(ctx, key, token, 15*time.Minute).Err(); err != nil {
		log.Fatalf("Gagal Menyimpan access token: %v", err)
	}

	if err := client.Set(ctx, refreshkey , refresh, 7*time.Hour).Err(); err != nil {
		log.Fatalf("gagal Menyimpan Refresh Token: %v", err)
	}

	log.Println("Token Berhasil di simpan")
}


func Logout(id string) (string, error) {
	ctx := context.Background()
	key := "JwtToken: "+id
	refreshkey := "RefreshToken: "+id

	if err := client.Del(ctx, key,refreshkey).Err(); err != nil {
		return "", err
	}

	return "berhasil hapus token", nil
}