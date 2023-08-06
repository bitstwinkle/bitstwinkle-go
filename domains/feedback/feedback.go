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

package feedback

import (
	"github.com/bitstwinkle/bitstwinkle-go/types/collections/more"
	"github.com/bitstwinkle/bitstwinkle-go/types/ctrl"
	"github.com/bitstwinkle/bitstwinkle-go/types/errors"
	"github.com/bitstwinkle/bitstwinkle-go/types/load"
	"github.com/bitstwinkle/bitstwinkle-go/types/ref"
	"github.com/bitstwinkle/bitstwinkle-go/types/state"
	"github.com/bitstwinkle/bitstwinkle-go/types/view/media"
)

type Comment struct {
	ID      string     `json:"id"`             //评论ID
	Handler ref.Collar `json:"handler"`        //处理者
	Comment string     `json:"comment"`        //评论
	Ctrl    ctrl.Ctrl  `json:"ctrl,omitempty"` //控制信息
	More    more.More  `json:"more,omitempty"` //扩展数据信息
	LastAt  string     `json:"last_at"`        //最后修改时间
}

type Feedback struct {
	Scope        ref.Scope    `json:"scope"`                   //所属业务域
	ID           string       `json:"id"`                      //唯一ID
	Submitter    ref.Collar   `json:"submitter"`               //提交者
	Acceptor     ref.Collar   `json:"acceptor"`                //受理者
	Feedback     string       `json:"feedback"`                //反馈信息
	MemoMore     more.More    `json:"memo_more,omitempty"`     //更多备注信息
	Media        *media.Media `json:"media,omitempty"`         //主图主视频
	MediaMore    media.More   `json:"media_more,omitempty"`    //更多图视频
	Status       state.Code   `json:"status"`                  //状态信息
	CommentArray []Comment    `json:"comment_array,omitempty"` //评论记录
	BirthAt      string       `json:"birth_at"`                //创建时间
	ModifiedAt   string       `json:"modified_at"`             //最后修改时间
}

type SubmitRequest struct {
	IdemID    string       `json:"idem_id"`              //[*]幂等ID
	Scope     ref.Scope    `json:"scope"`                //[*]所属业务域
	Submitter ref.Collar   `json:"submitter"`            //[*]提交者
	Acceptor  ref.Collar   `json:"acceptor"`             //[*]受理者
	Feedback  string       `json:"feedback"`             //[*]备注信息
	MemoMore  more.Array   `json:"memo_more,omitempty"`  //[-]更多备注信息
	Media     *media.Media `json:"media,omitempty"`      //[-]主图主视频
	MediaMore media.Array  `json:"media_more,omitempty"` //[-]更多图视频
	Status    state.Code   `json:"status"`               //[-]指定自定义状态信息
}

type LoadRequest struct {
	By        load.ByCode `json:"by"` // scope|submitter|acceptor
	Scope     *ref.Scope  `json:"scope"`
	Submitter *ref.Collar `json:"submitter,omitempty"`
	Acceptor  *ref.Collar `json:"acceptor,omitempty"`
	Page      *load.Page  `json:"page"`
}

type AdvanceRequest struct {
	IdemID     string     `json:"idem_id"`        //[*]幂等ID
	FeedbackID string     `json:"feedback_id"`    //[*]反馈ID
	Handler    ref.Collar `json:"handler"`        //[*]处理者
	Status     state.Code `json:"status"`         //[*]推进到的状态
	Comment    string     `json:"comment"`        //[*]评论
	More       more.Array `json:"more,omitempty"` //[-]扩展数据信息
}

type Service interface {
	Submit(req SubmitRequest) (feedback *Feedback, err *errors.Error)
	Load(req LoadRequest) (arr []Feedback, paging *load.Paging, err *errors.Error)
	Advance(req AdvanceRequest) (feedback *Feedback, err *errors.Error)
}
