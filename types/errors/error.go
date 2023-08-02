/*
 *  * Copyright (C) 2023 The Developer bitstwinkle
 *  *
 *  * Licensed under the Apache License, Version 2.0 (the "License");
 *  * you may not use this file except in compliance with the License.
 *  * You may obtain a copy of the License at
 *  *
 *  *      http://www.apache.org/licenses/LICENSE-2.0
 *  *
 *  * Unless required by applicable law or agreed to in writing, software
 *  * distributed under the License is distributed on an "AS IS" BASIS,
 *  * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  * See the License for the specific language governing permissions and
 *  * limitations under the License.
 */

package errors

import "fmt"

type Type = int32

const (
	Coding      Type = -999 // code error
	System      Type = -444 // system error
	Application Type = -777 // application error
)

type Error struct {
	Type    Type   `bson:"t" json:"t"`
	Code    string `bson:"code"  json:"code"`
	Message string `bson:"message" json:"message"`

	_err error // Native error
}

func Of(t Type, code string, message string, nativeErr ...error) *Error {
	err := &Error{
		Type:    t,
		Code:    code,
		Message: message,
	}
	if len(nativeErr) > 0 {
		err._err = nativeErr[0]
	}
	return err
}

func Assert(expect string, actual string, nativeErr ...error) *Error {
	return Of(Coding,
		"assert",
		fmt.Sprintf("assert failed, expect: %s, but %s", expect, actual),
		nativeErr...)
}

func Sys(code string, message string, nativeErr ...error) *Error {
	return Of(System,
		code,
		message,
		nativeErr...)
}

func Verify(message string, nativeErr ...error) *Error {
	return Of(Application, "verify", message, nativeErr...)
}

func (e *Error) Error() string {
	if e._err != nil {
		return fmt.Sprintf("[%s] %s. ******** %s ********", e.Code, e.Message, e._err.Error())
	}
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func (e *Error) IsApplication() bool {
	return e.Type == Application
}

func (e *Error) IsCoding() bool {
	return e.Type == Coding
}

func (e *Error) IsSystem() bool {
	return e.Type == System
}

func (e *Error) Native() error {
	return e._err
}
