/**
 *@Author luojunying
 *@Date 2022-01-11 18:40
 */
package main

import "fmt"

//餐馆
type Restaurant interface {
	getFood()
}

//北京烤鸭

type bjky struct{}

func (b *bjky) getFood() {
	fmt.Println("北京烤鸭....")
}

//海底捞

type hdl struct {}
func (h *hdl) getFood() {
	fmt.Println("海底捞....")
}

func NewRestaurant(name string) Restaurant {
	switch name {
	case "bjky":
		return &bjky{}
	case "hdl":
		return &hdl{}
	}
	return nil
}



