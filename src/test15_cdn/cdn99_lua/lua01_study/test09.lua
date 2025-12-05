#!/usr/bin/env lua

-- Lua 函数. https://www.runoob.com/lua/lua-functions.html

--[=[=

optional_function_scope function function_name( argument1, argument2, argument3..., argumentn)
    function_body
    return result_params_comma_separated
end

=]=]

local function mult_ret(a, b, c)
    -- do something ...
    return a - 1, b * 2, c + 3
end
a ,b, c = mult_ret(1,2,3)
print(a)
print(b)
print(c)
print(mult_ret(11,22,33))


local function find_max(tab)
    local mi = 1
    local m = tab[1]
    for index, value in pairs(tab) do
        if value > m then
            mi = index
            m = value
        end
    end
    return mi, m
end

tab1 = {12,312,432,43,543,645,65,8}
print(find_max(tab1))

-- 函数的可变参数
function add(...) 
local s = 0  
  for i, v in ipairs{...} do   --> {...} 表示一个由所有变长参数构成的数组  
    s = s + v  
  end  
  return s  
end  
print(add(3,4,5))  --->12
print(add(3,4,5,6,7))  --->25

--[=[=

| 特性     | pairs         | ipairs        |
| ------ | ------------- | ------------- |
| 遍历范围   | 所有键值对（数字+字符串） | 连续整数索引，从 1 开始 |
| 顺序     | 不保证           | 按 1,2,3… 顺序   |
| 遇到 nil | 不影响           | 遇到第一个 nil 停止  |
| 典型用途   | 遍历普通表、字典      | 遍历数组/列表       |


. 使用建议
数组/列表 → 用 ipairs
字典/键值对 → 用 pairs
混合表 → 需要注意，可能同时用到两者
=]=]

-- 在lua中不等于使用 ~= 
print(1 ~= 1) -- false
print(1 ~= 2) -- true
print(true ~= true) -- false
print(true ~= false) -- true


-- 逻辑运算符 and，or， not 
print(true and false)
print(true or false)
print(not true)
print(not false)

-- 连接运算符 .. 
-- #	一元运算符，返回字符串或表的长度。	#"Hello" 返回 5
print(#"hello")
print(#"hello123")
print(#tab1)

--[=[=
运算符优先级,从高到低的顺序：
^
not    - (unary)
*      /       %
+      -
..
<      >      <=     >=     ~=     ==
and
or

=]=]

s1 = '定义字符串的方式1'
s2 = "定义字符串的方式2"
-- 使用[[]]来定义多行字符串，如：[[多行字符串]]
local multilineString = [[
This is a multiline string.
It can contain multiple lines of text.
No need for escape characters.
定义多行字符串
的方式
]]

print(s1)
print(s2)
print(multilineString)

print(string.len(s1))
print(string.len(s2))
print(string.len(multilineString))
print(utf8.len(multilineString))
print(string.len("a"))
print(utf8.len("a"))
print(string.len("你")) -- 3
print(utf8.len("你")) -- 1
print(string.len("你a"))
print(utf8.len("你a"))


-- 实现一个字符串分割方法 split(str, separator)
-- string.find(s: string|number, pattern: string|number, init?: integer, plain?: boolean),返回的第一个参数为子串开始位置，返回的第二个参数为子串的结束位置

s1 = "as,dfg,hjkl"
print(s1)
print(string.find(s1, ","))
print(string.find(s1, ",",4+1))
print(string.find(s1, ",",7+1))

local strArr = {}
strArr[#strArr+1] = string.sub(s1, 1, 3-1)
strArr[#strArr+1] = string.sub(s1, 3+1, 7-1)
strArr[#strArr+1] = string.sub(s1, 7+1, #s1)
for index, value in ipairs(strArr) do
    print("i="..index..",substr="..value)
end

if nil then
    print("nil")
end

if not nil then
    print("not nil")
end

local function split(str, separator)
    if type(separator) ~= "string" then
        error("separator must be a string", 2)
    end

    local arr = {}
    local p = 1
    while true do
        local startI, endI = string.find(str, separator, p)
        if not startI then
            -- startI, endI 都返回nil之后说明 str 在p位置后已经没有 separator 对应的子串了，
            -- 此时将 p 位置后的所有内容作为分割的最后一个元素加入arr
            arr[#arr+1] = string.sub(str, p, #str)
            break
        end
        -- print("start="..startI.."end"..endI..",p="..p)
        arr[#arr+1] = string.sub(str, p, startI - 1)
        p = endI + 1
    end
    return arr
end

aaa = split("as,dfg,hjkl", ",")
for index, value in ipairs(aaa) do
    print("i="..index..",substr="..value)
end

print()

aaa2 = split("asedfgehjkl", "e")
for index, value in ipairs(aaa2) do
    print("i="..index..",substr="..value)
end


-- 不能将 `string` 赋给参数 `<T:table>`。
-- for index, value in ipairs(s1) do
    -- print("i="..index..",char="..value)
-- end

