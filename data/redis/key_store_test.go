package redis_test

import (
	"testing"
	"time"

	goredis "github.com/go-redis/redis"
	"github.com/keratin/authn-server/data/redis"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestKeyStore(t *testing.T) {
	client := goredis.NewClient(&goredis.Options{
		Addr: "127.0.0.1:6379",
		DB:   12,
	})
	secret := []byte("32bigbytesofsuperultimatesecrecy")

	t.Run("empty remote storage", func(t *testing.T) {
		client.FlushDB()
		store, err := redis.NewKeyStore(client, time.Hour, time.Second, secret)
		require.NoError(t, err)

		assert.NotEmpty(t, store.Keys())
		assert.Len(t, store.Keys(), 1)
		assert.Equal(t, store.Key(), store.Keys()[0])
	})

	t.Run("multiple servers", func(t *testing.T) {
		client.FlushDB()
		store1, err := redis.NewKeyStore(client, time.Hour, time.Second, secret)
		require.NoError(t, err)
		key1 := store1.Key()
		assert.NotEmpty(t, key1)

		store2, err := redis.NewKeyStore(client, time.Hour, time.Second, secret)
		require.NoError(t, err)
		assert.Equal(t, key1, store2.Key())
		assert.Len(t, store2.Keys(), 1)
		assert.Equal(t, key1, store2.Keys()[0])
	})
}