package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewLogoutCommand() *cobra.Command {
	login := &cobra.Command{
		Use:   "logout",
		Short: "Logout from redissyncer",
		Run:   logoutCommandFunc,
	}
	return login
}

func logoutCommandFunc(cmd *cobra.Command, args []string) {
	cmd.Println("exec logout")
	viper.Set("token", "")
	viper.WriteConfig()
	cmd.Println("Logout successful")
}
