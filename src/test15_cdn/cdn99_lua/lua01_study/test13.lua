#!/usr/bin/env lua


-- Lua 模块与包. https://www.runoob.com/lua/lua-modules-packages.html


--[=[=
require 函数
Lua提供了一个名为require的函数用来加载模块。要加载一个模块，只需要简单地调用就可以了。例如：

require("<模块名>") 或者 require "<模块名>"
=]=]

require("my_module")

print("my_module.AAA = "..my_module.AAA)
print("my_module.add1 = "..my_module.add1(11,22))
print("my_module.mult = "..my_module.mult(11,22))


