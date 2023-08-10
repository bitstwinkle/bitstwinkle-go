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

package volume

import (
	"github.com/bitstwinkle/bitstwinkle-go/types/errors"
	"github.com/bitstwinkle/bitstwinkle-go/types/times"
	"time"
)

type Volume struct {
	Total     int64     `bson:"total" json:"total"`         //历史总销量
	Daily     int64     `bson:"daily" json:"daily"`         //日销量
	Weekly    int64     `bson:"weekly" json:"weekly"`       //周销量
	Monthly   int64     `bson:"monthly" json:"monthly"`     //月销量
	Quarterly int64     `bson:"quarterly" json:"quarterly"` //季销量
	Yearly    int64     `bson:"yearly" json:"yearly"`       //年销量
	LastAt    time.Time `bson:"last_at" json:"last_at"`     //上次更新时间
}

func (m *Volume) Init(vol Volume) {
	m.Total = vol.Total
	m.Daily = vol.Daily
	m.Weekly = vol.Weekly
	m.Monthly = vol.Monthly
	m.Quarterly = vol.Quarterly
	m.Yearly = vol.Yearly
	m.LastAt = vol.LastAt
	if m.LastAt.IsZero() {
		m.LastAt = time.Now()
	}
}

func (m *Volume) Verify() *errors.Error {
	if m.Total < 0 ||
		m.Daily < 0 ||
		m.Weekly < 0 ||
		m.Monthly < 0 ||
		m.Quarterly < 0 ||
		m.Yearly < 0 {
		return errors.Verify("all must > 0")
	}
	if m.LastAt.IsZero() {
		m.LastAt = time.Now()
	}
	return nil
}

func (m *Volume) Inc(q int64) {
	m.Total += q
	now := time.Now()
	if times.SameDay(now, m.LastAt) {
		m.Daily += q
	} else {
		m.Daily = q
	}
	if times.SameWeek(now, m.LastAt) {
		m.Weekly += q
	} else {
		m.Weekly = q
	}
	if times.SameMonth(now, m.LastAt) {
		m.Monthly += q
	} else {
		m.Monthly = q
	}
	if times.SameQuarter(now, m.LastAt) {
		m.Quarterly += q
	} else {
		m.Quarterly = q
	}
	if times.SameYear(now, m.LastAt) {
		m.Yearly += q
	} else {
		m.Yearly = q
	}
	m.LastAt = now
}
