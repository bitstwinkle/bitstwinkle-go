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

package pud

import (
	"github.com/bitstwinkle/bitstwinkle-go/types/collections/more"
	"github.com/bitstwinkle/bitstwinkle-go/types/ctrl"
	"github.com/bitstwinkle/bitstwinkle-go/types/load"
	"github.com/bitstwinkle/bitstwinkle-go/types/ref"
	"github.com/bitstwinkle/bitstwinkle-go/types/view/media"
	"time"
)

// DollarCode 积分编码| Scope 内唯一
type DollarCode = string

// DollarLead 通过发行方和Dollar 编码可以定位到一种积分
type DollarLead struct {
	Publisher ref.Collar `bson:"publisher" json:"publisher"` //发行方
	Code      DollarCode `bson:"code" json:"code"`           //CODE
}

// DollarID 对应的ID
type DollarID = string

type Dollar struct {
	Scope       ref.Scope  `bson:"scope" json:"scope"`             //SCOPE
	Publisher   ref.Collar `bson:"publisher" json:"publisher"`     //发行方
	Code        DollarCode `bson:"code" json:"code"`               //CODE
	ID          DollarID   `bson:"id" json:"id"`                   //币定义ID
	Issuance    int64      `bson:"issuance" json:"issuance"`       //发行量
	Balance     int64      `bson:"balance" json:"balance"`         //余额
	Circulation int64      `bson:"circulation" json:"circulation"` //流通量
	Name        string     `bson:"name" json:"name"`               //NAME
	Info        more.More  `bson:"info" json:"info"`               //展示信息
	Media       media.More `bson:"media" json:"media"`             //图文视频信息
	Available   bool       `bson:"available" json:"available"`     //是否通行
	Ctrl        *ctrl.Ctrl `bson:"ctrl" json:"ctrl"`               //控制信息
	BirthAt     time.Time  `bson:"birth_at" json:"birth_at"`       //创建时间
	ModifiedAt  time.Time  `bson:"modified_at" json:"modified_at"` //最后修改时间
}

type DollarCreateRequest struct {
	Scope     ref.Scope    `bson:"scope" json:"scope"`         //SCOPE
	Publisher ref.Collar   `bson:"publisher" json:"publisher"` //发行方
	Code      DollarCode   `bson:"code" json:"code"`           //CODE
	Name      string       `bson:"name" json:"name"`           //NAME
	Info      *more.Input  `bson:"info" json:"info"`           //展示信息
	Media     *media.Input `bson:"media" json:"media"`         //图文视频信息
	Ctrl      *ctrl.Ctrl   `bson:"ctrl" json:"ctrl"`           //控制信息
}

type DollarSetRequest struct {
	Scope        ref.Scope        `bson:"scope" json:"scope"` //SCOPE
	DollarID     string           `bson:"dollar_id" json:"dollar_id"`
	Lead         *DollarLead      `bson:"lead" json:"lead"`
	NameSet      *ctrl.StringSet  `bson:"name_set" json:"name_set"`           //NAME
	InfoSet      *more.Set        `bson:"info_set" json:"info_set"`           //展示信息
	MediaSet     *media.Set       `bson:"media_set" json:"media_set"`         //图文视频信息
	CtrlSet      *ctrl.Set        `bson:"ctrl_set" json:"ctrl_set"`           //控制信息
	AvailableSet *ctrl.BooleanSet `bson:"available_set" json:"available_set"` //可用控制
}

// DollarIncCircRequest 增加发行量
type DollarIncCircRequest struct {
	Scope    ref.Scope   `bson:"scope" json:"scope"`         //SCOPE
	DollarID string      `bson:"dollar_id" json:"dollar_id"` //[id|lead]
	Lead     *DollarLead `bson:"lead" json:"lead"`           //[id|lead]
	Issuance int64       `bson:"issuance" json:"issuance"`   //发行量
}

// DollarGetRequest 获取Dollar 信息
type DollarGetRequest struct {
	Scope    ref.Scope   `bson:"scope" json:"scope"` //SCOPE
	By       load.ByCode `bson:"by" json:"by"`
	DollarID string      `bson:"dollar_id" json:"dollar_id"`
	Lead     *DollarLead `bson:"lead" json:"lead"`
}
