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
	"github.com/bitstwinkle/bitstwinkle-go/types/errors"
	"github.com/go-resty/resty/v2"
	"log"
	"net/http"
)

func NewClient(tokenPri string) *resty.Client {
	client := resty.New().
		SetPreRequestHook(func(_ *resty.Client, request *http.Request) error {
			err := Signature(request, tokenPri)
			if err != nil {
				log.Println(err)
				return err
			}
			return nil
		})
	return client
}

func RequireToken(initKey string) (Token, *errors.Error) {
	var token Token
	_, err := resty.New().R().
		SetResult(&token).
		SetQueryParam("init_key", initKey).
		Get("/vn/token")
	if err != nil {
		return Token{}, errors.Sys("system err: " + err.Error())
	}
	return token, nil
}
