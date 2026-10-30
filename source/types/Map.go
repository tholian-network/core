package types

import "sort"
import "sync"

type Map[Type any] struct {
	data  map[string]Type
	mutex *sync.RWMutex
}

func NewMap[Type any]() *Map[Type] {

	return &Map[Type]{
		data:  make(map[string]Type),
		mutex: &sync.RWMutex{},
	}

}

func (self *Map[Type]) Contains(key string) bool {

	var result bool

	self.mutex.RLock()

	_, ok := self.data[key]

	if ok == true {
		result = true
	}

	self.mutex.RUnlock()

	return result

}

func (self *Map[Type]) Get(key string) (Type, bool) {

	self.mutex.RLock()

	value, ok := self.data[key]

	self.mutex.RUnlock()

	return value, ok

}

func (self *Map[Type]) GetAll() []Type {

	self.mutex.RLock()

	result := make([]Type, 0)

	for _, value := range self.data {
		result = append(result, value)
	}

	self.mutex.RUnlock()

	return result

}

func (self *Map[Type]) Keys() []string {

	self.mutex.RLock()

	result := make([]string, 0)

	for key, _ := range self.data {
		result = append(result, key)
	}

	sort.Strings(result)

	self.mutex.RUnlock()

	return result

}

func (self *Map[Type]) Remove(key string) bool {

	var result bool

	self.mutex.Lock()

	_, ok := self.data[key]

	if ok == true {
		delete(self.data, key)
		result = true
	}

	self.mutex.Unlock()

	return result

}

func (self *Map[Type]) Set(key string, value Type) {

	self.mutex.Lock()

	self.data[key] = value

	self.mutex.Unlock()

}

func (self *Map[Type]) Length() int {

	var result int

	self.mutex.RLock()

	result = len(self.data)

	self.mutex.RUnlock()

	return result

}

