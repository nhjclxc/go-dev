#!/usr/bin/env lua


local tab1 = {1,2,3, nil}
local tab2 = {6,7,8, 89, 91}

local metaTab = {}

tab1 = setmetatable(tab1, metaTab)
tab2 = setmetatable(tab2, metaTab)

metaTab["__add"] = function (t1, t2)

    local maxLen = #t1
    if maxLen < #t2 then
        maxLen = #t2
    end

    local res = {}
    for i = 1, maxLen do
        local a = t1[i] or 0
        local b = t2[i] or 0
        res[#res+1] = a + b
    end

    return res
end

local ress = tab1 + tab2

for index, value in ipairs(ress) do
    print(index.." -> "..value)
end

print(getmetatable(tab1))
print(getmetatable(tab2))
print(getmetatable(metaTab))


local mytable = setmetatable({tencent="Tencent"},{})
print(getmetatable(metaTab))