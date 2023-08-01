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
)

type Category struct {
	Scope ref.Scope `json:"scope"` //所属业务域
	Lead  ref.Lead  `json:"lead"`  //业务链接: 唯一
	ID    string    `json:"id"`    //抽象类目ID
	Tree  struct {
		Stair  int      `json:"stair"`  //所属层级
		Parent string   `json:"parent"` //父亲ID
		Path   []string `json:"path"`   //全路径
	} `json:"tree"` //层级结构
	Name       string     `json:"name"`        //类目名称
	Media      media.More `json:"media"`       //更多图视频
	Info       more.More  `json:"info"`        //更多信息
	Ctrl       ctrl.Ctrl  `json:"ctrl"`        //控制信息
	Seq        int64      `json:"seq"`         //在上级类目中的排序
	BrithAt    string     `json:"brith_at"`    //创建时间
	ModifiedAt string     `json:"modified_at"` //最后修改时间
}

type CreateRequest struct {
	IdemID   string    `json:"idem_id"`   //[*]幂等ID
	Scope    ref.Scope `json:"scope"`     //所属业务域
	Lead     ref.Lead  `json:"lead"`      //业务链接: 唯一
	ParentID string    `json:"parent_id"` //父抽象类目ID,顶层使用$
	Name     string    `json:"name"`      //类目名称
	Info     *struct {
		Alias string     `json:"alias"` //[-]别名
		Code  string     `json:"code"`  //[-]编码
		More  more.Array `json:"more"`  //更多信息
	} `json:"info"` //展示信息
	Media *struct {
		Logo    *media.Media `json:"logo,omitempty"`    //LOGO
		Primary *media.Media `json:"primary,omitempty"` //主图视频
		More    more.Array   `json:"more"`              //更多图视频
	} `json:"media"` //图文视频
	Ctrl ctrl.Ctrl `json:"ctrl"` //控制信息
	Seq  int64     `json:"seq"`  //在上级类目中的排序
}

type SetRequest struct {
	IdemID    string        `json:"idem_id"`             //[*]幂等ID
	BrandID   string        `json:"brand_id"`            //[*]ID
	Available *ctrl.Boolean `json:"available,omitempty"` //[-]是否可用设置
	MediaSet  *media.Set    `json:"media_set,omitempty"` //[-]MEDIA SET
	InfoSet   *more.Set     `json:"info_set,omitempty"`  //[-]INFO SET
	CtrlSet   *ctrl.Set     `json:"ctrl_set,omitempty"`  //[-]CTRL SET
}

type GetRequest struct {
	By         load.By   `json:"by"`          //BY: category_id|lead
	CategoryID string    `json:"category_id"` //[category_id]
	Lead       *ref.Lead `json:"category_id"` //[lead]
}

type LoadRequest struct {
	Scope     ref.Scope     `json:"scope"`               //[*]所属业务域
	CtrlTag   []string      `json:"ctrl_tag,omitempty"`  //控制标
	Available *ctrl.Boolean `json:"available,omitempty"` //是否只返回有效,默认true
	Keyword   *ctrl.String  `json:"keyword,omitempty"`   //关键词信息
	Page      load.Page     `json:"page"`                //分页信息
}

type Service interface {
	Create(req CreateRequest) (*Category, *errors.Error)
	Set(req SetRequest) (*Category, *errors.Error)
	Get(req GetRequest) (*Category, *errors.Error)
	Remove(CategoryID string) *errors.Error
	Load(req LoadRequest) ([]*Category, load.Paging, *errors.Error)
}
