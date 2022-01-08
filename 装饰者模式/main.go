/**
 *@Author luojunying
 *@Date 2022-01-08 20:59
 */
package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"time"
)

//正常思维版本
func WrapLogger1(fun func(int) float64, logger *log.Logger) func(int) float64 {
	return func(i int) (result float64) {
		defer func(t time.Time) {
			logger.Printf("took=%v,n=%v,result=%v", time.Since(t), i, result)
		}(time.Now())
		return fun(i)
	}
}

type PiFunc func(int) float64
//稍微难理解的版本
func WrapLogger2(fun PiFunc, logger *log.Logger) PiFunc {
	return func(n int) float64 {
		fn := func(n int) (result float64) {
			defer func(t time.Time) {
				logger.Printf("took=%v,n=%v,result=%v", time.Since(t), n, result)
			}(time.Now())
			return fun(n)
		}
		return fn(n)
	}
}

//算π
func Pi(n int) float64 {
	ch := make(chan float64)
	for k := 0; k < n; k++ {
		go func(ch chan float64, k float64) {
			ch <- 4 * math.Pow(-1, k ) / (2 *k + 1)
		}(ch, float64(k))
	}
	fmt.Println("start.")
	result := 0.0

	for i := 0; i < n; i++ {
		result += <-ch
	}
	return result
}

func main() {
	f := WrapLogger1(Pi, log.New(os.Stdout, "[test] ", 1))
	res := f(10000)
	fmt.Println(res)
}