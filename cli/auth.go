package cli

import (
	"github.com/howeyc/gopass"
	"github.com/lastbackend/cli/cmd/config"
	"github.com/lastbackend/cli/cmd/context"
	"github.com/lastbackend/cli/libs/io"
	"github.com/lastbackend/cli/libs/errors"
)

func Auth() {

	ctx := context.Get()

	token, err := ctx.Storage.Token.Get()
	if err != nil {
		ctx.Debug.Error(err)
		io.Error(err.Error())
		return
	}

	ctx.Debug.Info("::token %s", token)

	if token != "" {
		return
	}

	Login()

	ctx.Debug.Info("::token %s", token)

}

func Login() {

	var msg = config.Get().Msg
	var ctx = context.Get()

	io.Println(msg.Login.StartMessage)

	login := io.GetString(msg.Login.EnterEmail)
	io.Print(msg.Login.EnterPass)
	bytePass, err := gopass.GetPasswd()
	if err != nil {
		ctx.Debug.Error(err)
		io.Error(err.Error())
		return
	}

	password := string(bytePass)

	h, token, err := ctx.API.Session.Get(login, password)
	if err != nil {
		ctx.Debug.Error(err)
		ctx.Debug.Info("\n REQUEST : %+v \n\n RESPONSE : %+v \n", h.Request, h.Response)
		io.Error(errors.Err(err))
		return
	}

	if token != "" {
		err = ctx.Storage.Token.Set(token)
		if err != nil {
			ctx.Debug.Error(err)
			io.Error(errors.Err(err))
			return
		}
	}

}

func Logout() {

	ctx := context.Get()

	err := ctx.Storage.Token.Delete()
	if err != nil {
		ctx.Debug.Error(err)
		io.Error(err.Error())
		return
	}
	io.Println(config.Get().Msg.Logout.CredentialsClear)

}

func Whoami() {

	ctx := context.Get()
	cfg := config.Get()

	token, err := context.Get().Storage.Token.Get()
	if err != nil {
		ctx.Debug.Error(err)
		io.Error(err.Error())
		return
	}

	if token == "" {
		io.Println(cfg.Msg.NotLoggined)
	}

	h, user, err := ctx.API.User.Get(token)
	if err != nil {
		ctx.Debug.Error(err)
		ctx.Debug.Info("\n REQUEST : %+v \n\n RESPONSE : %+v \n", h.Request.Header, h.Response)
		io.Error(errors.Err(err))
		return
	}

	ctx.Debug.Info("::user %#v", user)

	io.Printf("%s %s ", cfg.Msg.WhoAmI.LogginedBy, user.Username)

}