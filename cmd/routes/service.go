package routes

import (
	"github.com/lastbackend/cli/cmd/context"
	"github.com/lastbackend/cli/libs/io/table"
	"github.com/lastbackend/cli/utils"
	"github.com/urfave/cli"
)

func Services(c *cli.Context) error {

	var ctx = context.Get()
	//	var msg = config.Get().Msg

	token, err := context.Get().Storage.Token.Get()
	if err != nil {
		ctx.Debug.Error(err)
		return err
	}

	h, services, err := ctx.API.Service.List(token)
	if err != nil {
		ctx.Debug.Error(err)
		ctx.Debug.Info("\n REQUEST : %+v \n\n RESPONSE : %+v \n", h.Request.Header, h.Response)
		return err
	}

	ctx.Debug.Info("Services : %#v", services)

	var header []string = []string{"UUID", "Name", "Status", "Provision", "Type", "Hub", "Owner", "Repository", "Branch"}
	var data [][]string

	for _, service := range *services {
		d := []string{
			service.UUID, service.Name, service.Status, utils.BToS(service.Provision), service.Source.Type,
			service.Source.Hub, service.Source.Owner, service.Source.Repository, service.Source.Branch}
		data = append(data, d)
		d = d[:0]
	}

	table.PrintTable(header, data, []string{})

	return nil

}
