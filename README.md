基于 gin 搭建的 Go Web 开发较通用脚手架模板
```
web_app
├─controllers           控制器目录,负责处理路由、参数校验、请求转发
│
├─dao                   负责数据与存储相关功能
│  ├─mysql           	MySQL
│  ├─redis           	Redis
│
├─logger            	日志初始化目录
│
├─logic                 逻辑层(service服务层),负责处理业务逻辑
│
├─models		模型目录,定义数据库表结构 
│       
├─pkg                	第三方扩展类库目录
│       
├─routes                路由配置目录
│
├─settings              配置加载目录
│
├─config.yaml           配置文件
│
├─main.go           	入口文件
│
├─README.md             README 文件
│
├─web_app.log           日志 文件
```

#### 配置文件


如果要修改配置文件，只要修改：`config.yaml` 为对应的格式文件即可

如： `config.yaml` -> `config.json`。修改 `settings/settings.go`文件：

```go
func Init() (err error) {
	viper.SetConfigFile("./config.json")
	·
	·
}
```

#### 路由

gin 创建路由时替换 gin 默认的 `logger` 及 `recovery`

```go
func Setup() (route *gin.Engine) {
	route = gin.New()
	route.Use(logger.GinLogger())
	route.Use(logger.GinRecovery(true))
	route.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world.")
	})
	return route
}
```

#### 记录日志

```go
zap.L().Info(”logger init success...“)
zap.L().Error("connect DB failed", zap.Error(err))
```

#### 引入外部包

写入 `pak` 目录

#### 优雅关机

配合supervisor重启服务
