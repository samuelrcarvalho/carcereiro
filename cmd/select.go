/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
)

// selectCmd represents the select command
var selectCmd = &cobra.Command{
	Use:   "select",
	Short: "Liberar os comandos select no banco de dados para determinado usuário.",
	Long: `Necessário informar database.table usuario.

carcereiro liberar select database.table usuario`,
	Run: func(cmd *cobra.Command, args []string) {
		retorno := validar(args)
		if retorno == 1 {
			fmt.Println("Argumentos inválido, --help para mais informações.")
		} else {
			dadosSelect := aplicarDados{
				args[0],
				args[1],
				"grantSelect",
			}
			mensagem := dadosSelect.aplicarAlteracao()
			fmt.Println(mensagem)
		}
	},
}

func init() {
	liberarCmd.AddCommand(selectCmd)
}

// validar mais de 3 parâmetros, 0 parâmetros e falta de ponto no database.table.
func validar(a []string) int {
	i := 0
	if len(a) == 0 {
		return 1
	}
	for _, _ = range a {
		if i > 1 {
			return 1
		}
		i++
	}
	if !strings.Contains(a[0], ".") {
		return 1
	}
	return 0
}
