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
	"github.com/bitstwinkle/bitstwinkle-go/types/load"
	"github.com/bitstwinkle/bitstwinkle-go/types/money"
	"github.com/bitstwinkle/bitstwinkle-go/types/ref"
	"time"
)

type TransferID = string

type TransferStatus = string

const (
	Initialized     TransferStatus = "initialized"      //初始化|付款方金额冻结
	Locked          TransferStatus = "locked"           //已锁定|已达收款方,且锁定,该状态可撤销可退款
	Completed       TransferStatus = "completed"        //已完成|终结态.不可撤销,不可退款
	Canceled        TransferStatus = "canceled"         //已取消|终结态
	Failed          TransferStatus = "failed"           //已失败|终结态,付款方金额解冻
	PartialRefunded TransferStatus = "partial_refunded" //部分退款|终结态
	Refunded        TransferStatus = "refunded"         //全额退款|终结态
)

type Transfer struct {
	Scope             ref.Scope      `bson:"scope" json:"scope"`                       //SCOPE
	ID                TransferID     `bson:"id" json:"id"`                             //转账单据号
	PayerAccountID    AccountID      `bson:"payer_account_id" json:"payer_account_id"` //付款方
	PayeeAccountID    AccountID      `bson:"payee_account_id" json:"payee_account_id"` //收款方
	Amount            money.Amount   `bson:"amount" json:"amount"`                     //[*]转账金额
	Issue             ref.Collar     `bson:"issue" json:"issue"`                       //[*]对应事件
	Memo              string         `bson:"memo" json:"memo"`                         //[-]备注
	Status            TransferStatus `bson:"status" json:"status"`                     //转账交易单状态
	Locked            bool           `bson:"locked" json:"locked"`                     //是否已锁定
	LockDuration      int64          `bson:"lock_duration" json:"lock_duration"`       //锁定时长
	TransferTimestamp time.Time      `bson:"transfer_timestamp" json:"transfer_timestamp"`
}

// TransApplyRequest 发起同类型积分之间转账,若未推进,则只是付款方对应金额锁定
type TransApplyRequest struct {
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
	With   *struct {
		// Advance [默认:false] 是否立即执行,双方账户变动.TRUE:立即执行; FALSE: 不立即执行,需要执行推进
		Advance bool `bson:"advance" json:"advance"`
		// Lock [Advance==true起作用,默认:true] 是否锁定,锁定期可以撤销和退款.TRUE:锁定; FALSE: 不锁定,直接到收款方对方账户
		Lock bool `bson:"lock" json:"lock"`
		// LockDuration [Lock==true起作用,默认1天] 锁定的秒数,即到达该时间即自动解锁,收款方账户可用,最大290年
		LockDuration int64 `bson:"lock_duration" json:"lock_duration"`
	} `bson:"with" json:"with"`
}

// TransAdvanceRequest 推进,即时到达对方账户
type TransAdvanceRequest struct {
	Scope   ref.Scope `bson:"scope" json:"scope"`       //[*]所属域
	TransID string    `bson:"trans_id" json:"trans_id"` //转账交易号
	With    *struct {
		// Lock [默认:ture]是否锁定,锁定期可以撤销和退款.TRUE:锁定; FALSE: 不锁定,直接到收款方对方账户
		Lock bool `bson:"lock" json:"lock"`
		// LockDuration [Lock==true起作用,默认1天] 锁定的秒数,即到达该时间即自动解锁,收款方账户可用,最大290年
		LockDuration int64 `bson:"lock_duration" json:"lock_duration"`
	} `bson:"with" json:"with"`
}

// TransLockAdvanceRequest 锁定解除,推进到转账金额收款方账户可用
type TransLockAdvanceRequest struct {
	Scope   ref.Scope `bson:"scope" json:"scope"`       //[*]所属域
	TransID string    `bson:"trans_id" json:"trans_id"` //转账交易号
}

// TransCancelRequest 撤销转账
type TransCancelRequest struct {
	Scope   ref.Scope `bson:"scope" json:"scope"`       //[*]所属域
	TransID string    `bson:"trans_id" json:"trans_id"` //转账交易号
}

// TransRefundRequest 退款(需在锁定期才可退款)
type TransRefundRequest struct {
	Scope   ref.Scope    `bson:"scope" json:"scope"`       //[*]所属域
	TransID string       `bson:"trans_id" json:"trans_id"` //转账交易号
	Amount  money.Amount `bson:"amount" json:"amount"`     //退款金额
}

// TransGetRequest 获取转账交易单数据
type TransGetRequest struct {
	Scope   ref.Scope `bson:"scope" json:"scope"`       //[*]所属域
	TransID string    `bson:"trans_id" json:"trans_id"` //转账交易号
}

type TransLoadRequest struct {
	Scope               ref.Scope    `bson:"scope" json:"scope"` //[*]所属域
	AccountLeadArray    []*ref.Lead  `bson:"account_lead_array" json:"account_lead_array"`
	TransferIDArray     []TransferID `bson:"transfer_id_array" json:"transfer_id_array"`
	PayerAccountIDArray []AccountID  `bson:"payer_account_id_array" json:"payer_account_id_array"`
	PayeeAccountIDArray []AccountID  `bson:"payee_account_id_array" json:"payee_account_id_array"`
	Page                *load.Page   `bson:"page" json:"page"`
}
