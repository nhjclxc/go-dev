# go设计模式

[go-patterns](https://github.com/tmrts/go-patterns)

[云原生AI](https://www.yuque.com/onex/zsxq/asqq8sfnh5i3kuyz)

[go设计模式](https://feizaonet.feishu.cn/wiki/QMtxw0Pzhimb0PkoZj1coWhxnrh)

[go数据结构与算法](https://github.com/TheAlgorithms/Go)

https://github.com/onexstack/design-pattern/tree/master

[go设计模式](https://github.com/lee501/go-patterns) -> [go数据结构与算法](https://github.com/0xAX/go-algorithms) -> [go游戏](https://github.com/hyper0x/awesome-go-China/blob/master/zh_CN/README.md#game-development)


好的，我帮你列一个 **设计模式清单**，按照 **《GoF 23 种经典设计模式》** 来整理，并兼顾 Go 语言的特点（有些模式在 Go 中实现方式会略有不同，因为 Go 没有继承，更多用组合和接口）。我按 **创建型、结构型、行为型** 分类列出：

---

## **一、创建型模式 (Creational Patterns)**

1. **单例模式（Singleton）**
2. **工厂方法模式（Factory Method）**
3. **抽象工厂模式（Abstract Factory）**
4. **建造者模式（Builder）**
5. **原型模式（Prototype）**

---

## **二、结构型模式 (Structural Patterns)**

6. **适配器模式（Adapter）**
7. **桥接模式（Bridge）**
8. **组合模式（Composite）**
9. **装饰器模式（Decorator）**
10. **外观模式（Facade）**
11. **享元模式（Flyweight）**
12. **代理模式（Proxy）**

---

## **三、行为型模式 (Behavioral Patterns)**

13. **责任链模式（Chain of Responsibility）**
14. **命令模式（Command）**
15. **解释器模式（Interpreter）**
16. **迭代器模式（Iterator）**
17. **中介者模式（Mediator）**
18. **备忘录模式（Memento）**
19. **观察者模式（Observer）**
20. **状态模式（State）**
21. **策略模式（Strategy）**
22. **模板方法模式（Template Method）**
23. **访问者模式（Visitor）**

---

💡 **备注**：

* Go 没有类继承，所以在实现很多模式时更多依赖 **接口(interface) + 组合(struct embedding)**。
* 有些模式在 Go 里使用 **闭包、函数变量、channel 或 goroutine** 会更自然。

---

如果你需要，我可以帮你 **生成一个按顺序的 Go 设计模式实现计划表**，每个模式都附带示例代码和适用场景，这样我们就可以按顺序实现。

你希望我帮你生成这个计划表吗？
