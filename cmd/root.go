/*
Copyright © 2021 Samuel Rios Carvalho <Samuel Rios Carvalho>

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
package cmd

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "carcereiro",
	Short: "CLI para liberação de acesso a MySQL",
	Long: `CLI para liberação de acesso a MySQL, 
facilitando assim o trabalho de administradores de sistema,
além de poder ser utilizada para automação.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.carcereiro)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.

		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		viper.SetConfigType("toml")
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".mon-liberabanco" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".carcereiro")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		//fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}

type aplicarDados struct {
	bancoTabela string
	usuario     string
	tipo        string
}

// aplicarAlteracao conecta no banco e executa o que tiver que executar, mas valida se usuário já existe.
func (aD aplicarDados) aplicarAlteracao() {

	var checkUsuario string

	// Multiplas tabelas
	tabelas := strings.Split(aD.bancoTabela, ",")

	for _, i := range tabelas {

		bancoETabela := strings.Split(i, ".")
		// temporario fixo
		db, err := sql.Open("mysql", viper.GetString("usuario")+":"+viper.GetString("senha")+"@tcp("+viper.GetString("host")+":"+viper.GetString("port")+")/"+bancoETabela[0])
		if err != nil {
			panic(err.Error())
		}
		defer db.Close()

		// Valida se usuário existe
		_ = db.QueryRow("SELECT user FROM mysql.user WHERE user = '" + aD.usuario + "';").Scan(&checkUsuario)

		if checkUsuario == "" {
			fmt.Println("Usuário " + aD.usuario + " não existe!")
		} else {
			_, err := db.Query(montaQuery(aD.tipo, bancoETabela[0], bancoETabela[1], aD.usuario))
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println(i + " OK!")
			}
		}

	}
}

// montaQuery cria a query para entregar para execução.
func montaQuery(tipo string, banco string, tabela string, usuario string) string {
	var saida string
	switch tipo {
	case "grantSelect":
		saida = "GRANT SELECT ON TABLE `" + banco + "`.`" + tabela + "` TO '" + usuario + "'@'%'"
	}
	return saida
}
