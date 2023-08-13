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
 *
 */

package payment

import (
	"github.com/bitstwinkle/bitstwinkle-go/domains/capital/account"
	"github.com/bitstwinkle/bitstwinkle-go/domains/capital/capital"
	"github.com/bitstwinkle/bitstwinkle-go/types/collections/more"
	"github.com/bitstwinkle/bitstwinkle-go/types/errors"
	"github.com/bitstwinkle/bitstwinkle-go/types/load"
	"github.com/bitstwinkle/bitstwinkle-go/types/money"
	"github.com/bitstwinkle/bitstwinkle-go/types/ref"
	"github.com/bitstwinkle/bitstwinkle-go/types/ww"
)

// ID Payment ID
type ID = string

type Status = string

const (
	Draft           Status = "draft"            //草稿态
	Initialized     Status = "initialized"      //初始化
	Executing       Status = "executing"        //执行中
	Completed       Status = "completed"        //已完成
	Canceled        Status = "canceled"         //已取消
	Failed          Status = "failed"           //已失败
	Refunding       Status = "refunding"        //退款中
	PartialRefunded Status = "partial_refunded" //部分退款
	Refunded        Status = "refunded"         //全额退款
	Ultimate        Status = "ultimate"         //终结态,不可做任何变更
)

type Job struct {
	JobID            string          `bson:"job_id" json:"job_id,omitempty"`                                   //支付子单
	PayerAccountID   account.ID      `bson:"payer_account_id" json:"payer_account_id,omitempty"`               //付款人账户ID
	PayeeAccountID   account.ID      `bson:"payee_account_id" json:"payee_account_id,omitempty"`               //收款人账户ID
	Channel          capital.Channel `bson:"channel" json:"channel"`                                           //渠道码
	Amount           money.Amount    `bson:"amount" json:"amount"`                                             //金额
	Status           Status          `bson:"status,omitempty" json:"status,omitempty"`                         //执行状态
	Paras            more.More       `bson:"paras,omitempty" json:"paras,omitempty"`                           //最终结果
	PayerAccountLead *ref.Lead       `bson:"payer_account_lead,omitempty" json:"payer_account_lead,omitempty"` //付款人账户
	PayeeAccountLead *ref.Lead       `bson:"payee_account_lead,omitempty" json:"payee_account_lead,omitempty"` //收款人账户
}

type JobDefine struct {
	PayerAccountLead ref.Lead        `bson:"payer_account_lead" json:"payer_account_lead"` //付款人账户
	PayeeAccountLead ref.Lead        `bson:"payee_account_lead" json:"payee_account_lead"` //收款人账户
	Channel          capital.Channel `bson:"channel" json:"channel"`                       //渠道码
	Amount           money.Amount    `bson:"amount" json:"amount"`                         //金额
}

type Payment struct {
	Scope    ref.Scope    `bson:"scope" json:"scope"`         //所属业务域
	Lead     ref.Lead     `bson:"lead" json:"lead"`           //业务对应领衔:该领衔KEY同查找账户用的KEY一致
	ID       ID           `bson:"id" json:"id"`               //支付单ID
	Payer    ref.Collar   `bson:"payer" json:"payer"`         //付款人
	Payee    ref.Collar   `bson:"payee" json:"payee"`         //收款人
	Issue    ref.Collar   `bson:"issue" json:"issue"`         //关联事务
	Amount   money.Amount `bson:"amount" json:"amount"`       //总支付金额
	Status   Status       `bson:"status" json:"status"`       //执行状态
	JobArray []*Job       `bson:"job_array" json:"job_array"` //执行子任务
}

type CreateRequest struct {
	Scope    ref.Scope    `bson:"scope" json:"scope"`         //所属业务域
	Lead     ref.Lead     `bson:"lead" json:"lead"`           //业务对应领衔[唯一,所以无需单独幂等]
	Payer    ref.Collar   `bson:"payer" json:"payer"`         //付款人
	Payee    ref.Collar   `bson:"payee" json:"payee"`         //收款人
	Issue    ref.Collar   `bson:"issue" json:"issue"`         //关联事务
	Amount   money.Amount `bson:"amount" json:"amount"`       //总支付金额
	JobArray []*JobDefine `bson:"job_array" json:"job_array"` //执行子任务

	With *struct {
		Prepare bool `bson:"prepare" json:"prepare"` //创建时初始化
	} `bson:"with" json:"with"` //携带动作
}

type GetRequest struct {
	Scope     ref.Scope   `bson:"scope" json:"scope"` //[*]所属业务域
	By        load.ByCode `bson:"by" json:"by"`
	PaymentID ID          `bson:"payment_id" json:"payment_id"` //[id|lead]ID
	Lead      *ref.Lead   `bson:"lead" json:"lead"`             //[id|lead]业务对应领衔
}

type PrepareRequest = GetRequest

type CancelRequest = GetRequest

// ConsultRequest 支付渠道咨询
type ConsultRequest = GetRequest

type ConsultResponse struct {
	Channel []struct {
		Channel capital.Channel `bson:"channel" json:"channel"` //渠道
		Balance money.Amount    `bson:"balance" json:"balance"` //余额
	} `bson:"channel" json:"channel"` //支付渠道信息
}

// RefundRequest 退款
type RefundRequest struct {
	Scope     ref.Scope    `bson:"scope" json:"scope"`           //[*]所属业务域
	PaymentID ID           `bson:"payment_id" json:"payment_id"` //[id|lead]ID
	Lead      *ref.Lead    `bson:"lead" json:"lead"`             //[id|lead]业务对应领衔
	Amount    money.Amount `bson:"amount" json:"amount"`         //总退款金额
	Reason    string       `bson:"reason" json:"reason"`         //退款原因
	Jobs      []struct {
		JobID  string       `bson:"job_id" json:"job_id"` //各子单ID
		Amount money.Amount `bson:"amount" json:"amount"` //退款金额
	} `bson:"jobs,omitempty" json:"jobs,omitempty"` //各子单分别退款金额
}

// UltimateRequest 终结确认
type UltimateRequest struct {
	Scope     ref.Scope `bson:"scope" json:"scope"`           //[*]所属业务域
	PaymentID ID        `bson:"payment_id" json:"payment_id"` //[id|lead]ID
	Lead      *ref.Lead `bson:"lead" json:"lead"`             //[id|lead]业务对应领衔
}

type Service interface {
	Create(permit *ww.Permit, req CreateRequest) (*Payment, *errors.Error)
	Get(permit *ww.Permit, req GetRequest) (*Payment, *errors.Error)
	Prepare(permit *ww.Permit, req PrepareRequest) (*Payment, *errors.Error)
	Cancel(permit *ww.Permit, req CreateRequest) (*Payment, *errors.Error)
	Consult(permit *ww.Permit, req ConsultResponse) (*ConsultResponse, *errors.Error)
	Refund(permit *ww.Permit, req RefundRequest) (*Payment, *errors.Error)
	Ultimate(permit *ww.Permit, req UltimateRequest) (ID, *errors.Error)
}
