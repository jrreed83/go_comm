package ringbuffer

import "errors"

type ringbuffer struct {
	readPtr  uint32
	writePtr uint32
	capacity uint32
	buffer   []byte
	length   uint32
}

func NewRingBuffer(capacity uint32) *ringbuffer {
	return &ringbuffer{readPtr: 0,
		writePtr: 0,
		capacity: capacity,
		buffer:   make([]byte, capacity),
		length:   0}
}

func (r *ringbuffer) write(x byte) error {
	if r.length == r.capacity {
		return errors.New("No space to write")
	}
	r.buffer[r.writePtr] = x
	r.writePtr = (r.writePtr + 1) % r.capacity
	r.length += 1
	return nil
}

func (r *ringbuffer) read() (byte, error) {
	if r.length == 0 {
		return 0, errors.New("No items to read")
	}
	x := r.buffer[r.readPtr]
	r.readPtr = (r.readPtr + 1) % r.capacity
	r.length += 1
	return x, nil
}
