package services

import (
	"time"
	"fmt"
	"math/rand"
	"github.com/gomodule/redigo/redis"
	"../utils"
)

//*********************************************************
// Defaults
//*********************************************************
const SessionDefaultTimeout = 3600
const RedisPort = ":6379"

//*********************************************************

func createSessionId() string {
	source := rand.NewSource(time.Now().UnixNano())
	return utils.RandomString(32, source)
}

//*********************************************************
// Public functions
//*********************************************************

func NewPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", RedisPort)
		},
	}
}

func Ping(c redis.Conn) error {
	_, e := c.Do("PING")
	if e != nil {
		return e
	}
	return nil
}

// Create a new session associated with a given token
// Note that when a token already has an associated session, calling
// this function again will create a new session, replacing/invalidating
// the old session.
func CreateSession(token string) (string, string) {
	cxn := MyPool.Get()
	sessionId := createSessionId()
	cxn.Do("SETEX", token, SessionDefaultTimeout, sessionId)
	return token, sessionId // TODO Return the actual result from Do
}

func GetSessions(limit int) ([][]string, error) {
	cxn := MyPool.Get()
	keys, err := redis.Strings(cxn.Do("KEYS", "*"))
	result := make([][]string, 0)
	if err != nil {
		return result, fmt.Errorf("GetSessions: could not retrieve sessions")
	}
	for _, k := range keys {
		strtup2 := make([]string, 2)
		strtup2[0] = k
		strtup2[1], _ = redis.String(cxn.Do("GET", k))
		result = append(result, strtup2)
	}
	return result, nil
}

func GetSessionByToken(t string) (string, error) {
	cxn := MyPool.Get()
	sessionId, e := redis.String(cxn.Do("GET", t))

	if e != nil {
		return "", fmt.Errorf("GetTokenBySession: could not find session for token %s", t)
	}

	cxn.Do("SETEX", t, SessionDefaultTimeout, sessionId)
	return sessionId, nil
}

func DeleteSession(token string) error {
	cxn := MyPool.Get()
	r, err := redis.Int(cxn.Do("DEL", token))
	if err != nil {
		return err
	}
	if r == 0 {
		return fmt.Errorf("DeleteSession: could not delete session %s", token)
	}
	return err
}
