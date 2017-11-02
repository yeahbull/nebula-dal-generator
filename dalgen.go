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
	"log"
	_ "github.com/go-sql-driver/mysql" // import your used driver
)

func main() {
	dalgen, _ := dalgen2.NewDalgenConfig("./samples/dal/tables/apps.xml")

	db, err := sqlx.Connect("mysql", "root:@/nebulaim?charset=utf8")
	if err != nil {
		log.Fatalln(err)
	}

	s, _ := dalgen2.NewTableSchema(db, "apps")

	dalgen2.GenDataObject(s)
	dalgen2.GenDAO(dalgen, s)
}
