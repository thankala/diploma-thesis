package stores

import (
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/redis/go-redis/v9"
	"github.com/thankala/diploma-thesis/common/components"
	"golang.org/x/net/context"
	"strconv"
)

var (
	defaultStore    = "redis"
	defaultAddr     = "localhost:6379"
	defaultPassword = "" // no password set
	defaultDb       = 0  // use default DB
)

type RedisOptFunc func(opts *redisOpts)

type redisOpts struct {
	store    string
	addr     string
	password string
	db       int
}

type RedisStore struct {
	client *redis.Client
	locker *redsync.Redsync
}

func NewRedisStore(opts ...RedisOptFunc) *RedisStore {
	options := defaultOpts()
	for _, opt := range opts {
		opt(options)
	}
	client := redis.NewClient(&redis.Options{
		Addr:     options.addr,
		Password: options.password,
		DB:       options.db,
	})
	pool := goredis.NewPool(client)
	return &RedisStore{
		client: client,
		locker: redsync.New(pool),
	}
}

func defaultOpts() *redisOpts {
	return &redisOpts{
		defaultStore,
		defaultAddr,
		defaultPassword,
		defaultDb,
	}
}

func WithStore(store string) RedisOptFunc {
	if store == "" {
		return func(opts *redisOpts) {}
	} else {
		return func(opts *redisOpts) {
			opts.store = store
		}
	}
}

func WithStoreAddr(storeAddr string) RedisOptFunc {
	if storeAddr == "" {
		return func(opts *redisOpts) {}
	} else {
		return func(opts *redisOpts) {
			opts.addr = storeAddr
		}
	}
}

func WithStorePassword(storePassword string) RedisOptFunc {
	if storePassword == "" {
		return func(opts *redisOpts) {}
	} else {
		return func(opts *redisOpts) {
			opts.password = storePassword
		}
	}
}

func WithStoreDb(storeDb string) RedisOptFunc {
	if storeDb == "" {
		return func(opts *redisOpts) {}
	} else {
		return func(opts *redisOpts) {
			db, err := strconv.Atoi(storeDb)
			if err != nil {
				return
			}
			opts.db = db
		}
	}
}

func (r *RedisStore) Store(key string, state []byte) error {
	return r.client.Set(context.TODO(), key, state, 0).Err()
}

func (r *RedisStore) Load(key string) ([]byte, error) {
	val, err := r.client.Get(context.TODO(), key).Result()
	return []byte(val), err
}

func (r *RedisStore) AcquireLock(key string) components.Mutex {
	return r.locker.NewMutex(key)
}

func (r *RedisStore) ReleaseLock(mutex components.Mutex) (bool, error) {
	return mutex.Unlock()
}
