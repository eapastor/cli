package routes

import (
	"github.com/howeyc/gopass"
	"github.com/lastbackend/cli/cmd/config"
	"github.com/lastbackend/cli/cmd/context"
	"github.com/lastbackend/cli/libs/io"
	"github.com/urfave/cli"
)

func Login(c *cli.Context) (err error) {

	var msg = config.Get().Msg
	var ctx = context.Get()

	io.Println(msg.Login.StartMessage)

	login := io.GetString(msg.Login.EnterEmail)
	io.Print(msg.Login.EnterPass)
	bytePass, err := gopass.GetPasswd()
	if err != nil {
		ctx.Debug.Error(err)
		return err
	}

	password := string(bytePass)

	h, token, err := ctx.API.Session.Get(login, password)
	if err != nil {
		ctx.Debug.Error(err)
		ctx.Debug.Info("\n REQUEST : %+v \n\n RESPONSE : %+v \n", h.Request, h.Response)
		return err
	}

	if token != "" {
		err = ctx.Storage.Token.Set(token)
		if err != nil {
			ctx.Debug.Error(err)
			return err
		}
	}

	return nil

}

func Logout(c *cli.Context) (err error) {

	ctx := context.Get()

	err = ctx.Storage.Token.Delete()
	if err != nil {
		ctx.Debug.Error(err)
		return err
	}
	io.Println(config.Get().Msg.Logout.CredentialsClear)

	return nil

}

func Whoami(c *cli.Context) (err error) {

	ctx := context.Get()

	token, err := context.Get().Storage.Token.Get()
	if err != nil {
		ctx.Debug.Error(err)
		return err
	}

	h, user, err := ctx.API.User.Get(token)
	if err != nil {
		ctx.Debug.Error(err)
		ctx.Debug.Info("\n REQUEST : %+v \n\n RESPONSE : %+v \n", h.Request.Header, h.Response)
		return err
	}

	ctx.Debug.Info("::user %#v", user)

	io.Printf("%s %s ", config.Get().Msg.WhoAmI.LogginedBy, user.Username)

	return nil

}