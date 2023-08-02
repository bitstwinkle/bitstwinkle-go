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

package vwh

import (
	"github.com/bitstwinkle/bitstwinkle-go/domains/commodity/sku"
	"github.com/bitstwinkle/bitstwinkle-go/domains/commodity/types/inventory"
	"github.com/bitstwinkle/bitstwinkle-go/domains/commodity/types/pricing"
	"github.com/bitstwinkle/bitstwinkle-go/domains/commodity/types/volume"
	"github.com/bitstwinkle/bitstwinkle-go/types/ctrl"
	"github.com/bitstwinkle/bitstwinkle-go/types/errors"
	"github.com/bitstwinkle/bitstwinkle-go/types/load"
	"github.com/bitstwinkle/bitstwinkle-go/types/money"
	"github.com/bitstwinkle/bitstwinkle-go/types/ref"
	"time"
)

type ID = string

type ImportOptions struct {
	Source  string           `bson:"source" json:"source"`   //源头VWH ID
	Pricing pricing.Strategy `bson:"pricing" json:"pricing"` //价格策略
}

type Vwh struct {
	Scope           ref.Scope      `bson:"scope" json:"scope"`                                           //所属业务域
	Lead            ref.Lead       `bson:"lead" json:"lead"`                                             //业务注入键值
	ID              string         `bson:"id" json:"id"`                                                 //虚拟商品库ID
	Name            string         `bson:"name" json:"name"`                                             //虚拟商品库名称
	Imported        bool           `bson:"imported" json:"imported"`                                     //是否引入的库
	ImportedOptions *ImportOptions `bson:"imported_options,omitempty" json:"imported_options,omitempty"` //引入配置
	Available       bool           `bson:"available" json:"available"`                                   //是否可用
	Ctrl            ctrl.Ctrl      `bson:"ctrl" json:"ctrl"`                                             //控制信息
	BirthAt         time.Time      `bson:"birth_at" json:"birth_at"`                                     //创建时间
	ModifiedAt      time.Time      `bson:"modified_at" json:"modified_at"`                               //最后更新时间
}

type Item struct {
	VwhID          string              `bson:"vwh_id" json:"vwh_id"`                   //所属商品库ID
	SkuID          string              `bson:"sku_id" json:"sku_id"`                   //对应SKU ID
	SpuID          string              `bson:"spu_id" json:"spu_id"`                   //对应SPU ID
	Inventory      inventory.Inventory `bson:"inventory" json:"inventory"`             //库存信息
	Volume         volume.Volume       `bson:"volume" json:"volume"`                   //销量信息
	Price          money.Amount        `bson:"price" json:"price"`                     //指定的销售价
	SuggestedPrice money.Amount        `bson:"suggested_price" json:"suggested_price"` //建议销售价
	Ctrl           ctrl.Ctrl           `bson:"ctrl" json:"ctrl"`                       //控制信息
	Available      bool                `bson:"available" json:"available"`             //是否上架
	BrithAt        time.Time           `bson:"brith_at" json:"brith_at"`               //创建时间
	ModifiedAt     time.Time           `bson:"modified_at" json:"modified_at"`         //最后更新时间
	Commodity      *sku.Sku            `bson:"commodity" json:"commodity"`             //SKU 信息
}

type CreateRequest struct {
	IdemID          string         `bson:"idem_id" json:"idem_id"`                                       //[*]幂等ID
	Scope           ref.Scope      `bson:"scope" json:"scope"`                                           //所属业务域
	Lead            ref.Lead       `bson:"lead" json:"lead"`                                             //业务注入键值
	Name            string         `bson:"name" json:"name"`                                             //虚拟商品库名称
	Imported        bool           `bson:"imported" json:"imported"`                                     //是否引入的库
	ImportedOptions *ImportOptions `bson:"imported_options,omitempty" json:"imported_options,omitempty"` //引入配置
	Ctrl            ctrl.Ctrl      `bson:"ctrl" json:"ctrl"`                                             //控制信息
}

