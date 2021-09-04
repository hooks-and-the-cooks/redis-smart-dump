package main

import (
	"github.com/redis-smart-dump/redis"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
)

func dumpKeys(host string, port string, logger *logrus.Logger) {
	err := LoadConfig()
	if err != nil {
		logger.WithFields(logrus.Fields{"error": err}).Fatal("Error while loading config file!")
	}

	redisClient := redis.InitRedisClient(host, port)
	RESPOutputFile, err := os.OpenFile(viper.GetString("log.respOutputFile"), os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		logger.Fatal("Not able to open output dump file!")
	}
	fileLogger := log.New(RESPOutputFile, "", 0)

	keys, err := redisClient.Keys(ctx, "*").Result()
	if err != nil {
		logger.WithFields(logrus.Fields{"error": err}).Fatal("Error while selecting all keys! exiting!")
	}

	logger.WithFields(
		logrus.Fields{"info": redis.GetKeyspaceInfo(redisClient, ctx), "count of keys": len(keys)}).
		Info("Info from Redis")

	for i, key := range keys {
		logger.WithFields(logrus.Fields{"iteration": i, "key": key}).
			Info("Dumping key from Redis")
		err := redis.DumpKeys(redisClient, key, ctx, fileLogger)
		if err != nil {
			logger.WithFields(logrus.Fields{"error": err, "key": key}).
				Fatal("Error while converting key to RESP! exiting!")
		} else {
			logger.WithFields(logrus.Fields{"iteration": i, "key": key}).
				Info("Successful Dumping key from Redis")
		}
	}
}
