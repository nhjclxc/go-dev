package test02_factory_method

import (
	"fmt"
	"testing"
)

/*
1. 模式概念
意图：定义一个创建对象的接口，让子类决定实例化哪一个类。
核心：将对象的创建逻辑从使用逻辑中分离开，客户端只依赖接口，不直接依赖具体实现类。
适用场景：当系统需要灵活地决定实例化哪一个具体类，或者需要延迟创建对象。
*/

// Shape 定义一个行为接口
type Shape interface {
	Draw()
}

// 使用多个对象实现该接口
type Circle struct {
	Name string
}

func (c Circle) Draw() {
	fmt.Printf("%s 画了一个圆 \n", c.Name)
}

type Rectangle struct {
	Name string
}

func (r Rectangle) Draw() {
	fmt.Printf("%s 画了一个矩形 \n", r.Name)
}

// ShapeFactory 定义工厂接口
type ShapeFactory interface {
	CreateShape() Shape
}

type CircleFactory struct{}

func (f CircleFactory) CreateShape() Shape {
	return Circle{Name: "yuanyuanyuan"}
}

type RectangleFactory struct{}

func (f RectangleFactory) CreateShape() Shape {
	return Circle{Name: "jujuju"}
}

func Test21(t *testing.T) {

	// 创建工厂

	cf := CircleFactory{}
	c := cf.CreateShape()
	c.Draw()

	rf := RectangleFactory{}
	r := rf.CreateShape()
	r.Draw()

}
