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

package puc

import (
	"github.com/bitstwinkle/bitstwinkle-go/types/errors"
	"github.com/bitstwinkle/bitstwinkle-go/types/load"
)

type Addr = string

type Service interface {
	WalletPreCreate(req WalletPreCreateRequest) (WalletPreCreateResponse, *errors.Error)
	WalletMnemonicCheck(req WalletMnemonicCheckRequest) (WalletMnemonicCheckResponse, *errors.Error)
	WalletCreateConfirm(req WalletCreateConfirmRequest) (*Wallet, *errors.Error)
	WalletAddAccount(req WalletAddAccRequest) (*Account, *errors.Error)
	WalletGet(req WalletGetRequest) (*Wallet, *errors.Error)
	WalletLoadAccount(req WalletLoadAccRequest) ([]*Account, *errors.Error)
	AccLoadTransfer(req AccLoadTransferRequest) ([]*Transfer, *load.Paging, *errors.Error)
	Airdrop(req AirdropRequest) (transferID string, err *errors.Error)
	AccGet(req AccGetRequest) (*Account, *errors.Error)
}
