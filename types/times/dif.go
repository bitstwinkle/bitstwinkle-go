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

package times

import "time"

func SameDay(a time.Time, b time.Time) bool {
	return a.Year() == b.Year() &&
		a.Month() == b.Month() &&
		a.Day() == b.Day()
}

func SameWeek(a time.Time, b time.Time) bool {
	aYear, aWeek := a.ISOWeek()
	bYear, bWeek := b.ISOWeek()
	return aYear == bYear &&
		aWeek == bWeek
}

func SameMonth(a time.Time, b time.Time) bool {
	return a.Year() == b.Year() &&
		a.Month() == b.Month()
}

func SameQuarter(a time.Time, b time.Time) bool {
	aQ := (a.Month()-1)/3 + 1
	bQ := (b.Month()-1)/3 + 1
	return a.Year() == b.Year() &&
		aQ == bQ
}

func SameYear(a time.Time, b time.Time) bool {
	return a.Year() == b.Year()
}
