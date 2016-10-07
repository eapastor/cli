package cli

import (
	"github.com/lastbackend/cli/cmd/context"
	"github.com/lastbackend/cli/libs/errors"
	"github.com/lastbackend/cli/libs/io"
	"github.com/lastbackend/cli/libs/io/table"
	"github.com/lastbackend/cli/utils"
)

type serviceS struct{}

var SERVICE serviceS

func (serviceS) Deploy(name, region string, memory uint, template, git, docker string) {

	var ctx = context.Get()

	var sBrunch, sHub, sRepo, sType, sOwner string

	token, err := context.Get().Storage.Token.Get()
	if err != nil {
		ctx.Debug.Error(err)
		io.Error(errors.Err(err))
		return
	}

	if git != "" {
		gitData, err := utils.GitUrlParse(git)
		if err != nil {
			ctx.Debug.Info("VCS url: %S, error: %#v", git, err)
			io.Error(errors.Service.BadVCSURL())
			return
		}
		ctx.Debug.Info("GitDataParse: %#v", gitData)
		sHub = gitData.Hub
		sRepo = gitData.Repo
		sType = gitData.Vendor
		sOwner = gitData.Owner
		sBrunch = "master" // TODO add brunches var
	} else if docker != "" {
		dockerData, err := utils.DockerNamespaceParse(docker)
		if err != nil {
			ctx.Debug.Info("VCS url: %S, error: %#v", docker, err)
			io.Error(errors.Service.BadDockerNamespace())
			return
		}
		sHub = dockerData.Hub
		sRepo = dockerData.Repo
		sType = dockerData.Vendor
		sOwner = dockerData.Owner
		sType = dockerData.Vendor
		sBrunch = dockerData.Brunch
	} else if template != "" {
		sHub = "jumpstart"
		sType = "jumpstart"
		/*		{
					"opts":{
						"name":"fight",
						"memory":128
						},
					"node":{
						"uuid":"",
						"region":"sa"
					},
					"source":{
					"id":"f93d4c09-7a91-4837-8aa1-a55a218408b4",
					"type":"jumpstart",
					"hub":"jumpstart",
					"repo":null,
					"branch":null
					}
				}
			}*/
	}

	if !utils.IsServiceName(name) {
		ctx.Debug.Info("Service name: %s", name)
		io.Error(errors.Service.BadServiceUUID())
		return
	}
	if region == "" {
		//TODO check region
		ctx.Debug.Info("Service region: %s", region)
		io.Error(errors.Service.BadRegion())
		return
	}
	if !utils.IsMemory(memory) {
		ctx.Debug.Info("Service memory: %s", memory)
		io.Error(errors.Service.BadMemory())
		return
	}

	h, service, err := ctx.API.Service.Create(token, name, memory, "", "", region, "", sType, sHub, sOwner, sRepo, sBrunch, false)
	if err != nil {
		ctx.Debug.Error(err)
		ctx.Debug.Info("\n REQUEST : %#v \n\n RESPONSE : %#v \n", h.Request, h.Response)
		io.Error(errors.Err(err))
		return
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

}

func (serviceS) List() {

	var ctx = context.Get()
	//	var msg = config.Get().Msg

	token, err := context.Get().Storage.Token.Get()
	if err != nil {
		ctx.Debug.Error(err)
		io.Error(errors.Err(err))
		return
	}

	h, services, err := ctx.API.Service.List(token)
	if err != nil {
		ctx.Debug.Error(err)
		ctx.Debug.Info("\n REQUEST : %#v \n\n RESPONSE : %#v \n", h.Request, h.Response)
		io.Error(errors.Err(err))
		return
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

}

func (serviceS) Get(serviceName string) {

	var ctx = context.Get()

	if !utils.IsServiceName(serviceName) {
		io.Error(errors.Service.BadServiceUUID())
		return
	}

	token, err := ctx.Storage.Token.Get()
	if err != nil {
		ctx.Debug.Error(err)
		io.Error(errors.Err(err))
		return
	}

	h, service, err := ctx.API.Service.Get(token, serviceName)
	if err != nil {
		ctx.Debug.Error(err)
		ctx.Debug.Info("\n REQUEST : %#v \n\n RESPONSE : %#v \n", h.Request, h.Response)
		io.Error(errors.Err(err))
		return
	}

	ctx.Debug.Info("Services : %#v", service)

	if service.UUID == "" {
		io.Printf("Service %s not found", serviceName)
		return
	}

	var header []string = []string{"UUID", "Name", "Status", "Provision", "Type", "Hub", "Owner", "Repository", "Branch"}
	var data [][]string

	d := []string{
		service.UUID, service.Name, service.Status, utils.BToS(service.Provision), service.Source.Type,
		service.Source.Hub, service.Source.Owner, service.Source.Repository, service.Source.Branch}
	data = append(data, d)

	table.PrintTable(header, data, []string{})

}

func (serviceS) Start(serviceName string) {

	var ctx = context.Get()

	if !utils.IsServiceName(serviceName) {
		io.Error(errors.Service.BadServiceUUID())
		return
	}

	token, err := ctx.Storage.Token.Get()
	if err != nil {
		ctx.Debug.Error(err)
		io.Error(errors.Err(err))
		return
	}

	h, err := ctx.API.Service.Start(token, serviceName)
	if err != nil {
		ctx.Debug.Error(err)
		ctx.Debug.Info("\n REQUEST : %#v \n\n RESPONSE : %#v \n", h.Request, h.Response)
		io.Error(errors.Err(err))
		return
	}

}

func (serviceS) Restart(serviceName string) {

	var ctx = context.Get()

	if !utils.IsServiceName(serviceName) {
		io.Error(errors.Service.BadServiceUUID())
		return
	}

	token, err := ctx.Storage.Token.Get()
	if err != nil {
		ctx.Debug.Error(err)
		io.Error(errors.Err(err))
		return
	}

	h, err := ctx.API.Service.Restart(token, serviceName)
	if err != nil {
		ctx.Debug.Error(err)
		ctx.Debug.Info("\n REQUEST : %#v \n\n RESPONSE : %#v \n", h.Request, h.Response)
		io.Error(errors.Err(err))
		return
	}

}

func (serviceS) Stop(serviceName string) {

	var ctx = context.Get()

	if !utils.IsServiceName(serviceName) {
		io.Error(errors.Service.BadServiceUUID())
		return
	}

	token, err := ctx.Storage.Token.Get()
	if err != nil {
		ctx.Debug.Error(err)
		io.Error(errors.Err(err))
		return
	}

	h, err := ctx.API.Service.Stop(token, serviceName)
	if err != nil {
		ctx.Debug.Error(err)
		ctx.Debug.Info("\n REQUEST : %#v \n\n RESPONSE : %#v \n", h.Request, h.Response)
		io.Error(errors.Err(err))
		return
	}

}

func (serviceS) Remove(serviceName string) {

	var ctx = context.Get()

	if !utils.IsServiceName(serviceName) {
		io.Error(errors.Service.BadServiceUUID())
		return
	}

	token, err := ctx.Storage.Token.Get()
	if err != nil {
		ctx.Debug.Error(err)
		io.Error(errors.Err(err))
		return
	}

	h, err := ctx.API.Service.Remove(token, serviceName)
	if err != nil {
		ctx.Debug.Error(err)
		ctx.Debug.Info("\n REQUEST : %#v \n\n RESPONSE : %#v \n", h.Request, h.Response)
		io.Error(errors.Err(err))
		return
	}

}
