package main

import (
	"encoding/json"

	"io/ioutil"
)

type Config struct {
	Time string `json:"Time"`
	IP   string `json:"IP"`
	Port string `json:"Port"`
	Db   string `json:"Db"`
	User string `json:"User"`
	Pass string `json:"Pass"`
	Sql  string `json:"Sql"`
}

var FileData []byte

func LoadConfig(file string) (*Config, error) {

	cfg := &Config{
		Time: "0 8-10 * * *",
		Sql:  "select count(1) from tb_o_sale",
	}

	var err error
	FileData, err = ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(FileData, &cfg)

	if err != nil {
		return nil, err
	}

	return cfg, nil

}
