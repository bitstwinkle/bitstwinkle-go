/*
 *  * Copyright (C) 2023 The Developer bitstwinkle
 *  *
 *  * Licensed under the Apache License, Version 2.0 (the "License");
 *  * you may not use this file except in compliance with the License.
 *  * You may obtain a copy of the License at
 *  *
 *  *      http://www.apache.org/licenses/LICENSE-2.0
 *  *
 *  * Unless required by applicable law or agreed to in writing, software
 *  * distributed under the License is distributed on an "AS IS" BASIS,
 *  * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  * See the License for the specific language governing permissions and
 *  * limitations under the License.
 */

package category

import (
	"github.com/bitstwinkle/bitstwinkle-go/types/collections/more"
	"github.com/bitstwinkle/bitstwinkle-go/types/ctrl"
	"github.com/bitstwinkle/bitstwinkle-go/types/errors"
	"github.com/bitstwinkle/bitstwinkle-go/types/load"
	"github.com/bitstwinkle/bitstwinkle-go/types/ref"
	"github.com/bitstwinkle/bitstwinkle-go/types/view/media"
	"time"
)

const Ref = "CATEGORY"

type ID = string

func GetCollar(categoryID ID) ref.Collar {
	return ref.Collar{
		Code: Ref,
		ID:   categoryID,
	}
}

const ROOT ID = "$"

type Category struct {
	Scope ref.Scope `bson:"scope" json:"scope"` //所属业务域
	Lead  ref.Lead  `bson:"lead" json:"lead"`   //业务链接: 唯一
	ID    ID        `bson:"id" json:"id"`       //抽象类目ID
	Tree  struct {
		Stair  int      `bson:"stair" json:"stair"`   //所属层级
		Parent string   `bson:"parent" json:"parent"` //父亲ID
		Path   []string `bson:"path" json:"path"`     //全路径
	} `json:"tree"` //层级结构
	Name       string     `bson:"name" json:"name"`               //类目名称
	Media      media.More `bson:"media" json:"media"`             //更多图视频
	Info       more.More  `bson:"info" json:"info"`               //更多信息
	Ctrl       *ctrl.Ctrl `bson:"ctrl" json:"ctrl"`               //控制信息
	Seq        int64      `bson:"seq" json:"seq"`                 //在上级类目中的排序
	BrithAt    time.Time  `bson:"brith_at" json:"brith_at"`       //创建时间
	ModifiedAt time.Time  `bson:"modified_at" json:"modified_at"` //最后修改时间
}

type CreateRequest struct {
	Scope    ref.Scope    `bson:"scope" json:"scope"`         //所属业务域
	Lead     ref.Lead     `bson:"lead" json:"lead"`           //业务链接: 唯一
	ParentID ID           `bson:"parent_id" json:"parent_id"` //父抽象类目ID,顶层使用$
	Name     string       `bson:"name" json:"name"`           //类目名称
	Info     *more.Input  `bson:"info" json:"info"`           //展示信息
	Media    *media.Input `bson:"media" json:"media"`         //图文视频
	Ctrl     *ctrl.Ctrl   `bson:"ctrl" json:"ctrl"`           //控制信息
	Seq      int64        `bson:"seq" json:"seq"`             //在上级类目中的排序
}

type SetRequest struct {
	IdemID       string           `bson:"idem_id" json:"idem_id"`             //[*]幂等ID
	CategoryID   ID               `bson:"category_id" json:"category_id"`     //[*]ID
	AvailableSet *ctrl.BooleanSet `bson:"available_set" json:"available_set"` //[-]是否可用设置
	MediaSet     *media.Set       `bson:"media_set" json:"media_set"`         //[-]MEDIA SET
	InfoSet      *more.Set        `bson:"info_set" json:"info_set"`           //[-]INFO SET
	CtrlSet      *ctrl.Set        `bson:"ctrl_set" json:"ctrl_set"`           //[-]CTRL SET
}

type GetRequest struct {
	By         load.ByCode `bson:"by" json:"by"`                   //BY: category_id|lead
	CategoryID ID          `bson:"category_id" json:"category_id"` //[category_id]
	Lead       *ref.Lead   `bson:"lead" json:"lead"`               //[lead]
}

type LoadRequest struct {
	Scope     ref.Scope        `bson:"scope" json:"scope"`         //[*]所属业务域
	CtrlTag   []string         `bson:"ctrl_tag" json:"ctrl_tag"`   //控制标
	Available *ctrl.BooleanSet `bson:"available" json:"available"` //是否只返回有效,默认true
	Keyword   *ctrl.StringSet  `bson:"keyword" json:"keyword"`     //关键词信息
	Page      *load.Page       `bson:"page" json:"page"`           //分页信息
}

type Service interface {
	Create(req CreateRequest) (*Category, *errors.Error)
	Set(req SetRequest) (*Category, *errors.Error)
	Get(req GetRequest) (*Category, *errors.Error)
	Remove(categoryID string) *errors.Error
	Load(req LoadRequest) ([]*Category, *load.Paging, *errors.Error)
}
