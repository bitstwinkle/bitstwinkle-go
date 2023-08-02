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

package vpu

import (
	"github.com/bitstwinkle/bitstwinkle-go/domains/category"
	"github.com/bitstwinkle/bitstwinkle-go/domains/commodity/spu"
	"github.com/bitstwinkle/bitstwinkle-go/domains/commodity/types/spec"
	"github.com/bitstwinkle/bitstwinkle-go/domains/commodity/vku"
	"github.com/bitstwinkle/bitstwinkle-go/types/collections/more"
	"github.com/bitstwinkle/bitstwinkle-go/types/ctrl"
	"github.com/bitstwinkle/bitstwinkle-go/types/ref"
	"github.com/bitstwinkle/bitstwinkle-go/types/view/label"
	"github.com/bitstwinkle/bitstwinkle-go/types/view/media"
	"time"
)

type ID = spu.ID

type Vpu struct {
	VN         string            `bson:"vn" json:"vn"`                   //所属价值网络
	Scope      ref.Scope         `bson:"scope" json:"scope"`             //所属业务域
	Category   category.Category `bson:"category_id" json:"category_id"` //所属类目
	ID         ID                `bson:"id" json:"id"`                   //VPU ID
	Title      string            `bson:"title" json:"title"`             //标题
	Info       more.More         `json:"info,omitempty"`                 //介绍
	Media      media.More        `json:"media"`                          //图片视频
	Ctrl       ctrl.Ctrl         `json:"ctrl"`                           //控制信息
	Label      label.Array       `json:"label"`                          //标签
	Spec       []spec.Definition `json:"spec"`                           //规格定义
	BirthAt    time.Time         `json:"birth_at"`                       //创建时间
	ModifiedAt time.Time         `json:"modified_at"`                    //最后修改时间

	SpuID          spu.ID          `json:"spu_id"`
	ExMedia        media.More      `json:"ex_media"`        //更多图片视频
	SpecDefinition spec.Definition `json:"spec_definition"` //规格定义
	Commodities    []vku.Vku       `json:"commodities"`     //商品集
}
