package main

import (
	"errors"
	"log"

	"strconv"

	"encoding/json"
	"io/ioutil"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

var db *xorm.Engine

type dbConfig struct {
	DBAddress  string `json:"db_address"`
	DBPort     int    `json:"db_port"`
	DBUser     string `json:"db_user"`
	DBPassword string `json:"db_passwd"`
	DBName     string `json:"db_name"`
	DBCharset  string `json:"charset"`
}

func (c dbConfig) String() string {
	return c.DBUser + ":" +
		c.DBPassword + "@tcp(" +
		c.DBAddress + ":" +
		strconv.Itoa(c.DBPort) + ")/" +
		c.DBName + "?charset=" +
		c.DBCharset + "&timeout=5s&parseTime=True&loc=Asia%2FChongqing"
}

func (c dbConfig) Check() error {
	if c.DBAddress == "" {
		return errors.New("invalid database address")
	}
	if c.DBCharset == "" {
		return errors.New("invalid database charset")
	}
	if c.DBName == "" {
		return errors.New("invalid database name")
	}
	if c.DBPassword == "" {
		return errors.New("invalid database password")
	}
	if c.DBPort > 65536 || c.DBPort < 1 {
		return errors.New("invalid database port")
	}
	if c.DBUser == "" {
		return errors.New("invalid database username")
	}
	return nil
}

func readConfigFormFile() (dbConfig, error) {
	bs, err := ioutil.ReadFile("db_config.json")
	if err != nil {
		log.Fatalln("read database config file db_config.json fail with error:", err)
	}
	var config dbConfig
	err = json.Unmarshal(bs, &config)
	if err != nil {
		log.Fatalln("read database config file db_config.json fail with error:", err)
	}
	if err := config.Check(); err == nil {
		return dbConfig{}, err
	}
	return config, nil
}

//DatabaseInit 初始化数据库连接
func DatabaseInit() {
	var err error
	config, err := readConfigFormFile()
	if err != nil {
		log.Fatalln(err)
	}
	db, err = xorm.NewEngine("mysql", config.String())
	if err != nil {
		log.Fatalln("connect to database fail with error:", err)
	}
	db.SetMapper(core.GonicMapper{})
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(100)
	db.Sync2(&User{})
}
