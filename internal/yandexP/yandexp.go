package yandexp

import (
	"fmt"
	"time"
)

const (
	DeliveryStatePending   DeliveryState = "pending"      // message pending
	DeliveryStateAck       DeliveryState = "acknowledged" // message received
	DeliveryStateProcessed DeliveryState = "precessed"    //sucseess
	DeliveryStateCanceled  DeliveryState = "canceled"     // canceled
)

type MyType int
type DeliveryState string
type CircularBuffer struct {
	Values  []float64 // current value buffer
	headIdx int       // index head (first not empty element)
	tailIdx int       // index tail (last empty element)
}
type IntSlice []int

type Stopwatch struct {
	current  time.Time
	durrList []time.Duration
}

// Start for record current time
func (sw *Stopwatch) Start() {
	sw.current = time.Now()
}

// SaveSplit for record durration between current and now
func (sw *Stopwatch) SaveSplit() {
	sw.durrList = append(sw.durrList, time.Since(sw.current))
}

// GetResults for return durrList
func (sw *Stopwatch) GetResults() []time.Duration {
	return sw.durrList
}

func (s *IntSlice) Add(v int) {
	*s = append(*s, v)
}

func Handle(num float64, add func(float64)) {
	add(num)
}

// GetCurrentSize return current length buffer
func (b CircularBuffer) GetCurrentSize() int {
	if b.tailIdx < b.headIdx {
		return b.tailIdx + cap(b.Values) - b.headIdx
	}
	return b.tailIdx - b.headIdx
}

// cGetValues return slice current values buffer, save sort type
func (b CircularBuffer) GetValues() (retValues []float64) {
	for i := b.headIdx; i != b.tailIdx; i = (i + 1) % cap(b.Values) {
		retValues = append(retValues, b.Values[i])
	}
	return
}

// AddValue addd new value to buffer
func (b *CircularBuffer) AddValue(v float64) {
	b.Values[b.tailIdx] = v
	b.tailIdx = (b.tailIdx + 1) % cap(b.Values)
	if b.tailIdx == b.headIdx {
		b.headIdx = (b.headIdx + 1) % cap(b.Values)
	}
}

// NewCircularBuffer - constructor type CircullarBuffer
func NewCircularBuffer(size int) CircularBuffer {
	return CircularBuffer{Values: make([]float64, size+1)}
}

func (b CircularBuffer) ForceSetValueByIdx(idx int, v float64) {
	// лучше не использовать такой приём на практике, когда параметр метода
	// не указатель, а значение
	b.Values[idx] = v
}

func (m MyType) String() string {
	return fmt.Sprintf("MyType: %d", m)
}

func Print(m MyType) {
	m = 5

	s := m.String()
	fmt.Println(s)
}

// IsValid check valid value of DeliveryState
func (s DeliveryState) IsValid() bool {
	switch s {
	case DeliveryStatePending, DeliveryStateAck, DeliveryStateProcessed, DeliveryStateCanceled:
		return true
	default:
		return false
	}
}

// String return string view tipe DeliveryState
func (s DeliveryState) String() string {
	return s.String()
}

func HenleDeliveryStatus(status DeliveryState) error {
	if !status.IsValid() {
		return fmt.Errorf("status: invalid")
	}

	// code

	return nil
}
