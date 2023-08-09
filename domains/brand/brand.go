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

package brand

import (
	"github.com/bitstwinkle/bitstwinkle-go/types/collections/more"
	"github.com/bitstwinkle/bitstwinkle-go/types/ctrl"
	"github.com/bitstwinkle/bitstwinkle-go/types/errors"
	"github.com/bitstwinkle/bitstwinkle-go/types/load"
	"github.com/bitstwinkle/bitstwinkle-go/types/ref"
	"github.com/bitstwinkle/bitstwinkle-go/types/view/media"
	"github.com/bitstwinkle/bitstwinkle-go/types/ww"
	"time"
)

type ID = string

type Brand struct {
	Scope      ref.Scope  `bson:"scope" json:"scope"`
	Lead       *ref.Lead  `bson:"lead" json:"lead"`
	ID         ID         `bson:"id" json:"id"`
	Title      string     `bson:"title" json:"title"`
	Info       more.More  `json:"info"`
	Media      media.More `bson:"media" json:"media"`
	Ctrl       *ctrl.Ctrl `bson:"ctrl" json:"ctrl"`
	BrithAt    time.Time  `bson:"brith_at" json:"brith_at"`
	ModifiedAt time.Time  `bson:"modified_at"json:"modified_at"`
}

type RegisterRequest struct {
	Scope ref.Scope `bson:"scope" json:"scope"` //[*]所属业务域
	Lead  *ref.Lead `bson:"lead" json:"lead"`
	Title string    `bson:"title" json:"title"` //[*]名称
	Info  *struct {
		Alias string     `bson:"alias" json:"alias"` //[-]品牌别名
		Code  string     `bson:"code" json:"code"`   //[-]品牌编码
		More  more.Array `bson:"more" json:"more"`   //品牌主介绍
	} `json:"info"` //展示信息
	Media *struct {
		Logo    *media.Media `bson:"logo" json:"logo"`      //LOGO
		Primary *media.Media `bson:"primary"json:"primary"` //主图视频
		More    media.Array  `bson:"more" json:"more"`      //更多图视频
	} `bson:"media" json:"media"` //图文视频
	Ctrl *ctrl.Ctrl `bson:"ctrl" json:"ctrl"` //控制信息
}

type SetRequest struct {
	Scope     ref.Scope        `json:"scope"`               //[*]所属业务域
	BrandID   string           `json:"brand_id"`            //[*]品牌ID
	Available *ctrl.BooleanSet `json:"available,omitempty"` //[-]是否可用设置
	MediaSet  *media.Set       `json:"media_set,omitempty"` //[-]MEDIA SET
	InfoSet   *more.Set        `json:"info_set,omitempty"`  //[-]INFO SET
	CtrlSet   *ctrl.Set        `json:"ctrl_set,omitempty"`  //[-]CTRL SET
}

type GetRequest struct {
	Scope   ref.Scope   `bson:"scope" json:"scope"`       //[*]Scope
	By      load.ByCode `bson:"by" json:"by"`             //BY: lead|id
	BrandID string      `bson:"brand_id" json:"brand_id"` //[by id]
	Lead    *ref.Lead   `bson:"lead" json:"lead"`         //[by lead]
}

type RemoveRequest = GetRequest

type LoadRequest struct {
	Scope        ref.Scope        `bson:"scope" json:"scope"`           //[*]SCOPE
	LeadArray    []ref.Lead       `bson:"lead_array" json:"lead_array"` //Lead info
	BrandIDArray []ID             `bson:"id_array" json:"id_array"`     //Brand ID info
	Available    *ctrl.BooleanSet `bson:"available" json:"available"`   //Available
	CtrlTag      []string         `bson:"ctrl_tag" json:"ctrl_tag"`     //Ctrl Tag
	Keyword      *ctrl.StringSet  `bson:"keyword" json:"keyword"`       //Key Word
	Page         *load.Page       `bson:"page" json:"page"`             //Page
}

type Service interface {
	Register(permit *ww.Permit, req RegisterRequest) (ID, *errors.Error)
	Get(permit *ww.Permit, req GetRequest) (*Brand, *errors.Error)
	Set(permit *ww.Permit, req SetRequest) (ID, *errors.Error)
	Remove(permit *ww.Permit, req RemoveRequest) (ID, *errors.Error)
	Load(permit *ww.Permit, req LoadRequest) ([]*Brand, *load.Paging, *errors.Error)
}
