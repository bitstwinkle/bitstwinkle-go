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
	"github.com/bitstwinkle/bitstwinkle-go/types/view/media"
	"time"
)

type Organization struct {
	Scope ref.Scope `bson:"scope" json:"scope"` //所属域
	Lead  ref.Lead  `bson:"lead" json:"lead"`   //业务链接: 唯一
	ID    string    `bson:"id" json:"id"`       //组织ID
	Tree  struct {
		Stair  int      `bson:"stair" json:"stair"`   //所属层级
		Parent string   `bson:"parent" json:"parent"` //父亲ID
		Path   []string `bson:"path" json:"path"`     //全路径
	} `bson:"tree" json:"tree"` //组织结构
	Title      string     `bson:"title" json:"title"`              //组织名称
	Info       more.More  `bson:"info" json:"info"`                //更多信息
	Media      media.More `bson:"media" bson:"media" json:"media"` //媒体信息
	Available  bool       `bson:"available" json:"available"`      //是否可用
	Ctrl       *ctrl.Ctrl `bson:"ctrl" json:"ctrl"`                //控制信息
	BirthAt    time.Time  `bson:"birth_at" json:"birth_at"`        //创建时间
	ModifiedAt time.Time  `bson:"modified_at" json:"modified_at"`  //最后更新时间
}

type Worker struct {
	Scope   ref.Scope `bson:"scope" json:"scope"`   //所属域
	OrgID   string    `bson:"org_id" json:"org_id"` //所属组织ID
	OrgTree struct {
		Stair  int      `bson:"stair" json:"stair"`   //所属层级
		Parent string   `bson:"parent" json:"parent"` //父亲ID
		Path   []string `bson:"path" json:"path"`     //全路径
	} `json:"org_tree"` //组织结构
	OrgLeader  bool       `bson:"org_leader" json:"org_leader"`   //是否该组织管理员,ORG唯一
	ID         string     `bson:"id" json:"id"`                   //工作者ID
	UsrID      string     `bson:"usr_id" json:"usr_id"`           //所对应的用户ID
	WorkNumb   string     `bson:"work_numb" json:"work_numb"`     //工号,顶层组织下唯一
	WorkID     string     `bson:"work_id" json:"work_id"`         //工作用ID,例如 tony.zs
	WorkAlias  string     `bson:"work_alias" json:"work_alias"`   //工作昵称
	Permission []string   `bson:"permission" json:"permission"`   //顶层组织内权限码
	Info       more.More  `bson:"info" json:"info"`               //更多信息
	Media      media.More `bson:"media" json:"media"`             //媒体信息
	Available  bool       `bson:"available" json:"available"`     //是否可用
	Ctrl       *ctrl.Ctrl `bson:"ctrl" json:"ctrl"`               //控制信息
	BirthAt    string     `bson:"birth_at" json:"birth_at"`       //创建时间
	ModifiedAt string     `bson:"modified_at" json:"modified_at"` //最后更新时间
}

type OrgRegisterRequest struct {
	Scope    ref.Scope `json:"scope"`     //[*]所属域
	Lead     ref.Lead  `json:"lead"`      //[*]业务链接
	ParentID string    `json:"parent_id"` //[|]父亲组织ID,顶层组织使用$
	Title    string    `json:"title"`     //[*]组织名称
	Leader   string    `json:"leader"`    //[*]组织管理员,user.ID
	Info     *struct {
		Alias string     `bson:"alias" json:"alias"`
		Intro string     `bson:"intro" json:"intro"`
		More  more.Array `bson:"more" json:"more"`
	} `bson:"info" json:"info"`
	Media *struct {
		Logo    *media.Media `bson:"logo" json:"logo"`
		Primary *media.Media `bson:"primary" json:"primary"`
		More    media.Array  `bson:"more" json:"more"`
	} `bson:"media" json:"media"`
	Ctrl *ctrl.Ctrl `json:"ctrl"` //[-]控制信息
}

type OrgGetRequest struct {
	By    load.ByCode `bson:"by" json:"by"`         //id|lead
	Scope ref.Scope   `bson:"scope" json:"scope"`   //[*]所属业务域
	OrgID string      `bson:"org_id" json:"org_id"` //[|]ID
	Lead  *ref.Lead   `bson:"lead" json:"lead"`     //[|]业务链接
}

