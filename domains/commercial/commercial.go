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

package commercial

import (
	"github.com/bitstwinkle/bitstwinkle-go/types/collections/more"
	"github.com/bitstwinkle/bitstwinkle-go/types/ctrl"
	"github.com/bitstwinkle/bitstwinkle-go/types/errors"
	"github.com/bitstwinkle/bitstwinkle-go/types/load"
	"github.com/bitstwinkle/bitstwinkle-go/types/ref"
	"github.com/bitstwinkle/bitstwinkle-go/types/times"
	"github.com/bitstwinkle/bitstwinkle-go/types/tree"
	"github.com/bitstwinkle/bitstwinkle-go/types/view/media"
	"github.com/bitstwinkle/bitstwinkle-go/types/ww"
	"time"
)

type ID = string

type Commercial struct {
	Scope        ref.Scope           `bson:"scope" json:"scope"`
	Lead         ref.Lead            `bson:"lead" json:"lead"`
	Tree         tree.Collar         `bson:"tree" json:"tree"`
	ID           ID                  `bson:"id" json:"id"`
	Name         string              `bson:"name" json:"name"`
	OpeningHours *times.OpeningHours `bson:"opening_hours" json:"opening_hours"`     //营业时间设置
	Info         more.More           `bson:"info,omitempty" json:"info,omitempty"`   //更多信息
	Media        media.More          `bson:"media,omitempty" json:"media,omitempty"` //媒体信息
	Available    bool                `bson:"available" json:"available"`             //是否可用
	Ctrl         *ctrl.Ctrl          `bson:"ctrl,omitempty" json:"ctrl,omitempty"`   //控制信息
	Seq          int64               `bson:"seq" json:"seq"`                         //在上级中的排序
	BirthAt      time.Time           `bson:"birth_at" json:"birth_at"`               //创建时间
	ModifiedAt   time.Time           `bson:"modified_at" json:"modified_at"`         //最后更新时间
}

type CreateRequest struct {
	Scope        ref.Scope           `bson:"scope" json:"scope"`
	Lead         ref.Lead            `bson:"lead" json:"lead"`
	ParentID     ID                  `bson:"parent_id" json:"parent_id"`
	Name         string              `bson:"name" json:"name"`
	OpeningHours *times.OpeningHours `bson:"opening_hours" json:"opening_hours"`     //营业时间设置
	Info         *more.Input         `bson:"info,omitempty" json:"info,omitempty"`   //更多信息
	Media        *media.Input        `bson:"media,omitempty" json:"media,omitempty"` //媒体信息
	Ctrl         *ctrl.Ctrl          `bson:"ctrl,omitempty" json:"ctrl,omitempty"`   //控制信息
	Seq          int64               `bson:"seq" json:"seq"`                         //在上级中的排序
}

type SetRequest struct {
	Scope        ref.Scope        `bson:"scope" json:"scope"`
	ID           ID               `bson:"id" json:"id"`
	Lead         *ref.Lead        `bson:"lead" json:"lead"`
	NameSet      *ctrl.StringSet  `bson:"name_set" json:"name_set"`
	InfoSet      *more.Set        `bson:"info_set,omitempty" json:"info_set,omitempty"`   //更多信息
	MediaSet     *media.Set       `bson:"media_set,omitempty" json:"media_set,omitempty"` //媒体信息
	CtrlSet      *ctrl.Set        `bson:"ctrl_set,omitempty" json:"ctrl_set,omitempty"`   //控制信息
	AvailableSet *ctrl.BooleanSet `bson:"available_set,omitempty" json:"available_set,omitempty"`
	SeqSet       *ctrl.Int64Set   `bson:"seq_set" json:"seq_set"` //在上级中的排序
}

type GetRequest struct {
	Scope ref.Scope   `bson:"scope" json:"scope"`
	By    load.ByCode `bson:"by" json:"by"` //id|lead
	ID    ID          `bson:"id" json:"id"`
	Lead  *ref.Lead   `bson:"lead" json:"lead"`
}

type LoadRequest struct {
	Scope     ref.Scope   `bson:"scope" json:"scope"`
	IdArray   []ID        `bson:"id_array" json:"id_array"`
	LeadArray []*ref.Lead `bson:"lead_array" json:"lead_array"`
	TagArray  []string    `bson:"tag_array" json:"tag_array"`
	Page      *load.Page  `bson:"page" json:"page"`
}

type Service interface {
	Create(permit *ww.Permit, req CreateRequest) (ID, *errors.Error)
	Set(permit *ww.Permit, req SetRequest) (ID, *errors.Error)
	Get(permit *ww.Permit, req GetRequest) (*Commercial, *errors.Error)
	Load(permit *ww.Permit, req LoadRequest) ([]*Commercial, *load.Paging, *errors.Error)
}
