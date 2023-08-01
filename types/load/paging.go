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

package load

import (
	"fmt"
)

const DefaultPagingSize = 100

const MaxPagingSize = 2000

const FirstPage = 1

type Paging struct {
	Size      int64 `bson:"size" json:"size"`             // Page Size, default 100
	Current   int64 `bson:"current" json:"current"`       // From One
	Total     int64 `bson:"total" json:"total"`           // The Page Count
	ItemTotal int64 `bson:"item_total" json:"item_total"` // The Item Count
}

func (paging Paging) ToString() string {
	return fmt.Sprintf("size: %d|current: %d|total: %d|item_total: %d",
		paging.Size, paging.Current, paging.Total, paging.ItemTotal)
}

func (paging Paging) WithItemTotal(itemTotal int64) {
	if paging.Size == 0 {
		fmt.Println("assert paging.Size > 0, but paging.Size == 0")
		return
	}

	if paging.ItemTotal == 0 {
		paging.Total = 0
		paging.Current = FirstPage
		return
	}
	paging.ItemTotal = itemTotal
	paging.Total = paging.ItemTotal / paging.Size
	if paging.ItemTotal%paging.Size != 0 {
		paging.Total += 1
	}

	if paging.Current > paging.Total {
		paging.Current = paging.Total
	}

	if paging.Current < 1 {
		paging.Current = 1
	}
}

func (paging Paging) Skip() int64 {
	return (paging.Current - 1) * paging.Size
}

func (paging Paging) Limit() int64 {
	return paging.Size
}

func PagingOf(size int64, current int64) Paging {
	if size <= 0 {
		size = DefaultPagingSize
	}
	if size > MaxPagingSize {
		size = MaxPagingSize
	}
	if current < FirstPage {
		current = FirstPage
	}
	return Paging{
		Size:      size,
		Current:   current,
		Total:     0,
		ItemTotal: 0,
	}
}

func PagingALL() Paging {
	return PagingOf(MaxPagingSize, FirstPage)
}
