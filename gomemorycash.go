package gomemorycash

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

// NewCash create new cash storage
func NewCash() CashMethods {
	return storage{make(map[string]interface{})}
}
