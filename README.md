# GroupManager
一个基于cqhttp的点开即用的管理群聊插件

[![License](https://img.shields.io/github/license/Yiwen-Chan/GroupManagerBot.svg)](https://raw.githubusercontent.com/Yiwen-Chan/GroupManagerBot/master/LICENSE)
[![QQ 群](https://img.shields.io/badge/qq%E7%BE%A4-1048452984-green.svg)](https://jq.qq.com/?_wv=1027&k=QMb7x1mM)

本项目符合 [OneBot](https://github.com/howmanybots/onebot) 标准

可基于以下项目与机器人框架/平台进行交互
| 项目地址 | 平台 | 核心作者 | 备注 |
| --- | --- | --- | --- |
| [richardchien/coolq-http-api](https://github.com/richardchien/coolq-http-api) | CKYU | richardchien | 可在 Mirai 平台使用 [mirai-native](https://github.com/iTXTech/mirai-native) 加载 |
| [Mrs4s/go-cqhttp](https://github.com/Mrs4s/go-cqhttp) | MiraiGo | Mrs4s |  |
| [yyuueexxiinngg/cqhttp-mirai](https://github.com/yyuueexxiinngg/cqhttp-mirai) | Mirai | yyuueexxiinngg |  |

## 开始使用

注意：本插件使用websocket与cqhttp项目进行交互，非反向ws

1.下载对应版本的release，可直接运行

2.第一次运行自动产生config.json，修改后再次运行

3.发送“群管系统”呼出菜单

## 配置相关

GroupManager设置(GroupManager\config.json)
```json
{
	"master": 12345678,
	"host": "127.0.0.1",
	"port": 2333
}
```
若使用cqhttp-mirai，则配置如下(miraiOK\plugins\CQHTTPMirai\setting.yml)
```yaml
...
  # 正向Websocket服务器
  ws:
    # 可选，是否启用正向Websocket服务器，默认不启用
    enable: true
    # 可选，上报消息格式，string 为字符串格式，array 为数组格式, 默认为string
    postMessageFormat: string
    # 可选，访问口令, 默认为空, 即不设置Token
    accessToken: ""
    # 监听主机
    wsHost: "127.0.0.1"
    # 监听端口
    wsPort: 2333
...
```
若使用go-cqhttp，则配置如下(go-cqhttp/config.json)
```json
...
	"ws_config": {
		"enabled": true,
		"host": "127.0.0.1",
		"port": 2333
	},
...
```
注意:以上仅列出了与ws相关设置，其他配置可自行摸索或找我配置
## 功能列表
- [x] 禁言
- [x] 解除禁言
- [x] 全员禁言
- [x] 解除全员禁言v
- [ ] 撤回
- [x] 踢出
- [x] 退出群聊
- [x] 修改群名片
- [x] 设置群头衔
- [ ] 同意好友添加
- [ ] 同意群聊邀请
- [x] 私聊转发
- [x] 群聊转发
- [ ] 敏感词
- [ ] 黑白名单

## 开源许可

[GPL-3.0](https://raw.githubusercontent.com/Yiwen-Chan/GroupManagerBot/master/LICENSE) © Yiwen-Chan

## 问题反馈

遇到问题、BUG、或有其他建议欢迎提issue
也欢迎加入 QQ 交流群 1048452984 来和大家讨论

## 感谢

[Richard Chien](https://github.com/richardchien): [CQHTTP](https://github.com/richardchien/coolq-http-api) ， [NoneBot](https://github.com/nonebot/nonebot) 和 [qq-bot-api](https://github.com/catsworld/qq-bot-api)
