
# 概述

go-zero 提供了一个强大的 conf 包用于加载配置。我们目前支持的 yaml, json, toml 3 种格式的配置文件，go-zero 通过文件后缀会自行加载对应的文件格式。


我们使用 github.com/zeromicro/go-zero/core/conf conf 包进行配置的加载。

第一步我们会定义我们的配置结构体，其中定义我们所有需要的依赖。

第二步接着根据配置编写我们对应格式的配置文件。

第三步通过 conf.MustLoad 加载配置。












