package test03_abstract_factory

/*
好 👍，我们来看 抽象工厂模式（Abstract Factory Pattern） 在 Go 语言中的实现。

1. 模式概念
意图：提供一个创建一系列相关或相互依赖对象的接口，而无需指定它们的具体类。
区别：
	工厂方法：一个工厂只负责创建一个产品。
	抽象工厂：一个工厂可以创建一组相关的产品（比如不同风格的一整套 UI 组件）。

2. 示例场景
假设我们要做一个跨平台的 UI 库，支持 Mac 风格 和 Windows 风格。
每个工厂要能生产两类产品：Button 和 TextField。
*/

// === 抽象产品接口 ===
type Button interface {
	Render1()
}
type TextField interface {
	Render1()
}

// === 具体产品接口 ===
