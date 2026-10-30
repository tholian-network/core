package types

import "sort"
import "sync"

type SliceMap[Type any] struct {
	data  map[string][]Type
	mutex *sync.RWMutex
}

func NewSliceMap[Type any]() *SliceMap[Type] {

	return &SliceMap[Type]{
		data:  make(map[string][]Type),
		mutex: &sync.RWMutex{},
	}

}

func (self *SliceMap[Type]) Contains(key string) bool {

	var result bool

	self.mutex.RLock()

	_, ok := self.data[key]

	if ok == true {
		result = true
	}

	self.mutex.RUnlock()

	return result

}

func (self *SliceMap[Type]) Add(key string, value Type) {

	self.mutex.Lock()

	_, ok := self.data[key]

	if ok == false {
		self.data[key] = make([]Type, 0)
	}

	self.data[key] = append(self.data[key], value)

	self.mutex.Unlock()

}

func (self *SliceMap[Type]) Get(key string) ([]Type, bool) {

	self.mutex.RLock()

	value, ok := self.data[key]

	self.mutex.RUnlock()

	return value, ok

}

func (self *SliceMap[Type]) GetAll() []Type {

	self.mutex.RLock()

	result := make([]Type, 0)

	for _, value := range self.data {
		result = append(result, value...)
	}

	self.mutex.RUnlock()

	return result

}

func (self *SliceMap[Type]) Keys() []string {

	self.mutex.RLock()

	result := make([]string, 0)

	for key, _ := range self.data {
		result = append(result, key)
	}

	sort.Strings(result)

	self.mutex.RUnlock()

	return result

}

func (self *SliceMap[Type]) Set(key string, value []Type) {

	self.mutex.Lock()

	self.data[key] = value

	self.mutex.Unlock()

}

func (self *SliceMap[Type]) Length() int {

	var result int

	self.mutex.RLock()

	result = len(self.data)

	self.mutex.RUnlock()

	return result

}

