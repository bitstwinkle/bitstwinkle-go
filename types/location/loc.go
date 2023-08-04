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

package location

import (
	"fmt"
	"github.com/bitstwinkle/bitstwinkle-go/types/errors"
)

type Loc struct {
	Lon float64 `bson:"lon" json:"lon"`
	Lat float64 `bson:"lat" json:"lat"`
}

func (loc *Loc) Verify() *errors.Error {
	if loc.Lon < -180 || loc.Lon > 180 {
		return errors.Assert("(-180 - 180)", fmt.Sprintf("%.4f", loc.Lon))
	}

	if loc.Lat < -90 || loc.Lat > 90 {
		return errors.Assert("(-90 - 90)", fmt.Sprintf("%.4f", loc.Lat))
	}
	return nil
}

func LocOf(lon float64, lat float64) Loc {
	return Loc{
		Lon: lon,
		Lat: lat,
	}
}

type Item struct {
	Code string `bson:"code" json:"code"`
	Seq  int    `bson:"seq" json:"seq"`
	Loc  Loc    `bson:"loc" json:"loc"`
}

func (m Item) GetCode() string {
	return m.Code
}

func (m Item) GetSeq() int {
	return m.Seq
}

func (m Item) Verify() *errors.Error {
	return m.Loc.Verify()
}

type Set struct {
	Yes bool `bson:"yes" json:"yes"`
	Loc *Loc `bson:"loc" json:"loc"`
}

type ExLocSet struct {
	Yes     bool     `bson:"yes" json:"yes"`
	ExLoc   []Item   `bson:"exLoc" json:"exLoc"`     //新增或者更新的
	Removed []string `bson:"removed" json:"removed"` //需要删除的
}
