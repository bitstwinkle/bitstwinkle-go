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
 *
 */

package commodities

import (
	"github.com/bitstwinkle/bitstwinkle-go/domains/commodity/sku"
	"github.com/bitstwinkle/bitstwinkle-go/domains/commodity/spu"
	"github.com/bitstwinkle/bitstwinkle-go/domains/commodity/vmc"
	"github.com/bitstwinkle/bitstwinkle-go/domains/commodity/vwh"
	"github.com/bitstwinkle/bitstwinkle-go/types/errors"
	"github.com/bitstwinkle/bitstwinkle-go/types/strs"
)

// VkuID 非实际ID,组合型虚拟ID
type VkuID string

func NewVkuID(vmcID vmc.ID, vwhID vwh.ID, skuID sku.ID) (VkuID, *errors.Error) {
	vid, err := strs.NumbStrCompress(vmcID, vwhID, skuID)
	if err != nil {
		return strs.EMPTY, err
	}
	return VkuID(vid), nil
}

func (vkuID VkuID) Parse() (vmc.ID, vwh.ID, sku.ID, *errors.Error) {
	arr, err := strs.NumbStrDecompress(string(vkuID), 3)
	if err != nil {
		return strs.EMPTY, strs.EMPTY, strs.EMPTY, err
	}
	return arr[0], arr[1], arr[2], nil
}

// VpuID 非实际ID,组合型虚拟ID
type VpuID string

func NewVpuID(vmcID vmc.ID, vwhID vwh.ID, spuID spu.ID) (VpuID, *errors.Error) {
	vid, err := strs.NumbStrCompress(vmcID, vwhID, spuID)
	if err != nil {
		return strs.EMPTY, err
	}
	return VpuID(vid), nil
}

func (vpuID VpuID) Parse() (vmc.ID, vwh.ID, spu.ID, *errors.Error) {
	arr, err := strs.NumbStrDecompress(string(vpuID), 3)
	if err != nil {
		return strs.EMPTY, strs.EMPTY, strs.EMPTY, err
	}
	return arr[0], arr[1], arr[2], nil
}
