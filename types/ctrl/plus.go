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

package ctrl

type BooleanSet struct {
	Yes   bool `bson:"yes" json:"yes"` //是否设置
	Value bool `json:"value"`          //对应值
}

type Int64Set struct {
	Yes   bool  `bson:"yes" json:"yes"` //是否设置
	Value int64 `json:"value"`          //对应值
}

type StringSet struct {
	Yes   bool   `bson:"yes" json:"yes"` //是否设置
	Value string `json:"value"`          //对应值
}

type StringArraySet struct {
	Yes   bool     `bson:"yes" json:"yes"`
	Value []string `bson:"value" json:"value"`
}

type TagSet = StringArraySet
