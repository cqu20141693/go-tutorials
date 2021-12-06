package cmd

import (
	"crypto"
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(tryCmd)
}

var tryCmd = &cobra.Command{
	Use:   "md5",
	Short: "Try and possibly fail at something",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := Md5Func(); err != nil {
			return err
		}
		return nil
	},
}

func Md5Func() error {
	md5 := crypto.MD5.New()
	md5.Write([]byte("test"))
	fmt.Printf("%x \n", md5.Sum(nil))
	return nil
}
