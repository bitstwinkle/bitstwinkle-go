/*
 *
 *  *  * Copyright (C) 2023 The Developer bitstwinkle
 *  *  *
 *  *  * Licensed under the Apache License, Version 2.0 (the "License");
 *  *  * you may not use this file except in compliance with the License.
 *  *  * You may obtain a copy of the License at
 *  *  *
 *  *  *      http://www.apache.org/licenses/LICENSE-2.0
 *  *  *
 *  *  * Unless required by applicable law or agreed to in writing, software
 *  *  * distributed under the License is distributed on an "AS IS" BASIS,
 *  *  * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  *  * See the License for the specific language governing permissions and
 *  *  * limitations under the License.

 */

package pickup

import (
	"github.com/bitstwinkle/bitstwinkle-go/domains/address"
	"github.com/bitstwinkle/bitstwinkle-go/types/collections/more"
	"github.com/bitstwinkle/bitstwinkle-go/types/ctrl"
	"github.com/bitstwinkle/bitstwinkle-go/types/errors"
	"github.com/bitstwinkle/bitstwinkle-go/types/load"
	"github.com/bitstwinkle/bitstwinkle-go/types/location"
	"github.com/bitstwinkle/bitstwinkle-go/types/ref"
	"github.com/bitstwinkle/bitstwinkle-go/types/times"
)

const POINT = "PICK_UP_POINT"

type Point struct {
	Scope        ref.Scope           `bson:"scope" json:"scope"`                 //所属业务域
	Lead         ref.Lead            `bson:"lead" json:"lead"`                   //业务领衔
	ID           string              `bson:"id" json:"id"`                       //自提点ID
	Title        string              `bson:"title" json:"title"`                 //自提点名称
	OpeningHours *times.OpeningHours `bson:"opening_hours" json:"opening_hours"` //营业时间设置
	Info         more.More           `bson:"info" json:"info"`                   //自提点更多信息
	Ctrl         ctrl.Ctrl           `bson:"ctrl" json:"ctrl"`                   //控制信息
	Available    bool                `bson:"available" json:"available"`         //是否可用
	AddressID    address.Address     `bson:"address_id" json:"address_id"`       //对应地址ID
}

type PointCreateRequest struct {
	IdemID       string                      `bson:"idem_id" json:"idem_id"`             //[*]幂等ID
	Scope        ref.Scope                   `bson:"scope" json:"scope"`                 //所属业务域
	Lead         ref.Lead                    `bson:"lead" json:"lead"`                   //业务领衔
	Title        string                      `bson:"title" json:"title"`                 //自提点名称
	OpeningHours *times.OpeningHours         `bson:"opening_hours" json:"opening_hours"` //营业时间设置
	Info         more.Array                  `bson:"info" json:"info"`                   //自提点更多信息
	Ctrl         ctrl.Ctrl                   `bson:"ctrl" json:"ctrl"`                   //控制信息
	Address      address.AddrRegisterRequest `bson:"address" json:"address"`             //对应地址请求
}

type PointSetRequest struct {
	Scope        ref.Scope       `bson:"scope" json:"scope"`         //所属业务域
	Lead         *ref.Lead       `bson:"lead" json:"lead"`           //[id|lead]业务领衔
	PointID      string          `bson:"point_id" json:"point_id"`   //[id|lead]自提点ID
	TitleSet     *ctrl.StringSet `bson:"title_set" json:"title_set"` //标题
	OpeningHours *struct {
		Yes   bool                `bson:"yes" json:"yes"` //是否设置营业时间
		Value *times.OpeningHours `bson:"value" json:"value"`
	} `bson:"opening_hours" json:"opening_hours"` //营业时间设置
	InfoSet      *more.Set        `bson:"info_set" json:"info_set"`           //设置展示信息
	AvailableSet *ctrl.BooleanSet `bson:"available_set" json:"available_set"` //是否可用设置
	CtrlSet      *ctrl.Set        `bson:"ctrl_set" json:"ctrl_set"`           //设置控制信息
}

// PointBindAddressRequest 首次设置地址
type PointBindAddressRequest struct {
	IdemID  string                      `bson:"idem_id" json:"idem_id"`   //[*]幂等ID
	Scope   ref.Scope                   `bson:"scope" json:"scope"`       //所属业务域
	Lead    *ref.Lead                   `bson:"lead" json:"lead"`         //[id|lead]业务领衔
	PointID string                      `bson:"point_id" json:"point_id"` //[id|lead]对应自提点ID
	Request address.AddrRegisterRequest `bson:"request" json:"request"`   //地址创建请求
}

type PointGetRequest struct {
	Scope   ref.Scope `bson:"scope" json:"scope"`       //所属业务域
	Lead    *ref.Lead `bson:"lead" json:"lead"`         //[id|lead]业务领衔
	PointID string    `bson:"point_id" json:"point_id"` //[id|lead]对应自提点ID
}

type PointLoadRequest struct {
	Scope        ref.Scope     `bson:"scope" json:"scope"`                   //所属业务域
	LeadArray    []ref.Lead    `bson:"lead_array" json:"lead_array"`         //[id|lead]业务领衔
	PointIDArray []string      `bson:"point_id_array" json:"point_id_array"` //[id|lead]对应自提点ID
	Loc          *location.Loc `bson:"loc" json:"loc"`                       //根据坐标找最近
	Page         *load.Page    `bson:"page" json:"page"`                     //分页信息
}

type Service interface {
	PointCreate(req PointCreateRequest) (*Point, *errors.Error)
	PointSet(req PointSetRequest) (*Point, *errors.Error)
	PointGet(req PointGetRequest) (*Point, *errors.Error)
	PointBindAddress(req PointBindAddressRequest) (*address.Address, *errors.Error)
	PointLoad(req PointLoadRequest) ([]*Point, load.Paging, *errors.Error)
}
