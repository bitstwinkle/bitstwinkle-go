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

package address

import (
	"github.com/bitstwinkle/bitstwinkle-go/types/collections/more"
	"github.com/bitstwinkle/bitstwinkle-go/types/ctrl"
	"github.com/bitstwinkle/bitstwinkle-go/types/load"
	"github.com/bitstwinkle/bitstwinkle-go/types/location"
	"github.com/bitstwinkle/bitstwinkle-go/types/ref"
	"github.com/bitstwinkle/bitstwinkle-go/types/view/label"
)

type Contact struct {
	Person string    `json:"person"` //联系人姓名
	Gender string    `json:"gender"` //性别
	Phone  string    `json:"phone"`  //手机号码
	More   more.More `json:"more"`   //更多信息
}

type Address struct {
	Scope      ref.Scope               `json:"scope"`                 //所属域
	ID         string                  `json:"id"`                    //ADDR ID
	Title      string                  `json:"title"`                 //名称
	Address    string                  `json:"address"`               //长地址信息
	AreaPath   []Area                  `json:"path"`                  //全路径
	Zone       Zone                    `json:"zone"`                  //所属小区
	Loc        *location.Loc           `json:"loc,omitempty"`         //坐标信息
	ExLoc      map[string]location.Loc `json:"ex_loc,omitempty"`      //扩展坐标信息
	Owner      ref.Collar              `json:"owner"`                 //拥有者
	Contact    Contact                 `json:"contact"`               //联系人信息
	LabelArray label.Array             `json:"label_array,omitempty"` //标签信息
	Ctrl       ctrl.Ctrl               `json:"ctrl"`                  //控制信息
	BirthAt    string                  `json:"birth_at"`              //创建时间
	ModifiedAt string                  `json:"modified_at"`           //最后更新时间
}

type AddrRegisterRequest struct {
	Scope      ref.Scope               `json:"scope"`                 //所属域
	ZoneID     string                  `json:"zone_id"`               //所属小区
	Title      string                  `json:"title"`                 //名称
	Loc        *location.Loc           `json:"loc,omitempty"`         //坐标信息
	ExLoc      map[string]location.Loc `json:"ex_loc,omitempty"`      //扩展坐标信息
	Owner      ref.Collar              `json:"owner"`                 //拥有者
	Contact    Contact                 `json:"contact"`               //联系人信息
	LabelArray label.Array             `json:"label_array,omitempty"` //标签信息
	Ctrl       ctrl.Ctrl               `json:"ctrl"`                  //控制信息
}

type AddrSetRequest struct {
	Scope      ref.Scope   `json:"scope"`                 //所属域
	AddrID     string      `json:"addr_id"`               //所属地址ID
	Contact    *Contact    `json:"contact,omitempty"`     //联系人信息
	LabelArray label.Array `json:"label_array,omitempty"` //标签信息
	CtrlSet    *ctrl.Set   `json:"ctrl_set,omitempty"`    //控制信息
}

type AddrLoadRequest struct {
	By    load.By     `json:"by"`    //BY: owner|suggest|
	Scope ref.Scope   `json:"scope"` //[*]所属域
	Owner *ref.Collar `json:"owner"` //[owner]拥有者
}
