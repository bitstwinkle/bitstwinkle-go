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
 *
 */

package coin

import (
	"github.com/bitstwinkle/bitstwinkle-go/types/collections/more"
	"github.com/bitstwinkle/bitstwinkle-go/types/view/media"
	"time"
)

type Coin = string

// Benchmark 游戏币的基准价值
type Benchmark struct {
	Provisions int64  `bson:"provisions" json:"provisions"`
	Currency   string `bson:"currency" json:"currency"`
}

// Information 游戏币定义
// 只有VN守护者才有发行游戏币的资格,单个价值网络中发行的游戏币种类不超过10种
type Information struct {
	VN          string      `bson:"vn" json:"vn"`                   //所属价值网络
	Coin        string      `bson:"coin" json:"coin"`               //COIN 编码|全网唯一
	Name        string      `bson:"name" json:"name"`               //COIN NAME
	Icon        media.Media `bson:"icon" json:"icon"`               //图标
	Issuance    int64       `bson:"issuance" json:"issuance"`       //发行量
	Circulation int64       `bson:"circulation" json:"circulation"` //流通量
	Info        more.More   `bson:"info" json:"info"`               //展示信息
	Media       media.More  `bson:"media" json:"media"`             //图文视频信息
	Available   bool        `bson:"available" json:"available"`     //是否通行
	BirthAt     time.Time   `bson:"birth_at" json:"birth_at"`       //创建时间
	ModifiedAt  time.Time   `bson:"modified_at" json:"modified_at"` //最后修改时间
}

// MintageRequest 铸币
type MintageRequest struct {
	Coin     Coin  `bson:"coin" json:"coin"`
	Quantity int64 `bson:"quantity" json:"quantity"`
}
