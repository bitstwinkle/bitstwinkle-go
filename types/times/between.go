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
	"fmt"
	"github.com/bitstwinkle/bitstwinkle-go/types/errors"
	"time"
)

type Between struct {
	Start time.Time `bson:"start" json:"start"`
	End   time.Time `bson:"end" json:"end"`
}

func (btw Between) Verify() *errors.Error {
	if btw.Start.After(btw.End) {
		return errors.Verify("invalid between, start after end")
	}
	return nil
}

// HM 时间和分钟定义 HOUR MINUTE
type HM struct {
	Hour   int `bson:"h" json:"h"`
	Minute int `bson:"m" json:"m"`
}

func (hm HM) Verify() *errors.Error {
	if hm.Hour < 0 || hm.Hour > 23 {
		return errors.Verify("invalid hour, must be [0-24)")
	}
	if hm.Minute < 0 || hm.Minute > 59 {
		return errors.Verify("invalid minute, must be [0-60)")
	}
	return nil
}

func (hm HM) String() string {
	return fmt.Sprintf("%02d:%02d", hm.Hour, hm.Minute)
}

func (hm HM) Key() string {
	return fmt.Sprintf("%02d%02d", hm.Hour, hm.Minute)
}

type HMBetween struct {
	Start HM `bson:"start" json:"start"`
	End   HM `bson:"end" json:"end"`
}

func (btw HMBetween) Verify() *errors.Error {
	if err := btw.Start.Verify(); err != nil {
		return errors.Verify("invalid start: " + err.Error())
	}
	if err := btw.End.Verify(); err != nil {
		return errors.Verify("invalid end: " + err.Error())
	}
	if (btw.Start.Hour > btw.End.Hour) ||
		(btw.Start.Hour == btw.End.Hour && btw.Start.Minute > btw.End.Minute) {
		return errors.Verify("invalid between, start after end")
	}

	return nil
}

func (btw HMBetween) String() string {
	return fmt.Sprintf("%s - %s", btw.Start.String(), btw.End.String())
}

// Key 时间班次号,可用于排班
func (btw HMBetween) Key() string {
	return fmt.Sprintf("%s_%s", btw.Start.Key(), btw.End.Key())
}
