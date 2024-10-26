package memoryCash

type storage struct {
	db map[string]interface{}
}

type CashMethods interface {
	Set(key string, value interface{})
	Get(key string) interface{}
	Delete(key string)
}

func (cash storage) Set(key string, value interface{}) {
	cash.db[key] = value
}

func (cash storage) Get(key string) interface{} {
	return cash.db[key]
}

func (cash storage) Delete(key string) {
	delete(cash.db, key)
}

// New create new cash storage
func New() CashMethods {
	return storage{make(map[string]interface{})}
}