type OrgSetRequest struct {
	Scope        ref.Scope        `bson:"scope" json:"scope"`                 //[*]所属域
	OrgID        string           `bson:"org_id" json:"org_id"`               //[*]组织ID
	InfoSet      *more.Set        `bson:"info_set" json:"info_set"`           //[|]展示信息
	AvailableSet *ctrl.BooleanSet `bson:"available_set" json:"available_set"` //是否生效
	CtrlSet      *ctrl.Set        `bson:"ctrl_set" json:"ctrl_set"`           //[|]控制信息
}

type OrgLoadRequest struct {
	ParentIDArray []string   `bson:"parent_id_array" json:"parent_id_array"`
	LeadArray     []ref.Lead `bson:"lead_array" json:"lead_array"`
	IDArray       []string   `bson:"id_array" json:"id_array"`
	Page          *load.Page `bson:"page" json:"page"`
}

type WorkerAddRequest struct {
	Scope      ref.Scope  `json:"scope"`      //[*]所属域
	OrgID      string     `json:"org_id"`     //[*]对应组织ID
	UsrID      string     `json:"usr_id"`     //[*]对应用户ID
	OrgLeader  bool       `json:"org_leader"` //[-]是否组织管理员
	WorkNumb   string     `json:"work_numb"`  //[-]工号,顶层组织下唯一
	WorkID     string     `json:"work_id"`    //[-]工作用ID,例如 tony.zs
	WorkAlias  string     `json:"work_alias"` //[-]工作昵称
	Permission []string   `json:"permission"` //[-]顶层组织内权限码
	Info       more.Array `json:"info"`       //[-]更多信息
	Ctrl       *ctrl.Ctrl `json:"ctrl"`       //[-]控制信息
}

type WorkerGetRequest struct {
	By       load.ByCode `json:"by"`        //id|org&user
	Scope    ref.Scope   `json:"scope"`     //[*]所属业务域
	WorkerID string      `json:"worker_id"` //[id]工作者ID
	OrgID    string      `json:"org_id"`    //[org&user]组织ID
	UsrID    string      `json:"usr_id"`    //[org&user]USER ID
}

type WorkerSetRequest struct {
	Scope         ref.Scope  `bson:"scope" json:"scope"`       //[*]所属域
	Worker        string     `bson:"worker" json:"worker"`     //[*]对应工作者ID
	InfoSet       *more.Set  `bson:"info_set" json:"info_set"` //[|]展示信息
	MediaSet      *media.Set `bson:"media_set" json:"media_set"`
	PermissionSet *struct {
		Yes        bool     `bson:"yes" json:"yes"`               //是否设置
		Permission []string `bson:"permission" json:"permission"` //[-]需要新增或更新的权限
		Removed    []string `bson:"removed" json:"removed"`       //[-]需要删除的权限
	} `bson:"permission_set" json:"permission_set"`
	AvailableSet *ctrl.BooleanSet `bson:"available_set" json:"available_set"` //是否生效
	CtrlSet      *ctrl.Set        `json:"ctrl_set,omitempty"`                 //控制信息
}

type WorkerLoadRequest struct {
	OrgIDArray  []string   `bson:"org_id_array" json:"org_id_array"`
	IDArray     []string   `bson:"id_array" json:"id_array"`
	UserIDArray []string   `bson:"user_id_array" json:"user_id_array"`
	Page        *load.Page `bson:"page" json:"page"`
}

type Service interface {
	OrgRegister(req OrgRegisterRequest) (orgID string, err *errors.Error)
	OrgGet(req OrgGetRequest) (*Organization, *errors.Error)
	OrgSet(req OrgSetRequest) *errors.Error
	OrgLoad(req OrgLoadRequest) ([]Organization, *load.Paging, *errors.Error)

	WorkerAdd(req WorkerAddRequest) (workerID string, err *errors.Error)
	WorkerGet(req WorkerGetRequest) (*Worker, *errors.Error)
	WorkerSet(req WorkerSetRequest) (workerID string, err *errors.Error)
	WorkerLoad(req WorkerLoadRequest) ([]Worker, *load.Paging, *errors.Error)
}
