#!/usr/bin/env lua

-- Lua 循环   https://www.runoob.com/lua/lua-loops.html
-- while(...) do ... end
-- for
-- repeat ... until
-- 循环控制：break, goto

-- 1、while 
print("------------ while --------------")
local a = 1
while a < 5 do
    print("a = ".. a)
    -- a += 1 未知符号`+=`, lua不支持+=
    a = a + 1
end


print("------------ while --------------")
print("------------ for --------------")
--[=[
for var=exp1,exp2,exp3 do  
    <执行体>  
end  
]=]
-- exp1是循环变量的初始化 i = 10
-- exp2是循环变量的终止条件 i == 1
-- exp3是每一次循环变量的操作 i = i -1, exp3默认是1即表示i=i+1
for i=10,1,-1 do
    print(i)
end

local function ff(i)
    return i*3
end

for ii = 1, ff(3) do
    print("ii = "..ii)
end

-- 泛型for循环
local a = {"one", "two", "three"}
for k, v in pairs(a) do print(k.." -> "..v) end


print("------------ repeat...until 循环 --------------")
--[=[
repeat
   statements
until( condition )
]=]
-- repeat循环至少执行一次
local ri = 0
repeat
    print('ri = '..ri)
    ri = ri + 1
until ri == 10

print("------------ break --------------")

local aaa = 0
local sum = 0
repeat
    aaa = aaa + 1
    sum = sum + aaa
    print("aaa="..aaa..",sum="..sum)
    if sum > 10 then
        break
    end
until aaa == 10


print("------------ goto --------------")

-- 定义goto标签 :: Label ::


local wa = 11
:: Label :: print("print label")
while wa > 5 do
    print("wa = "..wa)
    wa = wa -1
    if wa % 3 == 0 then
        goto Label
    end
end
-- 执行 goto Label之后代码接着执行while wa > 5 do,所以循环变量的控制必须在goto语句之前，否则永远无法更新循环变量


print("---")
-- 模拟 continue 语句
local i = 11
while i > 5 do
    i = i - 1
    if i % 3 == 0 then
        goto continue
    end
    -- do something ...
    print("i = "..i)
    -- ::continue:: 标签必须写在最后面，以便跳过do something ...
    ::continue::
end
