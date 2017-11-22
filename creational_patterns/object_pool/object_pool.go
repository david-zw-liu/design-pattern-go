package object_pool

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Generic type
type Object interface{}

// Defined need to override methods
type IObjectPool interface {
	Validate(Object) bool
	Expire(Object)
	Create() Object
}

// Base object pool
type ObjectPool struct {
	expirationTime time.Duration
	lockList       map[Object]time.Time
	unlockList     map[Object]time.Time
	override       IObjectPool
}

// Get a object from pool
func (op *ObjectPool) CheckOut() Object {
	var result Object

	now := time.Now()

	if len(op.unlockList) > 0 {
		for v, t := range op.unlockList {
			if t.Add(op.expirationTime).Before(now) {
				delete(op.unlockList, v)
			} else {
				if op.override.Validate(v) {
					delete(op.unlockList, v)
					op.lockList[v] = now
					return v
				}

				delete(op.unlockList, v)
				op.override.Expire(v)
			}
		}
	}

	result = op.override.Create()
	op.lockList[result] = now

	return result
}

// Release object to pool
func (op *ObjectPool) CheckIn(o Object) {
	delete(op.lockList, o)
	op.unlockList[o] = time.Now()
}

// Define a typed pool embedding base pool
type IntgerPool struct {
	nextVal int64
	*ObjectPool
}

var integerPoolInstance *IntgerPool // singleton
var once sync.Once                  // async lock

// Create a singleton pool
func CreateIntegerPool(expirationTime time.Duration) *IntgerPool {
	once.Do(func() {
		integerPoolInstance = &IntgerPool{
			ObjectPool: &ObjectPool{
				expirationTime: expirationTime,
				unlockList:     map[Object]time.Time{},
				lockList:       map[Object]time.Time{},
			},
		}
		integerPoolInstance.ObjectPool.override = integerPoolInstance
	})

	return integerPoolInstance
}

// Check value / object is valid
func (ip *IntgerPool) Validate(i Object) bool {
	val, ok := i.(int64)
	if !ok {
		return false
	}

	return val%2 == 0
}

// Expire a value / object
func (ip *IntgerPool) Expire(i Object) {
	val, ok := i.(int64)

	if !ok {
		panic("Unsupported type.")
	}

	fmt.Printf("Expiring value: %v...\n", val)
}

// Create new value / object
func (ip *IntgerPool) Create() Object {
	result := ip.nextVal

	atomic.AddInt64(&ip.nextVal, 1)

	return result
}
