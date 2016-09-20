package routes

import (
	"errors"
	"github.com/lastbackend/cli/cmd/context"
	"github.com/lastbackend/cli/libs/io/table"
	"github.com/lastbackend/cli/utils"
	"github.com/urfave/cli"
)

func Deploy(c *cli.Context) error {

	var ctx = context.Get()

	var sBrunch, sHub, sRepo, sType, sOwner string

	token, err := context.Get().Storage.Token.Get()
	if err != nil {
		ctx.Debug.Error(err)
		return err
	}

	switch c.Command.FullName() {
	case "deploy git":
		var repo string = c.String(`repo`)
		if repo == "" {
			return errors.New("Bad VCS url")
		}
		gitData, err := utils.GitUrlParse(repo)
		if err != nil {
			ctx.Debug.Info("VCS url: %S, error: %#v", repo, err)
			return errors.New("Bad VCS url")
		}
		ctx.Debug.Info("GitDataParse: %#v", gitData)
		sHub = gitData.Hub
		sRepo = gitData.Repo
		sType = gitData.Vendor
		sOwner = gitData.Owner
		sBrunch = c.String(`brunch`)
	case "deploy docker":
		var image string = c.String(`image`)
		if image == "" {
			return errors.New("Bad VCS url")
		}
		dockerData, err := utils.DockerNamespaceParse(image)
		if err != nil {
			ctx.Debug.Info("VCS url: %S, error: %#v", image, err)
			return errors.New("Bad Docker namespace")
		}
		sHub = dockerData.Hub
		sRepo = dockerData.Repo
		sType = dockerData.Vendor
		sOwner = dockerData.Owner
		sType = dockerData.Vendor
		sBrunch = dockerData.Brunch
	case "deploy jumpstart":
	}

	name := c.String(`name`)
	if !utils.IsServiceName(name) {
		ctx.Debug.Info("Service name: %s", name)
		return errors.New("Bad service name")
	}
	reg := c.String(`region`)
	if reg == "" {
		//TODO check region
		ctx.Debug.Info("Service region: %s", reg)
		return errors.New("Bad service region")
	}
	memory := c.Uint(`memory`)
	if !utils.IsMemory(memory) {
		ctx.Debug.Info("Service memory: %s", memory)
		return errors.New("Bad service memory")
	}

	h, service, err := ctx.API.Service.Create(token, name, memory, "", "", reg, "", sType, sHub, sOwner, sRepo, sBrunch, false)
	if err != nil {
		ctx.Debug.Error(err)
		ctx.Debug.Info("\n REQUEST : %#v \n\n RESPONSE : %#v \n", h.Request, h.Response)
		return err
	}

	ctx.Debug.Info("Services : %#v", service)

	var header []string = []string{"UUID", "Name", "Status", "Provision", "Type", "Hub", "Owner", "Repository", "Branch"}
	var data [][]string

	d := []string{
		service.UUID, service.Name, service.Status, utils.BToS(service.Provision), service.Source.Type,
		service.Source.Hub, service.Source.Owner, service.Source.Repository, service.Source.Branch}
	data = append(data, d)
	d = d[:0]

	table.PrintTable(header, data, []string{})

	return nil

}

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
		ctx.Debug.Info("\n REQUEST : %#v \n\n RESPONSE : %#v \n", h.Request, h.Response)
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
