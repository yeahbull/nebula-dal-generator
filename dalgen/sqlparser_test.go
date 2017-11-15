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
	"testing"
	"fmt"
	"github.com/jmoiron/sqlx"
)


func TestTSqlparser(t *testing.T) {
	var sql = "select id, access_hash, first_name, last_name, username from users where id in (:id1) and id=:id2"
	q, names, err := compileNamedQuery([]byte(sql), sqlx.QUESTION)
	fmt.Printf("q: %v, name: %v, err: %v\n", q, names, err)

	//stmt, _ := sqlparser.Parse(sql)
	//
	//fmt.Println(sqlparser.String(stmt))
	//// sqlparser.Parse()
	//
	//got := GetBindvars(stmt)
	//fmt.Println("got: ", got)
}
