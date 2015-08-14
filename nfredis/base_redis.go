package nfredis

import (
	"jfcsrv/nfconst"
	"time"

	"github.com/garyburd/redigo/redis"
)

var (
	pool *redis.Pool
)

func getConn() (c redis.Conn) {
	if pool != nil {
		return pool.Get()
	} else {
		connstr := nfconst.JCfg.RedisServer + ":" + nfconst.JCfg.RedisPort
		dbIndex := nfconst.JCfg.RedisDbIndex
		pool = &redis.Pool{
			// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
			MaxIdle:     1,
			MaxActive:   10,
			IdleTimeout: 180 * time.Second,
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", connstr)
				if err != nil {
					return nil, err
				}
				// 选择db
				c.Do("SELECT", dbIndex)
				return c, nil
			},
		}
	}
	return pool.Get()
}

func Exists(key interface{}) (ret int, err error) {
	c := getConn()
	reply, err1 := c.Do("EXISTS", key)
	defer c.Close()
	ii, err2 := redis.Int(reply, err1)
	return ii, err2
}

func Hset(key interface{}, field interface{}, value interface{}) (ret int, err error) {
	c := getConn()
	reply, err1 := c.Do("HSET", key, field, value)
	defer c.Close()
	ii, err2 := redis.Int(reply, err1)
	return ii, err2
}

func Hget(key interface{}, field interface{}) (ret string, err error) {
	c := getConn()
	reply, err1 := c.Do("HGET", key, field)
	defer c.Close()
	ss, err2 := redis.String(reply, err1)
	return ss, err2
}

func Hmset(args ...interface{}) (ret string, err error) {
	c := getConn()
	reply, err1 := c.Do("HMSET", args...)
	defer c.Close()
	ss, err2 := redis.String(reply, err1)
	return ss, err2
}

func Hgetall(key interface{}) (ret []interface{}, err error) {
	c := getConn()
	reply, err1 := c.Do("HGETALL", key)
	values, err2 := redis.Values(reply, err1)
	return values, err2
}
