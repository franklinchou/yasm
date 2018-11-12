package services

import (
	"time"
	"math/rand"
	"github.com/go-redis/redis"
	"../utils"
)

//*********************************************************
// Application defaults
//*********************************************************
const SessionDefaultTimeout = 120
const RedisHost = "http://localhost:8000"
const RedisPassword = ""


//*********************************************************

type RedisClient interface {
	Connect() (*redis.Client)
}

func connect() *redis.Client {
	redisOptions := redis.Options{
		Addr:     RedisHost,
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
	connect().Set(sessionId, token, expiration)
	return sessionId
}
