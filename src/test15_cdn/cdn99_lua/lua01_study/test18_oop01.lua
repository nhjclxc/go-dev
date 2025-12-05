#!/usr/bin/env lua

-- Lua 面向对象(OOP). https://www.twle.cn/l/yufei/lua53/lua-basic-object-oriented.html

local Person = {}
Person.__index = Person -- 自索引

function Person:new(name, age, extra)
    local obj = extra or {} -- 创建实例
    -- 要想在函数里面使用self变量，方法必须定义为 “表:方法” 的格式
    -- ":"的作用是方法调用和隐式传入 self
    -- setmetatable(obj, self) 等价于 setmetatable(obj, Person)，在new里面建议使用self，这样更佳灵活
    setmetatable(obj, Person)  -- 设置当前创建的这个实例obj的元表指向Person表
    obj.name = name
    obj.age = age
    return obj
end

function Person:getName()
    return self.name
end

-- 如果不使用:定义方法，那么该方法就没有self变量，必须要手动传入
-- : 自动把对象 p1 作为第一个参数 self 传入方法
-- . 则不会自动传入，需要手动写
function Person.getAge(self)
    return self.age
end

--[[
此时：
    p = { name = "zhangsan" }
    元表为 Person
    且 Person.__index = Person

    p1.getName()的过程
        1、先在p1上找getName()，发现找不到
        2、看p1的元表（setmetatable(obj, self)），这里的self就是Person
        3、找到Person表，在Person里面查找，找到了Person:getName()
        5、在Person:getName()里面的self默认是当前调用的元表，即p1
]]

local p1 = Person:new("zhangsan", 18, {["addr"]="aaa",["phone"]="132"})

print(p1)
print(p1:getName())
print(p1.name)
print(p1.age)
print(p1.addr)
print(p1.phone)
-- print(p1.getAge()) -- attempt to index a nil value (local 'self')
print(p1.getAge(p1))

