#!/usr/bin/env lua

-- Lua 协程 (coroutine). https://www.twle.cn/l/yufei/lua53/lua-basic-coroutine.html

--[[
下列是 Lua 中 协程(corouting) 的常用函数
方法	描述
coroutine.create()	创建 coroutine，返回 coroutine， 参数是一个函数
coroutine.resume()	重启 coroutine (开启协程)
coroutine.yield()	挂起 coroutine，将 coroutine 设置为挂起状态
coroutine.status()	查看 coroutine 的状态
coroutine的状态有三种：dead，suspend，running
coroutine.wrap（）	创建 coroutine，返回一个函数
调用 wrap 函数，就进入 coroutine，和 create 功能一样
coroutine.running()	返回正在运行的 coroutine，一个 coroutine 就是一个线程
当使用 running 的时候，就是返回一个 corouting 的线程号
]]

local co = coroutine.create(function (i)
    print("协程输出："..i)
end)


coroutine.resume(co, 666)   -- 通过coroutine.resume开启create创建的协程并传递参数
print(coroutine.status(co))  -- dead

print("----create & resume end------")

-- 使用wrap创建一个协程
local cow = coroutine.wrap(function (i, j, k)
    print("使用wrap创建一个协程: i = "..i..", j="..j..",k="..k)
end)
-- 开启协程wrap创建的协程
cow(3,5,9)


local co2 = coroutine.create(function ()
    local ii = 10
    for i = 1, ii do

        print("i = "..i)

        if i % 3 == 0 then
            print("status = "..coroutine.status(co2)..", running = "..coroutine.running()..", i = "..i)
        end
        -- coroutine.yield()
    end
end)

coroutine.resume(co2)
coroutine.resume(co2)
coroutine.resume(co2)


print(coroutine.status(co2))   -- suspended
print(coroutine.running())

print("----------")