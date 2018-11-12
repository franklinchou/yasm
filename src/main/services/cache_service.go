package services

import (
	"time"
	"math/rand"
	"github.com/gomodule/redigo/redis"
	"../utils"
)

//*********************************************************
// Application defaults
//*********************************************************
const SessionDefaultTimeout = 120
const RedisPort = ":6379"

//*********************************************************

func NewPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", RedisPort)
			return c, err
		},
	}
}

func createSessionId() string {
	source := rand.NewSource(time.Now().UnixNano())
	return utils.RandomString(32, source)
}

func Ping(c redis.Conn) error {
	_, e := c.Do("PING")
	if e != nil {
		return e
	}
	return nil
}

func CreateSession(token string) string {
	expiration := time.Duration(SessionDefaultTimeout)
	sessionId := createSessionId()
	cxn := MyPool.Get()
	cxn.Do("SETEX", token, expiration, sessionId)
	return sessionId
}
