#!/usr/bin/env lua

-- Lua table(表). https://www.runoob.com/lua/lua-tables.html

-- 初始化表
mytable = {}
print(mytable)

-- 指定值
mytable[1]= "Lua"
print(mytable)

-- 移除引用
mytable = nil
print(mytable)
-- lua 垃圾回收会释放内存

local tt = {}
tt[1] = 'key是数字'
tt['a'] = 'key是字符串'
print(tt[1])
print(tt['a'])


-- 测试tab的赋值为地址赋值
tab = {}

tab['k'] = 'val'
print(tab['k'])

local t1 = tab
print(t1['k'])
local t2 = tab
print(t2['k'])

t1['kkk'] = 'valvalval'
print(t1['kkk'])
print(t2['kkk'])
-- t1['kkk']和 t2['kkk'] 都输出 'valvalval'

--[=[=
table的操作
按照指定符号拼接表中的所有元素：table.concat(list, separator)
表的指定位置插入数据：table.insert
移除表中某个位置的元素：table.remove
对表进行升序排序：table.sort
对表进行降序排序：table.sort(t, function(a, b) return a > b end)
=]=]
print()

--[=[= 打印tab =]=]
local function printTab(tab)
    for index, value in pairs(tab) do
        print("index="..index..", value="..value)
    end
end
t2 = {11,22,33,44,55,66,77,88,99}
printTab(t2)

print(table.concat(t2, "->"))

table.insert(t2, 888)
printTab(t2)
print()
table.insert(t2, 6, 666)
printTab(t2)

aa = table.remove(t2, 1)
print(aa)

t3 = {213,324,3,54,654,6,7,6,23,2}
printTab(t3)

table.sort(t3)
table.sort(t3, function (a, b) return a < b end)
printTab(t3)

table.sort(t3, function (a, b) return a > b end)
printTab(t3)

-- 把 table 展开成多个值
print(table.unpack(t3))

print("------------")


-- table 的 key 本身就是去重的，但是value不会去重，类似map

-- table按value去重
function table.unique(tab)
    local retTab = {}
    local valueTab = {}
    for key, value in pairs(tab) do
        if not valueTab[value] then
            valueTab[value] = true
            retTab[key] = value
        end
    end
    return retTab
end

local tt = {11,22,33,55,66,11,88,11}
printTab(tt)
tt = table.unique(tt)
print()
printTab(tt)

print(tt[111])
print(tt[111] == true)
print(not tt[111])
