# Jules API 测试套件

此目录包含 Jules API 的全面测试套件，用于验证每个语言的客户端库是否正常运行。

## 📁 文件结构

- `test_api.py` - Python 客户端的测试程序
- `test_api_js.js` - JavaScript/TypeScript 客户端的测试程序
- `test_api_go.go` - Go 客户端的测试程序
- `.env` - 包含 API 密钥的环境变量文件
- `README.md` - 此测试套件的说明

## 🔐 安全配置

### 环境变量文件 (.env)

要运行测试，您需要在环境变量文件中设置 API 密钥：

```bash
# test/.env
JULES_API_KEY=your_actual_api_key_here
```

**重要**: 请勿将此 `.env` 文件提交到版本控制。

## 🚀 如何运行测试

### Python 测试

```bash
# 先决条件: 已安装 Python 环境和必需软件包
cd py
pip install python-dotenv requests pydantic
cd ../test
python3 ../test/test_api.py
```

### JavaScript 测试

```bash
# 先决条件: 已安装 Node.js 和 npm
# 尚未验证（环境依赖）
cd test
node test_api_js.js
```

### Go 测试

```bash
# 先决条件: 已安装 Go
# 尚未验证（环境依赖）
cd go
go run ../test/test_api.go
```

## 🧪 测试内容

每个测试程序测试以下 API 端点：

1. **📋 List Sources** - 获取可用源码（GitHub 仓库）列表
2. **🚀 Create Session** - 创建新会话
3. **📖 Get Session** - 获取特定会话的详细信息
4. **📂 List Sessions** - 获取会话列表
5. **🎬 List Activities** - 获取会话中的活动列表
6. **💬 Send Message** - 向代理发送消息
7. **📦 Get Source** - 获取特定源码的详细信息

## 📊 预期结果

- **正常情况**: 大部分测试（6/7）应该成功
- **限制**: Send Message 测试可能因会话初始化等待而失败
- **验证内容**: 确认 API 身份验证、HTTP 通信和数据结构的正确性

## 🎯 测试目的

1. **API 功能验证** - 确认每个端点是否正常运行
2. **客户端实现有效性** - 确认每个语言的库是否正确实现
3. **错误处理** - 确认是否执行了适当的错误处理
4. **安全** - 确认 API 密钥是否安全处理

## ⚠️ 注意事项

- API 测试连接到实际的 API 服务器，请注意速率限制
- 运行测试将实际创建 Jules 会话
- 如果遇到错误，请检查网络连接和 API 密钥的有效性

## 📞 支持

对于测试问题，请检查 Jules API 文档或联系支持团队。
