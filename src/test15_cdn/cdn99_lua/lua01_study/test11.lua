#!/usr/bin/env lua

-- Lua 迭代器. https://www.runoob.com/lua/lua-iterators.html
-- pairs(tab)，索引无序的迭代器遍历，可以遍历表中所有的 key 可以返回 nil
-- ipairs(tab)，索引有序的迭代器遍历，不能返回 nil,如果遇到 nil 则退出


--[=[=
泛型 for 在自己内部保存迭代函数，实际上它保存三个值：迭代函数、状态常量、控制变量。

泛型 for 迭代器提供了集合的 key/value 对，语法格式如下：
for k, v in pairs(t) do
    print(k, v)
end
k, v为变量列表；pairs(t)为表达式列表。
=]=]

local tab= { 
[1] = "a", 
[3] = "b", 
[4] = "c" 
} 
for i,v in pairs(tab) do        -- 输出 "a" ,"b", "c"  ,
    print( tab[i] ) 
end 
print()
for i,v in ipairs(tab) do    -- 输出 "a" ,k=2时断开 
    print( tab[i] ) 
end
