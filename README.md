# CWXUDaXueXi
无锡学院信安专业青年大学习自动化脚本
* 自动完成每周青年大学习
* 自动截图提交班级小管家
* 自动根据身份匹配信息
* 定时任务(每周一9点自动运行)
* 快捷配置
* 可适配其他专业
# 配置config.json
配置文件信息基于抓包\
* identity字段表示团支书(或老师)身份，即收班级小管家作业的人的系统身份，填student或teacher，student表示班委，teacher表示老师
* Cookie字段从江苏共青团获取
![img](https://github.com/MengTL4/CWXUDaXueXi/blob/main/IMG/1.png)
* Authorization字段和member_id字段从班级小管家获取
![img](https://github.com/MengTL4/CWXUDaXueXi/blob/main/IMG/2.png)
# 编译命令
`go build`
# 运行截图
![img](https://github.com/MengTL4/CWXUDaXueXi/blob/main/IMG/3.gif)
