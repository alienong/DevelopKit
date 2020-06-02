/**
 * @Author: alienongwlx@gmail.com
 * @Description: Gin User Redis Session
 * @Version: 1.0.0
 * @Date: 2020/4/24 14:09
 */

package session

import (
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"time"
)

type RedisSessionService struct {
	store *redis.Store
}

func (s *RedisSessionService) NewRedisSessionService(redisAdr, key string) (*RedisSessionService, error) {
	store, err := redis.NewStore(10, "tcp", redisAdr, "", []byte(key))
	if err != nil {
		return nil, err
	}
	store.Options(sessions.Options{
		Path:   "/",
		MaxAge: int(30 * time.Minute),
	})
	return &RedisSessionService{&store}, nil
}

func (s *RedisSessionService) UseRedisSession(engine *gin.Engine) {
	engine.Use(sessions.Sessions("s", *s.store))
}

func (s *RedisSessionService) Set(key, value interface{}, c *gin.Context) error {
	session := sessions.Default(c)
	session.Set(key, value)
	err := session.Save()
	return err
}

func (s *RedisSessionService) Get(key interface{}, c *gin.Context) (interface{}, error) {
	session := sessions.Default(c)
	if session.Get(key) == nil {
		return nil, errors.New("Session writing failed")
	}
	value := session.Get(key)
	session.Save()
	return value, nil
}
