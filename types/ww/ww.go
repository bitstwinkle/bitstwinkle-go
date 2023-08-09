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

package ww

// VN The Value Net
type VN = string

// JD The Value Node
type JD = string

// Living The Value Living, user, cat, dog ...
type Living = string

// Permit : Passage permit
type Permit struct {
	VN     VN `bson:"vn" json:"vn"`
	JD     JD `bson:"jd" json:"jd"`
	Living `bson:"living" json:"living"`
}

func (p *Permit) String() string {
	return "bitstwinkle://" + p.VN + "/" + p.JD + "/" + p.Living
}
