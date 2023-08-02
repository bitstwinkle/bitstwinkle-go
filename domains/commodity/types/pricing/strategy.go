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

package pricing

import (
	"fmt"
	"github.com/bitstwinkle/bitstwinkle-go/types/errors"
	"github.com/bitstwinkle/bitstwinkle-go/types/math"
	"github.com/bitstwinkle/bitstwinkle-go/types/money"
)

type Method = string

const (
	Plus     Method = "plus"     //加法
	Multiply Method = "multiply" //乘法
	//Expression Method = "expression" //表达式 FUTURE
)

type Strategy struct {
	Method   Method            `bson:"method" json:"method"`                         //策略方法
	Fitting  *money.AmountSet  `bson:"fitting" json:"fitting"`                       //结果拟合,即将尾数设置为指定值
	Plus     *PlusStrategy     `bson:"plus,omitempty" json:"plus,omitempty"`         //[method==plus]
	Multiply *MultiplyStrategy `bson:"multiply,omitempty" json:"multiply,omitempty"` //[method==multiply]
}

func (s Strategy) Verify() *errors.Error {
	switch s.Method {
	case Plus:
		if s.Plus == nil {
			return errors.Verify("require plus")
		}
		if err := s.Plus.Verify(); err != nil {
			return err
		}
	case Multiply:
		if s.Multiply == nil {
			return errors.Verify("require multiply")
		}
		if err := s.Multiply.Verify(); err != nil {
			return err
		}
	default:
		return errors.Verify("invalid method: " + s.Method)
	}
	if s.Fitting != nil && s.Fitting.Yes {
		if s.Fitting.Amount <= 0 || s.Fitting.Amount >= 10 {
			return errors.Verify(fmt.Sprintf("invalid fitting: <=0 || >= 10, it is %d", s.Fitting.Amount))
		}
	}
	return nil
}

func (s Strategy) Calculate(datum money.Amount) (money.Amount, *errors.Error) {
	newPrice := datum
	var err *errors.Error
	switch s.Method {
	case Plus:
		newPrice, err = s.Plus.Calculate(datum)
	case Multiply:
		newPrice, err = s.Multiply.Calculate(datum)
	default:
		err = errors.Assert("right method", s.Method)
	}
	if err != nil {
		return datum, err
	}
	newPrice = s.fitting(newPrice)
	return newPrice, nil
}

func (s Strategy) fitting(price money.Amount) money.Amount {
	if s.Fitting == nil || !s.Fitting.Yes {
		return price
	}
	fit := s.Fitting.Amount
	newPrice := price
	if newPrice%10 != fit {
		newPrice += fit - newPrice%10
	}
	return newPrice
}

type Calculator interface {
	Verify() *errors.Error
	Calculate(refPrice money.Amount) (money.Amount, *errors.Error)
}

type PlusStrategy struct {
	Amount money.Amount `bson:"amount" json:"amount"` //加的金额
}

func (cal *PlusStrategy) Verify() *errors.Error {
	return nil
}

func (cal *PlusStrategy) Calculate(refPrice money.Amount) (money.Amount, *errors.Error) {
	newPrice := cal.Amount + refPrice
	if newPrice <= 0 {
		return refPrice, errors.Assert("price > 0", fmt.Sprintf("%d[%d]", newPrice, refPrice))
	}
	return newPrice, nil
}

type MultiplyStrategy struct {
	Ratio math.Ratio `bson:"ratio" json:"ratio"` //乘的比例
}

func (cal *MultiplyStrategy) Calculate(refPrice money.Amount) (money.Amount, *errors.Error) {
	return cal.Ratio.Multi(refPrice), nil
}

func (cal *MultiplyStrategy) Verify() *errors.Error {
	return cal.Ratio.Verify()
}
