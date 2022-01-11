/**
 *@Author luojunying
 *@Date 2022-01-11 18:58
 */
package main

import "fmt"

type Lunch interface {
	Cook()
}

type Rise struct {}

func (r *Rise) Cook() {
	fmt.Println("is is rise")
}

type Tomato struct {}

func (t *Tomato) Cook() {
	fmt.Println("is is tomato")
}

type LunchFactory interface {
	CreateFood() Lunch
	CreateVegetable() Lunch
}


func NewSimpleLunchFactory() LunchFactory {
	return &SimpleLunchFactory{}
}

type SimpleLunchFactory struct {}

func (s *SimpleLunchFactory) CreateFood() Lunch {
	return &Rise{}
}

func (s *SimpleLunchFactory) CreateVegetable() Lunch {
	return &Tomato{}
}

/**
优点: 抽象工厂模式除了具有工厂方法模式的优点外，最主要的优点就是可以在类的内部对产品族进行约束。所谓的产品族，一般或多或少的都存在一定的关联，抽象工厂模式就可以在类内部对产品族的关联关系进行定义和描述，而不必专门引入一个新的类来进行管理。

缺点: 产品族的扩展将是一件十分费力的事情，假如产品族中需要增加一个新的产品，则几乎所有的工厂类都需要进行修改。所以使用抽象工厂模式时，对产品等级结构的划分是非常重要的
 */