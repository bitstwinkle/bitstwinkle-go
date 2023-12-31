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

type Para struct {
	Key   string `bson:"key" json:"key"`
	Value string `bson:"val" json:"val"`
}

// Ctrl 数据控制器,用于给数据进行控制打标或者设置参数
type Ctrl struct {
	Tags  []string `bson:"tags,omitempty" json:"tags,omitempty"`
	Paras []Para   `bson:"paras,omitempty" json:"paras,omitempty"`
}

func New() *Ctrl {
	return &Ctrl{
		Tags:  []string{},
		Paras: []Para{},
	}
}

func (c *Ctrl) WithTag(tag ...string) *Ctrl {
	for _, it := range tag {
		exist := false
		if len(c.Tags) > 0 {
			for _, t := range c.Tags {
				if t == it {
					exist = true
					break
				}
			}
		}
		if exist {
			continue
		}
		c.Tags = append(c.Tags, it)
	}
	return c
}

func (c *Ctrl) WithPara(key string, value string) *Ctrl {
	c.Paras = append(c.Paras, Para{
		Key:   key,
		Value: value,
	})
	return c
}

type Set struct {
	Yes      bool     `json:"yes"`                     //是否设置
	AddTag   []string `json:"add_ctrl_tag,omitempty"`  //[|]控制标中需要新增的标
	RmvTag   []string `json:"rmv_ctrl_tag,omitempty"`  //[|]控制标中需要删除的标
	CtrlPara []Para   `json:"ctrl_para,omitempty"`     //[|]控制参数中需要新增或者更新的参数
	RmvPara  []string `json:"rmv_ctrl_para,omitempty"` //[|]控制参数中需要删除的参数
}
