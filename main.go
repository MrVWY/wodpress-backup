package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os/exec"
)

// wp export --dir=/tmp/ --user=Hayley --post_type=post --start_date=2020-06-01 --end_date=2020-06-05
//--allow-root

type Config struct {
	Dir       string `yaml:"Dir"`
	User      string `yaml:"User"`
	PostType  string `yaml:"Post_type"`
	StartDate string `yaml:"Start_date"`
	EndDate   string `yaml:"End_date"`
}

func WbInitConf(conf Config) {
	date, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		panic("get file date is fail")
	}
	err = yaml.Unmarshal(date, &conf)
	if err != nil {
		panic("Init conf is fail")
	}
	WbCLI(conf)
}

func WbCLI(conf Config) {
	//cmd := exec.Command("wp", "export", "--dir= ", "--user= ", "--post_type= ", "--start_date= ", "--end_date= ", "--allow-root")
	cmd := exec.Command("wp", "export", fmt.Sprintf("--dir=%s", conf.Dir),
		fmt.Sprintf("--user=%s", conf.User),
		fmt.Sprintf("--post_type=%s", conf.PostType),
		fmt.Sprintf("--start_date%s", conf.StartDate),
		fmt.Sprintf("--end_date=%s", conf.EndDate), "--allow-root")
	err := cmd.Run()
	if err != nil {
		panic("run the wp export CLI is fail")
	}
}

func main() {
	var conf Config
	WbCLI(conf)
}
