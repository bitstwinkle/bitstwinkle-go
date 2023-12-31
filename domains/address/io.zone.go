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
	"time"
)

type ZoneID = string
type Zone struct {
	ID     ZoneID `bson:"id" json:"id"`
	AreaID AreaID `bson:"area_id "json:"area_id"`
	Tree   struct {
		Stair  int      `bson:"stair" json:"stair"`
		Parent string   `bson:"parent" json:"parent"`
		Path   []AreaID `bson:"path" json:"path"`
	} `bson:"tree" json:"area_tree"`
	ExLink     ref.LinkMore            `bson:"ex_link" json:"ex_link"`
	Title      string                  `bson:"title" json:"title"`
	Address    string                  `bson:"address" json:"address"`
	Path       []Area                  `bson:"path" json:"path"`
	Loc        *location.Loc           `bson:"loc,omitempty" json:"loc,omitempty"`
	ExLoc      map[string]location.Loc `bson:"ex_loc" json:"ex_loc,omitempty"`
	BirthAt    time.Time               `bson:"birth_at" json:"birth_at"`
	ModifiedAt time.Time               `bson:"modified_at" json:"modified_at"`
}

type ZoneLoadRequest struct {
	AreaIDArray []AreaID   `bson:"area_id_array" json:"area_id_array"` //Area IDs
	ZoneIDArray []ZoneID   `bson:"zone_id_array" json:"zone_id_array"` //Zone IDs
	Page        *load.Page `bson:"page" json:"page"`
	With        *struct {
		Deep bool `bson:"deep" json:"deep"`
	} `bson:"with" json:"with"`
}
