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

package dalgen

import (
	"encoding/xml"
	"fmt"
	"github.com/golang/glog"
	"io/ioutil"
	"os"
)

const (
	RESULT_TYPE_DEFAULT = ""
	RESULT_TYPE_STRUCT  = "struct"
	RESULT_TYPE_MAP     = "map"
	RESULT_TYPE_INT     = "int"
	RESULT_TYPE_STRING  = "string"
)

const (
	RESULT_SET_DEFAULT = ""
	RESULT_SET_ONE     = "single"
	RESULT_SET_LIST    = "list"
)

type Coloum struct {
	ColName      string `xml:"name,attr"`
	ColType      string `xml:"type,attr"`
	ColVarName   string `xml:"varname,attr"`
	DefaultValue string `xml:"defaultValue,attr"`
}

type Operation struct {
	Name       string `xml:"name,attr"`
	ResultType string `xml:"result_type,attr"`
	ResultSet  string `xml:"result_set,attr"`
	Sql        string `xml:"sql"`
}

type DalgenConfig struct {
	FilePath string

	// 表名
	TableName string `xml:"sqlname,attr"`
	// 列
	Cols []Coloum `xml:"column"`
	// 操作
	Ops []Operation `xml:"operation"`
}

func NewDalgenConfig(configFile, xmlFileName string) (dalgen *DalgenConfig, err error) {
	file, err := os.Open(fmt.Sprintf("%s/tables/%s.xml", configFile, xmlFileName)) // For read access.
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		glog.Errorf("error: %v", err)
		return
	}
	v := &DalgenConfig{}
	err = xml.Unmarshal(data, v)
	if err != nil {
		glog.Errorf("error: %v", err)
		return
	}

	dalgen = v
	dalgen.FilePath = configFile
	return
}
