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

package strs

import (
	"bitstwinkle-go/types/errors"
	"time"
)

const (
	TimeLayout = "2006-01-02 15:04:05"
	DtLayout   = "2006-01-02"
)

// DtOf YYYY-MM-DD
func DtOf(str string) (time.Time, *errors.Error) {
	dt, err := time.Parse(DtLayout, str)
	if err != nil {
		return time.Time{}, errors.Verify("invalid datetime format: " + str)
	}
	return dt, nil
}

// TimeOf YYYY-MM-DD HH:mm:SS
func TimeOf(str string) (time.Time, error) {
	dt, err := time.Parse(TimeLayout, str)
	if err != nil {
		return time.Time{}, err
	}
	return dt, nil
}

func TimeTo(t time.Time) string {
	return t.Format(TimeLayout)
}

func DtTo(t time.Time) string {
	return t.Format(DtLayout)
}

func GetTimeLocation(code string) *time.Location {
	loc, _ := time.LoadLocation(code)
	return loc
}

func GetPlatformCreateTime() time.Time {
	return time.Date(2021, 10, 22, 8, 8, 8, 8, GetTimeLocation("Asia/Shanghai"))
}

func GetForever() time.Time {
	return GetPlatformCreateTime().AddDate(1000, 0, 0)
}

func NewTimestampAutoID() uint64 {
	return uint64(time.Now().UnixNano())
}
