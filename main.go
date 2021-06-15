/*
Copyright © 2021 Samuel Rios Carvalho <samuel.rios.carvalho@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"carcereiro/cmd"
	"fmt"

	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigType("toml")              // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("$HOME/.carcereiro") // call multiple times to add many search paths
	err := viper.ReadInConfig()              // Find and read the config file
	if err != nil {                          // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	cmd.Execute()
}
