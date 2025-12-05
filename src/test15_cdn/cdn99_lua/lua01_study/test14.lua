#!/usr/bin/env lua

-- Lua 面向对象 https://www.runoob.com/lua/lua-object-oriented.html


-- Lua 中面向对象 对象由属性和方法组成。
-- Lua 中的类可以通过 table + function 模拟出来。

-- 创建一个table来模拟类
Person = {}

-- 通过new方法类创建对象
function Person:new(name, age)
    local obj = {}  -- 创建一个新的表作为对象
    setmetatable(obj, self)  -- 设置元表，使其成为 Person 的实例
    self.__index = self  -- 设置索引元方法，指向 Person
    -- 初始化对象属性
    obj['name'] = name
    obj.age = age
    return obj
end

-- 绑定方法
function Person:sayHello()
    print("Hello, my name is " .. self.name)
end

-- 使用
local oobbjj = Person:new("zhangsan", 18)
print(oobbjj['name'])
print(oobbjj.age)
oobbjj:sayHello()



-- 定义矩形类
Rectangle = {area = 0, perimeter = 0, length = 0, breadth = 0}

function Rectangle:new(length, breadth)
    local obj = {}
    
end

