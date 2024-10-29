package infrastructure

type DataDriver struct {
}

func NewSessionDriver() *DataDriver {
	return &DataDriver{}
}

func (driver *DataDriver) Set(key string, value interface{}) error {
	return nil
}

func (driver *DataDriver) Get(key string) (interface{}, error) {
	return nil, nil
}

func (driver *DataDriver) Delete(key string) error {
	return nil
}
