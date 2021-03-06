// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/prantoran/relayserver/api"
	"github.com/prantoran/relayserver/config"
	"github.com/spf13/cobra"
)

var (
	ebconf config.App
)

// listenCmd represents the listen command
var listenCmd = &cobra.Command{
	Use:   "listen",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("listen called")
		listen()
	},
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("en persistentprerune")
		conf, err := loadConfig()
		if err != nil {
			return err
		}
		ebconf = *conf
		return nil
	},
}

func loadConfig() (*config.App, error) {
	conf := config.New()

	if err := conf.Init(); err != nil {
		return conf, err
	}

	if err := conf.ReadConsule(); err != nil {
		return conf, errors.Wrap(err, "can not read consule")
	}

	conf.Load()
	conf.Prt()
	return conf, nil
}

func listen() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/demo", api.DemoClient(&api.DemoConf{
		Msg: "demo works",
	}))
	http.Handle("/", r)
	fmt.Println("listening on port", ebconf.Self.Port)
	log.Fatal(http.ListenAndServe(":"+ebconf.Self.Port, r))

}

func init() {
	rootCmd.AddCommand(listenCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listenCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listenCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
