package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"strconv"
)

func stringToRedisCmd(k, val string) []string {
	return []string{"SET", k, val}
}

func hashToRedisCmd(k string, val map[string]string) []string {
	cmd := []string{"HSET", k}
	for k, v := range val {
		cmd = append(cmd, k, v)
	}
	return cmd
}

func setToRedisCmd(k string, val []string) []string {
	cmd := []string{"SADD", k}
	return append(cmd, val...)
}

func listToRedisCmd(k string, val []string) []string {
	cmd := []string{"RPUSH", k}
	return append(cmd, val...)
}

func zSetToRedisCmd(k string, val []redis.Z) []string {
	cmd := []string{"ZADD", k}

	for _, v := range val {
		cmd = append(cmd, fmt.Sprintf("%v", v.Score), fmt.Sprintf("%v", v.Member))
	}
	return cmd
}

func RESPSerializer(cmd []string) string {
	s := ""
	s += "*" + strconv.Itoa(len(cmd)) + "\r\n"
	for _, arg := range cmd {
		s += "$" + strconv.Itoa(len(arg)) + "\r\n"
		s += arg + "\r\n"
	}
	return s
}

func DumpKeys(client *redis.Client, key string, ctx context.Context, logger *log.Logger) error {
	var redisCmd []string

	var keyType string

	typeCmd := client.Type(ctx, key)
	if typeCmd.Err() != nil {
		return typeCmd.Err()
	}

	keyType = typeCmd.Val()

	switch keyType {
	case "string":
		getCmd := client.Get(ctx, key)
		if getCmd.Err() != nil {
			return getCmd.Err()
		}
		redisCmd = stringToRedisCmd(key, getCmd.Val())

	case "list":
		lRangeCmd := client.LRange(ctx, key, 0, -1)
		if lRangeCmd.Err() != nil {
			return lRangeCmd.Err()
		}
		redisCmd = listToRedisCmd(key, lRangeCmd.Val())

	case "set":
		setCmd := client.SMembers(ctx, key)
		if setCmd.Err() != nil {
			return setCmd.Err()
		}
		redisCmd = setToRedisCmd(key, setCmd.Val())

	case "hash":
		hashCmd := client.HGetAll(ctx, key)
		if hashCmd.Err() != nil {
			return hashCmd.Err()
		}
		redisCmd = hashToRedisCmd(key, hashCmd.Val())

	case "zset":
		zRangeCmd := client.ZRangeWithScores(ctx, key, 0, -1)
		if zRangeCmd.Err() != nil {
			return zRangeCmd.Err()
		}
		redisCmd = zSetToRedisCmd(key, zRangeCmd.Val())

	default:
		return fmt.Errorf("Key %s is of unreconized type %s", key, keyType)
	}

	logger.Print(RESPSerializer(redisCmd))

	return nil
}

func GetKeyspaceInfo(client *redis.Client, ctx context.Context) *redis.StringCmd {
	return client.Info(ctx, "Keyspace")
}
