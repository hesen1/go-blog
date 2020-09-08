package util

import (
	"sync"
)

// Promise 数据结构
type Promise struct {
	wg  sync.WaitGroup
	res string
	err error
}

// NewPromise 新建一个Promise
func NewPromise(f func() (string, error)) *Promise {
	p := &Promise{}
	p.wg.Add(1)
	go func() {
			p.res, p.err = f()
			p.wg.Done()
	}()
	return p
}

// Then then方法
func (p *Promise) Then(r func(string), e func(error)) {
	go func() {
			p.wg.Wait()
			if p.err != nil {
					e(p.err)
					return
			}
			r(p.res)
	}()
}
