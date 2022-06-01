# 抖音极简版

我们是第三届字节跳动青训营 HeartDancer小组 组号：112358


# 项目启动

请使用go build运行项目，请不要使用go run

```shell
go build

./douyin.exe

```

# 成员分工

| 成员           | 负责内容                                                     |
| -------------- | ------------------------------------------------------------ |
| 朱朝阳         | 接口：视频流、投稿、用户登录、粉丝列表    其他：数据库设计、项目结构初始化、对接云存储 |
| 赵语云、石木子 | 项目的环境搭建：MySQL、Redis                                 |
| 吴振宇         | 接口：赞操作、点赞列表                                       |
| 陈义才         | 接口：评论操作、评论列表                                     |
| 冀虹           | 接口：用户注册、用户信息、发布列表                           |
| 申永燕         | 接口：关注操作、关注列表                                     |

# 其他环境

获取视频封面用到了`ffmpeg`，需要此环境[ffbinaries.com/downloads](https://link.juejin.cn/?target=https%3A%2F%2Fffbinaries.com%2Fdownloads)

windows环境选择**windows-64**下载解压得到一个 .exe 文件，放置到 GOPATH 下的 bin 目录即可 

# 其他

项目介绍：https://bytedance.feishu.cn/docx/doxcnbgkMy2J0Y3E6ihqrvtHXPg

接口文档：https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18902556

