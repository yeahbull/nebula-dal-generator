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

/*
import (
	"github.com/xwb1989/sqlparser"
)

type Section struct {
	Columns map[string]map[string][]*sqlparser.ColName // map[table]map[column][]{ColName}
	Tables  map[string][]struct {                      // map[table][]{TableName/Subquery}
		Table    *sqlparser.TableName
		Subquery *Section
	}
}

func NewSection() *Section {
	return &Section{
		Columns: make(map[string]map[string][]*sqlparser.ColName),
		Tables: make(map[string][]struct {
			Table    *sqlparser.TableName
			Subquery *Section
		}),
	}
}

func (s *Section) tableColumns(table string) map[string][]*sqlparser.ColName {
	cols, has := s.Columns[table]
	if !has {
		cols = make(map[string][]*sqlparser.ColName)
		s.Columns[table] = cols
	}

	return cols
}

func (s *Section) AddColumn(table string, col *sqlparser.ColName) {
	tabName := string(col.Qualifier)
	if tabName == "" {
		tabName = table
	}

	cols := s.tableColumns(tabName)
	if colName := string(col.Name); colName != "?" {
		cols[colName] = append(cols[colName], col)
	}
}

func (s *Section) AddTableName(as string, tab *sqlparser.TableName) string {
	if as == "" {
		as = string(tab.Name)
	}
	s.Tables[as] = append(s.Tables[as], struct {
		Table    *sqlparser.TableName
		Subquery *Section
	}{
		Table: tab,
	})

	return as
}

func (s *Section) AddSubquery(as string) *Section {
	newsec := NewSection()
	s.Tables[as] = append(s.Tables[as], struct {
		Table    *sqlparser.TableName
		Subquery *Section
	}{
		Subquery: newsec,
	})

	return newsec
}
*/

/*
func (s *Section) Inspect(node sqlparser.SQLNode) {
	var table string
	sqlparser.Inspect(node, func(node sqlparser.SQLNode) bool {
		switch node := node.(type) {
		case *sqlparser.TableName:
			table = s.AddTableName("", node)
		case *sqlparser.AliasedTableExpr:
			switch tab := node.Expr.(type) {
			case *sqlparser.TableName:
				table = s.AddTableName(string(node.As), tab)
			case *sqlparser.Subquery:
				s.AddSubquery(string(node.As)).Inspect(tab)
			}

			return false // stop inspect subnodes
		case *sqlparser.ColName:
			s.AddColumn(table, node)
		}

		fmt.Printf("node: %v\n", node)

		return true
	})
}

func (s *Section) modelTable(v Visitor, tab *sqlparser.TableName) (*Table, error) {
	tabname := string(tab.Name)
	model := v.Models[tabname]
	if model == nil {
		return nil, fmt.Errorf("model %s hasn't been registered", tabname)
	}

	return model, nil
}

func (s *Section) replace(v Visitor) error {
	for tabalias, cols := range s.Columns {
		tabnodes, has := s.Tables[tabalias]
		if !has {
			return fmt.Errorf("table alias %s not found in sql", tabalias)
		}

		tabnode := tabnodes[0]
		if tabnode.Table == nil {
			// don't replace columns of subquery section because of alias
			continue
		}

		model, err := s.modelTable(v, tabnode.Table)
		if err != nil {
			return err
		}

		for field, cols2 := range cols {
			for _, col := range cols2 {
				col.Name = []byte(model.Fields.DefGet(field, "").(string))
			}
		}
	}

	for _, tabnodes := range s.Tables {
		for _, tabnode := range tabnodes {
			if tabnode.Table != nil {
				if string(tabnode.Table.Name) == "DUAL" {
					continue
				}

				model, err := s.modelTable(v, tabnode.Table)
				if err != nil {
					return err
				}

				tabnode.Table.Name = []byte(model.Name)
			} else {
				if err := tabnode.Subquery.replace(v); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

*/
