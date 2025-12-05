#!/usr/bin/env lua

-- Lua 变量 https://www.runoob.com/lua/lua-variables.html
-- Lua 变量有三种类型：全局变量、局部变量、表中的域。
-- 变量的默认值均为 nil。

-- 不使用local修饰的变量即为全局变量
g = 111
local l = 222
function opt()
    print("in opt before fun g = " .. g)
    print("in opt before fun l = " .. l)
    g = 666
    l = 888
    print("in opt after fun g = " .. g)
    print("in opt after fun l = " .. l)
end

print("opt before fun g = " .. g)
print("opt before fun l = " .. l)

opt()

print("opt after fun g = " .. g)
print("opt after fun l = " .. l)


-- 同时对多个变量进行赋值
-- 如果要对多个变量赋值必须依次对每个变量赋值。
local a,b,c = 1,2,3
print("a = "..a..", b="..b..",c="..c)

-- lua会先计算等号=左边对值，在进行赋值操作，所以可以使用=来进行变量的交换
b,a = a,b
print("a = "..a..", b="..b..",c="..c)
