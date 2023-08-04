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

package mall

import (
	"github.com/bitstwinkle/bitstwinkle-go/domains/address"
	"github.com/bitstwinkle/bitstwinkle-go/types/collections/more"
	"github.com/bitstwinkle/bitstwinkle-go/types/ctrl"
	"github.com/bitstwinkle/bitstwinkle-go/types/errors"
	"github.com/bitstwinkle/bitstwinkle-go/types/load"
	"github.com/bitstwinkle/bitstwinkle-go/types/location"
	"github.com/bitstwinkle/bitstwinkle-go/types/ref"
	"github.com/bitstwinkle/bitstwinkle-go/types/times"
	"github.com/bitstwinkle/bitstwinkle-go/types/view/label"
	"github.com/bitstwinkle/bitstwinkle-go/types/view/media"
	"time"
)

const MALL = "MALL"

type ID = string

type Mall struct {
	Scope        ref.Scope           `bson:"scope" json:"scope"`                 //所属业务域
	Lead         ref.Lead            `bson:"lead" json:"lead"`                   //业务领衔
	ID           ID                  `bson:"id" json:"id"`                       //MLL ID
	Title        string              `bson:"title" json:"title"`                 //MLL 名称
	OpeningHours *times.OpeningHours `bson:"opening_hours" json:"opening_hours"` //营业时间设置
	Info         more.More           `bson:"info" json:"info"`                   //更多信息
	Media        media.More          `bson:"media" json:"media"`                 //更多媒体信息
	Label        label.Array         `bson:"label" json:"label"`                 //标签
	Ctrl         *ctrl.Ctrl          `bson:"ctrl" json:"ctrl"`                   //控制参数
	Available    bool                `bson:"available" json:"available"`         //是否可用
	BrithAt      time.Time           `bson:"brith_at" json:"brith_at"`           //创建时间
	ModifiedAt   time.Time           `bson:"modified_at" json:"modified_at"`     //最后修改时间

	Address *address.Address `bson:"address" json:"address"` //地址信息
}

type CreateRequest struct {
	IdemID       string              `bson:"idem_id" json:"idem_id"`             //[*]幂等ID
	Scope        ref.Scope           `bson:"scope" json:"scope"`                 //所属业务域
	Lead         *ref.Lead           `bson:"lead" json:"lead"`                   //业务领衔
	Title        string              `bson:"title" json:"title"`                 //MALL名称
	OpeningHours *times.OpeningHours `bson:"opening_hours" json:"opening_hours"` //营业时间设置
	Info         *struct {
		Alias string    `bson:"alias" json:"alias"` //简称
		Intro string    `bson:"intro" json:"intro"` //简介
		More  more.More `bson:"more" json:"more"`   //更多信息
	} `bson:"info" json:"info"` //更多信息
	Media *struct {
		Logo    *media.Media `bson:"logo" json:"logo"`       //LOGO
		Primary *media.Media `bson:"primary" json:"primary"` //主图
		More    media.More   `bson:"media" json:"media"`     //更多媒体信息
	} `bson:"media" json:"media"` //更多信息
	Label label.Array `bson:"label" json:"label,omitempty"` //标签
	Ctrl  *ctrl.Ctrl  `bson:"ctrl" json:"ctrl"`             //控制参数
}

type SetRequest struct {
	Scope        ref.Scope       `bson:"scope" json:"scope"`         //所属业务域
	Lead         *ref.Lead       `bson:"lead" json:"lead"`           //业务领衔
	MallID       ID              `bson:"mall_id" json:"mall_id"`     //对应MALL ID
	TitleSet     *ctrl.StringSet `bson:"title_set" json:"title_set"` //标题
	OpeningHours *struct {
		Yes   bool                `bson:"yes" json:"yes"` //是否设置营业时间
		Value *times.OpeningHours `bson:"value" json:"value"`
	} `bson:"opening_hours" json:"opening_hours"` //营业时间设置
	InfoSet      *more.Set        `bson:"info_set" json:"info_set"`           //设置展示信息
	MediaSet     *media.Set       `bson:"media_set" json:"media_set"`         //设置媒体信息
	LabelSet     *label.Set       `bson:"label_set" json:"label_set"`         //设置标签信息
	AvailableSet *ctrl.BooleanSet `bson:"available_set" json:"available_set"` //是否可用设置
	CtrlSet      *ctrl.Set        `bson:"ctrl_set" json:"ctrl_set"`           //设置控制信息
}

type GetRequest struct {
	Scope  ref.Scope `bson:"scope" json:"scope"`     //[*]所属业务域
	Lead   *ref.Lead `bson:"lead" json:"lead"`       //[id|lead]业务领衔
	MallID ID        `bson:"mall_id" json:"mall_id"` //[id|lead]对应MALL ID
	With   *struct {
		Address bool `bson:"address" json:"address"` //是否携带地址信息,默认不携带
	} `bson:"with" json:"with"` //携带信息
}

type LoadRequest struct {
	Scope     ref.Scope     `bson:"scope" json:"scope"`           //所属业务域
	LeadArray []ref.Lead    `bson:"lead_array" json:"lead_array"` //通过领衔查找
	IDArray   []ID          `bson:"id_array" json:"id_array"`     //店铺ID列表
	Loc       *location.Loc `bson:"loc" json:"loc"`               //通过地址就近
	Page      *load.Page    `bson:"page" json:"page"`             //分页信息
}

// AddressBindRequest 首次设置地址
type AddressBindRequest struct {
	IdemID  string                      `bson:"idem_id" json:"idem_id"` //[*]幂等ID
	Scope   ref.Scope                   `bson:"scope" json:"scope"`     //所属业务域
	Lead    *ref.Lead                   `bson:"lead" json:"lead"`       //[id|lead]业务领衔
	MallID  ID                          `bson:"mall_id" json:"mall_id"` //[id|lead]MALL ID
	Request address.AddrRegisterRequest `bson:"request" json:"request"` //地址创建请求
}

type Service interface {
	Create(req CreateRequest) (*Mall, *errors.Error)
	Set(req SetRequest) (*Mall, *errors.Error)
	Get(req GetRequest) (*Mall, *errors.Error)
	Load(req LoadRequest) ([]Mall, load.Paging, *errors.Error)
	AddressBind(req AddressBindRequest) (*address.Address, *errors.Error)
}
