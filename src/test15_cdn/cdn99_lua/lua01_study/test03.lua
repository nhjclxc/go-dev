#!/usr/bin/env/ lua

-- Lua数据类型 https://www.runoob.com/lua/lua-data-types.html
-- Lua 中有 8 个基本类型分别为：nil、boolean、number、string、userdata、function、thread 和 table。

print(type("Hello world"))      --> string
print(type('Hello world'))      --> string
print(type(Hello))      --> nil
print(type(10.4*3))             --> number
print(type(print))              --> function
print(type(type))               --> function
print(type(true))               --> boolean
print(type(nil))                --> nil
print(type(type(X)))            --> string



-- nil
print("------------nil--------------")
tab = { key1 = "val1", key2 = "val2", "val3" }  -- table 在输出时无序
for k, v in pairs(tab) do
    print(k .. " - " .. v)
end

-- 用nil的删除功能将数据置空
tab.key1 = nil
for k, v in pairs(tab) do
    print(k .. " - " .. v)
end

print(X)
print(type(X))
print(type(type(X)))  -- type(X)返回字符串
print(X == "nil")
print(type(X) == "nil")

print(x == nil)
print(x == "nil")

print("------------nil--------------")

print("------------bool--------------")

print(true)
print(false)
print(type(true))
print(type(false))

-- Lua 把 false 和 nil 看作是 false，其他的都为 true，数字 0 也是 true:
-- if false or nil then
if nil then
    print("nil是true")
else
    print("nil是false")
end
print(type(false))

if 0 then
    print("0是true")
else
    print("0 是false")
end

if 1 then
    print("1 是true")
else
    print("1 是false")
end

print("------------bool--------------")


-- number

-- string

-- table

print("------------ table --------------")
-- table 的索引可以是数字、字符串或表类型
local tab1 = {}
print(tab1)

tab1.key1 = "val1"
print(tab1)

local tab2 = {"aa", "bb", "cc"}
print(tab2)

-- 变量前加local表示这个变量是局部变量，局部变量的作用域是块级的
-- 变量前不加local表示全局变量


-- 数组的索引可以是数字或者是字符串。
a = {}
a["key11"] = "value11"
key = 10
a[key] = 22
a[key] = a[key] + 11
for k, v in pairs(a) do
    print(k .. " : " .. v)
end
-- 在 Lua 中，.. 是 字符串连接符（concatenation operator），用于把两个或多个字符串拼接在一起。
-- 在 Lua 中，+ 是 数字相加，类型必须是数字

-- 在 Lua 里表的默认初始索引一般以 1 开始
tab3 = {"a", "b", "c", "d", "e"}

for key, val in pairs(tab3) do 
    print(key .. " -> " .. val)
end

print("------------ table --------------")
print("------------ function --------------")
-- 函数是一等值（first-class value），可以被赋值给变量、作为参数传递、返回值，并形成闭包。
-- 在 Lua 中function函数也可以是一个变量，可以使用闭包（closure）特性

function add(a ,b)
    return a + b
end

print(add(1,2))
local addtmp = add
print(addtmp(11,22))
tab3["add_func"] = add  -- 把函数放到tab里面进行使用
print(tab3["add_func"](111,222))

-- Lua 函数可以访问 它定义时所在作用域的局部变量，即使这个函数在外部被调用，也能记住原来的变量，这就是 闭包。
function makeCounter()
    local count = 0  -- 局部变量
    return function() -- 匿名函数闭包
        count = count + 1
        return count
    end
end

local counter1 = makeCounter()
print(counter1())  --> 1
print(counter1())  --> 2
print(counter1())  --> 3

local counter2 = makeCounter()
print(counter2())  --> 1  （counter2 是另一个闭包，独立计数）

-- function作为函数返回值
function apply(func, a, b)
    return func(a,b)    
end

local a = apply(add, 3, 5)
print("a = " .. a)

function fib(num)
    if num == 1 or num ==2 then
        return 1
    end
    return fib(num -1 ) + fib(num -2)
end

print(fib(1))
print(fib(2))
print(fib(3))
print(fib(4))
print(fib(5))

function factorial1(num)
    if num == 0 then
        return 1
    end
    return num * factorial1(num -1)
end

print("------------ function --------------")


-- 出事啊户tab时，索引不能时数字和字符串，后续设置值的时候可以使用
-- tabb = {1 : 111, "22" : "23"}
local tab111 = {22,"23"} -- 混合初始化
local taba = {}
taba[1] = "aaa"
taba["b"] = "bbb"
for key, val in pairs(taba) do 
    print(key .. " -> ".. val)
end