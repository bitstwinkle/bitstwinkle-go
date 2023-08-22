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

package puc

import (
	"github.com/bitstwinkle/bitstwinkle-go/types/load"
	"github.com/bitstwinkle/bitstwinkle-go/types/money"
	"github.com/bitstwinkle/bitstwinkle-go/types/ref"
)

type Account struct {
	Lead    ref.Lead     `bson:"lead" json:"lead"`
	Addr    Addr         `bson:"addr" json:"addr"`
	Coin    Coin         `bson:"coin" json:"coin"`
	Balance money.Amount `bson:"balance" json:"balance"`
	//PublicKey  string       `bson:"public_key" json:"public_key"`
	//Mnemonic   string       `bson:"mnemonic" json:"mnemonic"`
	//PrivateKey string       `bson:"private_key" json:"private_key"`
	//BirthAt    time.Time    `bson:"birth_at" json:"birth_at"`
}

type Transfer struct {
	From   string       `bson:"from" json:"from"`
	To     string       `bson:"to" json:"to"`
	Amount money.Amount `bson:"amount" json:"amount"`
}

type AccLoadTransferRequest struct {
	LeadArray []*ref.Lead `bson:"lead_array" json:"lead_array"`
	Page      *load.Page  `bson:"page" json:"page"`
}

//
//type ID = string
//
//type Key struct {
//	Owner   ref.Collar  `bson:"owner" json:"owner"`     //账户所有人
//	Defined ref.Defined `bson:"defined" json:"defined"` //业务自定义
//}
//
//type Options struct {
//	Injectable       bool         `bson:"injectable"`        //是否可充值
//	Extractable      bool         `bson:"extractable"`       //是否可提现
//	InAble           bool         `bson:"in_able"`           //是否可转入
//	OutAble          bool         `bson:"out_able"`          //是否可转出
//	OverdraftAble    bool         `bson:"overdraft_able"`    //是否可透支
//	OverdraftBalance money.Amount `bson:"overdraft_balance"` //可透支金额
//}
//
//type Account struct {
//	Scope     ref.Scope      `bson:"scope" json:"scope"`       //所属域
//	Coin      coin.Coin      `bson:"coin" json:"coin"`         //游戏币
//	Owner     ref.Collar     `bson:"owner" json:"owner"`       //账户所有人
//	Defined   ref.Defined    `bson:"defined" json:"defined"`   //业务自定义
//	Currency  money.Currency `bson:"currency" json:"currency"` //币种
//	Balance   money.Amount   `bson:"balance" json:"balance"`   //余额
//	Title     string         `bson:"title" json:"title"`       //账户名
//	Options   *Options       `bson:"options"`
//	Available bool           `bson:"available" json:"available"`           //是否可用
//	Ctrl      ctrl.Ctrl      `bson:"ctrl,omitempty" json:"ctrl,omitempty"` //控制信息
//	ID        ID             `bson:"id" json:"id"`                         //账户ID
//	//Running    *Running       `bson:"running"`                              //流水情况
//	BirthAt    time.Time `bson:"birth_at" json:"birth_at"`       //创建时间
//	ModifiedAt time.Time `bson:"modified_at" json:"modified_at"` //最后更新时间
//}
//
//type InjectRequest struct {
//	Scope      ref.Scope    `bson:"scope" json:"scope"`             //[*]所属域
//	AccountID  ID           `bson:"account_id" json:"account_id"`   //[id|key]所属账户ID
//	AccountKey *Key         `bson:"account_key" json:"account_key"` //[id|key]所属账户KEY
//	Amount     money.Amount `bson:"amount" json:"amount"`           //[*]注入金额
//	Issue      *ref.Ref     `bson:"issue" json:"issue"`             //[*]对应事件
//	Memo       string       `bson:"memo" json:"memo"`               //[-]备注
//}
//
//type TransRequest struct {
//	Scope      ref.Scope    `bson:"scope" json:"scope"`             //[*]所属域
//	AccountID  ID           `bson:"account_id" json:"account_id"`   //[id|key]所属账户ID
//	AccountKey *Key         `bson:"account_key" json:"account_key"` //[id|key]所属账户KEY
//	Amount     money.Amount `bson:"amount" json:"amount"`           //[*]注入金额
//	Issue      *ref.Ref     `bson:"issue" json:"issue"`             //[*]对应事件
//	Memo       string       `bson:"memo" json:"memo"`               //[-]备注
//}
