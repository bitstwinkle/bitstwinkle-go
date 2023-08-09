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

type AreaID = string

type Area struct {
	ID   AreaID `bson:"id" json:"id"` //Area ID
	Tree struct {
		Stair  int      `bson:"stair" json:"stair"`   //Level
		Parent string   `bson:"parent" json:"parent"` //Parent ID
		Path   []AreaID `bson:"path" json:"path"`     //FULL PATH
	} `bson:"tree" json:"tree"` //The Area Tree
	ExLink     ref.LinkMore  `bson:"ex_link,omitempty" json:"ex_link,omitempty"` //Third Link INfo
	Title      string        `bson:"title" json:"title"`                         //Name
	Address    string        `bson:"address" json:"address"`                     //Full Address String
	Path       []Area        `bson:"path" json:"path"`                           //Full Path
	Loc        *location.Loc `bson:"loc,omitempty" json:"loc,omitempty"`         //Location Info
	Available  bool          `bson:"available" json:"available"`                 //Available
	BirthAt    time.Time     `bson:"birth_at" json:"birth_at"`                   //Created Time
	ModifiedAt time.Time     `bson:"modified_at" json:"modified_at"`             //Last Modify
}

type AreaLoadRequest struct {
	AreaIDArray []AreaID   `bson:"area_id_array" json:"area_id_array"` //Area IDs
	Page        *load.Page `bson:"page" json:"page"`                   //page info
}
