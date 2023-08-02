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
	"github.com/bitstwinkle/bitstwinkle-go/domains/commodity/types/inventory"
	"github.com/bitstwinkle/bitstwinkle-go/domains/commodity/types/spec"
	"github.com/bitstwinkle/bitstwinkle-go/domains/commodity/types/volume"
	"github.com/bitstwinkle/bitstwinkle-go/domains/commodity/vmc"
	"github.com/bitstwinkle/bitstwinkle-go/domains/commodity/vwh"
	"github.com/bitstwinkle/bitstwinkle-go/types/ctrl"
	"github.com/bitstwinkle/bitstwinkle-go/types/load"
	"github.com/bitstwinkle/bitstwinkle-go/types/money"
	"github.com/bitstwinkle/bitstwinkle-go/types/ref"
	"time"
)

type Vku struct {
	VN             string              `bson:"vn" json:"vn"`                           //所属价值网络
	Scope          ref.Scope           `bson:"scope" json:"scope"`                     //所属业务域
	VkuID          VkuID               `bson:"vku_id" json:"vku_id"`                   //虚拟ID
	Spec           []spec.Spec         `bson:"spec" json:"spec"`                       //规格信息
	Inventory      inventory.Inventory `bson:"inventory" json:"inventory"`             //库存信息
	Volume         volume.Volume       `bson:"volume" json:"volume"`                   //销量信息
	Price          money.Amount        `bson:"price" json:"price"`                     //指定的销售价
	SuggestedPrice money.Amount        `bson:"suggested_price" json:"suggested_price"` //建议销售价
	Ctrl           ctrl.Ctrl           `bson:"ctrl" json:"ctrl"`                       //控制信息
	Available      bool                `bson:"available" json:"available"`             //是否上架
	BrithAt        time.Time           `bson:"brith_at" json:"brith_at"`               //创建时间
	ModifiedAt     time.Time           `bson:"modified_at" json:"modified_at"`         //最后更新时间

	Related struct {
		VmcID vmc.ID `bson:"vmc_id" json:"vmc_id"`
		VwhID vwh.ID `bson:"vwh_id" json:"vwh_id"`
		SkuID sku.ID `bson:"sku_id" json:"sku_id"`
		SpuID spu.ID `bson:"spu_id" json:"spu_id"`
	} `bson:"related" json:"related"`
	Template *Vpu `bson:"template" json:"template"` //VPU 信息
}

type VkuLoadRequest struct {
	Scope      ref.Scope  `bson:"scope" json:"scope"`               //所属业务域
	VwhIDArray []vwh.ID   `bson:"vwh_id_array" json:"vwh_id_array"` //虚拟商品库ID
	VmcIDArray []vmc.ID   `bson:"vmc_id_array" json:"vmc_id_array"` //销售渠道ID
	SpuIDArray []spu.ID   `bson:"spu_id_array" json:"spu_id_array"` //SPU ID
	Keyword    string     `bson:"keyword" json:"keyword"`           //关键词
	Page       *load.Page `bson:"page" json:"page"`                 //分页信息

	WithTemplate *ctrl.BooleanSet `bson:"with_template" json:"with_template"` //是否加载SPU
}
