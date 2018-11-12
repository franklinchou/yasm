package services

import (
	"time"
	"math/rand"
	"github.com/go-redis/redis"
	"../utils"
	"fmt"
)

//*********************************************************
// Application defaults
//*********************************************************
const SessionDefaultTimeout = 120
const RedisHost = "localhost"
const RedisPort = 6379
const RedisPassword = ""

//*********************************************************

type RedisClient interface {
	Connect() (*redis.Client)
}

func Connect() *redis.Client {
	RedisAddress := fmt.Sprintf("%s:%d", RedisHost, RedisPort)
	redisOptions := redis.Options{
		Addr:     RedisAddress,
		Password: RedisPassword,
		DB:       0,
	}
	return redis.NewClient(&redisOptions)
}

func createSessionId() string {
	source := rand.NewSource(time.Now().UnixNano())
	return utils.RandomString(32, source)
}

func CreateSession(token string) string {
	expiration := time.Duration(SessionDefaultTimeout)
	sessionId := createSessionId()
	Connect().Set(sessionId, token, expiration)
	return sessionId
}
