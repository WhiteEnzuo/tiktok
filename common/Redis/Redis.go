package Redis

import "github.com/garyburd/redigo/redis"

type Redis struct {
	Host   string
	Port   string
	DB     string
	Server redis.Conn
}

func (r *Redis) getConnect() error {
	conn, err := redis.Dial("tcp", r.Host+":"+r.Port)
	if err != nil {
		return err
	}
	r.Server = conn
	_, err = conn.Do("select", r.DB)
	if err != nil {
		return err
	}
	return nil
}
func NewRedis(Host, Port, DB string) *Redis {
	r := new(Redis)
	r.Host = Host
	r.Port = Port
	r.DB = DB
	err := r.getConnect()
	if err != nil {
		return nil
	}
	return r
}

func (r *Redis) Set(key, val string) (do interface{}, err error) {
	return r.Server.Do("set", key, val)
}
func (r *Redis) Get(key string) (do interface{}, err error) {
	reply, err := r.Server.Do("get", key)
	if err != nil {
		return nil, err
	}
	return string(reply.([]byte)), nil
}
func (r *Redis) Keys(key string) (do interface{}, err error) {
	reply, err := r.Server.Do("keys", key)
	if err != nil {
		return nil, err

	}
	item := reply.([]interface{})
	temp := make([]string, len(item))
	for i, bytes := range item {
		temp[i] = string(bytes.([]byte))
	}
	return temp, nil
}
