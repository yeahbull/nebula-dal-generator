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

package dataobject

type AppsDO struct {
     Id int32 `db:"Id"`
 ApiId int32 `db:"ApiId"`
 ApiHash string `db:"ApiHash"`
 Title string `db:"Title"`
 ShortName string `db:"ShortName"`
 CreatedAt string `db:"CreatedAt"`
 DeletedAt string `db:"DeletedAt"`

}
got:  map[app_id:{} api_hash:{} title:{} short_name:{}]
{id int32 { false} NO PRI { false} auto_increment select,insert,update,references }    title
{api_id int32 { false} NO UNI { false}  select,insert,update,references }    title
{api_hash string {utf8_general_ci true} NO  { false}  select,insert,update,references }    title
{title string {utf8_general_ci true} NO  { false}  select,insert,update,references }    title
{short_name string {utf8_general_ci true} NO  { false}  select,insert,update,references }    title
{created_at string { false} NO  {CURRENT_TIMESTAMP true}  select,insert,update,references }    title
{deleted_at string { false} NO  {0000-00-00 00:00:00 true}  select,insert,update,references }    title
&{title string {utf8_general_ci true} NO  { false}  select,insert,update,references }
{id int32 { false} NO PRI { false} auto_increment select,insert,update,references }    short_name
{api_id int32 { false} NO UNI { false}  select,insert,update,references }    short_name
{api_hash string {utf8_general_ci true} NO  { false}  select,insert,update,references }    short_name
{title string {utf8_general_ci true} NO  { false}  select,insert,update,references }    short_name
{short_name string {utf8_general_ci true} NO  { false}  select,insert,update,references }    short_name
{created_at string { false} NO  {CURRENT_TIMESTAMP true}  select,insert,update,references }    short_name
{deleted_at string { false} NO  {0000-00-00 00:00:00 true}  select,insert,update,references }    short_name
&{short_name string {utf8_general_ci true} NO  { false}  select,insert,update,references }
{id int32 { false} NO PRI { false} auto_increment select,insert,update,references }    app_id
{api_id int32 { false} NO UNI { false}  select,insert,update,references }    app_id
{api_hash string {utf8_general_ci true} NO  { false}  select,insert,update,references }    app_id
{title string {utf8_general_ci true} NO  { false}  select,insert,update,references }    app_id
{short_name string {utf8_general_ci true} NO  { false}  select,insert,update,references }    app_id
{created_at string { false} NO  {CURRENT_TIMESTAMP true}  select,insert,update,references }    app_id
{deleted_at string { false} NO  {0000-00-00 00:00:00 true}  select,insert,update,references }    app_id
<nil>
