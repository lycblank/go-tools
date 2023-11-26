package goroutine

import (
    "fmt"
    "reflect"
    "sync"
)

var defaultGoroutine Goroutine = NewGoroutine()

func Go(fun interface{}, args ...interface{}) {
    defaultGoroutine.Go(fun, args)
}

type Goroutine interface {
    Go(fun interface{}, args ...interface{})
    Wait()
}

// NewGoroutine new goroutine manager
func NewGoroutine() Goroutine {
    return &goroutine{}
}

type goroutine struct {
    wait sync.WaitGroup
}

// Go add a goroutine and not panic
func (g *goroutine) Go(fun interface{}, args ...interface{}) {
    if fun == nil {
        panic("fun is nil")
    }

    f := reflect.ValueOf(fun)
    for f.Kind() == reflect.Ptr {
        f = f.Elem()
    }
    if f.Kind() != reflect.Func {
        panic(fmt.Sprintf("fun(%+v) is not func", fun))
    }
    funcName := f.Type().Name()

    g.wait.Add(1)
    go func() {
        defer g.handlePanic(funcName, args...)
        var params []reflect.Value
        for _, arg := range args {
            params = append(params, reflect.ValueOf(arg))
        }
        f.Call(params)
    }()
}

// Wait all goroutine exit
func (g *goroutine) Wait() {
    g.wait.Wait()
}

// handlePanic deal goroutine panic
func (g *goroutine) handlePanic(funName string, args ...interface{}) {
    g.wait.Done()
    if r := recover(); r != nil {
        fmt.Printf("handle panic[%+v], fun: %s, args: %+v\n", r, funName, args)
    }
}
