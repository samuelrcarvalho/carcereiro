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
package cmd

import (
	//"fmt"
	"fmt"
	"log"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configuração inicial de conexão dos bancos de dados.",
	Long: `Antes de iniciar a execução do Carcereiro, é necessário a
configuração do CLI para que as conexões sejam feitas automaticamente,
sem a necessidade ao executar todo comando digitar as credenciais.`,
	Run: func(cmd *cobra.Command, args []string) {
		configurar := Falaai()
		configurar.Configure()
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)
}

type dadosBanco struct {
	nome    string
	host    string
	port    string
	usuario string
	senha   string
}

// falaai solicita dados para configurado banco e retorna os dados no tipo dadosBanco
func falaai() dadosBanco {
	dados := dadosBanco{}
	fmt.Print("Nome do seu banco: ")
	fmt.Scanln(&dados.nome)
	fmt.Print("Host: ")
	fmt.Scanln(&dados.host)
	fmt.Print("Port: ")
	fmt.Scanln(&dados.port)
	fmt.Print("Usuario: ")
	fmt.Scanln(&dados.usuario)
	fmt.Print("Senha: ")
	fmt.Scanln(&dados.senha)
	return dados
}

// configure armazena as informações do usuário em um arquivo de texto.
func (d dadosBanco) configure() {
	texto := []byte("nome = " + d.nome + "\nhost = " + d.host + "\nport = " + d.port + "\nusuario = " + d.usuario + "\nsenha = " + d.senha)
	home, _ := homedir.Dir()
	err := os.WriteFile(home+"/.carcereiro", texto, 0600)
	if err != nil {
		log.Fatal(err)
	}
}
