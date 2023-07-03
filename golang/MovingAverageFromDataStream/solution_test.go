package MovingAverageFromDataStream

import "testing"

func TestMovingAverage_Next(t *testing.T) {
	q := Constructor(3)
	result := q.Next(1)
	if result != 1.0 {
		t.Errorf("1 nerror: %f", result)
	}
	result = q.Next(10)
	if result != 5.5 {
		t.Errorf("2 err: %f", result)
	}
	result = q.Next(3)
	if result != 4.666667 {
		t.Errorf("3 err: %f", result)
	}
	result = q.Next(5)
	if result != 6.0 {
		t.Errorf("4 err: %f", result)
	}
}
