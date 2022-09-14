package proc

import (
	"sync"
)

const (
	OMZChannel = "OMZChannel"
)

var chanContext *ChanContext

var once sync.Once

func GetChanInstance() *ChanContext {
	once.Do(func() {
		chanContext = &ChanContext{}
		chanContext.data = make(map[string]interface{})
	})

	return chanContext
}

type ChanContext struct {
	data map[string]interface{}
}

func (o *ChanContext) Put(key string, value interface{}) {
	o.data[key] = value
}

func (o *ChanContext) Get(key string) (interface{}, bool) {
	val, exists := o.data[key]
	return val, exists
}
