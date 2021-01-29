# GinBlog

#### 创建入口main.go

#### gomod 初始化 
```
    go mod init GinBlog
```

#### 安装Gin框架
```
    go get -u github.com/gin-gonic/gin 
```
### 目录框架
+ config	//管理配置参数
+ model    //数据库存储模型
+ api          //控制器,在里面分版本做
+ middleware //用于做中间件,配置跨域中间件
+ routers   //路由接口
+ utils         //做公共工具的包
+ log       //用于管理日志
+ view      //存放html
+static     //托管静态资源

### 添加配置文件
1. 安装goini  `go get gopkg.in/ini.v1`
2. 新建config.ini配置文件
3. 在util下新建setting.go配置工具类


### 初始化项目
1. 建立route.go
2. 再main中引入route的Init方法
3.启动测试链接

