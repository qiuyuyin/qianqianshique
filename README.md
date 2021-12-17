# 千千诗阙

**系统基础架构设计**

:star2:项目目前已经上线：**qianqianshique.com**

## 项目架构设计

### 1、前台：

**Vue+vuetify+vue-router+vuex**

**日间**

![image-20211213205351594](https://yili979.oss-cn-beijing.aliyuncs.com/imgvX7orTk1DtfyaWS.png)

![image-20211213205407432](https://yili979.oss-cn-beijing.aliyuncs.com/imgvjx8tgRMi4DdSOW.png)

![image-20211213205440937](https://yili979.oss-cn-beijing.aliyuncs.com/img5CEscunpxKmkO2e-20211213214924313.png)

**夜间**

![image-20211213205503604](https://yili979.oss-cn-beijing.aliyuncs.com/imguh6q1cDFPB8IWmn.png)

**登录页面：**

![image-20211101145939956](https://i.loli.net/2021/11/01/REbwdxkzuh34O82.png)

**注册页面：**

![image-20211101150000028](https://i.loli.net/2021/11/01/fugE3loCvINRFtp.png)

移动端：

<img src="https://yili979.oss-cn-beijing.aliyuncs.com/imgiwsN57OCnpGo9h2.jpg" alt="Screenshot_2021-12-13-20-58-00-849_com.android.br" style="zoom:33%;" />



<img src="https://yili979.oss-cn-beijing.aliyuncs.com/imgu5FfSBbHzC9yNem.jpg" alt="Screenshot_2021-12-13-20-56-59-946_com.android.br" style="zoom:33%;" />



<img src="https://yili979.oss-cn-beijing.aliyuncs.com/img3wouMi2tKyGxJLI.jpg" alt="Screenshot_2021-12-13-20-57-04-119_com.android.br" style="zoom:33%;" />



### 2、后台：

**go系列实现： gin(路由) + gorm(orm层交互) + qmgo(mongoDB调用) + zap(配置管理) + oss(对象存储) + swagger (接口文档)**：

```txt
├── api
├── config
├── config.yaml
├── core
├── docs
├── global
├── go.mod
├── go.sum
├── initialize
├── log
├── main
├── main.go
├── model
├── resource
├── router
├── service
├── test
└── utils
```



### 3、数据库：

由于数据数量较大，使用关系型数据库会导致数据库不易管理，并且诗词数据中存在较多分枝，使用NoSql可以更高效开发。

**mongoDB：**

![image-20211213203859491](https://yili979.oss-cn-beijing.aliyuncs.com/imgoNEelgjPCfbJqHU.png)

**mysql：**

![image-20211213203821874](https://yili979.oss-cn-beijing.aliyuncs.com/imgPKcjsIaT2RmQp96.png)

本项目仅提供学习，请勿商用
