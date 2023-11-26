package goroutine

import (
    "testing"
)

func TestGo(t *testing.T) {
    g := NewGoroutine()
    a := 1
    testFunc := func() {
        a += 1
        panic("test")
    }
    g.Go(testFunc)
    g.Wait()
    if a != 2 {
        t.Errorf("test go failed. expect 2， acual is %d", a)
    }
}

func TestWait(t *testing.T) {

}
