package length

import (
	"encoding/json"
	"fmt"

	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "length <string>",
	Short:   "Length of a string",
	Aliases: []string{"len"},
	Args:    cobra.ExactArgs(1),
	Run: func(c *cobra.Command, args []string) {
		l := len(args[0])
		if root.RootCmdFlagJson {
			outJson, err := json.Marshal(l)
			if err != nil {
				panic(err)
			}
			fmt.Println(string(outJson))
		} else {
			fmt.Println(l)
		}
	},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
