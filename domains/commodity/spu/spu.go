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

/*
 * 杭州菩公英科技有限公司版权所有
 * 作者: 川谷
 * 时间: 2023/8/2
 * --------------------------------------
 * ******* 给生命以时光,给岁月以欢畅 ********
 * --------------------------------------
 */

package spu

import (
	"github.com/bitstwinkle/bitstwinkle-go/types/collections/more"
	"github.com/bitstwinkle/bitstwinkle-go/types/ctrl"
	"github.com/bitstwinkle/bitstwinkle-go/types/ref"
	"github.com/bitstwinkle/bitstwinkle-go/types/view/label"
	"github.com/bitstwinkle/bitstwinkle-go/types/view/media"
)

type Spu struct {
	Scope      ref.Scope   `json:"scope"`           //所属业务域
	CategoryID string      `json:"category_id"`     //所属类目
	Title      string      `json:"title"`           //标题
	Alias      []string    `json:"alias,omitempty"` //别名
	Info       more.More   `json:"info,omitempty"`  //介绍
	Media      media.More  `json:"media"`           //图片视频
	Ctrl       ctrl.Ctrl   `json:"ctrl"`            //控制信息
	Label      label.Array `json:"label"`           //标签
	//Spec       []*spec.DefItem `json:"spec"`            //规格定义
}
