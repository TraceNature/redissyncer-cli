package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
	"interactioncli/httpquerry"
)

func NewLoginCommand() *cobra.Command {
	login := &cobra.Command{
		Use:   "login <username password>",
		Short: "login redissyncer server",
		Run:   loginCommandFunc,
	}
	return login
}

func loginCommandFunc(cmd *cobra.Command, args []string) {
	if len(args) != 2 {
		cmd.PrintErr(cmd.Usage())
		return
	}

	resp, err := httpquerry.Login(viper.GetString("syncserver"), args[0], args[1])

	if err != nil {
		cmd.PrintErr(err)
		return
	}

	code := gjson.Get(resp, "code").String()

	if code == "2000" {
		viper.Set("token", gjson.Get(resp, "data.token").String())
		viper.WriteConfig()
	}

	cmd.Println(resp)
}
