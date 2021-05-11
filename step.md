# 参考文档
- [youtube go api](https://www.youtube.com/playlist?list=PLKmlCa2HUPq-K7hIyHGbDoYs6YZBM8yA-)
- [youtube go api code](https://github.com/herusdianto/gorm_crud_example)
- [go api example](https://hellokoding.com/crud-restful-apis-with-go-modules-wire-gin-gorm-and-mysql/)



### 一、初始化项目 及目录机构
- git init
- go mod init 

### 二、初始化配置文件
- 添加配置文件（yaml）
- viper读取配置配文件
    ``` golang
    //Init config
    func Init() {
        viper.SetConfigName("config") //  设置配置文件名 (不带后缀)
        viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
        viper.AddConfigPath("config") // 比如添加当前目录
        //viper.AddConfigPath("/workspace/appName1") // 可以多次调用添加路径
        err := viper.ReadInConfig() // 搜索路径，并读取配置数据
        if err != nil {
            panic(fmt.Errorf("fatal error config file: %s ", err))
        }
    }
    ```
- main.go config.Init()

### 三、初始化日志文件
- 添加配置，logrus初始化日志
    ```golang
    func Init() {
        // 设置日志格式为json格式
        if viper.GetString("env") == "production" {
            logrus.SetFormatter(&logrus.JSONFormatter{})
            file, err := os.OpenFile(viper.GetString("log.logger_file"), os.O_CREATE|os.O_WRONLY, 0666)
            if err == nil {
                logrus.SetOutput(file)
            } else {
                panic("日志文件创建失败")
            }
            // 设置日志级别为warn以上
            logrus.SetLevel(logrus.WarnLevel)
            //设置行号
            logrus.SetReportCaller(true)
        } else {
            // 设置日志级别为info以上
            logrus.SetLevel(logrus.InfoLevel)
            logrus.SetFormatter(&logrus.TextFormatter{
                TimestampFormat: "2006-01-02 15:04:05",
            })
        }
    }
    ```
- main.go log.Init()

### 四、初始化db
- 添加配置，gorm初始化mysql
    ```golang
    var DB *gorm.DB
    var err error

    func Init() {
        dns := fmt.Sprintf(
            "%s:%s@/%s?charset=utf8mb4&parseTime=True&loc=Local",
            viper.GetString("db.username"),
            viper.GetString("db.password"),
            viper.GetString("db.dbname"),
        )
        DB, err = gorm.Open(viper.GetString("db.type"), dns)
        if err != nil {
            logrus.Panic("连接数据库失败", err)
        }
        DB.DB().SetMaxIdleConns(10)
        DB.DB().SetMaxOpenConns(100)
        // 启用Logger，显示详细日志
        DB.LogMode(true)
    }

    func CloseDB() {
        err := DB.Close()
        if err != nil {
            logrus.Warn("数据库连接关闭失败")
        }
    }
    ```
- main.go database.Init()


### 五、初始化redis
- 添加配置，redisgo初始化
    ```golang
    var RedisClient *redis.Pool
    func Init() {
        // 从配置文件获取redis的ip以及db
        RedisHost := viper.GetString("redis.addr")
        RedisPassword := viper.GetString("redis.password")
        redisDB := viper.GetInt("redis.db")
        MaxIdle := viper.GetInt("redis.MaxIdle")
        MaxActive := viper.GetInt("redis.MaxActive")
        IdleTimeout := viper.GetInt("redis.IdleTimeout")
        // 建立连接池
        RedisClient = &redis.Pool{
            MaxIdle:     MaxIdle,   /*最大的空闲连接数*/
            MaxActive:   MaxActive, /*最大的激活连接数*/
            IdleTimeout: time.Duration(IdleTimeout) * time.Second,
            Dial: func() (redis.Conn, error) {
                c, err := redis.Dial("tcp", RedisHost, redis.DialPassword(RedisPassword))
                if err != nil {
                    return nil, err
                }
                // 选择db
                c.Do("SELECT", redisDB)
                return c, nil
            },
        }
    }
    ```
- main.go cache.Init()

### 六、测试mq
- 添加配置，go-nsq发收消息测试
    ```golang
    import (
        "fmt"
        "strconv"
        "time"

        "github.com/nsqio/go-nsq"
        "github.com/spf13/viper"
    )

    // nsq发布消息
    func InitProducer(msgBody string) {
        // 默认配置
        var Host = viper.GetString("nsq.addr")
        var Topic = viper.GetString("nsq.topic")
        // 新建生产者
        p, err := nsq.NewProducer(Host, nsq.NewConfig())
        if err != nil {
            panic(err)
        }
        // 发布消息
        if err := p.Publish(Topic, []byte(msgBody)); err != nil {
            panic(err)
        }
    }

    // nsq订阅消息
    type ConsumerT struct{}

    func (*ConsumerT) HandleMessage(msg *nsq.Message) error {
        fmt.Println(string(msg.Body))
        return nil
    }

    func InitConsumer() {
        var Host = viper.GetString("nsq.addr")
        var Topic = viper.GetString("nsq.topic")
        var Channel = viper.GetString("nsq.channel")
        // 新建一个消费者
        cfg := nsq.NewConfig()
        //设置重连时间
        cfg.LookupdPollInterval = time.Second
        c, err := nsq.NewConsumer(Topic, Channel, cfg)

        if err != nil {
            panic(err)
        }
        // 添加消息处理
        c.AddHandler(&ConsumerT{})
        // 建立连接
        if err := c.ConnectToNSQD(Host); err != nil {
            panic(err)
        }
        /*
            //建立多个nsqd连接
            if err := c.ConnectToNSQDs([]string{"127.0.0.1:4150", "127.0.0.1:4152"}); err != nil {
                panic(err)
            }
            //建立一个nsqd连接
            if err := c.ConnectToNSQD("127.0.0.1:4150"); err != nil {
                panic(err)
            }
        */
    }

    func TestNsq() {
        InitConsumer()
        for i := 0; i < 10; i++ {
            InitProducer("hello " + strconv.Itoa(i))
        }
    }

    ```
- main.go queue.TestNsq()



- 创建model
- 创建repository
- 创建service
- controller


### 七、安装自动注入 wire
- go get github.com/google/wire/cmd/wire


### 七、路由模块



### 八、中间件模块

### 八、错误码模块

### 八、参数校验模块

### 八、参数校验模块



### 依赖
- [类型转换](https://github.com/spf13/cast)
- [web框架](gin)
- [orm](grom)
- [依赖注入](wire)
- [配置读取](https://github.com/spf13/viper)