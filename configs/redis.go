package configs

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func KeepToRedis(token,refresh string)  {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	ctx := context.Background()

	if err := rdb.Set(ctx,"JwtToken",token, 15*time.Minute).Err();err != nil {
		log.Fatalf("Gagal Menyimpan access token: %v",err)
	}

	if err := rdb.Set(ctx,"RefreshToken",refresh, 7*time.Hour).Err();err != nil {
		log.Fatalf("gagal Menyimpan Refresh Token: %v", err)
	}

	log.Println("Token Berhasil di simpan")

	// tokenjwt, err := rdb.Get(ctx, "JwtToken").Result()
	// if err != nil {
	// 	log.Fatalf("Gagal Mengambil Token JWT: %v", err.Error())
	// }

	// log.Printf("Token JWT: %s", tokenjwt)
}
