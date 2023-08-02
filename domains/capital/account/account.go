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

package account

import (
	"github.com/bitstwinkle/bitstwinkle-go/domains/capital/capital"
	"github.com/bitstwinkle/bitstwinkle-go/types/collections/more"
	"github.com/bitstwinkle/bitstwinkle-go/types/ctrl"
	"github.com/bitstwinkle/bitstwinkle-go/types/errors"
	"github.com/bitstwinkle/bitstwinkle-go/types/load"
	"github.com/bitstwinkle/bitstwinkle-go/types/money"
	"github.com/bitstwinkle/bitstwinkle-go/types/ref"
	"time"
)

type ID = string

// Certificate 对应支付通道侧凭证
type Certificate = more.More

// Running 流水信息
type Running struct {
	In  money.BigAmount `bson:"in" json:"in"`   //入账总金额
	Out money.BigAmount `bson:"out" json:"out"` //出账总金额
}

// Account 虚拟账户
type Account struct {
	Scope       ref.Scope       `bson:"scope" json:"scope"`             //所属业务域
	Lead        ref.Lead        `bson:"lead" json:"lead"`               //业务唯一KEY
	ID          ID              `bson:"id" json:"id"`                   //账户ID
	Title       string          `bson:"title" json:"title"`             //名称备注
	Channel     capital.Channel `bson:"channel" json:"channel"`         //对应支付通道
	Certificate Certificate     `bson:"certificate" json:"certificate"` //凭证信息
	Available   bool            `bson:"available" json:"available"`     //是否可用
	BirthAt     time.Time       `bson:"birth_at" json:"birth_at"`       //创建时间
	ModifiedAt  time.Time       `bson:"modified_at" json:"modified_at"` //最后修改时间

	Running *Running `bson:"running" json:"running"` //总流水信息
}

type Direction = int

const (
	IN  Direction = 1
	OUT Direction = -1
)

type Bill struct {
	Scope     ref.Scope    `bson:"scope" json:"scope"`           //所属业务域
	AccountID ID           `bson:"account_id" json:"account_id"` //对应账户ID
	Direction Direction    `bson:"direction" json:"direction"`   //资金方向IN:1|OUT:-1
	Payer     ref.Ref      `bson:"payer" json:"payer"`           //付款方
	Payee     ref.Ref      `bson:"payee" json:"payee"`           //收款方
	Issue     ref.Ref      `bson:"issue" json:"issue"`           //对应事务
	Amount    money.Amount `bson:"amount" json:"amount"`         //账单金额
	Title     string       `bson:"title" json:"title"`           //账单标题
	Memo      string       `bson:"memo" json:"memo"`             //账单备注
	BirthAt   time.Time    `bson:"birth_at" json:"birth_at"`     //创建时间
}

type CreateRequest struct {
	Scope       ref.Scope       `bson:"scope" json:"scope"`             //[*]所属业务域
	Lead        ref.Lead        `bson:"lead" json:"lead"`               //[*]业务唯一KEY
	Title       string          `bson:"title" json:"title"`             //[*]名称备注
	Channel     capital.Channel `bson:"channel" json:"channel"`         //[*]对应支付通道
	Certificate Certificate     `bson:"certificate" json:"certificate"` //[*]凭证信息
}

type SetRequest struct {
	Scope     ref.Scope        `bson:"scope" json:"scope"`                             //[*]所属业务域
	Lead      ref.Lead         `bson:"lead" json:"lead"`                               //[id|lead]业务唯一KEY
	AccountID ID               `bson:"account_id" json:"account_id"`                   //[id|lead]账户ID
	TitleSet  *ctrl.StringSet  `bson:"title_set,omitempty" json:"title_set,omitempty"` //设置名称备注
	CertSet   *more.Set        `bson:"cert_set,omitempty" json:"cert_set,omitempty"`   //设置凭证信息
	Available *ctrl.BooleanSet `bson:"available,omitempty" json:"available,omitempty"` //是否可用
}

type GetRequest struct {
	By        load.By   `bson:"by" json:"by"`                 //BY:id|lead
	AccountID ID        `bson:"account_id" json:"account_id"` //账户ID
	Lead      *ref.Lead `bson:"lead" json:"lead"`             //账户领衔

	WithRunning bool `bson:"with_running" json:"with_running"` //携带资金流水,默认false
}

type BillLoadRequest struct {
	Scope     ref.Scope   `bson:"scope" json:"scope"`           //[*]所属业务域
	AccountID ID          `bson:"account_id" json:"account_id"` //[id|lead]账户ID
	Lead      *ref.Lead   `bson:"lead" json:"lead"`             //[id|lead]账户领衔
	Issue     *ref.Collar `bson:"issue" json:"issue"`           //事务相关
	Between   *struct {
		Start *time.Time `bson:"start" json:"start"` //起始时间
		End   *time.Time `bson:"end" json:"end"`     //终点时间
	} `bson:"between" json:"between"` //时间区间
	Direction *ctrl.IntSet `bson:"direction" json:"direction"` //资金方向
	Page      *load.Page   `bson:"page" json:"page"`           //分页信息

	WithRunning bool `bson:"with_running" json:"with_running"` //携带资金流水,默认 true
}

type BillLoadResponse struct {
	Paging  load.Paging `bson:"paging" json:"paging"`
	Items   []Bill      `bson:"items" json:"items"`
	Running *struct {
		Total    Running `bson:"total" json:"total"`         //查询条件对应流水
		ThisPage Running `bson:"this_page" json:"this_page"` //本页流水统计
	} `bson:"running" json:"running"` //流水统计信息

}

type Service interface {
	Create(req CreateRequest) (*Account, *errors.Error)
	Set(req SetRequest) (*Account, *errors.Error)
	Get(req GetRequest) (*Account, *errors.Error)
	BillLoad(req BillLoadRequest) (*BillLoadResponse, *errors.Error)
}
