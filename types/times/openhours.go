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

package times

import (
	"github.com/bitstwinkle/bitstwinkle-go/types/errors"
	"time"
)

const (
	OpeningHoursNoLimit = "no_limit" //不限
	OpeningHoursByWeek  = "week"     //按周设置
)

type OpeningHours struct {
	Mode        string           `bson:"mode" json:"mode"`                 //设置模式
	WeekSetting *OHByWeekSetting `bson:"week_setting" json:"week_setting"` //基于周设置
}

func (oh OpeningHours) Verify() *errors.Error {
	switch oh.Mode {
	case OpeningHoursNoLimit:
		return nil
	case OpeningHoursByWeek:
		if oh.WeekSetting == nil {
			return errors.Verify("require week_setting")
		}
		return oh.WeekSetting.Verify()
	}
	return errors.Verify("invalid mode: " + oh.Mode)
}

func (oh OpeningHours) GetOpening(dt time.Time) (bool, OHItem) {
	switch oh.Mode {
	case OpeningHoursByWeek:
		return oh.WeekSetting.GetOpening(dt)
	}
	return OHNoLimitSetting{}.GetOpening(dt)
}

func (oh OpeningHours) Next(benchmark time.Time, size int) []OHItem {
	switch oh.Mode {
	case OpeningHoursByWeek:
		return oh.WeekSetting.Next(benchmark, size)
	}
	return OHNoLimitSetting{}.Next(benchmark, size)
}

type OHItem struct {
	Date    time.Time   `bson:"date" json:"date"`       //对应日期
	Opening []HMBetween `bson:"opening" json:"opening"` //营业的时间
}

type OHSetting interface {
	Verify() *errors.Error                       //进行校验
	GetOpening(dt time.Time) (bool, OHItem)      //某天可用的营业时间
	Next(benchmark time.Time, size int) []OHItem //基于参考时间获取接下来的size个营业时间
}

type OHNoLimitSetting struct{}

func (ohl OHNoLimitSetting) Verify() *errors.Error {
	return nil
}

func (ohl OHNoLimitSetting) GetOpening(dt time.Time) (bool, OHItem) {
	return true, OHItem{
		Date: dt,
		Opening: []HMBetween{
			{
				Start: HM{
					Hour:   0,
					Minute: 0,
				},
				End: HM{
					Hour:   23,
					Minute: 59,
				},
			},
		},
	}
}

func (ohl OHNoLimitSetting) Next(benchmark time.Time, size int) []OHItem {
	if size <= 0 {
		return []OHItem{}
	}
	ohItemArr := make([]OHItem, size)
	dt := benchmark
	for i := 0; i < size; i++ {
		_, item := ohl.GetOpening(dt)
		ohItemArr[i] = item
		dt = dt.Add(24 * time.Hour)
	}
	return ohItemArr
}

type OHByWeekSetting struct {
	Exclude        []time.Weekday      `bson:"exclude" json:"exclude"`                 //排除日期
	DefaultSetting []HMBetween         `bson:"default_setting" json:"default_setting"` //默认设置
	SpecialSetting map[int][]HMBetween `bson:"special_setting" json:"special_setting"` //特殊指定
}

func (ohw OHByWeekSetting) Verify() *errors.Error {
	if len(ohw.Exclude) > 0 {
		for _, wd := range ohw.Exclude {
			if wd < time.Sunday || wd > time.Saturday {
				return errors.Verify("invalid exclude: must be [0-6]")
			}
		}
	}
	if len(ohw.DefaultSetting) > 0 {
		for _, hmb := range ohw.DefaultSetting {
			if err := hmb.Verify(); err != nil {
				return errors.Verify("invalid default_setting: " + err.Error())
			}
		}
	}
	if len(ohw.SpecialSetting) > 0 {
		for k, v := range ohw.SpecialSetting {
			if k < int(time.Sunday) || k > int(time.Saturday) {
				return errors.Verify("invalid special_setting: must be [0-6]")
			}
			if len(v) > 0 {
				for _, hmb := range v {
					if err := hmb.Verify(); err != nil {
						return errors.Verify("invalid special_setting: " + err.Error())
					}
				}
			}

		}
	}
	return nil
}

func (ohw OHByWeekSetting) GetOpening(dt time.Time) (bool, OHItem) {
	wd := dt.Weekday()
	if len(ohw.Exclude) > 0 {
		for _, exc := range ohw.Exclude {
			if exc == wd {
				return false, OHItem{}
			}
		}
	}
	if len(ohw.SpecialSetting) > 0 {
		hmBetween, found := ohw.SpecialSetting[int(wd)]
		if found {
			return true, OHItem{
				Date:    dt,
				Opening: hmBetween,
			}
		}
	}
	if len(ohw.DefaultSetting) > 0 {
		return true, OHItem{
			Date:    dt,
			Opening: ohw.DefaultSetting,
		}
	}
	return OHNoLimitSetting{}.GetOpening(dt)
}

func (ohw OHByWeekSetting) Next(benchmark time.Time, size int) []OHItem {
	if size <= 0 {
		return []OHItem{}
	}
	ohItemArr := make([]OHItem, size)
	dt := benchmark
	i := 0
	dti := 0
	for i < size {
		open, ohi := ohw.GetOpening(dt)
		if open {
			ohItemArr[i] = ohi
			i += 1
		}
		dt = dt.Add(24 * time.Hour)
		dti += 1
		if dti > 3*365 {
			break
		}
	}
	return ohItemArr[:i]
}
