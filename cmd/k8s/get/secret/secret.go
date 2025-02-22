package secret

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	k8s_get_cmd "github.com/sikalabs/slu/cmd/k8s/get"
	rootcmd "github.com/sikalabs/slu/cmd/root"
	"github.com/sikalabs/slu/utils/k8s"

	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var CmdFlagName string
var CmdFlagNamespace string
var CmdFlagKey string

var Cmd = &cobra.Command{
	Use:     "secret",
	Short:   "Get data from Secret",
	Aliases: []string{"sec"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		clientset, defaultNamespace, _ := k8s.KubernetesClient()

		namespace := defaultNamespace
		if CmdFlagNamespace != "" {
			namespace = CmdFlagNamespace
		}

		secretClient := clientset.CoreV1().Secrets(namespace)

		secret, err := secretClient.Get(context.TODO(), CmdFlagName, metav1.GetOptions{})
		if err != nil {
			log.Fatal(err)
		}

		if CmdFlagKey != "" {
			if rootcmd.RootCmdFlagJson {
				outJson, err := json.Marshal(string(secret.Data[CmdFlagKey]))
				if err != nil {
					panic(err)
				}
				fmt.Println(string(outJson))
			} else {
				fmt.Println(string(secret.Data[CmdFlagKey]))
			}
		} else {
			if rootcmd.RootCmdFlagJson {
				out := make(map[string]string)
				for key, val := range secret.Data {
					out[key] = string(val)
				}
				outJson, err := json.Marshal(out)
				if err != nil {
					panic(err)
				}
				fmt.Println(string(outJson))
			} else {
				for key, val := range secret.Data {
					fmt.Printf("KEY:   %s\nVALUE: %s\n---\n", key, val)
				}
			}
		}
	},
}

func init() {
	k8s_get_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&CmdFlagName,
		"secret-name",
		"s",
		"default",
		"Secret Name",
	)
	Cmd.MarkFlagRequired("secret-name")
	Cmd.Flags().StringVarP(
		&CmdFlagNamespace,
		"namespace",
		"n",
		"",
		"Kubernetes Namespace",
	)
	Cmd.Flags().StringVarP(
		&CmdFlagKey,
		"key",
		"k",
		"",
		"Get only specific key from data",
	)
}
