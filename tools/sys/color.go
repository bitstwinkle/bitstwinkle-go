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

package sys

import (
	"fmt"
	"github.com/gookit/color"
	"strings"
)

func Info(arg ...any) {
	color.Style{color.FgBlue}.Println(doFormat(arg...))
}

func Success(arg ...any) {
	color.Style{color.FgGreen}.Println(doFormat(arg...))
}

func Warn(arg ...any) {
	color.Style{color.FgYellow}.Println(doFormat(arg...))
}

func Error(arg ...any) {
	color.Style{color.FgRed}.Println(doFormat(arg...))
}

func doFormat(args ...any) string {
	if len(args) == 0 {
		return ""
	}
	buf := strings.Builder{}
	lenIdx := 0
	for _, item := range args {
		str := fmt.Sprintf("%v", item)
		lenIdx += len(str)
		buf.WriteString(str)
	}
	for i := lenIdx; i < 100; i++ {
		buf.WriteString(" ")
	}
	msg := buf.String()
	return "[ " + msg + " ]"
}
