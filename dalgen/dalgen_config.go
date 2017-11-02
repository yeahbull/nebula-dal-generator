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
	"os"
	"fmt"
	"io/ioutil"
)

type Coloum struct {
	ColName 		string	`xml:"name,attr"`
	ColType 		string	`xml:"type,attr"`
	ColVarName 		string	`xml:"varname,attr"`
	DefaultValue 	string	`xml:"defaultValue,attr"`
}

type Operation struct {
	Name string		`xml:"name,attr"`
	Sql  string		`xml:"sql"`
}

type DalgenConfig struct {
	// XMLName xml.Name 	`xml:"table"`
	// 表名
	TableName string	`xml:"sqlname,attr"`
	// 列
	Cols []Coloum		`xml:"column"`
	// 操作
	Ops  []Operation	`xml:"operation"`
}

func NewDalgenConfig(configFile string) (dalgen *DalgenConfig, err error) {
	file, err := os.Open(configFile) // For read access.
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	v := &DalgenConfig{}
	err = xml.Unmarshal(data, v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	dalgen = v
	return
}

//func (d* DalgenConfig) ToDataObject() string {
//	return ""
//}
//
//func (d* DalgenConfig) ToDAO() string {
//	return ""
//}
//
