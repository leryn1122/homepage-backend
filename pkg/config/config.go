package config

import (
	"flag"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

type Config struct {
	ServerConfig ServerConfig `toml:"server"`
	DbConfig     DbConfig     `toml:"database"`
}

type ServerConfig struct {
	Name string `toml:"name"`
	Host string `toml:"host"`
	Port int    `toml:"port"`
}

type DbConfig struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Scheme   string `toml:"scheme"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	Driver   string `toml:"driver"`
}

var Configuration *Config

func init() {
	filename := flag.String("c", "etc/config.toml", "Path to the config")
	flag.Parse()
	logrus.Infof("config file located at: %s", *filename)

	var content []byte
	content, err := ioutil.ReadFile(*filename)
	if err != nil {
		logrus.Fatalf("fail to read configuration: %v", err)
		panic(err)
	}

	content = []byte(os.ExpandEnv(string(content)))
	config := &Config{}
	err = toml.Unmarshal(content, config)
	if err != nil {
		logrus.Fatalf("fail to unmarshal yaml: %v", err)
	}
	Configuration = config
}
