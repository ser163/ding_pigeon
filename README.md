# ding_pigeon
  一个用来与钉钉机器人接口的程序
## 使用说明

您可以使用以下命令来运行此程序：
* 通过管道发送文本消息：`echo "Hello World" | dpigeon`
* 发送文本消息：`dpigeon -type=text -content="Hello World"`

* 发送链接消息：`dpigeon -type=link -title="Link Title" -content="Link Content" -url="https://example.com" -pic="https://example.com/pic.jpg"`

* 发送 Markdown 消息：`dpigeon -type=markdown -title="Markdown Title" -content="# Markdown Content"`

## 配置说明
  将config.yaml.bak 改名为config.yaml
```yaml
webhook: https://oapi.dingtalk.com/robot/send?access_token=YOUR_ACCESS_TOKEN_HERE
secret: YOUR_SECRET_HERE
```
## 编译
Linux编译
```shell
go build -ldflags "-s -w" -o dpingeon main.go
```
windows下交叉编译
```shell
GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o dpingeon main.go
```