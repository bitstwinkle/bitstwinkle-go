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
	"time"
)

type PoiID = string
type POI struct {
	ID         PoiID         `bson:"id" json:"id"`                       //ADDR ID
	Title      string        `bson:"title" json:"title"`                 //Name
	Address    string        `bson:"address" json:"address"`             //Full Address String
	AreaPath   []Area        `bson:"area_path"json:"area_path"`          //Full Area Path
	Zone       Zone          `bson:"zone" json:"zone"`                   //Zone Info
	Loc        *location.Loc `bson:"loc,omitempty" json:"loc,omitempty"` //Location
	Available  bool          `bson:"available" json:"available"`         //Available
	BirthAt    time.Time     `bson:"birth_at" json:"birth_at"`           //创建时间
	ModifiedAt time.Time     `bson:"modified_at" json:"modified_at"`     //最后更新时间
}

type PoiLoadRequest struct {
	AreaIDArray []AreaID   `bson:"area_id_array" json:"area_id_array"` //Area IDs
	ZoneIDArray []ZoneID   `bson:"zone_id_array" json:"zone_id_array"` //Zone IDs
	PoiIDArray  []PoiID    `bson:"poi_id_array" json:"poi_id_array"`   //POI IDs
	Page        *load.Page `bson:"page" json:"page"`                   //Paging
	With        *struct {
		Deep bool `bson:"deep" json:"deep"`
	} `bson:"with" json:"with"`
}
