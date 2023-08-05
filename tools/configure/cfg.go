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

package configure

import (
	"github.com/bitstwinkle/bitstwinkle-go/types/errors"
	"github.com/bitstwinkle/bitstwinkle-go/types/strs"
	"github.com/spf13/viper"
	"time"
)

func GetString(key string, def ...string) string {
	ok := exists(key)
	if ok {
		return viper.GetString(key)
	}
	if len(def) > 0 {
		return def[0]
	}
	return strs.EMPTY
}

func MustGetString(key string) (string, *errors.Error) {
	val := GetString(key)
	if len(val) == 0 {
		return val, errors.Sys("Miss configure: " + key + "! " +
			"This parameter should be set in an environment variable, startup parameter, or configuration file.")
	}
	return val, nil
}

func GetBool(key string, def ...bool) bool {
	ok := exists(key)
	if ok {
		return viper.GetBool(key)
	}
	if len(def) > 0 {
		return def[0]
	}
	return false
}

func GetInt(key string, def ...int) int {
	ok := exists(key)
	if ok {
		return viper.GetInt(key)
	}
	if len(def) > 0 {
		return def[0]
	}
	return 0
}

func GetUint64(key string, def ...uint64) uint64 {
	ok := exists(key)
	if ok {
		return viper.GetUint64(key)
	}
	if len(def) > 0 {
		return def[0]
	}
	return 0
}

func GetDuration(key string, def ...time.Duration) time.Duration {
	ok := exists(key)
	if ok {
		return viper.GetDuration(key)
	}
	if len(def) > 0 {
		return def[0]
	}
	return 0
}

var gConfigs = make(map[string]any)

func register(key string, val any) {
	gConfigs[key] = val
}

// Dump Used to output all used configurations
func Dump(call func(key string, val any)) {
	for k, v := range gConfigs {
		call(k, v)
	}
}

func exists(key string) bool {
	val := viper.Get(key)
	if val != nil {
		register(key, val)
	}
	return val != nil
}

func init() {
	viper.AddConfigPath("./_bin/")
	viper.AddConfigPath("./")
	viper.SetConfigName(GetString("sys.mode", "local"))
	viper.SetConfigType("yaml")
	_ = viper.ReadInConfig()
	viper.AutomaticEnv()
}
