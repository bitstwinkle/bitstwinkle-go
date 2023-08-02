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

package vmc

import (
	"github.com/bitstwinkle/bitstwinkle-go/domains/commodity/vwh"
	"github.com/bitstwinkle/bitstwinkle-go/types/ctrl"
	"github.com/bitstwinkle/bitstwinkle-go/types/errors"
	"github.com/bitstwinkle/bitstwinkle-go/types/load"
	"github.com/bitstwinkle/bitstwinkle-go/types/ref"
	"time"
)

type ID = string

type Vmc struct {
	Scope      ref.Scope `bson:"scope" json:"scope"`             //所属业务域
	Lead       ref.Lead  `bson:"lead" json:"lead"`               //业务注入键值
	ID         string    `bson:"id" json:"id"`                   //虚拟销售渠道ID
	Name       string    `bson:"name" json:"name"`               //虚拟销售渠道名称
	Available  bool      `bson:"available" json:"available"`     //是否可用
	Ctrl       ctrl.Ctrl `bson:"ctrl" json:"ctrl"`               //控制信息
	BirthAt    time.Time `bson:"birth_at" json:"birth_at"`       //创建时间
	ModifiedAt time.Time `bson:"modified_at" json:"modified_at"` //最后更新时间
}

type Item struct {
	VmcID      string    `bson:"vmc_id" json:"vmc_id"`           //所属销售渠道ID
	VwhID      string    `bson:"vwh_id" json:"vwh_id"`           //所引商品库ID
	SkuID      string    `bson:"sku_id" json:"sku_id"`           //对应SKU ID
	Saleable   bool      `bson:"saleable" json:"saleable"`       //是否上架
	Seq        int64     `bson:"seq" json:"seq"`                 //排序序号
	BrithAt    time.Time `bson:"brith_at" json:"brith_at"`       //创建时间
	ModifiedAt time.Time `bson:"modified_at" json:"modified_at"` //最后更新时间
	Commodity  *vwh.Item `bson:"commodity" json:"commodity"`     //商品信息
}

type CreateRequest struct {
	IdemID string     `bson:"idem_id" json:"idem_id"` //[*]幂等ID
	Scope  ref.Scope  `bson:"scope" json:"scope"`     //所属业务域
	Lead   ref.Lead   `bson:"lead" json:"lead"`       //业务注入键值
	Name   string     `bson:"name" json:"name"`       //虚拟销售渠道名称
	Ctrl   *ctrl.Ctrl `bson:"ctrl" json:"ctrl"`       //控制信息
}

type SetRequest struct {
	IdemID       string           `bson:"idem_id" json:"idem_id"`                                 //[*]幂等ID
	Scope        ref.Scope        `bson:"scope" json:"scope"`                                     //[*]所属业务域
	VmcID        string           `bson:"vmc_id" json:"vmc_id"`                                   //[vmc_id|lead]对应ID
	Lead         *ref.Lead        `bson:"lead" json:"lead"`                                       //[vmc_id|lead]业务注入键值
	NameSet      *ctrl.StringSet  `bson:"name_set" json:"name_set"`                               //设置名称
	AvailableSet *ctrl.BooleanSet `bson:"available_set,omitempty" json:"available_set,omitempty"` //设置是否可用
	CtrlSet      *ctrl.Set        `bson:"ctrl_set,omitempty" json:"ctrl_set,omitempty"`           //控制参数设置
}

type GetRequest struct {
	By            load.By   `bson:"by" json:"by"`                         //BY: vmc_id | lead
	VmcID         string    `bson:"vmc_id" json:"vmc_id"`                 //[vmc_id|lead]对应ID
	Lead          *ref.Lead `bson:"lead" json:"lead"`                     //[vmc_id|lead]业务注入键值
	WithCommodity bool      `bson:"with_commodity" json:"with_commodity"` //默认TRUE
}

type LoadRequest struct {
	Scope     ref.Scope        `bson:"scope" json:"scope"`         //所属业务域
	Owner     *ref.CollarSet   `bson:"owner" json:"owner"`         //Lead.owner
	Tag       *ctrl.TagSet     `bson:"tag" json:"tag"`             //根据控制标
	Available *ctrl.BooleanSet `bson:"available" json:"available"` //是否只返回有效或者无效
	Page      load.Page        `bson:"page" json:"page"`           //分页信息
}

type ItemSetRequest struct {
	IdemID      string           `bson:"idem_id" json:"idem_id"`                       //[*]幂等ID
	Scope       ref.Scope        `bson:"scope" json:"scope"`                           //所属业务域
	VmcID       string           `bson:"vmc_id" json:"vmc_id"`                         //对应ID
	VwhID       string           `bson:"vwh_id" json:"vwh_id"`                         //对于商品库ID
	SkuID       string           `bson:"sku_id" json:"sku_id"`                         //对应SKU ID
	SaleableSet *ctrl.BooleanSet `bson:"saleable_set" json:"saleable_set"`             //是否可销售设置
	SeqSet      *ctrl.Int64Set   `bson:"seq_set" json:"seq_set"`                       //排序信息
	CtrlSet     *ctrl.Ctrl       `bson:"ctrl_set,omitempty" json:"ctrl_set,omitempty"` //控制信息
}

type ItemGetRequest struct {
	By            load.By   `bson:"by" json:"by"`                         //BY:item_key[*]
	Scope         ref.Scope `bson:"scope" json:"scope"`                   //所属业务域
	VmcID         string    `bson:"vmc_id" json:"vmc_id"`                 //对应VMC ID
	VwhID         string    `bson:"vwh_id" json:"vwh_id"`                 //对应VWH ID
	SkuID         string    `bson:"sku_id" json:"sku_id"`                 //对应SKU ID
	WithCommodity bool      `bson:"with_commodity" json:"with_commodity"` //是否携带商品信息
}

type ItemLoadRequest struct {
	Scope      ref.Scope        `bson:"scope" json:"scope"`               //所属业务域
	VmcIDArray []string         `bson:"vmc_id_array" json:"vmc_id_array"` //对应VMC ID
	VwhIDArray []string         `bson:"vwh_id_array" json:"vwh_id_array"` //对应VWH ID
	SpuIDArray []string         `bson:"spu_id_array" json:"spu_id_array"` //对应SPU ID
	Saleable   *ctrl.BooleanSet `bson:"saleable" json:"saleable"`         //是否已上架
	Tag        []string         `bson:"tag" json:"tag"`                   //控制标
	Keyword    *ctrl.StringSet  `bson:"keyword" json:"keyword"`           //关键词
	Page       *load.Page       `bson:"page" json:"page"`                 //分页信息
}

type Service interface {
	Create(req CreateRequest) (*Vmc, *errors.Error)
	Set(req SetRequest) (*Vmc, *errors.Error)
	Get(req GetRequest) (*Vmc, *errors.Error)
	ItemSet(req ItemSetRequest) (*Item, *errors.Error)
	ItemGet(req ItemGetRequest) (*Item, *errors.Error)
	ItemLoad(req ItemLoadRequest) ([]Item, load.Paging, *errors.Error)
}
