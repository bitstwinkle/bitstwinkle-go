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

package math

import "bitstwinkle-go/types/errors"

type Ratio struct {
	Num int64 `bson:"num" json:"num"` // Numerator
	Den int64 `bson:"den" json:"den"` // Denominator
}

func (r Ratio) Verify() *errors.Error {
	if r.Num < 0 {
		return errors.Verify("invalid numerator, it is <= zero")
	}
	if r.Den <= 0 {
		return errors.Verify("invalid denominator, it is <= zero")
	}
	return nil
}

func (r Ratio) Multi(val int64) int64 {
	if r.Den == 0 {
		return 0
	}
	left := (val * r.Num) % r.Den
	if left > 0 {
		return val*r.Num/r.Den + 1
	}

	return val * r.Num / r.Den
}

func Of(num int64, den int64) Ratio {
	r := Ratio{
		Num: num,
		Den: den,
	}
	return r
}
