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

type AddrID = string

type Contact struct {
	Person string    `bson:"person" json:"person"` //Name of contact person
	Gender string    `bson:"gender" json:"gender"` //Gender
	Phone  string    `bson:"phone" json:"phone"`   //Phone Number
	More   more.More `bson:"more" json:"more"`     //More Info
}

type Address struct {
	Scope       ref.Scope     `bson:"scope" json:"scope"`                                 //Scope Info
	Lead        ref.Lead      `bson:"lead" json:"lead"`                                   //Lead Info
	ID          AddrID        `bson:"id" json:"id"`                                       //ADDR ID
	LbsProvider string        `bson:"lbs_provider" json:"lbs_provider"`                   //Provider: INNER|GD|BAIDU|TENCENT|GOOGLE|...
	LbsPoiID    string        `bson:"lbs_poi_id" json:"lbs_poi_id"`                       //POI ID
	Title       string        `bson:"title" json:"title"`                                 //Address Shot String
	AddrPath    []string      `bson:"addr_path"json:"addr_path"`                          //Full Path
	Loc         *location.Loc `bson:"loc,omitempty" json:"loc,omitempty"`                 //Location Info
	Contact     *Contact      `bson:"contact,omitempty" json:"contact,omitempty"`         //Contact Info
	LabelArray  label.Array   `bson:"label_array,omitempty" json:"label_array,omitempty"` //Label Info
	Ctrl        *ctrl.Ctrl    `bson:"ctrl,omitempty" json:"ctrl,omitempty"`               //Ctrl Info
	BirthAt     time.Time     `bson:"birth_at" json:"birth_at"`                           //Created Time
	ModifiedAt  time.Time     `bson:"modified_at" json:"modified_at"`                     //Modified Time
}

type AddrRegisterRequest struct {
	Scope       ref.Scope     `bson:"scope" json:"scope"`                 //Scope Info
	Lead        *ref.Lead     `bson:"lead" json:"lead"`                   //Lead Info
	LbsProvider string        `bson:"lbs_provider" json:"lbs_provider"`   //Provider: INNER|GD|BAIDU|TENCENT|GOOGLE|...
	LbsPoiID    string        `bson:"lbs_poi_id" json:"lbs_poi_id"`       //POI ID
	Title       string        `bson:"title" json:"title"`                 //Address Shot String
	AddrPath    []string      `bson:"addr_path"json:"addr_path"`          //Full Path
	Loc         *location.Loc `bson:"loc,omitempty" json:"loc,omitempty"` //坐标信息
	Contact     *struct {
		Person string     `bson:"person" json:"person"` //Name of contact person
		Gender string     `bson:"gender" json:"gender"` //Gender
		Phone  string     `bson:"phone" json:"phone"`   //Phone Number
		More   more.Array `bson:"more" json:"more"`     //More Info
	} `bson:"contact" json:"contact"` //Contact Info
	LabelArray label.Array `bson:"label_array,omitempty" json:"label_array,omitempty"` //Label Info
	Ctrl       *ctrl.Ctrl  `bson:"ctrl,omitempty" json:"ctrl,omitempty"`               //Ctrl Info
}

type AddrSetRequest struct {
	Scope    ref.Scope       `bson:"scope" json:"scope"`         //[*]Scope Info
	Lead     *ref.Lead       `bson:"lead" json:"lead"`           //[id|lead]业务领衔
	AddrID   AddrID          `bson:"addr_id" json:"addr_id"`     //[id|lead]所属地址ID
	TitleSet *ctrl.StringSet `bson:"title_set" json:"title_set"` //Set Title
	LocSet   *location.Set   `bson:"loc_set" json:"loc_set"`     //Set Location
	Contact  *struct {
		Person string     `bson:"person" json:"person"` //Name of contact person
		Gender string     `bson:"gender" json:"gender"` //Gender
		Phone  string     `bson:"phone" json:"phone"`   //Phone Number
		More   more.Array `bson:"more" json:"more"`     //More Info
	} `bson:"contact" json:"contact"` //Contact Info
	LabelArray *label.Set `bson:"label_array" json:"label_array"` //Label Info
	CtrlSet    *ctrl.Set  `bson:"ctrl_set" json:"ctrl_set"`       //Ctrl Info
}

type AddrLoadRequest struct {
	Scope       ref.Scope     `bson:"scope" json:"scope"`           //[*]Scope
	LeadArray   []ref.Lead    `bson:"lead_array" json:"lead_array"` //By Lead
	AddrIDArray []AddrID      `bson:"id_array" json:"id_array"`     //By IDs
	Loc         *location.Loc `bson:"loc" json:"loc"`               //By Loc
	Page        *load.Page    `bson:"page" json:"page"`             //Page Info
}
