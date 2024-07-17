package main

var (
	// 版本信息
	appVersion bool // 控制是否显示版本
	APPVersion = "v0.0.2"
	BuildTime  = "2006-01-02 15:04:05"
	GitCommit  = "xxxxxxxxxxx"
	ConfigFile = "config.yaml"
	config     *Config
)

type Log struct {
	Level string `yaml:"level"`
	File  string `yaml:"file"`
}
type Harbor struct {
	URL      string `yaml:"url"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type DingRobot struct {
	AccessToken string `yaml:"accessToken"` // 钉钉机器人token
	Secret      string `yaml:"secret"`      // 钉钉机器人secret
}

type Config struct {
	Log       Log       `yaml:"log"`
	Harbor    Harbor    `yaml:"harbor"`
	DingRobot DingRobot `yaml:"dingRobot"`
	BeforTime string    `yaml:"beforTime"`
}
