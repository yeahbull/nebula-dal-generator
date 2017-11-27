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
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"strings"
)

/*
	mysql> show full fields from apps;
	+------------+--------------+-----------------+------+-----+---------------------+----------------+---------------------------------+---------+
	| Field      | Type         | Collation       | Null | Key | Default             | Extra          | Privileges                      | Comment |
	+------------+--------------+-----------------+------+-----+---------------------+----------------+---------------------------------+---------+
	| id         | int(11)      | NULL            | NO   | PRI | NULL                | auto_increment | select,insert,update,references |         |
	+------------+--------------+-----------------+------+-----+---------------------+----------------+---------------------------------+---------+

	mysql> show  create  table  apps
	+-------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
	| Table | Create Table                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
	+-------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
	| apps  | CREATE TABLE `apps` (
	  `id` int(11) NOT NULL AUTO_INCREMENT,
	  `api_id` int(11) NOT NULL,
	  `api_hash` varchar(256) NOT NULL,
	  `title` varchar(128) NOT NULL,
	  `short_name` varchar(128) NOT NULL,
	  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	  `deleted_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
	  PRIMARY KEY (`id`),
	  UNIQUE KEY `api_id` (`api_id`)
	) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COMMENT='apps' |
	+-------+-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------+
*/

type FieldSchema struct {
	Field      string         `db:"Field"`
	Type       string         `db:"Type"`
	Collation  sql.NullString `db:"Collation"`
	Null       string         `db:"Null"`
	Key        string         `db:"Key"`
	Default    sql.NullString `db:"Default"`
	Extra      string         `db:"Extra"`
	Privileges string         `db:"Privileges"`
	Comment    string         `db:"Comment"`
}

type TableSchema struct {
	Name    string `db:"Table"`        //  表名
	Comment string `db:"Create Table"` //  字段信息
	Fields  []FieldSchema
}

func GetTableSchema(db *sqlx.DB, tableName string) string {
	type Table struct {
		Name      string `db:"Table"`        //  表名
		CreateDDL string `db:"Create Table"` //  字段信息
	}

	t := &Table{}
	err := db.Get(t, "show  create  table  apps")
	if err != nil {
		log.Println("1: ", err)
		return ""
	}

	// idx := strings. t.CreateDDL
	idx := strings.Index(t.CreateDDL, "COMMENT='")
	if idx == -1 {
		log.Println("2: ", idx)
		return ""
	}

	s := t.CreateDDL[idx+len("COMMENT='"):]
	idx = strings.Index(s, "'")
	// log.Println("3: ", idx, ", ", s)

	// log.Println(s[:idx])
	return s[:idx]
}

func NewTableSchema(db *sqlx.DB, dalgen *DalgenConfig) (*TableSchema, error) {
	comment := GetTableSchema(db, dalgen.TableName)

	fields := []FieldSchema{}

	err := db.Select(&fields, fmt.Sprintf("show full fields from %s", dalgen.TableName))
	if err != nil {
		log.Println("4: ", err)
		return nil, err
	}

	// 转成go类型
	for i, v := range fields {
		fields[i].Type = ToGoType(v.Type)
	}

	t := &TableSchema{}
	t.Name = dalgen.TableName
	t.Comment = comment
	t.Fields = fields

	// log.Println("t ==> ", t)

	return t, err
}

func (t *TableSchema) GetFieldSchema(fldName string) (fld *FieldSchema) {
	for i, v := range t.Fields {
		// fmt.Println(v, "  ", fldName)
		if v.Field == fldName {
			fld = &t.Fields[i]
			return
		}
	}
	return
}

func ToGoType(name string) (t string) {
	var s string
	idx := strings.Index(name, "(")
	if idx > 0 {
		s = name[:idx]
	} else {
		s = name
	}

	switch s {
	case "tinyint":
		t = "int8"
	case "int":
		t = "int32"
	case "bigint":
		t = "int64"
	case "varchar":
		t = "string"
	case "timestamp":
		t = "string"
	case "text":
		t = "string"
	case "blob":
		// fmt.Println("s ==> ", s)
		t = "[]byte"
	default:
		t = "string"
	}

	return t
}
