# Agenda Client & Service

> 课程《服务计算》作业七：用 Go 完成 Agenda 客户端和 RESTful API 服务端同步开发

[![Build Status](https://travis-ci.org/ishoping/service_agenda.svg?branch=master)](https://travis-ci.org/ishoping/service_agenda)
 
## 下载镜像
docker pull ishoping/service_agenda:1.0
 
## 项目管理与团队协作

### TODO list

#### 任务要求

##### 重构、或新建 agenda 项目，根目录必须包含
- [x] cli 目录
- [x] service 目录
- [x] .travis —— ``.travis.yml``
- [x] apiary.apib —— ``apiary.apib``
- [x] dockerfile —— ``Dockerfile``
- [x] LICENSE
- [x] README.md
- [x] README-yourid.md 记录你的工作摘要（个人评分依据）
##### API 开发
- [x] 使用 API Blueprint 设计
- [x] 资源 URL 命名符合 RESTful 设计标准
- [x] 资源 CRUD 基本完整
##### API 客户端开发
- [x] 可用命令 5 个以上
- [x] 必须有 XXX-test.go 文件
##### 服务端开发
- [x] 使用 sqlite3 作为数据库
- [x] 建议使用课程提供的服务端框架
- [x] 必须有 XXX-test.go 文件
##### 容器镜像制作
- [x] 在 docker hub 上生成镜像
- [x] base 镜像 go-1.8
- [x] 需要加载 sqlite3
- [x] 同时包含客户端与服务器
##### README.md
- [x] 有 build pass 标签
- [x] 有简短使用说明
- [x] 有系统测试的结果（包含如何下载镜像，如何启动服务器，如何使用命令行，cli 的 mock 测试结果， 综合系统测试结果）
##### README-yourid.md
- [x] fork 项目的位置
- [x] 个人工作摘要（每次提交）
- [x] 项目小结
 
### 如何启动服务器
 
![image](https://github.com/ishoping/service_agenda/blob/master/result_image/start_server.png)

### cli的mock测试结果
#### 更改用户属性

![image](https://github.com/ishoping/service_agenda/blob/master/result_image/cli_mock_chusr_result.png)

#### 创建会议

![image](https://github.com/ishoping/service_agenda/blob/master/result_image/cli_mock_cm_result.png)

#### 删除会议

![image](https://github.com/ishoping/service_agenda/blob/master/result_image/cli_mock_dm_result.png)

#### 列出所有用户

![image](https://github.com/ishoping/service_agenda/blob/master/result_image/cli_mock_listAllUsers_result.png)

#### 查询会议

![image](https://github.com/ishoping/service_agenda/blob/master/result_image/cli_mock_qm_result.png)

#### 注册

![image](https://github.com/ishoping/service_agenda/blob/master/result_image/cli_mock_register_result.png)

#### 删除用户

![image](https://github.com/ishoping/service_agenda/blob/master/result_image/cli_mock_userdelete_result.png)


### 综合测试结果
#### 更改用户属性

![image](https://github.com/ishoping/service_agenda/blob/master/result_image/zonghe_chusr_result.png)

#### 创建会议

![image](https://github.com/ishoping/service_agenda/blob/master/result_image/zonghe_cm_result.png)

#### 删除会议

![image](https://github.com/ishoping/service_agenda/blob/master/result_image/zonghe_dm_result.png)

#### 列出所有用户

![image](https://github.com/ishoping/service_agenda/blob/master/result_image/zonghe_listAllUsers_result.png)

#### 查询会议

![image](https://github.com/ishoping/service_agenda/blob/master/result_image/zonghe_qm_result.png)

#### 注册

![image](https://github.com/ishoping/service_agenda/blob/master/result_image/zhonghe_register_result.png)

#### 删除用户

![image](https://github.com/ishoping/service_agenda/blob/master/result_image/zonghe_userdelete_result.png)

