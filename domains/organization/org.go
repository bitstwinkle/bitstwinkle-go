/*
 *  * Copyright (C) 2023 The Developer bitstwinkle
 *  *
 *  * Licensed under the Apache License, Version 2.0 (the "License");
 *  * you may not use this file except in compliance with the License.
 *  * You may obtain a copy of the License at
 *  *
 *  *      http://www.apache.org/licenses/LICENSE-2.0
 *  *
 *  * Unless required by applicable law or agreed to in writing, software
 *  * distributed under the License is distributed on an "AS IS" BASIS,
 *  * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  * See the License for the specific language governing permissions and
 *  * limitations under the License.
 */

package organization

import (
	"github.com/bitstwinkle/bitstwinkle-go/types/collections/more"
	"github.com/bitstwinkle/bitstwinkle-go/types/ctrl"
	"github.com/bitstwinkle/bitstwinkle-go/types/errors"
	"github.com/bitstwinkle/bitstwinkle-go/types/load"
	"github.com/bitstwinkle/bitstwinkle-go/types/ref"
)

type Organization struct {
	Scope ref.Scope `json:"scope"` //所属域
	Lead  ref.Lead  `json:"lead"`  //业务链接: 唯一
	ID    string    `json:"id"`    //组织ID
	Tree  struct {
		Stair  int      `json:"stair"`  //所属层级
		Parent string   `json:"parent"` //父亲ID
		Path   []string `json:"path"`   //全路径
	} `json:"tree"` //组织结构
	Name       string    `json:"name"`        //组织名称
	Alias      string    `json:"alias"`       //组织别名
	Logo       string    `json:"logo"`        //对应LOGO
	Memo       string    `json:"memo"`        //备注信息
	Info       more.More `json:"info"`        //更多信息
	Ctrl       ctrl.Ctrl `json:"ctrl"`        //控制信息
	BirthAt    string    `json:"birth_at"`    //创建时间
	ModifiedAt string    `json:"modified_at"` //最后更新时间
}

type Worker struct {
	Scope   ref.Scope `json:"scope"`  //所属域
	OrgID   string    `json:"org_id"` //所属组织ID
	OrgTree struct {
		Stair  int      `json:"stair"`  //所属层级
		Parent string   `json:"parent"` //父亲ID
		Path   []string `json:"path"`   //全路径
	} `json:"org_tree"` //组织结构
	OrgLeader  bool      `json:"org_leader"`  //是否该组织管理员,ORG唯一
	ID         string    `json:"id"`          //工作者ID
	UsrID      string    `json:"usr_id"`      //所对应的用户ID
	WorkNumb   string    `json:"work_numb"`   //工号,顶层组织下唯一
	WorkID     string    `json:"work_id"`     //工作用ID,例如 tony.zs
	WorkAlias  string    `json:"work_alias"`  //工作昵称
	Permission []string  `json:"permission"`  //顶层组织内权限码
	BirthAt    string    `json:"birth_at"`    //创建时间
	ModifiedAt string    `json:"modified_at"` //最后更新时间
	Info       more.More `json:"info"`        //更多信息
	Ctrl       ctrl.Ctrl `json:"ctrl"`        //控制信息
}

type OrgRegisterRequest struct {
	Scope    ref.Scope  `json:"scope"`           //[*]所属域
	Lead     ref.Lead   `json:"lead"`            //[*]业务链接
	ParentID string     `json:"parent_id"`       //[|]父亲组织ID,顶层组织使用$
	Name     string     `json:"name"`            //[*]组织名称
	Leader   string     `json:"leader"`          //[*]组织管理员,user.ID
	Alias    string     `json:"alias,omitempty"` //[-]别名
	Logo     string     `json:"logo,omitempty"`  //[-]LOGO
	Memo     string     `json:"memo,omitempty"`  //[-]备注信息
	Info     more.Array `json:"info,omitempty"`  //[-]更多信息
	Ctrl     ctrl.Ctrl  `json:"ctrl,omitempty"`  //[-]控制信息
}

type OrgGetRequest struct {
	By    load.ByCode `json:"by"`     //id|lead
	Scope ref.Scope   `json:"scope"`  //[*]所属业务域
	UsrID string      `json:"usr_id"` //[|]用户ID
	Lead  *ref.Lead   `json:"lead"`   //[|]业务链接
}

type OrgSetRequest struct {
	Scope   ref.Scope `json:"scope"`              //[*]所属域
	OrgID   string    `json:"org_id"`             //[*]组织ID
	InfoSet *more.Set `json:"info_set,omitempty"` //[|]展示信息
	CtrlSet *ctrl.Set `json:"ctrl_set,omitempty"` //[|]控制信息
}

type WorkerAddRequest struct {
	Scope      ref.Scope  `json:"scope"`          //[*]所属域
	OrgID      string     `json:"org_id"`         //[*]对应组织ID
	UsrID      string     `json:"usr_id"`         //[*]对应用户ID
	OrgLeader  bool       `json:"org_leader"`     //[-]是否组织管理员
	WorkNumb   string     `json:"work_numb"`      //[-]工号,顶层组织下唯一
	WorkID     string     `json:"work_id"`        //[-]工作用ID,例如 tony.zs
	WorkAlias  string     `json:"work_alias"`     //[-]工作昵称
	Permission []string   `json:"permission"`     //[-]顶层组织内权限码
	Info       more.Array `json:"info,omitempty"` //[-]更多信息
	Ctrl       ctrl.Ctrl  `json:"ctrl,omitempty"` //[-]控制信息
}

type WorkerGetRequest struct {
	By       load.ByCode `json:"by"`        //id|org&user
	Scope    ref.Scope   `json:"scope"`     //[*]所属业务域
	WorkerID string      `json:"worker_id"` //[id]工作者ID
	OrgID    string      `json:"org_id"`    //[org&user]组织ID
	UsrID    string      `json:"usr_id"`    //[org&user]USER ID
}

type WorkerSetRequest struct {
	Scope         ref.Scope `json:"scope"`              //[*]所属域
	Worker        string    `json:"worker"`             //[*]对应工作者ID
	InfoSet       *more.Set `json:"info_set,omitempty"` //[|]展示信息
	Permission    []string  `json:"permission"`         //[-]需要新增或更新的权限
	RmvPermission []string  `json:"rmv_permission"`     //[-]需要删除的权限
	CtrlSet       *ctrl.Set `json:"ctrl_set,omitempty"` //控制信息
}

type Service interface {
	OrgRegister(req OrgRegisterRequest) (orgID string, err *errors.Error)
	OrgGet(req OrgGetRequest) (*Organization, *errors.Error)
	OrgSet(req OrgSetRequest) *errors.Error
	OrgSwitch(scope ref.Scope, orgID string, available bool) *errors.Error

	WorkerAdd(req WorkerAddRequest) (workerID string, err *errors.Error)
	WorkerGet(req WorkerGetRequest) (*Worker, *errors.Error)
	WorkerSet(req WorkerSetRequest) (workerID string, err *errors.Error)
	WorkerSwitch(scope ref.Scope, workerID string, available bool) *errors.Error
}
