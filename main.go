package main

import (
	"context"
	"log"

	"github.com/AkifhanIlgaz/key-aero-api/cfg"
)

func main() {
	config, err := cfg.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not read environment variables", err)
	}

	ctx := context.TODO()

	mongoClient, err := connectToMongo(ctx, config.MongoUri)
	if err != nil {
		log.Fatal("Could not connect to mongo", err)
	}
	defer mongoClient.Disconnect(ctx)

	redisClient, err := connectToRedis(ctx, config.RedisUri)
	if err != nil {
		log.Fatal("Could not connect to redis", err)
	}
	defer redisClient.Close()
}

