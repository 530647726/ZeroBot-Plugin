<img src="https://socialify.git.ci/Yiwen-Chan/ZeroBot-Plugin/image?forks=1&issues=1&language=1&owner=1&pulls=1&stargazers=1&theme=Light" alt="ZeroBot-Plugin" width="640" height="320" />

# ZeroBot-Plugin

![Badge](https://img.shields.io/badge/OneBot-v11-black)
[![License](https://img.shields.io/github/license/Yiwen-Chan/ZeroBot-Plugin.svg)](https://raw.githubusercontent.com/Yiwen-Chan/ZeroBot-Plugin/master/LICENSE)
[![反馈群](https://img.shields.io/badge/反馈群-1048452984-green.svg)](https://jq.qq.com/?_wv=1027&k=QMb7x1mM)

### 功能
- 群管
- 涩图
- 点歌
- TODO...

### 使用方法

本项目符合 [OneBot](https://github.com/howmanybots/onebot) 标准，可基于以下项目与机器人框架/平台进行交互
| 项目地址 | 平台 | 核心作者 | 备注 |
| --- | --- | --- | --- |
| [Yiwen-Chan/OneBot-YaYa](https://github.com/Yiwen-Chan/OneBot-YaYa) | [先驱](https://www.xianqubot.com/) | kanri |  |
| [richardchien/coolq-http-api](https://github.com/richardchien/coolq-http-api) | CKYU | richardchien | 可在 Mirai 平台使用 [mirai-native](https://github.com/iTXTech/mirai-native) 加载 |
| [Mrs4s/go-cqhttp](https://github.com/Mrs4s/go-cqhttp) | [MiraiGo](https://github.com/Mrs4s/MiraiGo) | Mrs4s |  |
| [yyuueexxiinngg/cqhttp-mirai](https://github.com/yyuueexxiinngg/cqhttp-mirai) | [Mirai](https://github.com/mamoe/mirai) | yyuueexxiinngg |  |
| [takayama-lily/onebot](https://github.com/takayama-lily/onebot) | [OICQ](https://github.com/takayama-lily/oicq) | takayama |  |

#### 本地编译
1. 下载安装 [Go](https://studygolang.com/dl/golang/go1.16.2.windows-amd64.msi) 环境
2. 下载安装 [TDM-GCC](https://github.com/jmeubank/tdm-gcc/releases)，并添加到环境变量
3. [clone](https://github.com/Yiwen-Chan/ZeroBot-Plugin/archive/master.zip) 本项目，本地解压
4. 编辑 main.go 文件，内容按需修改
5. 双击点击 build.bat 文件
6. 运行框架，并同时运行本插件

#### 利用 Actions 编译 (推荐)
1. 点击右上角 Fork 本项目，并转跳到自己 Fork 的仓库
2. 点击仓库上方的 Actions 按钮，确认使用 Actions
3. 编辑 main.go 文件，内容按需修改，返回仓库
4. 点击 Actions 按钮，等待编译完成，在 Actions 里下载编译好的文件
5. 运行框架，并同时运行本插件

