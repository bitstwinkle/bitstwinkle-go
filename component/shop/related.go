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

package shop

import (
	"github.com/bitstwinkle/bitstwinkle-go/domains/category"
	"github.com/bitstwinkle/bitstwinkle-go/types/ref"
)

func GetCollar(shopID ID) ref.Collar {
	return ref.Collar{
		Code: SHOP,
		ID:   shopID,
	}
}

// GenCategoryLead 生成类目使用的Lead
func GenCategoryLead(shopID ID) ref.Lead {
	return ref.LeadOfOneToMany(GetCollar(shopID), SHOP)
}

// GenCartLead 生成购物车使用的Lead
func GenCartLead(shopID ID) ref.Lead {
	return ref.LeadOfOneToOne(GetCollar(shopID))
}

// GenCapitalAccountLead 生成资金账户用Lead
func GenCapitalAccountLead(shopID ID) ref.Lead {
	return ref.LeadOfOneToMany(GetCollar(shopID), SHOP)
}

// GenOrderLead 生成订单用Lead
func GenOrderLead(shopID ID) ref.Lead {
	return ref.LeadOfOneToMany(GetCollar(shopID), SHOP)
}

// GenVwhLead 生成VWH用Lead
func GenVwhLead(shopID ID) ref.Lead {
	return ref.LeadOfOneToMany(GetCollar(shopID), SHOP)
}

// GenVmcLead 生成VMC用Lead
func GenVmcLead(categoryID category.ID) ref.Lead {
	return ref.LeadOfOneToOne(category.GetCollar(categoryID))
}
