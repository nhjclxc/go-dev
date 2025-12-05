#!/usr/bin/env lua


-- Lua 的模块是由变量、函数等已知元素组成的 table，因此创建一个模块很简单，就是创建一个 table，然后把需要导出的常量、函数放入其中，最后返回这个 table 就行。以下为创建自定义模块 module.lua，文件代码格式如下：
-- 每一个模块的入口文件必须是init.lua, init.lua所在的文件夹名称就是这个模块名称，如当前模块就是my_module
-- init.lua 文件不一过多代码，init.lua只用来导出变量和函数等

-- 模块内部必须要包含模块名， utils是my_module的字模块
local utils = require("my_module.utils")   -- ★ 关键点

my_module = {}

my_module.AAA = "AAA变量"

function my_module.add1(a, b)
    print("这是一个自定义的add函数，非通用")
    return a + b + 1
end

function my_module.mult(a, b)
    print("这是一个自定义的mult函数")
    return utils.mult(a, b)
end

-- 导出当前模块
return my_module