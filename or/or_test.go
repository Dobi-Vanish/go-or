package or

import (
	"testing"
	"time"
)

func TestOrEmpty(t *testing.T) {
	ch := Or()
	select {
	case <-ch:
	default:
		t.Error("ожидался закрытый канал, но чтение заблокировалось")
	}
}

func TestOrSingle(t *testing.T) {
	c := make(chan interface{})
	closed := make(chan struct{})

	go func() {
		<-Or(c)
		close(closed)
	}()

	select {
	case <-closed:
		t.Error("Or вернул закрытый канал раньше времени")
	case <-time.After(50 * time.Millisecond):
	}

	close(c)

	select {
	case <-closed:
	case <-time.After(50 * time.Millisecond):
		t.Error("Or не закрыл канал после закрытия входного")
	}
}

func TestOrMultiple(t *testing.T) {
	ch1 := make(chan interface{})
	ch2 := make(chan interface{})
	ch3 := make(chan interface{})

	done := Or(ch1, ch2, ch3)
	closed := make(chan struct{})

	go func() {
		<-done
		close(closed)
	}()

	close(ch2)

	select {
	case <-closed:
	case <-time.After(50 * time.Millisecond):
		t.Error("Or не закрылся после закрытия одного из каналов")
	}
	close(ch1)
	close(ch3)
}

func TestOrWithAlreadyClosed(t *testing.T) {
	ch1 := make(chan interface{})
	close(ch1)

	ch2 := make(chan interface{})
	done := Or(ch1, ch2)

	select {
	case <-done:
	case <-time.After(50 * time.Millisecond):
		t.Error("Or не закрылся, хотя один из каналов уже был закрыт")
	}
}

func TestOrDoesNotCloseInputs(t *testing.T) {
	ch := make(chan interface{})
	_ = Or(ch)

	select {
	case <-ch:
		t.Error("Or не должен закрывать входной канал")
	default:
	}
}
