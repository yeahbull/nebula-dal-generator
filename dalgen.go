/*
 *  Copyright (c) 2017, https://github.com/nebulaim
 *  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	dalgen2 "github.com/nebulaim/nebula-dal-generator/dalgen"
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"flag"
	"github.com/BurntSushi/toml"
	"fmt"
	"github.com/golang/glog"
	"os"
)

var XmlFile = flag.String("xml", "", "Dalgen model xml config file name")
var IniFile = flag.String("config", "./dalgen.toml", "Dalgen config file name")

type Config struct {
	DalFilePath string
	MysqlDsn 	string
}

func init() {
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "false")
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func help() {
	fmt.Println("./dalgen --xml=model_name [--config=./dalgen.toml]")
}

func main() {
	flag.Parse()

	// TODO(@benqi): 没设置'--xml'选项，加载xml目录下所有配置
	if *XmlFile == "" {
		glog.Error("flags error, need flag: --xml=")
		help()
		return
	}

	// 载入配置文件
	config := Config{}
	if _, err := toml.DecodeFile(*IniFile, &config); err != nil {
		glog.Error("Load dalgen config ", *IniFile, ", error: ", err)
		return
	}

	// 检查目录是否存在
	if _, err := PathExists(config.DalFilePath); err != nil {
		glog.Error("dal filepath not exist: ", *IniFile, ", error: ", err)
		return
	}

	// TODO(@benqi): 检查dal_file_path目录下dataobject、dao和tables是否存在

	// 数据库
	db, err := sqlx.Connect("mysql", "root:@/nebulaim?charset=utf8")
	if err != nil {
		glog.Error("Connect database error: ", err)
		return
	}

	// 加载配置文件
	// xmlFilePath := config.DalFilePath + "/" + *XmlFile
	dalgenConfig, err := dalgen2.NewDalgenConfig(config.DalFilePath, *XmlFile)
	if err != nil {
		glog.Errorf("NewDalgenConfig - (%s, %s) error: %s", config.DalFilePath, *XmlFile, err)
		return
	}

	s, _ := dalgen2.NewTableSchema(db, dalgenConfig)
	dalgen2.GenDataObject(dalgenConfig, s)
	dalgen2.GenDAO(dalgenConfig, s)
}
