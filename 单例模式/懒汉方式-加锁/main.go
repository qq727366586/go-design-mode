/**
 *@Author luojunying
 *@Date 2022-01-09 23:04
 */
package main

import "sync"

type singleton struct {}

var ins *singleton

var mt sync.Mutex

func NewSingleton() *singleton {
	mt.Lock()
	defer mt.Unlock()
	if ins == nil {
		ins = &singleton{}
	}
	return ins
}

//缺点: 虽然可以做到单例模式,但是并发情况下会造成网络拥塞问题
