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

package rest

import (
	"encoding/json"
	"github.com/bitstwinkle/bitstwinkle-go/network/security"
	"github.com/bitstwinkle/bitstwinkle-go/tools/configure"
	"github.com/bitstwinkle/bitstwinkle-go/tools/sys"
	"github.com/bitstwinkle/bitstwinkle-go/types/errors"
	"github.com/bitstwinkle/bitstwinkle-go/types/strs"
	"github.com/go-resty/resty/v2"
	"log"
	"net/http"
	"sync"
	"time"
)

type H = map[string]string

const (
	//DefaultBitstwinkleURL = "https://api.pugc.tech"
	DefaultBitstwinkleURL = "http://localhost:8080"
)

func NewClient() *resty.Client {
	cli := resty.New().
		SetBaseURL(gBaseUrl).
		SetPreRequestHook(doSignature).
		OnAfterResponse(doWrapResponse)
	if sys.RunMode.IsRd() {
		cli.EnableTrace()
	}
	return cli
}

func doSignature(_ *resty.Client, request *http.Request) error {
	rightNow := time.Now()
	if gToken.AccessTokenExpire.Before(rightNow) {
		err := doTurn()
		if err != nil {
			log.Println(err)
			return err
		}
	} else {
		if rightNow.Sub(gToken.AccessTokenExpire).Minutes() < 10 {
			go func() {
				err := doTurn()
				if err != nil {
					log.Println(err)
				}
			}()
		}
	}
	err := security.Signature(request, gToken.AccessTokenPub, gToken.AccessToken)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func doWrapResponse(_ *resty.Client, response *resty.Response) error {
	if response == nil {
		return nil
	}
	strTokenExpire := response.Header().Get(security.HeaderTokenExpire)
	if strTokenExpire != strs.EMPTY {
		tokenExpire := strs.Int64Of(strTokenExpire, 0)
		newTokenExpireTime := time.UnixMilli(tokenExpire)
		if gToken.AccessTokenExpire.Before(newTokenExpireTime) {
			gTokenMutex.Lock()
			gToken.AccessTokenExpire = newTokenExpireTime
			gTokenMutex.Unlock()
		}
	}
	if response.IsSuccess() {
		return nil
	}

	var respErr errors.Error
	err := json.Unmarshal(response.Body(), &respErr)
	if err != nil {
		return errors.Verify("[invalid - data]invalid response")
	}
	return &respErr
}

var (
	gDoTurnCtrlMutex sync.Mutex
	gDoTurnLastCall  time.Time
)

// doTurn Obtain a new token pair through refresh token
func doTurn() *errors.Error {
	gDoTurnCtrlMutex.Lock()
	defer gDoTurnCtrlMutex.Unlock()

	if time.Since(gDoTurnLastCall) < 30*time.Second {
		log.Println("Function doTurn called too soon")
		return nil
	}

	refreshTokenPubSign, err := security.Sign(gToken.RefreshTokenPub, gToken.RefreshToken)
	if err != nil {
		return err
	}
	var token security.Token
	cli := resty.New().
		SetBaseURL(gBaseUrl).
		OnAfterResponse(doWrapResponse)
	if sys.RunMode.IsRd() {
		cli.EnableTrace()
	}
	_, nErr := cli.R().
		SetResult(&token).
		SetFormData(H{
			security.RefreshTokenPub:     gToken.RefreshTokenPub,
			security.RefreshTokenPubSign: refreshTokenPubSign,
		}).
		Post("/" + security.Security + "/" + security.Turn)
	if nErr != nil {
		return errors.Sys("system err: " + nErr.Error())
	}
	gTokenMutex.Lock()
	defer gTokenMutex.Unlock()
	gToken.Clone(token)
	gDoTurnLastCall = time.Now()
	return nil
}

var gBaseUrl string
var gToken security.Token
var gTokenMutex sync.Mutex

func init() {
	var err *errors.Error
	gBaseUrl = configure.GetString("bitstwinkle.url", DefaultBitstwinkleURL)

	gToken.RefreshTokenPub, err = configure.MustGetString("bitstwinkle.auth.refresh.token.pub")
	if err != nil {
		sys.Exit(err)
		return
	}
	gToken.RefreshToken, err = configure.MustGetString("bitstwinkle.auth.refresh.token.pri")
	if err != nil {
		sys.Exit(err)
		return
	}
	sys.Info("Bitstwinkle URL: ", gBaseUrl)
	sys.Info("Bitstwinkle Refresh Token Pub: ", gToken.RefreshTokenPub)
}
