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
	"bytes"
	"fmt"
	"github.com/golang/glog"
	"github.com/xwb1989/sqlparser"
	"io/ioutil"
	"text/template"
	"github.com/jmoiron/sqlx"
)

//Id        int32		`db:"id"`
type TemplateField struct {
	Name  string
	Type  string
	DbMap string
}

type TemplateDO struct {
	Name   string
	Fields []TemplateField
}

func GenDataObject(dalgen *DalgenConfig, schema *TableSchema) {
	do := TemplateDO{}

	do.Name = ToCamel(schema.Name)
	for _, v := range schema.Fields {
		f := TemplateField{}
		f.Name = ToCamel(v.Field)
		f.Type = v.Type
		f.DbMap = fmt.Sprintf("`db:\"%s\"`", v.Field)
		do.Fields = append(do.Fields, f)
	}

	// glog.Info(do)
	t := template.Must(template.ParseFiles("./tpl/gen_do.tpl"))
	// fmt.Println(t.Name())
	// t.Execute(os.Stdout, do)

	var buf bytes.Buffer
	t.Execute(&buf, do)
	err := ioutil.WriteFile(fmt.Sprintf("./%s/dataobject/%s_do.go", dalgen.FilePath, schema.Name), buf.Bytes(), 0666)
	if err != nil {
		glog.Fatal("GenDAO error: ", err)
	}
}

//func (dao *AppsDAO) Insert(do *do.AppsDO) (id int32, err error) {
//	// TODO(@benqi): sqlmap
//	var sql = "INSERT INTO apps (api_id, api_hash, title, short_name) VALUES (:api_id, :api_hash, :title, :short_name)"
//	r, err := dao.db.NamedExec(sql, do)
//	if err != nil {
//		glog.Error("AppsDAO/Insert error: ", err)
//		return
//	}
//
//	id2, _ := r.LastInsertId()
//	id = int32(id2)
//	do.Id = id
//	return
//}

//type TemplateFunc struct {
//	Name string
//	Type string
//	DbMap string
//}

type Param struct {
	// TODO(@benqi): 驼峰参数名, 参数名从user_id ==> userId
	Name      string
	Type      string
	FieldName string
}

type Func struct {
	TableName string
	QueryType string

	FuncName string
	Params   []Param
	QueryParams   []Param
	Sql      string
	CompiledByNamedSql      string
	ParamHasList string
}

type TemplateDAO struct {
	Name  string
	Funcs []Func
	HasList string
}

func GenDAO(dalgen *DalgenConfig, schema *TableSchema) {
	dao := TemplateDAO{}
	dao.Name = ToCamel(schema.Name)
	dao.HasList = "false"
	// TODO(@benqi): 表名字段名等合法性检查

	for _, v := range dalgen.Ops {
		stmt, err := sqlparser.Parse(v.Sql)
		if err != nil {
			glog.Error("sql: ", v.Sql, ", error: ", err)
			// fmt.Println(err)
			// t.Error(err)
			continue
		}

		f := Func{}
		f.TableName = dao.Name
		f.FuncName = v.Name

		f.Sql = sqlparser.String(stmt)
		f.CompiledByNamedSql, _, err = compileNamedQuery([]byte(f.Sql), sqlx.QUESTION)
		if err != nil {
			glog.Error("compileNamedQuery: ", v.Sql, ", error: ", err)
			continue
		}

		switch stmt.(type) {
		case *sqlparser.Select:
			if v.ResultSet == "list" {
				f.QueryType = "SELECT_STRUCT_LIST"
			} else if v.ResultSet == "map"{
				f.QueryType = "SELECT_STRUCT_MAP"
			} else {
				f.QueryType = "SELECT_STRUCT_SINGLE"
			}
		case *sqlparser.Insert:
			f.QueryType = "INSERT"
		case *sqlparser.Update:
			f.QueryType = "UPDATE"
		case *sqlparser.Delete:
			f.QueryType = "DELETE"
		default:
			continue
		}

		got := GetBindvarList(stmt)
		// fmt.Println("got: ", got)
		for _, k := range got {
			// var isParam bool
			isParam := false
			// = false
			for _, p2 := range v.Parsms.OpParams {
				switch p2.ParamType {
				case PARAM_TYPE_INT32LIST, PARAM_TYPE_INT64LIST, PARAM_TYPE_STRINGLIST, PARAM_TYPE_UINT32LIST, PARAM_TYPE_UINT64LIST:
					dao.HasList = "true"
					f.ParamHasList = "true"
				}

				if p2.ParamName == k {
					p := Param{
						Name:      ToCamel(p2.ParamName),
						Type:      p2.ParamType,
						FieldName: p2.ParamName,
					}

					found := false
					for _, v := range f.Params {
						if v.Name == p.Name {
							found = true
						}
					}
					if !found {
						f.Params = append(f.Params, p)
					}

					f.QueryParams = append(f.QueryParams, p)

					isParam = true
					break
				}
			}

			if !isParam {
				fld := schema.GetFieldSchema(k)
				if fld != nil {
					// fmt.Println(fld)
					p := Param{
						Name:      ToCamel(fld.Field),
						Type:      fld.Type,
						FieldName: fld.Field,
					}

					found := false
					for _, v := range f.Params {
						if v.Name == p.Name {
							found = true
						}
					}
					if !found {
						f.Params = append(f.Params, p)
					}
					f.QueryParams = append(f.QueryParams, p)
				} else {
					glog.Errorf("Not find fld: %s", k)
				}
			}
		}
		// fmt.Printf("Process - func: [%s], query: [%s]\n", f.FuncName, f.Sql)
		dao.Funcs = append(dao.Funcs, f)
	}

	// f3, _ := file.Create(fmt.Sprintf("./%s_dao.go", dao.Name))

	var buf bytes.Buffer
	t := template.Must(template.ParseFiles("./tpl/gen_dao.tpl"))
	// fmt.Println(t.Name())
	// b := bytes.Buffer{}
	t.Execute(&buf, dao)
	err := ioutil.WriteFile(fmt.Sprintf("./%s/dao/mysql_dao/%s_dao.go", dalgen.FilePath, schema.Name), buf.Bytes(), 0666)
	if err != nil {
		glog.Fatal("GenDAO error: ", err)
	}
}
