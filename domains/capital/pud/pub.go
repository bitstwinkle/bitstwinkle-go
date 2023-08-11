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
	"github.com/bitstwinkle/bitstwinkle-go/types/collections/more"
	"github.com/bitstwinkle/bitstwinkle-go/types/ctrl"
	"github.com/bitstwinkle/bitstwinkle-go/types/errors"
	"github.com/bitstwinkle/bitstwinkle-go/types/load"
	"github.com/bitstwinkle/bitstwinkle-go/types/money"
	"github.com/bitstwinkle/bitstwinkle-go/types/ref"
	"github.com/bitstwinkle/bitstwinkle-go/types/view/media"
	"github.com/bitstwinkle/bitstwinkle-go/types/ww"
	"time"
)

// DollarCode 积分编码| Scope 内唯一
type DollarCode = string

type DollarLead struct {
	Publisher ref.Collar `bson:"publisher" json:"publisher"` //发行方
	Code      DollarCode `bson:"code" json:"code"`           //CODE
}

type Dollar struct {
	Scope       ref.Scope  `bson:"scope" json:"scope"`             //SCOPE
	Publisher   ref.Collar `bson:"publisher" json:"publisher"`     //发行方
	Code        DollarCode `bson:"code" json:"code"`               //CODE
	ID          string     `bson:"id" json:"id"`                   //币定义ID
	Issuance    int64      `bson:"issuance" json:"issuance"`       //发行量
	Balance     int64      `bson:"balance" json:"balance"`         //余额
	Circulation int64      `bson:"circulation" json:"circulation"` //流通量
	Name        string     `bson:"name" json:"name"`               //NAME
	Info        more.More  `bson:"info" json:"info"`               //展示信息
	Media       media.More `bson:"media" json:"media"`             //图文视频信息
	Available   bool       `bson:"available" json:"available"`     //是否通行
	Ctrl        *ctrl.Ctrl `bson:"ctrl" json:"ctrl"`               //控制信息
	BirthAt     time.Time  `bson:"birth_at" json:"birth_at"`       //创建时间
	ModifiedAt  time.Time  `bson:"modified_at" json:"modified_at"` //最后修改时间
}

type AccountID = string

type Options struct {
	Injectable       bool         `bson:"injectable" json:"injectable"`               //是否可充值
	Extractable      bool         `bson:"extractable" json:"extractable"`             //是否可提现
	InAble           bool         `bson:"in_able" json:"in_able"`                     //是否可转入
	OutAble          bool         `bson:"out_able" json:"out_able"`                   //是否可转出
	OverdraftAble    bool         `bson:"overdraft_able" json:"overdraft_able"`       //是否可透支
	OverdraftBalance money.Amount `bson:"overdraft_balance" json:"overdraft_balance"` //可透支金额
}

// Running 流水信息
type Running struct {
	In  money.BigAmount `bson:"in" json:"in"`   //入账总金额
	Out money.BigAmount `bson:"out" json:"out"` //出账总金额
}

type Account struct {
	Scope      ref.Scope    `bson:"scope" json:"scope"`                   //SCOPE
	Defined    ref.Defined  `bson:"defined" json:"defined"`               //业务自定义
	Lead       ref.Lead     `bson:"lead" json:"lead"`                     //Lead
	ID         AccountID    `bson:"id" json:"id"`                         //AccountID
	Dollar     DollarCode   `bson:"dollar" json:"dollar"`                 //积分
	Balance    money.Amount `bson:"balance" json:"balance"`               //余额
	Tip        string       `bson:"tip" json:"tip"`                       //账户备注名
	Options    *Options     `bson:"options" json:"options"`               //账户选项
	Available  bool         `bson:"available" json:"available"`           //是否可用
	Ctrl       *ctrl.Ctrl   `bson:"ctrl,omitempty" json:"ctrl,omitempty"` //控制信息
	Running    *Running     `bson:"running"`                              //流水情况
	BirthAt    time.Time    `bson:"birth_at" json:"birth_at"`             //创建时间
	ModifiedAt time.Time    `bson:"modified_at" json:"modified_at"`       //最后更新时间
}

type DollarCreateRequest struct {
	Scope     ref.Scope    `bson:"scope" json:"scope"`         //SCOPE
	Publisher ref.Collar   `bson:"publisher" json:"publisher"` //发行方
	Code      DollarCode   `bson:"code" json:"code"`           //CODE
	Name      string       `bson:"name" json:"name"`           //NAME
	Info      *more.Input  `bson:"info" json:"info"`           //展示信息
	Media     *media.Input `bson:"media" json:"media"`         //图文视频信息
	Ctrl      *ctrl.Ctrl   `bson:"ctrl" json:"ctrl"`           //控制信息
}

