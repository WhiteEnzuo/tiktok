package Redis

import "github.com/garyburd/redigo/redis"

type Redis struct {
	Host   string
	Port   string
	Server redis.Conn
}

func (r *Redis) getConnect() error {
	conn, err := redis.Dial("tcp", r.Host+":"+r.Port)
	if err != nil {
		return err
	}
	r.Server = conn
	return nil
}
func NewRedis(Host, Port string) *Redis {
	r := new(Redis)
	r.Host = Host
	r.Port = Port
	err := r.getConnect()
	if err != nil {
		return nil
	}
	return r
}
