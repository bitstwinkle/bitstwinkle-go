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

package security

import (
	"github.com/bitstwinkle/bitstwinkle-go/types/collections/more"
	"github.com/bitstwinkle/bitstwinkle-go/types/errors"
	"time"
)

type Token struct {
	TokenPub    string    `bson:"token_pub" json:"token_pub"`
	TokenExpire time.Time `bson:"token_expire" json:"token_expire"`
	Token       string    `bson:"token" json:"token"`
}

func (token Token) Clone(newToken Token) {
	token.TokenPub = newToken.TokenPub
	token.TokenExpire = newToken.TokenExpire
	token.Token = newToken.Token
}

func (token Token) IsAccessExpired() bool {
	return token.TokenExpire.After(time.Now())
}

type Passport struct {
	Code        string    `bson:"code" json:"code"`
	ID          string    `bson:"id" json:"id"`
	Certificate more.More `bson:"certificate" json:"certificate"`
}

type Service interface {
	// Turn Obtain a new token pair through refresh token
	// secretPub: The public key used in network transmission
	// sign: Parameters that are signed using secretKey
	Turn(secretPub string, nonce string, timestamp time.Time, sign string) (*Token, *errors.Error)

	// Auth Verifies the user and returns the TOKEN
	Auth(passport Passport) (*Token, *errors.Error)
}
