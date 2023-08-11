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
	"github.com/bitstwinkle/bitstwinkle-go/domains/category"
	"github.com/bitstwinkle/bitstwinkle-go/domains/commodity/spu"
	"github.com/bitstwinkle/bitstwinkle-go/domains/commodity/types/spec"
	"github.com/bitstwinkle/bitstwinkle-go/domains/commodity/vmc"
	"github.com/bitstwinkle/bitstwinkle-go/domains/commodity/vwh"
	"github.com/bitstwinkle/bitstwinkle-go/types/collections/more"
	"github.com/bitstwinkle/bitstwinkle-go/types/ctrl"
	"github.com/bitstwinkle/bitstwinkle-go/types/load"
	"github.com/bitstwinkle/bitstwinkle-go/types/ref"
	"github.com/bitstwinkle/bitstwinkle-go/types/view/label"
	"github.com/bitstwinkle/bitstwinkle-go/types/view/media"
	"time"
)

type Vpu struct {
	Scope      ref.Scope         `bson:"scope" json:"scope"`             //所属业务域
	Category   category.Category `bson:"category_id" json:"category_id"` //所属类目
	VpuID      VpuID             `bson:"vpu_id" json:"vpu_id"`           //VPU ID
	Title      string            `bson:"title" json:"title"`             //标题
	Info       more.More         `json:"info,omitempty"`                 //介绍
	Media      media.More        `json:"media,omitempty"`                //图片视频
	Ctrl       *ctrl.Ctrl        `json:"ctrl,omitempty"`                 //控制信息
	Label      label.Array       `json:"label,omitempty"`                //标签
	Spec       []spec.Definition `json:"spec"`                           //规格定义
	BirthAt    time.Time         `json:"birth_at"`                       //创建时间
	ModifiedAt time.Time         `json:"modified_at"`                    //最后修改时间

	Related struct {
		VmcID vmc.ID `bson:"vmc_id" json:"vmc_id"`
		VwhID vwh.ID `bson:"vwh_id" json:"vwh_id"`
		SpuID spu.ID `bson:"spu_id" json:"spu_id"`
	} `bson:"related" json:"related"`
	Commodities []*Vku `bson:"commodities,omitempty" json:"commodities,omitempty"` //商品集
}

type VpuLoadRequest struct {
	Scope      ref.Scope  `bson:"scope" json:"scope"`               //所属业务域
	VwhIDArray []vwh.ID   `bson:"vwh_id_array" json:"vwh_id_array"` //虚拟商品库ID
	VmcIDArray []vmc.ID   `bson:"vmc_id_array" json:"vmc_id_array"` //销售渠道ID
	Keyword    string     `bson:"keyword" json:"keyword"`           //关键词
	Page       *load.Page `bson:"page" json:"page"`                 //分页信息

	WithCommodities *ctrl.BooleanSet `bson:"with_commodities" json:"with_commodities"` //是否加载商品集
}
