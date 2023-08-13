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

package media

// Mix 分析更改
// *media.More: 新的More
// bool: 是否有变化
func Mix(moreM More, setM *Set) (More, bool) {
	if setM == nil {
		return moreM, false
	}
	if moreM == nil {
		return OfArray(setM.Newest), true
	}
	newMore := NewMore()
	setMore := OfArray(setM.Newest)
	for k, v := range setMore {
		newMore[k] = v
	}
	if len(setM.Removed) > 0 {
		for _, k := range setM.Removed {
			delete(newMore, k)
		}
	}
	return newMore, true
}
