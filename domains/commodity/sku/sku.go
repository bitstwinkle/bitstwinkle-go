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

package sku

import (
	"github.com/bitstwinkle/bitstwinkle-go/domains/commodity/spu"
	"github.com/bitstwinkle/bitstwinkle-go/domains/commodity/types/spec"
	"github.com/bitstwinkle/bitstwinkle-go/types/ctrl"
	"github.com/bitstwinkle/bitstwinkle-go/types/errors"
	"github.com/bitstwinkle/bitstwinkle-go/types/load"
	"github.com/bitstwinkle/bitstwinkle-go/types/ref"
	"time"
)

type ID = string

type Sku struct {
	Scope      ref.Scope   `bson:"scope" json:"scope"`                 //所属业务域
	SpuID      spu.ID      `bson:"spu_id" json:"spu_id"`               //SPU ID
	Spec       []spec.Spec `bson:"spec" json:"spec"`                   //规格信息
	Ctrl       *ctrl.Ctrl  `bson:"ctrl" json:"ctrl,omitempty"`         //控制信息
	BirthAt    time.Time   `bson:"birth_at" json:"birth_at"`           //创建时间
	ModifiedAt time.Time   `bson:"modified_at" json:"modified_at"`     //最后更新时间
	Spu        *spu.Spu    `bson:"spu,omitempty" json:"spu,omitempty"` //SPU DETAIL
}

type CreateRequest struct {
	Scope ref.Scope    `bson:"scope" json:"scope"`   //所属业务域
	SpuID spu.ID       `bson:"spu_id" json:"spu_id"` //SPU ID
	Spec  []spec.Value `bson:"spec" json:"spec"`     //规格定义
	Ctrl  *ctrl.Ctrl   `bson:"ctrl" json:"ctrl"`     //控制信息
}

type SetRequest struct {
	Scope        ref.Scope        `bson:"scope" json:"scope"`                 //所属业务域
	SkuID        ID               `bson:"sku_id" json:"sku_id"`               //SKU ID
	SpecSet      *spec.Set        `bson:"spec_set" json:"spec_set"`           //规格定义
	AvailableSet *ctrl.BooleanSet `bson:"available_set" json:"available_set"` //是否可用
	CtrlSet      *ctrl.Set        `bson:"ctrl_set" json:"ctrl_set"`           //控制信息
}

type GetRequest struct {
	By    load.By      `bson:"by" json:"by"`         //BY: sku_id|spu_id
	Scope ref.Scope    `bson:"scope" json:"scope"`   //[*]所属业务域
	SkuID ID           `bson:"sku_id" json:"sku_id"` //[sku_id]对应SkU ID
	SpuID spu.ID       `bson:"spu_id" json:"spu_id"` //[spu_id]对应SPU ID
	Spec  []spec.Value `bson:"spec" json:"spec"`     //[spu_id]SPEC
}

type LoadRequest struct {
	SpuIDArray []string         `bson:"spu_id_array" json:"spu_id_array"` //SPU ID
	WithSpu    *ctrl.BooleanSet `bson:"with_spu" json:"with_spu"`         //是否携带SPU信息
	Page       *load.Page       `bson:"page" json:"page"`                 //分页信息
}

type Service interface {
	Create(req CreateRequest) (*Sku, *errors.Error)
	Set(req SetRequest) (*Sku, *errors.Error)
	Get(req GetRequest) (*Sku, *errors.Error)
	Load(req LoadRequest) ([]Sku, load.Paging, *errors.Error)
}
