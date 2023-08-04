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
)

type ID = string

type Matter struct {
	Ref      ref.Ref      `json:"ref"`      //对象引用
	Price    money.Amount `json:"price"`    //单价
	Quantity int64        `json:"quantity"` //数量
	Amount   money.Amount `json:"amount"`   //金额
}

type Record struct {
	Code    string            `json:"code"`
	Message string            `json:"msg"`
	Paras   map[string]string `json:"paras,omitempty"`
	LastAt  string            `json:"last_at"`
}

type Status = string

type Order struct {
	Scope       ref.Scope     `json:"scope"`                  //所属业务域
	Defined     ref.Defined   `json:"defined"`                //业务自定义键值
	ID          ID            `json:"id"`                     //订单ID
	Purchaser   ref.Ref       `json:"purchaser"`              //购买者
	Provider    ref.Ref       `json:"provider"`               //提供者
	Title       string        `json:"title"`                  //标题
	Amount      money.Amount  `json:"amount"`                 //订单金额
	MatterArray []Matter      `json:"matter_array"`           //订单关联商品事物
	Address     *address.Area `json:"address"`                //配送地址
	Memo        string        `json:"memo"`                   //备注信息
	Status      state.Code    `json:"status"`                 //主状态
	ExStatus    state.Code    `json:"ex_status"`              //各流程自定义状态
	Info        more.More     `json:"info"`                   //更多信息数据
	Ctrl        ctrl.Ctrl     `json:"ctrl"`                   //控制参数
	RecordArray []Record      `json:"record_array,omitempty"` //订单变化记录
	BirthAt     string        `json:"birth_at"`               //创建时间
	ModifiedAt  string        `json:"modified_at"`            //最后修改时间
}

type CreateRequest struct {
	Scope       *ref.Scope   `bson:"scope" json:"scope"`     //所属业务域
	Defined     *ref.Defined `bson:"defined" json:"defined"` //业务自定义键值
	Purchaser   *ref.Collar  `bson:"purchaser" json:"purchaser"`
	Provider    *ref.Collar  `bson:"provider" json:"provider"`
	AddrID      string       `bson:"addr_id" json:"addr_id"`
	Title       string       `bson:"title" json:"title"`
	Amount      money.Amount `bson:"amount" json:"amount"`
	Memo        string       `bson:"memo" json:"memo"`
	ExStatus    state.Code   `bson:"ex_status" json:"ex_status"`
	Ctrl        *ctrl.Ctrl   `bson:"ctrl" json:"ctrl"`
	MatterArray []struct {
		Ref      ref.Collar   `bson:"ref" json:"ref"`           //对象引用
		Price    money.Amount `bson:"price" json:"price"`       //单价
		Quantity int64        `bson:"quantity" json:"quantity"` //数量
		Amount   money.Amount `bson:"amount" json:"amount"`     //金额
	}
}

type GetRequest struct {
	By      load.By `bson:"by" json:"by"`             //BY: order_id
	OrderID string  `bson:"order_id" json:"order_id"` //[order_id] ID
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
