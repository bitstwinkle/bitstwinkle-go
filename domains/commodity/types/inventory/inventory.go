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

package inventory

import (
	"github.com/bitstwinkle/bitstwinkle-go/types/errors"
	"github.com/bitstwinkle/bitstwinkle-go/types/times"
	"time"
)

const Unlimited int64 = 9999999999

type Quota struct {
	Limited bool  `bson:"limited" json:"limited"` //是否限制
	Quota   int64 `bson:"value" json:"value"`     //限制的值
}

func (q Quota) Verify() *errors.Error {
	if q.Limited && q.Quota <= 0 {
		return errors.Verify("the quota is limited, but the quota size is not set")
	}
	return nil
}

func (q Quota) InitBalance() int64 {
	if !q.Limited {
		return Unlimited
	}
	return q.Quota
}

func (q Quota) Same(other Quota) bool {
	return q.Limited == other.Limited && q.Quota == other.Quota
}

// Plan 库存配额
type Plan struct {
	Total     Quota `bson:"total" json:"total"`         //总配额
	Daily     Quota `bson:"daily" json:"daily"`         //日配额
	Weekly    Quota `bson:"weekly" json:"weekly"`       //周配额
	Monthly   Quota `bson:"monthly" json:"monthly"`     //月配额
	Quarterly Quota `bson:"quarterly" json:"quarterly"` //季配额
	Yearly    Quota `bson:"yearly" json:"yearly"`       //年配额
}

func (p Plan) Dict() map[string]Quota {
	return map[string]Quota{
		"total":     p.Total,
		"daily":     p.Daily,
		"weekly":    p.Weekly,
		"monthly":   p.Monthly,
		"quarterly": p.Quarterly,
		"yearly":    p.Yearly,
	}
}

func (p Plan) Verify() *errors.Error {
	dict := p.Dict()
	for k, v := range dict {
		if err := v.Verify(); err != nil {
			return errors.Verify("invalid plan." + k + ": " + err.Error())
		}
	}
	return nil
}

func (p Plan) Diff(newPlan Plan) map[string]Quota {
	pDict := p.Dict()
	oDict := newPlan.Dict()
	diff := make(map[string]Quota)
	for pK, pV := range pDict {
		oV := oDict[pK]
		if pV.Same(oV) {
			continue
		}
		diff[pK] = oV
	}
	return diff
}

// Inventory 库存信息
type Inventory struct {
	Plan      Plan      `bson:"plan" json:"plan"`           //配额计划信息
	Total     int64     `bson:"total" json:"total"`         //历史总库存
	Daily     int64     `bson:"daily" json:"daily"`         //当日剩余库存
	Weekly    int64     `bson:"weekly" json:"weekly"`       //周剩余库存
	Monthly   int64     `bson:"monthly" json:"monthly"`     //月剩余库存
	Quarterly int64     `bson:"quarterly" json:"quarterly"` //季度剩余库存
	Yearly    int64     `bson:"yearly" json:"yearly"`       //年剩余库存
	LastAt    time.Time `bson:"last_at" json:"last_at"`     //最后更新时间
}

// Of 通过配额初始化库存信息
func Of(p Plan) (Inventory, *errors.Error) {
	if err := p.Verify(); err != nil {
		return Inventory{}, err
	}
	return Inventory{
		Plan:      p,
		Total:     p.Total.InitBalance(),
		Daily:     p.Daily.InitBalance(),
		Weekly:    p.Weekly.InitBalance(),
		Monthly:   p.Monthly.InitBalance(),
		Quarterly: p.Quarterly.InitBalance(),
		Yearly:    p.Yearly.InitBalance(),
		LastAt:    time.Now(),
	}, nil
}

func (inv Inventory) Dict() map[string]int64 {
	return map[string]int64{
		"total":     inv.Total,
		"daily":     inv.Daily,
		"weekly":    inv.Weekly,
		"monthly":   inv.Monthly,
		"quarterly": inv.Quarterly,
		"yearly":    inv.Yearly,
	}
}

func (inv Inventory) doSetQuota(quota int64, whichInv *int64, whichQuota *Quota) int64 {
	if quota < 0 {
		quota = 0
	}
	var inc int64 = 0
	if !whichQuota.Limited {
		whichQuota.Limited = true
		inc = quota
	} else {
		inc = quota - whichQuota.Quota
	}

	whichQuota.Quota = quota
	*whichInv += inc
	return inc
}

func (inv Inventory) SetTotalQuota(quota int64) int64 {
	return inv.doSetQuota(quota, &inv.Total, &inv.Plan.Total)
}

func (inv Inventory) SetDailyQuota(quota int64) int64 {
	return inv.doSetQuota(quota, &inv.Daily, &inv.Plan.Daily)
}

func (inv Inventory) SetWeeklyQuota(quota int64) int64 {
	return inv.doSetQuota(quota, &inv.Weekly, &inv.Plan.Weekly)
}

func (inv Inventory) SetMonthlyQuota(quota int64) int64 {
	return inv.doSetQuota(quota, &inv.Monthly, &inv.Plan.Monthly)
}

func (inv Inventory) SetQuarterlyQuota(quota int64) int64 {
	return inv.doSetQuota(quota, &inv.Quarterly, &inv.Plan.Quarterly)
}

func (inv Inventory) SetYearlyQuota(quota int64) int64 {
	return inv.doSetQuota(quota, &inv.Yearly, &inv.Plan.Yearly)
}

func (inv Inventory) Use(q int64) {
	inv.Total -= q
	now := time.Now()
	if times.SameDay(inv.LastAt, now) {
		inv.Daily -= q
	}
	if times.SameWeek(inv.LastAt, now) {
		inv.Weekly -= q
	}
	if times.SameMonth(inv.LastAt, now) {
		inv.Monthly -= q
	}
	if times.SameQuarter(inv.LastAt, now) {
		inv.Quarterly -= q
	}
	if times.SameYear(inv.LastAt, now) {
		inv.Yearly -= q
	}
}

// Diff 比较两个库存信息的变化,并返回更新后的数据
func (inv Inventory) Diff(newInv Inventory) (map[string]Quota, map[string]int64) {
	quotaDiff := inv.Plan.Diff(newInv.Plan)
	invDict := inv.Dict()
	newDict := newInv.Dict()
	diff := make(map[string]int64)
	for iK, iV := range invDict {
		nV := newDict[iK]
		if iV == nV {
			continue
		}
		diff[iK] = nV
	}
	return quotaDiff, diff
}
