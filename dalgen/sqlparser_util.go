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
	"github.com/xwb1989/sqlparser"
)

// GetBindvars returns a map of the bind vars referenced in the statement.
// TODO(sougou); This function gets called again from vtgate/planbuilder.
// Ideally, this should be done only once.
func GetBindvarList(stmt sqlparser.Statement) []string {
	bindvars := []string{}
	_ = sqlparser.Walk(func(node sqlparser.SQLNode) (kontinue bool, err error) {
		switch node := node.(type) {
		case *sqlparser.SQLVal:
			if node.Type == sqlparser.ValArg {
				bindvars = append(bindvars, string(node.Val[1:]))
			}
		case sqlparser.ListArg:
			bindvars = append(bindvars, string(node[2:]))
		}
		return true, nil
	}, stmt)
	return bindvars
}
