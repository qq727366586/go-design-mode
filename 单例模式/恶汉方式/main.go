/**
 *@Author luojunying
 *@Date 2022-01-09 23:10
 */
package main

type singleton struct {}

var ins = &singleton{}

func NewSingleton() *singleton {
	return ins
}

//缺点:第一次类装载时会变慢