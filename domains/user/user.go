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

package user

import (
	"github.com/bitstwinkle/bitstwinkle-go/types/collections/more"
	"github.com/bitstwinkle/bitstwinkle-go/types/ctrl"
	"github.com/bitstwinkle/bitstwinkle-go/types/errors"
	"github.com/bitstwinkle/bitstwinkle-go/types/load"
	"github.com/bitstwinkle/bitstwinkle-go/types/ref"
)

type User struct {
	ID       string `json:"id"`               //用户ID
	Nickname string `json:"nickname"`         //用户昵称
	Name     string `json:"name,omitempty"`   //真实姓名
	Avatar   string `json:"avatar,omitempty"` //头像
	Gender   string `json:"gender,omitempty"` //[-]性别 F女|M男|U未知
	Passport map[string]struct {
		Code   string `json:"code"`    //认证通道编码
		ID     string `json:"id"`      //该通道内的ID
		LastAt string `json:"last_at"` //登记时间
	} `json:"passport"` //护照列表
	Info       more.More `json:"info"`        //更多信息
	Ctrl       ctrl.Ctrl `json:"ctrl"`        //控制信息
	BirthAt    string    `bson:"birth_at"`    //创建时间
	ModifiedAt string    `bson:"modified_at"` //最后修改时间
}

type PassportRequest struct {
	Code string     `json:"code"` //认证通道编码
	ID   string     `json:"id"`   //该通道内的ID
	More more.Array `json:"more"` //通道参数
}

type RegisterRequest struct {
	IdemID   string    `json:"idem_id"`          //[*]幂等ID
	Scope    ref.Scope `json:"scope"`            //[*]所属业务域
	Nickname string    `json:"nickname"`         //[*]昵称
	Name     string    `json:"name,omitempty"`   //[-]名称
	Avatar   string    `json:"avatar,omitempty"` //[-]头像
	Gender   string    `json:"gender,omitempty"` //[-]性别 F|M|U
	Passport struct {
		Code string     `json:"code"` //认证通道编码
		ID   string     `json:"id"`   //该通道内的ID
		More more.Array `json:"more"` //通道参数
	} `json:"passport"` //绑定首个通道信息
	Info more.More `json:"more"` //更多信息
	Ctrl ctrl.Ctrl `json:"ctrl"` //控制信息
}

type GetRequest struct {
	By       load.ByCode `json:"by"`     //id|passport
	Scope    ref.Scope   `json:"scope"`  //[*]所属业务域
	UsrID    string      `json:"usr_id"` //[|]用户ID
	Passport *struct {
		Code string `json:"code"` //认证通道编码
		ID   string `json:"id"`   //该通道内的ID
	} //[|]通道信息
}

type BindPassportRequest struct {
	Scope    ref.Scope `json:"scope"`  //[*]所属业务域
	UsrID    string    `json:"usr_id"` //[*]用户ID
	Passport struct {
		Code string     `json:"code"` //认证通道编码
		ID   string     `json:"id"`   //该通道内的ID
		More more.Array `json:"more"` //通道参数
	} `json:"passport"` //绑定首个通道信息
}

type UnBindPassportRequest struct {
	Scope    ref.Scope `json:"scope"`  //[*]所属业务域
	UsrID    string    `json:"usr_id"` //[*]用户ID
	Passport struct {
		Code string `json:"code"` //认证通道编码
		ID   string `json:"id"`   //该通道内的ID
	} `json:"passport"` //绑定首个通道信息
}

type SetRequest struct {
	Scope   ref.Scope `json:"scope"`              //[*]所属业务域
	UsrID   string    `json:"usr_id"`             //[*]用户ID
	InfoSet *more.Set `json:"info_set,omitempty"` //[|]展示信息
	CtrlSet *ctrl.Set `json:"ctrl_set,omitempty"` //[|]控制信息
}

type LoadRequest struct {
	Scope    *ref.Scope `json:"scope,omitempty"` //[*]所属域
	Passport *struct {
		Code string `json:"code"` //认证通道编码
	} `json:"passport,omitempty"` //绑定的通道信息
	Tag     []string  `json:"tag,omitempty"`      //控制标
	IDArray []string  `json:"id_array,omitempty"` //ID列表
	Keyword string    `json:"keyword,omitempty"`  //关键词
	Page    load.Page `json:"page"`
}

type Service interface {
	Register(req RegisterRequest) (usrID string, err *errors.Error)
	Get(req GetRequest) (*User, *errors.Error)
	BindPassport(req BindPassportRequest) (usrID string, err *errors.Error)
	UnBindPassport(req UnBindPassportRequest) (usrID string, err *errors.Error)
	Set(req SetRequest) (usrID string, err *errors.Error)
	Load(req LoadRequest) ([]User, load.Pagination, *errors.Error)
}
