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
	"time"
)

type Contact struct {
	Person string     `bson:"person" json:"person"` //联系人姓名
	Gender string     `bson:"gender" json:"gender"` //性别
	Phone  string     `bson:"phone" json:"phone"`   //手机号码
	More   more.Array `bson:"more" json:"more"`     //更多信息
}

type Address struct {
	Scope      ref.Scope               `bson:"scope" json:"scope"`                                 //所属域
	Lead       ref.Lead                `bson:"lead" json:"lead"`                                   //业务领衔
	ID         string                  `bson:"id" json:"id"`                                       //ADDR ID
	Title      string                  `bson:"title" json:"title"`                                 //名称
	Address    string                  `bson:"address" json:"address"`                             //长地址信息
	AreaPath   []Area                  `bson:"area_path"json:"area_path"`                          //全路径
	Zone       Zone                    `bson:"zone" json:"zone"`                                   //所属小区
	Loc        *location.Loc           `bson:"loc,omitempty" json:"loc,omitempty"`                 //坐标信息
	ExLoc      map[string]location.Loc `bson:"ex_loc,omitempty" json:"ex_loc,omitempty"`           //扩展坐标信息
	Contact    *Contact                `bson:"contact,omitempty" json:"contact,omitempty"`         //联系人信息
	LabelArray label.Array             `bson:"label_array,omitempty" json:"label_array,omitempty"` //标签信息
	Ctrl       ctrl.Ctrl               `bson:"ctrl" json:"ctrl"`                                   //控制信息
	BirthAt    time.Time               `bson:"birth_at" json:"birth_at"`                           //创建时间
	ModifiedAt time.Time               `bson:"modified_at" json:"modified_at"`                     //最后更新时间
}

type AddrRegisterRequest struct {
	Scope      ref.Scope       `bson:"scope" json:"scope"`                                 //所属域
	Lead       *ref.Lead       `bson:"lead" json:"lead"`                                   //业务领衔
	ZoneID     string          `bson:"zone_id" json:"zone_id"`                             //所属小区
	Title      string          `bson:"title" json:"title"`                                 //名称
	Loc        *location.Loc   `bson:"loc,omitempty" json:"loc,omitempty"`                 //坐标信息
	ExLoc      []location.Item `bson:"ex_loc,omitempty" json:"ex_loc,omitempty"`           //扩展坐标信息
	Contact    *Contact        `bson:"contact" json:"contact"`                             //联系人信息
	LabelArray label.Array     `bson:"label_array,omitempty" json:"label_array,omitempty"` //标签信息
	Ctrl       *ctrl.Ctrl      `bson:"ctrl,omitempty" json:"ctrl,omitempty"`               //控制信息
}

type AddrSetRequest struct {
	Scope      ref.Scope          `bson:"scope" json:"scope"`             //[*]所属域
	Lead       *ref.Lead          `bson:"lead" json:"lead"`               //[id|lead]业务领衔
	AddrID     string             `bson:"addr_id" json:"addr_id"`         //[id|lead]所属地址ID
	TitleSet   *ctrl.StringSet    `bson:"title_set" json:"title_set"`     //设置名称
	LocSet     *location.Set      `bson:"loc_set" json:"loc_set"`         //设置坐标
	ExLoc      *location.ExLocSet `bson:"ex_loc" json:"ex_loc"`           //扩展坐标信息
	Contact    *Contact           `bson:"contact" json:"contact"`         //联系人信息
	LabelArray label.Array        `bson:"label_array" json:"label_array"` //标签信息
	CtrlSet    *ctrl.Set          `bson:"ctrl_set" json:"ctrl_set"`       //控制信息
}

type AddrLoadRequest struct {
	By      load.By       `bson:"by" json:"by"`             //BY: owner|suggest|
	Scope   ref.Scope     `bson:"scope" json:"scope"`       //[*]所属域
	Lead    []ref.Lead    `bson:"lead" json:"lead"`         //[id|lead]业务领衔
	IDArray []string      `bson:"id_array" json:"id_array"` //[id|lead]业务领衔
	Loc     *location.Loc `bson:"loc" json:"loc"`           //就近查询
	Page    *load.Page    `bson:"page" json:"page"`         //分页信息
}
