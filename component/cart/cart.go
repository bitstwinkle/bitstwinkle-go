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

package cart

import (
	"github.com/bitstwinkle/bitstwinkle-go/domains/commodity/commodities"
	"github.com/bitstwinkle/bitstwinkle-go/types/errors"
	"github.com/bitstwinkle/bitstwinkle-go/types/ref"
	"time"
)

type ID = string

type Item struct {
	CommodityID commodities.VkuID `bson:"commodity_id" json:"commodity_id"` //对应商品 ID
	Quantity    int32             `bson:"quantity" json:"quantity"`         //数量
	Seq         int32             `bson:"seq" json:"seq"`                   //购物车内排序
	Selected    bool              `bson:"selected" json:"selected"`         //是否已选中
	Available   bool              `bson:"available" json:"available"`       //该商品是否可用
	LastAt      time.Time         `bson:"last_at" json:"last_at"`           //登记时间
}

type Cart struct {
	Scope       ref.Scope `bson:"scope" json:"scope"`       //所属业务域
	Lead        ref.Lead  `bson:"lead" json:"lead"`         //业务领衔
	ID          ID        `bson:"id" json:"id"`             //唯一ID
	Commodities []*Item   `bson:"commodities"`              //购物车中商品
	BirthAt     time.Time `bson:"birth_at" json:"birth_at"` //创建时间
}

// PutRequest 向购物车放置商品
type PutRequest struct {
	Scope       ref.Scope         `bson:"scope" json:"scope"`               //[*]所属业务域
	Lead        ref.Lead          `bson:"lead" json:"lead"`                 //[lead|id]业务领衔
	ID          ID                `bson:"id" json:"id"`                     //[lead|id]唯一ID
	CommodityID commodities.VkuID `bson:"commodity_id" json:"commodity_id"` //商品ID
	Quantity    int32             `bson:"quantity" json:"quantity"`         //数量
}

// DeductRequest 从购物车扣减商品数量
type DeductRequest = PutRequest

// SelectedSwitchRequest 切换选中
type SelectedSwitchRequest struct {
	Scope       ref.Scope         `bson:"scope" json:"scope"`               //[*]所属业务域
	Lead        ref.Lead          `bson:"lead" json:"lead"`                 //[lead|id]业务领衔
	ID          ID                `bson:"id" json:"id"`                     //[lead|id]唯一ID
	CommodityID commodities.VkuID `bson:"commodity_id" json:"commodity_id"` //商品ID
}

// GetRequest 获取购物车内容
type GetRequest struct {
	Scope ref.Scope `bson:"scope" json:"scope"` //[*]所属业务域
	Lead  ref.Lead  `bson:"lead" json:"lead"`   //[lead|id]业务领衔
	ID    ID        `bson:"id" json:"id"`       //[lead|id]唯一ID
}

// CleanRequest 清空购物车
type CleanRequest = GetRequest

type Service interface {
	Put(req PutRequest) (*Cart, *errors.Error)
	Deduct(req DeductRequest) (*Cart, *errors.Error)
	SelectedSwitch(req SelectedSwitchRequest) (*Cart, *errors.Error)
	Get(req GetRequest) (*Cart, *errors.Error)
	Clean(req CleanRequest) *errors.Error
}
