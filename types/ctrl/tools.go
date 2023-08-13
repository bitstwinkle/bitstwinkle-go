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

package ctrl

// Mix 分析更改
// *Ctrl: 新的Ctrl
// bool: 是否有变化
func Mix(ctrlM *Ctrl, setM *Set) (*Ctrl, bool) {
	if setM == nil {
		return ctrlM, false
	}
	if ctrlM == nil {
		return &Ctrl{
			Tags:  setM.AddTag,
			Paras: setM.CtrlPara,
		}, true
	}
	newCtrl := &Ctrl{
		Tags:  ctrlM.Tags,
		Paras: ctrlM.Paras,
	}

	// # Handle Tag
	newTagDict := make(map[string]bool)
	if len(newCtrl.Tags) > 0 {
		for _, t := range newCtrl.Tags {
			newTagDict[t] = true
		}
	}
	if len(setM.AddTag) > 0 {
		for _, t := range setM.AddTag {
			newTagDict[t] = true
		}
	}
	if len(setM.RmvTag) > 0 {
		for _, rt := range setM.RmvTag {
			delete(newTagDict, rt)
		}
	}
	newTagArr := make([]string, len(newTagDict))
	i := 0
	for k, _ := range newTagDict {
		newTagArr[i] = k
		i += 1
	}
	newCtrl.Tags = newTagArr

	// # Handle Para
	newParaDict := make(map[string]Para)
	if len(newCtrl.Paras) > 0 {
		for _, fPara := range newCtrl.Paras {
			newParaDict[fPara.Key] = fPara
		}
	}
	if len(setM.CtrlPara) > 0 {
		for _, uPara := range setM.CtrlPara {
			newParaDict[uPara.Key] = uPara
		}
	}
	if len(setM.RmvPara) > 0 {
		for _, rp := range setM.RmvPara {
			delete(newParaDict, rp)
		}
	}
	newParaArr := make([]Para, len(newParaDict))
	i = 0
	for _, v := range newParaDict {
		newParaArr[i] = v
		i += 1
	}
	newCtrl.Paras = newParaArr
	return newCtrl, true
}
