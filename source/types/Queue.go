package types

import "sync"

type Queue[Type comparable] struct {
	data  []Type
	mutex *sync.RWMutex
}

func NewQueue[Type comparable]() *Queue[Type] {

	return &Queue[Type]{
		data:  make([]Type, 0),
		mutex: &sync.RWMutex{},
	}

}

func (queue *Queue[Type]) Add(value Type) bool {

	var result bool

	queue.mutex.Lock()

	found := false

	for _, other := range queue.data {

		if other == value {
			found = true
			break
		}

	}

	if found == false {
		queue.data = append(queue.data, value)
	}

	queue.mutex.Unlock()

	return result

}

func (queue *Queue[Type]) Contains(value Type) bool {

	var result bool

	queue.mutex.RLock()

	for _, other := range queue.data {

		if other == value {
			result = true
			break
		}

	}

	queue.mutex.RUnlock()

	return result

}

func (queue *Queue[Type]) Pop() (Type, bool) {

	queue.mutex.Lock()
	defer queue.mutex.Unlock()

	var zero Type

	if len(queue.data) > 0 {

		index := len(queue.data) - 1
		value := queue.data[index]
		queue.data = queue.data[0:index]

		return value, true

	} else {
		return zero, false
	}

}

func (queue *Queue[Type]) Shift() (Type, bool) {

	queue.mutex.Lock()
	defer queue.mutex.Unlock()

	var zero Type

	if len(queue.data) > 0 {

		value := queue.data[0]
		queue.data = queue.data[1:]

		return value, true

	} else {
		return zero, false
	}

}

func (queue *Queue[Type]) Remove(value Type) bool {

	var result bool

	queue.mutex.Lock()

	for index, other := range queue.data {

		if other == value {
			queue.data = append(queue.data[0:index], queue.data[index+1:]...)
			result = true
		}

	}

	queue.mutex.Unlock()

	return result

}

func (queue *Queue[Type]) Length() int {

	var result int

	queue.mutex.RLock()

	result = len(queue.data)

	queue.mutex.RUnlock()

	return result

}

