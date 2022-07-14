package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"io/ioutil"
	"os"
	"strings"
)

const (
	FlagName = "flag-name"
	ArgIndex = "arg-index"
)

func main() {
	var (
		flagName      string
		argumentIndex int
	)

	cmd := &cobra.Command{
		Use: "option",
		RunE: func(cmd *cobra.Command, arg []string) error {
			var value string

			data, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				return err
			}
			args := strings.Split(strings.TrimSpace(string(data)), " ")

			if cmd.Flags().Changed(FlagName) {
				fset := pflag.NewFlagSet("option", pflag.ContinueOnError)
				fset.StringVar(&value, flagName, "", "")

				if err := fset.Parse(args); err != nil {
					return err
				}
			}

			if cmd.Flags().Changed(ArgIndex) {
				if len(args) > argumentIndex {
					value = args[argumentIndex]
				}
			}

			fmt.Print(value)
			return nil
		},
	}

	cmd.Flags().StringVarP(&flagName, FlagName, "f", "", "the name of flag whose value you want to get")
	cmd.Flags().IntVarP(&argumentIndex, ArgIndex, "a", 0, "the index of argument whose value you want to get")
	cmd.MarkFlagsMutuallyExclusive(FlagName, ArgIndex)

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
