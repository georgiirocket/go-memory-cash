package memory_cash

type Storage struct {
	db map[string]interface{}
}

type Methods interface {
	Set(key string, value interface{})
	Get(key string) interface{}
	Delete(key string)
}

func (cash Storage) Set(key string, value interface{}) {
	cash.db[key] = value
}

func (cash Storage) Get(key string) interface{} {
	return cash.db[key]
}

func (cash Storage) Delete(key string) {
	delete(cash.db, key)
}

// New create new cash storage
func New() Methods {
	return Storage{make(map[string]interface{})}
}