type DollarSetRequest struct {
	Scope        ref.Scope        `bson:"scope" json:"scope"` //SCOPE
	DollarID     string           `bson:"dollar_id" json:"dollar_id"`
	Lead         *DollarLead      `bson:"lead" json:"lead"`
	NameSet      *ctrl.StringSet  `bson:"name_set" json:"name_set"`           //NAME
	InfoSet      *more.Set        `bson:"info_set" json:"info_set"`           //展示信息
	MediaSet     *media.Set       `bson:"media_set" json:"media_set"`         //图文视频信息
	CtrlSet      *ctrl.Set        `bson:"ctrl_set" json:"ctrl_set"`           //控制信息
	AvailableSet *ctrl.BooleanSet `bson:"available_set" json:"available_set"` //可用控制
}

// IncCirculationRequest 增加发行量
type IncCirculationRequest struct {
	Scope    ref.Scope   `bson:"scope" json:"scope"`         //SCOPE
	DollarID string      `bson:"dollar_id" json:"dollar_id"` //[id|lead]
	Lead     *DollarLead `bson:"lead" json:"lead"`           //[id|lead]
	Issuance int64       `bson:"issuance" json:"issuance"`   //发行量
}

type DollarGetRequest struct {
	Scope    ref.Scope   `bson:"scope" json:"scope"` //SCOPE
	By       load.ByCode `bson:"by" json:"by"`
	DollarID string      `bson:"dollar_id" json:"dollar_id"`
	Lead     *DollarLead `bson:"lead" json:"lead"`
}

// InjectRequest 向账户充值, 从Dollar.balance 扣除
type InjectRequest struct {
	Scope      ref.Scope `bson:"scope" json:"scope"` //[*]所属域
	AccountKey struct {
		AccountID AccountID `bson:"account_id" json:"account_id"` //[id|key]所属账户ID
		Lead      *ref.Lead `bson:"lead" json:"lead"`             //[id|key]所属账户lead
	} `bson:"account_key" json:"account_key"`
	Amount money.Amount `bson:"amount" json:"amount"` //[*]注入金额
	Issue  ref.Collar   `bson:"issue" json:"issue"`   //[*]对应事件
	Memo   string       `bson:"memo" json:"memo"`     //[-]备注
}

// ExtractRequest 从账户提取, 向Dollar.balance 回收
type ExtractRequest struct {
	Scope      ref.Scope `bson:"scope" json:"scope"` //[*]所属域
	AccountKey struct {
		AccountID AccountID `bson:"account_id" json:"account_id"` //[id|key]所属账户ID
		Lead      *ref.Lead `bson:"lead" json:"lead"`             //[id|key]所属账户lead
	} `bson:"account_key" json:"account_key"`
	Amount money.Amount `bson:"amount" json:"amount"` //[*]提取金额
	Issue  ref.Collar   `bson:"issue" json:"issue"`   //[*]对应事件
	Memo   string       `bson:"memo" json:"memo"`     //[-]备注
}

// TransRequest 同类型积分之间转账
type TransRequest struct {
	Scope           ref.Scope `bson:"scope" json:"scope"` //[*]所属域
	PayerAccountKey struct {
		AccountID AccountID `bson:"account_id" json:"account_id"` //[id|key]所属账户ID
		Lead      *ref.Lead `bson:"lead" json:"lead"`             //[id|key]所属账户lead
	} `bson:"payer_account_key" json:"payer_account_key"`
	PayeeAccountKey struct {
		AccountID AccountID `bson:"account_id" json:"account_id"` //[id|key]所属账户ID
		Lead      *ref.Lead `bson:"lead" json:"lead"`             //[id|key]所属账户lead
	} `bson:"payee_account_key" json:"payee_account_key"`
	Amount money.Amount `bson:"amount" json:"amount"` //[*]转账金额
	Issue  ref.Collar   `bson:"issue" json:"issue"`   //[*]对应事件
	Memo   string       `bson:"memo" json:"memo"`     //[-]备注
}

type Service interface {
	DollarCreate(permit *ww.Permit, req DollarCreateRequest) (string, *errors.Error)
	DollarSet(permit *ww.Permit, req DollarSetRequest) (string, *errors.Error)
	IncCirculation(permit *ww.Permit, req IncCirculationRequest) (string, *errors.Error)
}
