package redis

type RedisOptions struct {
	Host     string `json:"host" yaml:"host" description:"Redis service host address"`
	Port     string `json:"port" yaml:"port" description:"Redis service port address"`
	Username string `json:"username" yaml:"username" description:"Redis service username"`
	Password string `json:"password" yaml:"password" description:"Redis service password"`
}

func NewRedisOptions() *RedisOptions {
	return &RedisOptions{
		Host:     "127.0.0.1",
		Port:     "6379",
		Username: "",
		Password: "",
	}
}
