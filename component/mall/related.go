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

package mall

import "github.com/bitstwinkle/bitstwinkle-go/types/ref"

func GetCollar(mallID ID) ref.Collar {
	return ref.Collar{
		Code: MALL,
		ID:   mallID,
	}
}

// GenShopLead 生成Shop使用的Lead
func GenShopLead(mallID ID) ref.Lead {
	return ref.LeadOfOneToMany(GetCollar(mallID), MALL)
}

// GenCartLead 生成购物车使用的Lead
func GenCartLead(mallID ID) ref.Lead {
	return ref.LeadOfOneToOne(GetCollar(mallID))
}

// GenCapitalAccountLead 生成资金账户用Lead
func GenCapitalAccountLead(mallID ID) ref.Lead {
	return ref.LeadOfOneToMany(GetCollar(mallID), MALL)
}
