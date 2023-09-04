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
	"github.com/bitstwinkle/bitstwinkle-go/types/errors"
	"github.com/bitstwinkle/bitstwinkle-go/types/ref"
	"github.com/bitstwinkle/bitstwinkle-go/types/strs"
	"time"
)

type Unique struct {
	Addr Addr      `bson:"addr" json:"addr"`
	Lead *ref.Lead `bson:"lead" json:"lead"`
}

func (uni Unique) Verify() *errors.Error {
	if uni.Addr == strs.EMPTY {
		if uni.Lead == nil {
			return errors.Verify("Neither addr nor lead must be set")
		}
		if err := uni.Lead.Verify(); err != nil {
			return err
		}
	}
	return nil
}

type Wallet struct {
	Addr    Addr      `bson:"addr" json:"addr"`
	Lead    ref.Lead  `bson:"lead" json:"lead"`
	BrithAt time.Time `bson:"brith_at" json:"brith_at"`

	AccArray []*Account `bson:"acc_array" json:"acc_array"`
}

type WalletPreCreateRequest struct {
	Scope ref.Scope `bson:"scope" json:"scope"` //[*]所属域
	Lead  ref.Lead  `bson:"lead" json:"lead"`
}

type WalletPreCreateResponse struct {
	Scope      ref.Scope `bson:"scope" json:"scope"` //[*]所属域
	Unique     Unique    `bson:"unique" json:"unique"`
	PublicKey  string    `bson:"public_key" json:"public_key"`
	Mnemonic   string    `bson:"mnemonic" json:"mnemonic"`
	PrivateKey string    `bson:"private_key" json:"private_key"`
}

type WalletMnemonicCheckRequest struct {
	Scope    ref.Scope `bson:"scope" json:"scope"` //[*]所属域
	Lead     ref.Lead  `bson:"lead" json:"lead"`
	Mnemonic string    `bson:"mnemonic" json:"mnemonic"`
}

type WalletMnemonicCheckResponse struct {
	Correct bool `bson:"correct" json:"correct"`
}

type WalletCreateConfirmRequest struct {
	Scope ref.Scope `bson:"scope" json:"scope"` //[*]所属域
	Lead  ref.Lead  `bson:"lead" json:"lead"`
}

type WalletAddAccRequest struct {
	Scope        ref.Scope `bson:"scope" json:"scope"` //[*]所属域
	WalletUnique Unique    `bson:"wallet_unique" json:"wallet_unique"`
	Name         string    `bson:"name" json:"name"`
	Coin         Coin      `bson:"coin" json:"coin"`
}

// WalletImportAccRequest todo
type WalletImportAccRequest struct {
	Scope        ref.Scope `bson:"scope" json:"scope"` //[*]所属域
	WalletUnique Unique    `bson:"wallet_unique" json:"wallet_unique"`
	Coin         Coin      `bson:"coin" json:"coin"`
}

type WalletGetRequest struct {
	Scope        ref.Scope `bson:"scope" json:"scope"` //[*]所属域
	WalletUnique Unique    `bson:"wallet_unique" json:"wallet_unique"`
	With         *struct {
		Account bool `bson:"account" json:"account"`
	} `bson:"with" json:"with"`
}

type WalletLoadAccRequest struct {
	Scope        ref.Scope `bson:"scope" json:"scope"` //[*]所属域
	WalletUnique Unique    `bson:"wallet_unique" json:"wallet_unique"`
}
