package main

import (
	"context"
	"github.com/my/repo/internal/config"
	"github.com/my/repo/internal/governor"
	"github.com/my/repo/internal/redis"
	"github.com/my/repo/internal/rest"
	"log/slog"
	"os"
	"os/signal"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	conf := config.NewConfig()
	newredis.NewRedis(conf)
	gov := governor.NewGovernor()
	r := rest.NewRest(conf, gov)
	go func() {
		if err := r.Start(ctx); err != nil {
			slog.Error("failed start rest server", err)
		}
	}()
	go func(cancelFunc context.CancelFunc) {
		shutdown := make(chan os.Signal, 1)
		signal.Notify(shutdown, os.Interrupt)
		<-shutdown
		cancelFunc()
	}(cancel)
	<-ctx.Done()
}

//// Получаем адрес Redis из переменной окружения
//redisAddr := os.Getenv("REDIS_ADDR")
//if redisAddr == "" {
//redisAddr = "localhost:6379"
//}
//
//// Подключаемся к Redis
//rdb := redis.NewClient(&redis.Options{
//Addr:     redisAddr, // Адрес Redis
//Password: "",        // Без пароля
//DB:       0,
//})
//
//// Проверяем соединение
//
//_, err := rdb.Ping(ctx).Result()
//if err != nil {
//log.Fatalf("Ошибка подключения к Redis: %v", err)
//}
//
//fmt.Println("Redis подключен!")
//
//err = rdb.ZAdd(ctx, "leader",
//redis.Z{Score: 150, Member: "Eraga"},
//redis.Z{Score: 100, Member: "SHeraga"},
//redis.Z{Score: 500, Member: "Araga"},
//redis.Z{Score: 50, Member: "ZHanaga"},
//redis.Z{Score: 1150, Member: "Inatay"},
//redis.Z{Score: 950, Member: "Samat"},
//).Err()
//if err != nil {
//log.Fatal(err)
//}
////ranks, err := rdb.ZRevRank(ctx, "leader", "Eraga").Result()
////if err != nil {
////	fmt.Println("лощибка ")
////	log.Fatal(err)
////}
//topPlayers, err := rdb.ZRevRangeWithScores(ctx, "leader", 0, 9).Result()
//if err != nil {
//log.Fatal(err)
//}
//for _, player := range topPlayers {
//fmt.Printf("oynwy: %v, upai: %v\n", player.Member, player.Score)
//
//}
//err = rdb.ZAdd(ctx, "leader",
//redis.Z{Score: 15000, Member: "Eraga"},
//redis.Z{Score: 10000, Member: "SHeraga"},
//redis.Z{Score: 500, Member: "Araga"},
//redis.Z{Score: 5000, Member: "ZHanaga"},
//redis.Z{Score: 1150, Member: "Inatay"},
//redis.Z{Score: 950, Member: "Samat"},
//).Err()
//if err != nil {
//log.Fatal(err)
//}
//topPlayers, err = rdb.ZRevRangeWithScores(ctx, "leader", 0, 9).Result()
//if err != nil {
//log.Fatal(err)
//}
//fmt.Println()
//for _, player := range topPlayers {
//
//fmt.Printf("oynwy: %v, upai: %v\n", player.Member, player.Score)
//
//}
//fmt.Println("Прочитанное значение из Redis:", ranks)
