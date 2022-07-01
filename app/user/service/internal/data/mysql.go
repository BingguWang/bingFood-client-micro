package data

import "fmt"

type Mysql struct {
    Host         string `mapstructure:"host" json:"host" yaml:"host"`                             // 服务器地址
    Port         string `mapstructure:"port" json:"port" yaml:"port"`                             // 端口
    Config       string `mapstructure:"config" json:"config" yaml:"config"`                       // 高级配置
    Dbname       string `mapstructure:"db_name" json:"dbname" yaml:"db_name"`                     // 数据库名
    Username     string `mapstructure:"username" json:"username" yaml:"username"`                 // 数据库用户名
    Password     string `mapstructure:"password" json:"password" yaml:"password"`                 // 数据库密码
    MaxIdleConns int    `mapstructure:"max_idle_conns" json:"maxIdleConns" yaml:"max_idle_conns"` // 空闲中的最大连接数
    MaxOpenConns int    `mapstructure:"max_open_conns" json:"maxOpenConns" yaml:"max_open_conns"` // 打开到数据库的最大连接数
    LogMode      string `mapstructure:"log_mode" json:"logMode" yaml:"log_mode"`                  // 是否开启Gorm全局日志
    LogZap       bool   `mapstructure:"log_zap" json:"logZap" yaml:"log_zap"`                     // 是否通过zap写入日志文件
}

func (m *Mysql) Dsn() string {
    return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", m.Username, m.Password, m.Host, m.Port, m.Dbname)
}

