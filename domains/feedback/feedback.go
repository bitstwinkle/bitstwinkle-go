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
	"github.com/bitstwinkle/bitstwinkle-go/types/ww"
	"time"
)

type ID = string

type Comment struct {
	ID      string     `json:"id"`             //评论ID
	Handler ref.Collar `json:"handler"`        //处理者
	Comment string     `json:"comment"`        //评论
	Ctrl    ctrl.Ctrl  `json:"ctrl,omitempty"` //控制信息
	More    more.More  `json:"more,omitempty"` //扩展数据信息
	LastAt  time.Time  `json:"last_at"`        //最后修改时间
}

type Feedback struct {
	Scope        ref.Scope  `json:"scope"`               //所属业务域
	ID           ID         `json:"id"`                  //唯一ID
	Submitter    ref.Collar `json:"submitter"`           //提交者
	Feedback     string     `json:"feedback"`            //反馈信息
	Info         more.More  `json:"memo_more,omitempty"` //更多备注信息
	Media        media.More `json:"media,omitempty"`     //主图主视频
	Status       state.Code `json:"status"`              //状态信息
	Ctrl         *ctrl.Ctrl `bson:"ctrl" json:"ctrl"`
	CommentArray []*Comment `json:"comment_array,omitempty"` //评论记录
	BirthAt      time.Time  `json:"birth_at"`                //创建时间
	ModifiedAt   time.Time  `json:"modified_at"`             //最后修改时间
}

type SubmitRequest struct {
	Scope     ref.Scope   `bson:"scope" json:"scope"`         //[*]Scope
	Lead      *ref.Lead   `bson:"lead" json:"lead"`           //[*]Lead
	Submitter ref.Collar  `bson:"submitter" json:"submitter"` //[*]Submitter
	Feedback  string      `bson:"feedback" json:"feedback"`   //[*]Feedback
	Info      more.Array  `bson:"info" json:"info"`           //[-]More Memo
	Media     media.Array `bson:"media" json:"media"`         //[-]Media
	Status    state.Code  `bson:"status" json:"status"`       //[-]Status
	Ctrl      *ctrl.Ctrl  `bson:"ctrl" json:"ctrl"`           //[-]Ctrl
}

type GetRequest struct {
	Scope      ref.Scope   `bson:"scope" json:"scope"` //[*]Scope
	By         load.ByCode `bson:"by" json:"by"`       //BY: lead|id
	FeedbackID ID          `bson:"feedback_id" json:"feedback_id"`
	Lead       *ref.Lead   `bson:"lead" json:"lead"`
}

type LoadRequest struct {
	Scope           ref.Scope       `bson:"scope" json:"scope"`           //[*]SCOPE
	LeadArray       []ref.Lead      `bson:"lead_array" json:"lead_array"` //Lead info
	Submitter       []*ref.Collar   `bson:"submitter" json:"submitter"`   //Submitter
	FeedbackIDArray []ID            `bson:"id_array" json:"id_array"`     //Brand ID info
	CtrlTag         []string        `bson:"ctrl_tag" json:"ctrl_tag"`     //Ctrl Tag
	Keyword         *ctrl.StringSet `bson:"keyword" json:"keyword"`       //Key Word
	Page            *load.Page      `bson:"page" json:"page"`             //Page
}

type AdvanceRequest struct {
	Scope      ref.Scope  `bson:"scope" json:"scope"` //[*]Scope
	FeedbackID string     `json:"feedback_id"`        //[*]反馈ID
	Handler    ref.Collar `json:"handler"`            //[*]处理者
	Status     state.Code `json:"status"`             //[*]推进到的状态
	Comment    string     `json:"comment"`            //[*]评论
	More       more.Array `json:"more,omitempty"`     //[-]扩展数据信息
}

type Service interface {
	Submit(permit *ww.Permit, req SubmitRequest) (feedbackID *ID, err *errors.Error)
	Get(permit *ww.Permit, req GetRequest) (feedbackM *Feedback, err *errors.Error)
	Load(permit *ww.Permit, req LoadRequest) (arr []*Feedback, paging *load.Paging, err *errors.Error)
	Advance(permit *ww.Permit, req AdvanceRequest) (feedbackID *ID, err *errors.Error)
}
