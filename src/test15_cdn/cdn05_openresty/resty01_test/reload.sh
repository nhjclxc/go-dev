# #!/bin/bash

# # 使用元素openresty启动时可以使用这个脚本加快重启（降少命令执行次数）

# # 传入参数（第 1 个参数是配置文件名）
# CONF_FILE=${1:-nginx.conf}

# echo "使用配置文件: $CONF_FILE"

# echo "==> Reloading..."

# openresty -p $(pwd) -c conf/$CONF_FILE -t
# if [ $? -ne 0 ]; then
#     echo "❌ 配置文件检查失败"
#     exit 1
# fi

# echo "==> Stopping old nginx/openresty..."
# nginx -s quit 2>/dev/null || true >> logs/error.log 2>&1
# pkill -9 nginx 2>/dev/null || true >> logs/error.log 2>&1

# echo "==> Starting openresty..."
# openresty -p $(pwd) -c conf/$CONF_FILE >> logs/error.log 2>&1
# # openresty -p $(pwd) -c conf/$CONF_FILE &> logs/error.log 2>&1
# # &> 是覆盖输出， >> 是追加输出

# echo "==> Done."

# # chmod +x reload.sh
# # ./reload.sh
# # ./reload.sh nginx_02.conf


