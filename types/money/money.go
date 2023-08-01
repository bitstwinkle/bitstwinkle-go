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

package money

type Currency string
type Amount = int64
type BigAmount int64

func NewBigAmount(x int64) BigAmount {
	return BigAmount(x)
}

func (c Currency) Verify() error {
	return nil
}

type Money struct {
	Currency Currency `bson:"c"`
	Amount   Amount   `bson:"a"`
}

type Converter interface {
	Convert(src Money, destCurrency Currency) (Money, error)
}

type CurrencyDefine struct {
	Code string
}

func CNY(amount Amount) Money {
	return Money{
		Currency: "CNY",
		Amount:   amount,
	}
}