type SetRequest struct {
	IdemID             string          `bson:"idem_id" json:"idem_id"`   //[*]幂等ID
	Scope              ref.Scope       `bson:"scope" json:"scope"`       //[*]所属业务域
	VwhID              string          `bson:"vwh_id" json:"vwh_id"`     //[vwh_id|lead]对应ID
	Lead               ref.Lead        `bson:"lead" json:"lead"`         //[vwh_id|lead]业务注入键值
	NameSet            *ctrl.StringSet `bson:"name_set" json:"name_set"` //设置名称
	ImportedOptionsSet *struct {
		Yes   bool           `bson:"yes" json:"yes"`     //是否设置
		Value *ImportOptions `bson:"value" json:"value"` //对应配置
	} `bson:"imported_options_set,omitempty" json:"imported_options_set,omitempty"` //设置引入策略
	AvailableSet *ctrl.BooleanSet `bson:"available_set,omitempty" json:"available_set,omitempty"` //设置是否可用
	CtrlSet      *ctrl.Set        `bson:"ctrl_set,omitempty" json:"ctrl_set,omitempty"`           //控制参数设置
}

type GetRequest struct {
	By    load.By   `bson:"by" json:"by"`         //BY: vwh_id|lead
	VwhID string    `bson:"vwh_id" json:"vwh_id"` //[vwh_id]
	Lead  *ref.Lead `bson:"lead" json:"lead"`     //[lead]
}

type LoadRequest struct {
	Scope     ref.Scope        `bson:"scope" json:"scope"`         //所属业务域
	Owner     *ref.CollarSet   `bson:"owner" json:"owner"`         //Lead.owner
	Tag       *ctrl.TagSet     `bson:"tag" json:"tag"`             //根据控制标
	Available *ctrl.BooleanSet `bson:"available" json:"available"` //是否只返回有效或者无效
	Page      load.Page        `bson:"page" json:"page"`           //分页信息
}

type SetItemRequest struct {
	IdemID       string    `bson:"idem_id" json:"idem_id"` //[*]幂等ID
	Scope        ref.Scope `bson:"scope" json:"scope"`     //所属业务域
	VwhID        string    `bson:"vwh_id" json:"vwh_id"`   //对应ID
	SkuID        string    `bson:"sku_id" json:"sku_id"`   //对应SKU ID
	InventorySet *struct {
		Yes   bool           `bson:"yes" json:"yes"`     //是否设置
		Value inventory.Plan `bson:"value" json:"value"` //配额信息
	} `bson:"inventory_set,omitempty" json:"inventory_set,omitempty"` //库存信息
	VolumeSet *struct {
		Yes   bool          `bson:"yes" json:"yes"`     //是否设置
		Value volume.Volume `bson:"value" json:"value"` //销量信息
	} `bson:"volume_set,omitempty" json:"volume_set,omitempty"` //初始销量[*只在第一次初始化是有效,后面更新时无效*]
	PriceSet          *money.AmountSet `bson:"price,omitempty" json:"price,omitempty"`                         //指定的销售价
	SuggestedPriceSet *money.Amount    `bson:"suggested_price_set,omitempty" json:"suggested_price,omitempty"` //建议销售价
	CtrlSet           *ctrl.Ctrl       `bson:"ctrl_set,omitempty" json:"ctrl_set,omitempty"`                   //控制信息
}

type GetItemRequest struct {
	By    load.By   `bson:"by" json:"by"`         //BY:item_key[*]
	Scope ref.Scope `bson:"scope" json:"scope"`   //所属业务域
	VwhID string    `bson:"vwh_id" json:"vwh_id"` //对应ID
	SkuID string    `bson:"sku_id" json:"sku_id"` //对应SKU ID
}

type LoadItemRequest struct {
	Scope         ref.Scope `bson:"scope" json:"scope"`                   //所属业务域
	VwhIDArray    []string  `bson:"vwh_id_array" json:"vwh_id_array"`     //对应VWH ID
	SpuIDArray    []string  `bson:"spu_id_array" json:"spu_id_array"`     //对应SKU ID
	Page          load.Page `bson:"page" json:"page"`                     //分页信息
	WithCommodity bool      `bson:"with_commodity" json:"with_commodity"` //是否携带商品详情
}

type Service interface {
	Create(req CreateRequest) (*Vwh, *errors.Error)
	Set(req SetRequest) (*Vwh, *errors.Error)
	Get(req GetRequest) (*Vwh, *errors.Error)
	SetItem(req SetItemRequest) (*Item, *errors.Error)
	GetItem(req GetItemRequest) (*Item, *errors.Error)
	LoadItem(req LoadItemRequest) ([]Item, load.Paging, *errors.Error)
}
