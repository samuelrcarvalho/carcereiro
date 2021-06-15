package cmd

import (
	"fmt"

	"github.com/spf13/viper"
)

func leConf() {

	fmt.Println(viper.Get("port"))
}
