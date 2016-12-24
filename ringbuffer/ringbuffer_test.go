package ringbuffer

import "testing"

func TestRingBuffer(t *testing.T) {
	r := NewRingBuffer(10)

	err := r.Write(1)
	if err != nil {
		t.Log(err)
	}
	err = r.Write(2)
	if err != nil {
		t.Log(err)
	}
	err = r.Write(3)
	if err != nil {
		t.Log(err)
	}

	val, err_ := r.Read()
	if err_ != nil {
		t.Log(err_)
	}
	t.Log(val)

	val, err_ = r.Read()
	if err_ != nil {
		t.Log(err_)
	}
	t.Log(val)

}
