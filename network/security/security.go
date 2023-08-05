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
	AccessTokenPub     string    `bson:"access_token_pub" json:"access_token_pub"`
	AccessTokenExpire  time.Time `bson:"access_token_expire" json:"access_token_expire"`
	AccessToken        string    `bson:"access_token" json:"access_token"`
	RefreshTokenPub    string    `bson:"refresh_token_pub" json:"refresh_token_pub"`
	RefreshTokenExpire time.Time `bson:"refresh_token_expire" json:"refresh_token_expire"`
	RefreshToken       string    `bson:"refresh_token" json:"refresh_token"`
}

func (token Token) Clone(newToken Token) {
	token.AccessTokenPub = newToken.AccessTokenPub
	token.AccessTokenExpire = newToken.AccessTokenExpire
	token.AccessToken = newToken.AccessToken
	token.RefreshTokenPub = newToken.RefreshTokenPub
	token.RefreshTokenExpire = newToken.RefreshTokenExpire
	token.RefreshToken = newToken.RefreshToken
}

func (token Token) IsAccessExpired() bool {
	return token.AccessTokenExpire.After(time.Now())
}

func (token Token) IsRefreshExpired() bool {
	return token.RefreshTokenExpire.After(time.Now())
}

type Passport struct {
	Code        string    `bson:"code" json:"code"`
	ID          string    `bson:"id" json:"id"`
	Certificate more.More `bson:"certificate" json:"certificate"`
}

type Service interface {
	// Turn Obtain a new token pair through refresh token
	// refreshTokenPub: The public key used in network transmission
	// refreshTokenPubSign: Local refresh token Private key Indicates the signature of the refresh token public key
	Turn(refreshTokenPub string, refreshTokenPubSign string) (*Token, *errors.Error)

	// Auth Verifies the user and returns the TOKEN
	Auth(passport Passport) (*Token, *errors.Error)
}
