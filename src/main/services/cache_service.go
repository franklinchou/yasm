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
// Models
//*********************************************************

type stringTuple struct {
	k, v string
}

//*********************************************************

func createSessionId() string {
	source := rand.NewSource(time.Now().UnixNano())
	return utils.RandomString(32, source)
}

func zip(k, v []string) ([]stringTuple, error) {
	if len(k) != len(v) {
		return nil, fmt.Errorf("zip: cannot zip data sets of different length")
	}
	r := make([]stringTuple, len(k), len(v))
	for i, key := range k {
		r[i] = stringTuple{key, v[i]}
	}
	return r, nil
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

func CreateSession(token string) (string, string) {
	cxn := MyPool.Get()
	expiration := time.Duration(SessionDefaultTimeout)
	sessionId := createSessionId()
	cxn.Do("SETEX", token, expiration, sessionId)
	return token, sessionId // TODO Return the actual result from Do
}

func GetSessions(limit int) ([]stringTuple, error) {
	cxn := MyPool.Get()
	iter := 0
	keys := make([]string, 0)
	values := make([]string, 0)
	for {
		if arr, err := redis.Values(cxn.Do("SCAN", iter)); err != nil {
			panic(err)
		} else {
			iter, _ = redis.Int(arr[0], nil)
			k, _ := redis.String(arr[1], nil)
			v, _ := redis.String(cxn.Do("MGET", arr[1], nil))
			keys = append(keys, k)
			values = append(values, v)
		}

		if iter == 0 {
			break
		}
	}
	return zip(keys, values)
}


func GetTokenBySession(sessionId string) (string, error) {
	c := MyPool.Get()
	token, e := redis.String(c.Do("MGET", sessionId))
	if e != nil {
		return "", fmt.Errorf("GetTokenBySession: could not find token for session %s", sessionId)
	}

	expiration := time.Duration(SessionDefaultTimeout)
	c.Do("SETEX", token, expiration, sessionId)
	return token, nil
}


func DeleteSession(sessionId string) error {
	cxn := MyPool.Get()
	r, err := redis.Int(cxn.Do("DEL", sessionId))
	if err != nil {
		return err
	}
	if r == 0 {
		return fmt.Errorf("DeleteSession: could not delete session %s", sessionId)
	}
	return err
}