/**
 *@Author luojunying
 *@Date 2022-01-09 23:00
 */
package main

//懒汉方式 旨在第一次初始化的时候加载
//缺点 非线程安全的, 当正在创建时,有现成来访问ins, 如果ins == nil, 就不是单例模式了

type singleton struct {}

var ins *singleton

func NewSingleton() *singleton {
	if ins == nil {
		ins = &singleton{}
	}
	return ins
}

