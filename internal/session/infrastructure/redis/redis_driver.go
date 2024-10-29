package infrastructure

type RedisDriver struct {
}

func NewSessionDriver() *RedisDriver {
	return &RedisDriver{}
}

func (driver *RedisDriver) Set(key string, value interface{}) error {
	return nil
}

func (driver *RedisDriver) Get(key string) (interface{}, error) {
	return nil, nil
}

func (driver *RedisDriver) Delete(key string) error {
	return nil
}
