

-- 如果要将一个函数在其他文件或模块使用，必须要使用一个table进行导出

local utils = {}

function utils.mult(a, b)
    return a * b
end

return utils
