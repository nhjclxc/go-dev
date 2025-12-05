#!/usr/bin/env lua

-- Lua 数组. https://www.runoob.com/lua/lua-arrays.html

-- arr1和arr2的创建方式等价
local arr1 = {11,22,33,44,55,66}
local arr12 = {[1]=11,[2]=22,[3]=33}
print(arr1[1], arr1[2], arr1[3])
print(arr12[1], arr12[2], arr12[3])

-- 建议使用ipairs遍历数组或列表, 索引有序
-- 建议使用pairs遍历切片或字典，索引无序
for index, val in pairs(arr12) do
    print("index = "..index..", val = "..val)
end
print()
for index, val in ipairs(arr12) do
    print("index = "..index..", val = "..val)
end

print("计算数组长度 = "..#arr1)

arr1[4] = 666
-- lua的索引从1开始计算
-- 使用for循环遍历数组,
for i = 1, #arr1 do
    print("i = "..i..",arr1["..i.."]="..arr1[i])
end

-- 默认为追加一个元素
arr1[#arr1+1] = 888
print("arr1.len="..#arr1..", arr1[#arr1]=",arr1[#arr1])

-- 删除某个位置的元素, 删除一个位置的元素之后，该元素后面的数据全部向前移动一个位置
print(arr1[3])
table.remove(arr1, 3)
print(arr1[3])
for i = 1, #arr1 do
    print("i = "..i..",arr1["..i.."]="..arr1[i])
end


-- lua 二维数组
local arr2 = {}  -- 先创建一个一维数组
for i=1, 3 do
    arr2[i] = {}  -- 再循环内部创建二维数组
   for j=1, 4 do
        arr2[i][j] = i*j
    end
end

for i=1, #arr2 do
   for j=1, #arr2[i] do
        print(arr2[i][j])
    end 
    print("---")
end