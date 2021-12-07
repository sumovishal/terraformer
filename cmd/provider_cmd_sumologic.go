// Copyright 2019 The Terraformer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package cmd

import (
	"errors"
	"fmt"
	sumologic_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/sumologic"
	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
	"os"
)

func newCmdSumologicImporter(options ImportOptions) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sumologic",
		Short: "Import current state to Terraform configuration from Sumo Logic",
		Long:  "Import current state to Terraform configuration from Sumo Logic",
		RunE: func(cmd *cobra.Command, args []string) error {
			accessKey := os.Getenv("SUMOLOGIC_ACCESSKEY")
			fmt.Println(accessKey)
			if len(accessKey) == 0 {
				return errors.New("Must set SUMOLOGIC_ACCESSKEY env var")
			}
			accessID := os.Getenv("SUMOLOGIC_ACCESSID")
			fmt.Println(accessID)
			if len(accessID) == 0 {
				return errors.New("Must set SUMOLOGIC_ACCESSID env var")
			}
			provider := newSumologicProvider()
			fmt.Println("provider=", provider)
			err := Import(provider, options, []string{accessKey, accessID})
			if err != nil {
				return err
			}
			return nil
		},
	}

	cmd.AddCommand(listCmd(newSumologicProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "users", "users=id1:id2:id4")
	fmt.Println("in newCmdSumologicImporter()")
	return cmd
}

func newSumologicProvider() terraformutils.ProviderGenerator {
	return &sumologic_terraforming.SumologicProvider{}
}
