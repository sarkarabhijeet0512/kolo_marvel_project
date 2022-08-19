package persistence

import (
	"kolo_marvel_project/utils"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

type NewRepositoryIn struct {
	fx.In

	Log  *logrus.Logger
	Pool *redis.Pool `name:"redisWorker"`
}

// RedisStore represents the cache with redis persistence
type RedisStore struct {
	pool              *redis.Pool
	defaultExpiration time.Duration
	log               *logrus.Logger
}

func NewRedisRepository(i NewRepositoryIn) (Repo CacheStore, err error) {

	Repo = &RedisStore{
		log:               i.Log,
		pool:              i.Pool,
		defaultExpiration: time.Hour * 24 * 7,
	}

	return
}

// NewRedisCacheWithPool returns a RedisStore using the provided pool
// until redigo supports sharding/clustering, only one host will be in hostList
func NewRedisCacheWithPool(pool *redis.Pool, log *logrus.Logger, defaultExpiration time.Duration) *RedisStore {
	return &RedisStore{pool, defaultExpiration, log}
}

// Set (see CacheStore interface)
func (c *RedisStore) Set(key string, value interface{}, expires time.Duration) error {
	conn := c.pool.Get()
	defer conn.Close()
	return c.invoke(conn.Do, key, value, expires)
}

// Add (see CacheStore interface)
func (c *RedisStore) Add(key string, value interface{}, expires time.Duration) error {
	conn := c.pool.Get()
	defer conn.Close()
	if exists(conn, key) {
		return ErrNotStored
	}
	return c.invoke(conn.Do, key, value, expires)
}

// Replace (see CacheStore interface)
func (c *RedisStore) Replace(key string, value interface{}, expires time.Duration) error {
	conn := c.pool.Get()
	defer conn.Close()
	if !exists(conn, key) {
		return ErrNotStored
	}
	err := c.invoke(conn.Do, key, value, expires)
	if value == nil {
		return ErrNotStored
	}

	return err

}

// Get (see CacheStore interface)
func (c *RedisStore) Get(key string, ptrValue interface{}) error {
	conn := c.pool.Get()
	defer conn.Close()
	raw, err := conn.Do("GET", key)
	if raw == nil {
		return ErrCacheMiss
	}
	item, err := redis.Bytes(raw, err)
	if err != nil {
		return err
	}
	return utils.Deserialize(item, ptrValue)
}

func exists(conn redis.Conn, key string) bool {
	retval, _ := redis.Bool(conn.Do("EXISTS", key))
	return retval
}

// Delete (see CacheStore interface)
func (c *RedisStore) Delete(key string) error {
	conn := c.pool.Get()
	defer conn.Close()
	if !exists(conn, key) {
		return ErrCacheMiss
	}
	_, err := conn.Do("DEL", key)
	return err
}

// Increment (see CacheStore interface)
func (c *RedisStore) Increment(key string, delta uint64) (uint64, error) {
	conn := c.pool.Get()
	defer conn.Close()
	// Check for existance *before* increment as per the cache contract.
	// redis will auto create the key, and we don't want that. Since we need to do increment
	// ourselves instead of natively via INCRBY (redis doesn't support wrapping), we get the value
	// and do the exists check this way to minimize calls to Redis
	val, err := conn.Do("GET", key)
	if val == nil {
		return 0, ErrCacheMiss
	}
	if err == nil {
		currentVal, err := redis.Int64(val, nil)
		if err != nil {
			return 0, err
		}
		sum := currentVal + int64(delta)
		_, err = conn.Do("SET", key, sum)
		if err != nil {
			return 0, err
		}
		return uint64(sum), nil
	}

	return 0, err
}

// Decrement (see CacheStore interface)
func (c *RedisStore) Decrement(key string, delta uint64) (newValue uint64, err error) {
	conn := c.pool.Get()
	defer conn.Close()
	// Check for existance *before* increment as per the cache contract.
	// redis will auto create the key, and we don't want that, hence the exists call
	if !exists(conn, key) {
		return 0, ErrCacheMiss
	}
	// Decrement contract says you can only go to 0
	// so we go fetch the value and if the delta is greater than the amount,
	// 0 out the value
	currentVal, err := redis.Int64(conn.Do("GET", key))
	if err == nil && delta > uint64(currentVal) {
		tempint, err := redis.Int64(conn.Do("DECRBY", key, currentVal))
		return uint64(tempint), err
	}
	tempint, err := redis.Int64(conn.Do("DECRBY", key, delta))
	return uint64(tempint), err
}

// Flush (see CacheStore interface)
func (c *RedisStore) Flush() error {
	conn := c.pool.Get()
	defer conn.Close()
	_, err := conn.Do("FLUSHALL")
	return err
}

func (c *RedisStore) invoke(f func(string, ...interface{}) (interface{}, error),
	key string, value interface{}, expires time.Duration) error {

	switch expires {
	case DEFAULT:
		expires = c.defaultExpiration
	case FOREVER:
		expires = time.Duration(0)
	}

	b, err := utils.Serialize(value)
	if err != nil {
		return err
	}

	if expires > 0 {
		_, err := f("SETEX", key, int32(expires/time.Second), b)
		return err
	}

	_, err = f("SET", key, b)
	return err

}
