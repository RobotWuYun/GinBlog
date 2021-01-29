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

### 数据库
1. 安装gorm `go get -u github.com/jinzhu/gorm`
2. 在model下建立 user,category,Article的model
3. 建立db入口文件
    + 引入驱动
    + 建立链接
    + 相关设置
    + 自动迁移
4. 奖励对象并映射数据库和json 
5. 将Model的结构体注册到自动迁移里
6. 再主函数调用数据库初始化函数

### 错误处理
1. 再util里建立errmsg.go,并建立msg,code之间的关系

### API接口
1. 在api下分类建立api源码

### 在route中注册地址
### 在model添加数据库操作方法
### 在api中写逻辑,调用model中的方法