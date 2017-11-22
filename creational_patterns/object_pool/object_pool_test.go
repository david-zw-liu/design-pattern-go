package object_pool

import (
	"testing"
	"time"
)

func TestIntegerPool(t *testing.T) {
	ip := CreateIntegerPool(100 * time.Microsecond)
	var value int64

	value = (ip.CheckOut()).(int64) // Get from pool
	if value != 0 {
		t.Errorf("Expected first time value is 0, but got %v", value)
	}
	ip.CheckIn(value) // Release value 0

	value = (ip.CheckOut()).(int64) // Get from pool
	if value != 0 {
		t.Errorf("Expected second time value after released is also 0, but got %v", value)
	}
	value = (ip.CheckOut()).(int64) // Get another by creating
	if value != 1 {
		t.Errorf("Expected third time value is 1, but got %v", value)
	}
	// Release value 1
	ip.CheckIn(value)

	// Wait 10 ms to get last release value 1
	// But value 1 is invalide (func Validate(Object) bool),
	// It would get new value 2
	value = (ip.CheckOut()).(int64) // Get another
	if value != 2 {
		t.Errorf("Expected forth time value is 2, but got %v", value)
	}
	ip.CheckIn(value)

	// Wait 200 ms to expire all value in unlockList (expirationTime: 100ms)
	// So it would create new value 3
	time.Sleep(200 * time.Microsecond)
	value = (ip.CheckOut()).(int64) // Get another
	if value != 3 {
		t.Errorf("Expected fifth time value is 3, but got %v", value)
	}
}
