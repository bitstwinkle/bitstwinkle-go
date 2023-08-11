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
	"github.com/bitstwinkle/bitstwinkle-go/types/errors"
	"github.com/bitstwinkle/bitstwinkle-go/types/load"
	"github.com/bitstwinkle/bitstwinkle-go/types/ww"
)

type Service interface {
	// DollarCreate 创建积分定义
	DollarCreate(permit *ww.Permit, req DollarCreateRequest) (dollarID DollarID, err *errors.Error)

	// DollarSet 设置积分定义
	DollarSet(permit *ww.Permit, req DollarSetRequest) (dollarID DollarID, err *errors.Error)

	// DollarGet 设置积分定义
	DollarGet(permit *ww.Permit, req DollarGetRequest) (dollar *Dollar, err *errors.Error)

	// DollarIncCirc 积分增加发行量
	DollarIncCirc(permit *ww.Permit, req DollarIncCircRequest) (dollarID DollarID, err *errors.Error)

	// AccountCreate 创建账户
	AccountCreate(permit *ww.Permit, req AccountCreateRequest) (accountID AccountID, err *errors.Error)

	// AccountSet 账户设置
	AccountSet(permit *ww.Permit, req AccountSetRequest) (accountID AccountID, err *errors.Error)

	// AccountGet 账户信息获取
	AccountGet(permit *ww.Permit, req AccountGetRequest) (account *Account, err *errors.Error)

	// AccountLoad 查找Account
	AccountLoad(permit *ww.Permit, req AccountLoadRequest) ([]*Account, *load.Paging, *errors.Error)

	// Airdrop 空投
	Airdrop(permit *ww.Permit, req AirdropRequest) (transferID TransferID, err *errors.Error)

	// Buyback 回收
	Buyback(permit *ww.Permit, req BuybackRequest) (transferID TransferID, err *errors.Error)

	// TransApply 发起转账申请
	TransApply(permit *ww.Permit, req TransApplyRequest) (transferID TransferID, err *errors.Error)

	// TransAdvance 转账推进
	TransAdvance(permit *ww.Permit, req TransAdvanceRequest) (transferID TransferID, err *errors.Error)

	// TransLockAdvance 转账锁定推进,使之收款方可用
	TransLockAdvance(permit *ww.Permit, req TransLockAdvanceRequest) (transferID TransferID, err *errors.Error)

	// TransCancel 在锁定期,撤销转账
	TransCancel(permit *ww.Permit, req TransCancelRequest) (transferID TransferID, err *errors.Error)

	// TransRefund 在锁定期退款
	TransRefund(permit *ww.Permit, req TransRefundRequest) (transferID TransferID, err *errors.Error)

	// TransGet 获取转账明细
	TransGet(permit *ww.Permit, req TransGetRequest) (transfer *Transfer, err *errors.Error)

	// TransLoad 获取转账明细列表
	TransLoad(permit *ww.Permit, req TransLoadRequest) ([]*Transfer, *load.Paging, *errors.Error)
}
