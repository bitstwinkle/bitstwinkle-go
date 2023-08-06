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
	"github.com/bitstwinkle/bitstwinkle-go/types/load"
	"github.com/bitstwinkle/bitstwinkle-go/types/location"
	"github.com/bitstwinkle/bitstwinkle-go/types/ref"
)

type Zone struct {
	ID     string `json:"id"`      //ZONE ID
	AreaID string `json:"area_id"` //AREA ID
	Tree   struct {
		Stair  int      `json:"stair"`  //所属层级
		Parent string   `json:"parent"` //父亲ID
		Path   []string `json:"path"`   //全路径
	} `json:"area_tree"`
	ExLink     ref.LinkMore            `json:"ex_link"`          //外部链接信息,例如地图
	Title      string                  `json:"title"`            //名称
	Address    string                  `json:"address"`          //长地址信息
	Path       []Area                  `json:"path"`             //全路径
	Loc        *location.Loc           `json:"loc,omitempty"`    //坐标信息
	ExLoc      map[string]location.Loc `json:"ex_loc,omitempty"` //扩展坐标信息
	BirthAt    string                  `json:"birth_at"`         //创建时间
	ModifiedAt string                  `json:"modified_at"`      //最后更新时间
}

type ZoneLoadRequest struct {
	By       load.ByCode `json:"by"`        //by: area_id
	AreaID   string      `json:"area_id"`   //[area_id]
	WithDeep bool        `json:"with_deep"` //[area_id]是否钻取
	Page     *load.Page  `json:"page"`      //分页信息
}
