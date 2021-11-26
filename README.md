### 1、前台：

**Vue + vuetify + vue-router + vuex + axios**

项目目前进度：

**日间模式：**

![image-20211126164615299](https://i.loli.net/2021/11/26/p8uMSCZyNOQAa6w.png)

![image-20211126164853992](https://i.loli.net/2021/11/26/cgtX1SvapUsVYPO.png)

![image-20211126164915633](https://i.loli.net/2021/11/26/sP73xjkXOQSLr68.png)

![image-20211126164943350](https://i.loli.net/2021/11/26/DThv5Kart6WGdXU.png)

**夜间模式**：

![image-20211126165121159](https://i.loli.net/2021/11/26/dEsLeOFauv5kiqC.png)

![image-20211126165147925](https://i.loli.net/2021/11/26/uaLz1Rbv3fC92FT.png)

### 2、后台：

**go系列实现： gin(路由) + gorm(orm层交互) + casbin(用户权限验证) + jwt(登录缓存) + zap(配置管理) + oss(对象存储) + swagger (接口文档)**：

```js
- poem_server
---api      	controller层
---config   	配置model
---core     	核心启动进程
---docs     	swagger api文档
---global   	全局调用配置
---initialize   初始化
---log          保存日志位置
---middleware   中间件
---model      	model层
---resource     保存配置文件位置
---router  	 	配置路由
---service      service层
---utils   		工具
-config.yaml 	配置文件
-main.go  		启动位置
```

### 3、数据库（目前并没有加入过多用户系统）

**mongoDB：**

![image-20211126165833994](https://i.loli.net/2021/11/26/kIYcpTgfmHhBGWL.png)

**mysql：**

![image-20211126165419925](https://i.loli.net/2021/11/26/nP3rzHcCJOEXTFq.png)
