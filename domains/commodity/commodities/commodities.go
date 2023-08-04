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

package commodities

import (
	"github.com/bitstwinkle/bitstwinkle-go/domains/category"
	"github.com/bitstwinkle/bitstwinkle-go/domains/commodity/types/inventory"
	"github.com/bitstwinkle/bitstwinkle-go/domains/commodity/types/spec"
	"github.com/bitstwinkle/bitstwinkle-go/domains/commodity/types/volume"
	"github.com/bitstwinkle/bitstwinkle-go/types/collections/more"
	"github.com/bitstwinkle/bitstwinkle-go/types/ctrl"
	"github.com/bitstwinkle/bitstwinkle-go/types/errors"
	"github.com/bitstwinkle/bitstwinkle-go/types/load"
	"github.com/bitstwinkle/bitstwinkle-go/types/money"
	"github.com/bitstwinkle/bitstwinkle-go/types/ref"
	"github.com/bitstwinkle/bitstwinkle-go/types/view/label"
	"github.com/bitstwinkle/bitstwinkle-go/types/view/media"
)

type CreateRequest struct {
	IdemID     string            `bson:"idem_id" json:"idem_id"`         //[*]幂等ID
	Scope      ref.Scope         `bson:"scope" json:"scope"`             //所属业务域
	CategoryID category.ID       `bson:"category_id" json:"category_id"` //所属类目
	Title      string            `bson:"title" json:"title"`             //标题
	Spec       []spec.Definition `bson:"spec" json:"spec"`               //规格定义
	Info       *struct {
		Alias string     `bson:"alias" json:"alias"` //别名,可以设置多个别名,通过","(半角逗号)分割
		Intro string     `bson:"intro" json:"intro"` //主介绍
		More  more.Array `bson:"more" json:"more"`   //更多信息内容
	} `bson:"info" json:"info"`
	Media *struct {
		Primary *media.Media `bson:"primary" json:"primary"` //主图
		More    media.Array  `bson:"more" json:"more"`       //更多图片视频
	} `bson:"media,omitempty" json:"media,omitempty"`
	Label       label.Array `bson:"label" json:"label"`       //标签
	Ctrl        ctrl.Ctrl   `bson:"ctrl" json:"ctrl"`         //控制信息
	VwhLead     *ref.Lead   `bson:"vwh_lead" json:"vwhLead"`  //商品库领衔信息
	VmcLead     *ref.Lead   `bson:"vmc_lead" json:"vmc_lead"` //销售区领衔信息
	Commodities []struct {
		Spec           []spec.Value        `bson:"spec" json:"spec"`                       //规格定义
		Inventory      inventory.Inventory `bson:"inventory" json:"inventory"`             //库存信息
		Volume         volume.Volume       `bson:"volume" json:"volume"`                   //销量信息
		Price          money.Amount        `bson:"price" json:"price"`                     //指定的销售价
		SuggestedPrice money.Amount        `bson:"suggested_price" json:"suggested_price"` //建议销售价
		Ctrl           *ctrl.Ctrl          `bson:"ctrl" json:"ctrl"`                       //控制信息
	} `bson:"commodities" json:"commodities"` //商品
}

type Service interface {
	// Create 创建商品()
	Create(req CreateRequest) (*Vpu, *errors.Error)

	// VpuLoad 分页查询VPU数据
	VpuLoad(req VpuLoadRequest) ([]Vpu, load.Paging, *errors.Error)

	// VkuLoad 分页查询VKU数据
	VkuLoad(req VkuLoadRequest) ([]Vku, load.Paging, *errors.Error)
}
