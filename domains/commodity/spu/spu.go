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

package spu

import (
	"github.com/bitstwinkle/bitstwinkle-go/domains/commodity/types/spec"
	"github.com/bitstwinkle/bitstwinkle-go/types/collections/more"
	"github.com/bitstwinkle/bitstwinkle-go/types/ctrl"
	"github.com/bitstwinkle/bitstwinkle-go/types/errors"
	"github.com/bitstwinkle/bitstwinkle-go/types/ref"
	"github.com/bitstwinkle/bitstwinkle-go/types/view/label"
	"github.com/bitstwinkle/bitstwinkle-go/types/view/media"
	"github.com/bitstwinkle/bitstwinkle-go/types/ww"
	"time"
)

type ID = string

type Spu struct {
	Scope      ref.Scope         `json:"scope"`          //所属业务域
	CategoryID string            `json:"category_id"`    //所属类目
	Title      string            `json:"title"`          //标题
	Info       more.More         `json:"info,omitempty"` //介绍
	Media      media.More        `json:"media"`          //图片视频
	Ctrl       *ctrl.Ctrl        `json:"ctrl"`           //控制信息
	Label      label.Array       `json:"label"`          //标签
	Spec       []spec.Definition `json:"spec"`           //规格定义
	BirthAt    time.Time         `json:"birth_at"`
	ModifiedAt time.Time         `json:"modified_at"`
}

type CreateRequest struct {
	Scope      ref.Scope         `json:"scope"`       //所属业务域
	CategoryID string            `json:"category_id"` //所属类目
	Title      string            `json:"title"`       //标题
	Spec       []spec.Definition `json:"spec"`        //规格定义
	Info       *struct {
		Alias string     `json:"alias,omitempty"` //别名,可以设置多个别名,通过","(半角逗号)分割
		Intro string     `json:"intro,omitempty"` //主介绍
		More  more.Array `json:"more,omitempty"`  //更多信息内容
	} `bson:"info" json:"info"`
	Media *struct {
		Primary *media.Media `json:"primary,omitempty"` //主图
		More    media.Array  `json:"more,omitempty"`    //更多图片视频
	} `bson:"media" json:"media"`
	Label label.Array `json:"label,omitempty"` //标签
	Ctrl  *ctrl.Ctrl  `json:"ctrl"`            //控制信息
}

type SetRequest struct {
	Scope        ref.Scope           `json:"scope"`         //所属业务域
	SpuID        string              `json:"spu_id"`        //对应SPU ID
	TitleSet     *ctrl.StringSet     `json:"title_set"`     //标题
	SpecSet      *spec.DefinitionSet `json:"spec_set"`      //设置规格
	InfoSet      *more.Set           `json:"info_set"`      //设置展示信息
	MediaSet     *media.Set          `json:"media_set"`     //设置媒体信息
	LabelSet     *label.Set          `json:"label_set"`     //设置标签信息
	AvailableSet *ctrl.BooleanSet    `json:"available_set"` //是否可用设置
	CtrlSet      *ctrl.Set           `json:"ctrl_set"`      //设置控制信息
}

type MoveRequest struct {
	Scope      ref.Scope `json:"scope"`       //所属业务域
	SpuID      string    `json:"spu_id"`      //对应SPU ID
	CategoryID string    `json:"category_id"` //迁移到的类目ID
}

type GetRequest struct {
	Scope ref.Scope `bson:"scope" json:"scope"`   //所属业务域
	SpuID ID        `json:"spu_id" json:"spu_id"` //对应SPU ID
}

type Service interface {
	Create(permit *ww.Permit, req CreateRequest) (*Spu, *errors.Error)
	Set(permit *ww.Permit, req SetRequest) (*Spu, *errors.Error)
	Move(permit *ww.Permit, req MoveRequest) *errors.Error
	Get(permit *ww.Permit, req GetRequest) (*Spu, *errors.Error)
}
