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

package order

import (
	"github.com/bitstwinkle/bitstwinkle-go/domains/address"
	"github.com/bitstwinkle/bitstwinkle-go/types/collections/more"
	"github.com/bitstwinkle/bitstwinkle-go/types/ctrl"
	"github.com/bitstwinkle/bitstwinkle-go/types/load"
	"github.com/bitstwinkle/bitstwinkle-go/types/money"
	"github.com/bitstwinkle/bitstwinkle-go/types/ref"
	"github.com/bitstwinkle/bitstwinkle-go/types/state"
	"github.com/bitstwinkle/bitstwinkle-go/types/view/media"
	"time"
)

type ID = string

type Matter struct {
	Ref         ref.Collar   `bson:"ref" json:"ref"`                 //The Matter Ref
	Title       string       `bson:"title" json:"title"`             //The Matter Title
	Price       money.Amount `bson:"price" json:"price"`             //Origin Price
	Quantity    int64        `bson:"quantity" json:"quantity"`       //Order Quantity
	Amount      money.Amount `bson:"amount" json:"amount"`           //Order Amount For This Matter
	Requirement more.Array   `bson:"requirement" json:"requirement"` //Other Requirement
}

type Record struct {
	Code    string            `json:"code"`
	Message string            `json:"msg"`
	Paras   map[string]string `json:"paras,omitempty"`
	LastAt  string            `json:"last_at"`
}

type Status = string

type Order struct {
	Scope       ref.Scope        `bson:"scope" json:"scope"`                                   //所属业务域
	Defined     ref.Defined      `bson:"defined" json:"defined"`                               //业务自定义键值
	ID          ID               `bson:"id" json:"id"`                                         //订单ID
	Purchaser   ref.Collar       `bson:"purchaser" json:"purchaser"`                           //购买者
	Provider    ref.Collar       `bson:"provider" json:"provider"`                             //提供者
	Title       string           `bson:"title" json:"title"`                                   //标题
	Amount      money.Amount     `bson:"amount" json:"amount"`                                 //订单金额
	MatterArray []*Matter        `bson:"matter_array" json:"matter_array"`                     //订单关联商品事物
	Address     *address.Address `bson:"address" json:"address"`                               //配送地址
	Info        more.More        `bson:"info,omitempty" json:"info,omitempty"`                 //更多信息数据
	Media       media.More       `bson:"media,omitempty" json:"media,omitempty"`               //更多媒体数据
	Ctrl        *ctrl.Ctrl       `bson:"ctrl,omitempty" json:"ctrl,omitempty"`                 //控制参数
	Status      state.Code       `bson:"status" json:"status"`                                 //主状态
	ExStatus    state.Code       `bson:"ex_status" json:"ex_status"`                           //各流程自定义状态
	RecordArray []*Record        `bson:"record_array,omitempty" json:"record_array,omitempty"` //订单变化记录
	BirthAt     time.Time        `bson:"birth_at" json:"birth_at"`                             //创建时间
	ModifiedAt  time.Time        `bson:"modified_at" json:"modified_at"`                       //最后修改时间
}

type CreateRequest struct {
	Scope       ref.Scope    `bson:"scope" json:"scope"`     //所属业务域
	Defined     ref.Defined  `bson:"defined" json:"defined"` //业务自定义键值
	Purchaser   ref.Collar   `bson:"purchaser" json:"purchaser"`
	Provider    ref.Collar   `bson:"provider" json:"provider"`
	AddrID      string       `bson:"addr_id" json:"addr_id"`
	Title       string       `bson:"title" json:"title"`
	Amount      money.Amount `bson:"amount" json:"amount"`
	Info        *more.Input  `bson:"info" json:"info"`
	Media       *media.Input `bson:"media" json:"media"`
	Ctrl        *ctrl.Ctrl   `bson:"ctrl" json:"ctrl"`
	ExStatus    state.Code   `bson:"ex_status" json:"ex_status"`
	MatterArray []*Matter    `bson:"matter_array" json:"matter_array"`
}

type GetRequest struct {
	Scope   ref.Scope   `bson:"scope" json:"scope"`
	By      load.ByCode `bson:"by" json:"by"`             //BY: id
	OrderID string      `bson:"order_id" json:"order_id"` //[order_id] ID
}

type SetRequest struct {
	OrderID string   `bson:"order_id" json:"order_id"` //Order ID
	InfoSet more.Set `bson:"info_set" json:"info_set"` //信息数据更新
	CtrlSet ctrl.Set `bson:"ctrl_set" json:"ctrl_set"` //控制数据更新
}

type AdvanceRequest struct {
	OrderID  string     `bson:"order_id" json:"order_id"`   //Order ID
	Status   state.Code `bson:"status" json:"status"`       //主状态
	ExStatus state.Code `bson:"ex_status" json:"ex_status"` //自定义状态
	Record   *struct {
		Code    string            `bson:"code" json:"code"`
		Message string            `bson:"message" json:"msg"`
		Paras   map[string]string `bson:"paras" json:"paras,omitempty"`
	} `bson:"record" json:"record"` //自定义记录
}
