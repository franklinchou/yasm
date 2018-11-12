package services

import (
	"time"
	"math/rand"
	"github.com/gomodule/redigo/redis"
	"../utils"
	"fmt"
)

//*********************************************************
// Application defaults
//*********************************************************
const SessionDefaultTimeout = 3600
const RedisPort = ":6379"

//*********************************************************

func createSessionId() string {
	source := rand.NewSource(time.Now().UnixNano())
	return utils.RandomString(32, source)
}

type stringTuple struct {
	k, v string
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
			c, err := redis.Dial("tcp", RedisPort)
			return c, err
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
