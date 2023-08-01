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
)

type Brand struct {
	ID         string     `json:"id"`          //品牌ID
	Name       string     `json:"name"`        //名称
	Code       string     `json:"code"`        //品牌编码
	Media      media.More `json:"ex_media"`    //更多图视频
	Info       more.More  `json:"info"`        //更多信息
	Ctrl       ctrl.Ctrl  `json:"ctrl"`        //控制信息
	BrithAt    string     `json:"brith_at"`    //创建时间
	ModifiedAt string     `json:"modified_at"` //最后更新时间
}

type RegisterRequest struct {
	IdemID string    `json:"idem_id"` //[*]幂等ID
	Scope  ref.Scope `json:"scope"`   //[*]所属业务域
	Name   string    `json:"name"`    //[*]名称
	Info   *struct {
		Alias string     `json:"alias"` //[-]品牌别名
		Code  string     `json:"code"`  //[-]品牌编码
		More  more.Array `json:"more"`  //品牌主介绍
	} `json:"info"` //展示信息
	Media *struct {
		Logo    *media.Media `json:"logo,omitempty"`    //LOGO
		Primary *media.Media `json:"primary,omitempty"` //主图视频
		More    more.Array   `json:"more"`              //更多图视频
	} `json:"media"` //图文视频
	Ctrl ctrl.Ctrl `json:"ctrl"` //控制信息
}

type SetRequest struct {
	IdemID    string        `json:"idem_id"`             //[*]幂等ID
	BrandID   string        `json:"brand_id"`            //[*]品牌ID
	Available *ctrl.Boolean `json:"available,omitempty"` //[-]是否可用设置
	MediaSet  *media.Set    `json:"media_set,omitempty"` //[-]MEDIA SET
	InfoSet   *more.Set     `json:"info_set,omitempty"`  //[-]INFO SET
	CtrlSet   *ctrl.Set     `json:"ctrl_set,omitempty"`  //[-]CTRL SET
}

type GetRequest struct {
	By      load.By `json:"by"`       //BY: brand_id[*]
	BrandID string  `json:"brand_id"` //[brand_id]
}

type LoadRequest struct {
	Scope     ref.Scope     `json:"scope"`               //[*]所属业务域
	CtrlTag   []string      `json:"ctrl_tag,omitempty"`  //控制标
	Available *ctrl.Boolean `json:"available,omitempty"` //是否只返回有效,默认true
	Keyword   *ctrl.String  `json:"keyword,omitempty"`   //关键词信息
	Page      load.Page     `json:"page"`                //分页信息
}

type Service interface {
	Register(req RegisterRequest) (*Brand, *errors.Error)
	Get(req GetRequest) (*Brand, *errors.Error)
	Set(req SetRequest) (*Brand, *errors.Error)
	Remove(brandID string) *errors.Error
	Load(req LoadRequest) ([]*Brand, load.Paging, *errors.Error)
}
