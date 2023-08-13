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

package pud

import (
	"github.com/bitstwinkle/bitstwinkle-go/types/ctrl"
	"github.com/bitstwinkle/bitstwinkle-go/types/load"
	"github.com/bitstwinkle/bitstwinkle-go/types/money"
	"github.com/bitstwinkle/bitstwinkle-go/types/ref"
	"time"
)

type AccountID = string

type Options struct {
	Injectable  bool `bson:"injectable" json:"injectable"`   //是否可充值
	Extractable bool `bson:"extractable" json:"extractable"` //是否可提现
	InAble      bool `bson:"in_able" json:"in_able"`         //是否可转入
	OutAble     bool `bson:"out_able" json:"out_able"`       //是否可转出
}

// Running 流水信息
type Running struct {
	In  money.BigAmount `bson:"in" json:"in"`   //入账总金额
	Out money.BigAmount `bson:"out" json:"out"` //出账总金额
}

type Account struct {
	Scope            ref.Scope    `bson:"scope" json:"scope"`                         //SCOPE
	Lead             ref.Lead     `bson:"lead" json:"lead"`                           //Lead
	ID               AccountID    `bson:"id" json:"id"`                               //AccountID
	DollarID         DollarID     `bson:"dollar_id" json:"dollar_id"`                 //积分定义ID
	DollarCode       DollarCode   `bson:"dollar_code" json:"dollar_code"`             //积分
	Balance          money.Amount `bson:"balance" json:"balance"`                     //总余额,包含锁定
	AvailableBalance money.Amount `bson:"available_balance" json:"available_balance"` //可用余额
	Tip              string       `bson:"tip" json:"tip"`                             //账户备注名
	Options          *Options     `bson:"options" json:"options"`                     //账户选项
	Available        bool         `bson:"available" json:"available"`                 //是否可用
	Ctrl             *ctrl.Ctrl   `bson:"ctrl,omitempty" json:"ctrl,omitempty"`       //控制信息
	Running          *Running     `bson:"running"`                                    //流水情况
	BirthAt          time.Time    `bson:"birth_at" json:"birth_at"`                   //创建时间
	ModifiedAt       time.Time    `bson:"modified_at" json:"modified_at"`             //最后更新时间

	Dollar *Dollar `bson:"dollar,omitempty" json:"dollar,omitempty"`
}

type AccountCreateRequest struct {
	Scope      ref.Scope   `bson:"scope" json:"scope"`                                 //SCOPE
	Lead       ref.Lead    `bson:"lead" json:"lead"`                                   //Lead
	DollarID   DollarID    `bson:"dollar_id" json:"dollar_id"`                         //积分定义ID
	DollarLead *DollarLead `bson:"dollar_lead,omitempty" json:"dollar_lead,omitempty"` //积分定义LEAD
	Tip        string      `bson:"tip" json:"tip"`                                     //账户备注名
	Options    *Options    `bson:"options" json:"options"`                             //账户选项
	Ctrl       *ctrl.Ctrl  `bson:"ctrl,omitempty" json:"ctrl,omitempty"`               //控制信息
}

type AccountSetRequest struct {
	Scope        ref.Scope        `bson:"scope" json:"scope"`                 //SCOPE
	AccountID    AccountID        `bson:"account_id" json:"account_id"`       //[id|key]所属账户ID
	Lead         *ref.Lead        `bson:"lead" json:"lead"`                   //[id|key]所属账户lead
	TipSet       *ctrl.StringSet  `bson:"tip_set" json:"tip_set"`             //账户备注名
	Options      *Options         `bson:"options" json:"options"`             //账户选项
	CtrlSet      *ctrl.Set        `bson:"ctrl_set" json:"ctrl_set"`           //控制信息
	AvailableSet *ctrl.BooleanSet `bson:"available_set" json:"available_set"` //是否可用
}

type AccountGetRequest struct {
	Scope     ref.Scope   `bson:"scope" json:"scope"`           //SCOPE
	By        load.ByCode `bson:"by" json:"by"`                 //BY:[id|lead] [*id]
	AccountID AccountID   `bson:"account_id" json:"account_id"` //ID
	Lead      *ref.Lead   `bson:"lead" json:"lead"`             //Lead
}

// AirdropRequest 向账户充值, 从Dollar.balance 扣除
type AirdropRequest struct {
	Scope     ref.Scope    `bson:"scope" json:"scope"`                     //[*]所属域
	AccountID AccountID    `bson:"account_id" json:"account_id"`           //[id|key]所属账户ID
	Lead      *ref.Lead    `bson:"lead" json:"lead"`                       //[id|key]所属账户lead
	Amount    money.Amount `bson:"amount" json:"amount"`                   //[*]注入金额
	Issue     *ref.Collar  `bson:"issue,omitempty" json:"issue,omitempty"` //[*]对应事件
	Memo      string       `bson:"memo,omitempty" json:"memo,omitempty"`   //[-]备注
}

// BuybackRequest 从账户提取, 向Dollar.balance 回收
type BuybackRequest struct {
	Scope     ref.Scope    `bson:"scope" json:"scope"`           //[*]所属域
	AccountID AccountID    `bson:"account_id" json:"account_id"` //[id|key]所属账户ID
	Lead      *ref.Lead    `bson:"lead" json:"lead"`             //[id|key]所属账户lead
	Amount    money.Amount `bson:"amount" json:"amount"`         //[*]提取金额
	Issue     *ref.Collar  `bson:"issue" json:"issue"`           //[*]对应事件
	Memo      string       `bson:"memo" json:"memo"`             //[-]备注
}

type AccountLoadRequest struct {
	Scope            ref.Scope   `bson:"scope" json:"scope"` //[*]所属域
	AccountIDArray   []AccountID `bson:"account_id_array" json:"account_id_array"`
	AccountLeadArray []*ref.Lead `bson:"account_lead_array" json:"account_lead_array"`
	DollarIDArray    []DollarID  `bson:"dollar_id_array" json:"dollar_id_array"`
	Page             *load.Page  `bson:"page" json:"page"`

	With *struct {
		Dollar bool `bson:"dollar" json:"dollar"`
	} `bson:"with" json:"with"`
}
