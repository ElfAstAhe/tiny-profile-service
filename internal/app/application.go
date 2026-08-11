package app

import (
	"errors"

	"github.com/ElfAstAhe/go-service-template/pkg/app"
	libcontainer "github.com/ElfAstAhe/go-service-template/pkg/container"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/logger"
	"github.com/ElfAstAhe/tiny-profile-service/internal/app/containers"
	"github.com/ElfAstAhe/tiny-profile-service/internal/config"
)

type Application struct {
	*app.BaseApplication
	conf *config.Config
	log  logger.Logger
}

var _ app.Application = (*Application)(nil)

//goland:noinspection SpellCheckingInspection
func NewApplication(opts ...Option) (*Application, error) {
	// create instance
	res := &Application{}
	// setup
	for _, opt := range opts {
		opt(res)
	}
	// orchestrator
	orch := containers.NewOrchestrator(res.conf, res.log)
	libcontainer.SetDefaultOrchestrator(orch)
	// embed
	res.BaseApplication = app.NewBaseApplication(
		app.WithOrchestrator(orch),
		app.WithLogger(res.log),
		app.WithCloseTimeout(res.conf.App.CloseTimeout),
		app.WithStopTimeout(res.conf.App.StopTimeout),
	)
	// orchestrator and containers
	err := errors.Join(
		// app container
		res.GetOrchestrator().Register(containers.NewAppContainer(res.GetOrchestrator(), res.log)),
		/*
			// tools container
			res.GetOrchestrator().Register(containers.NewToolsContainer(res.GetOrchestrator(), res.log)),
			// client container
			res.GetOrchestrator().Register(containers.NewClientContainer(res.GetOrchestrator(), res.log)),
			// postgres container
			res.GetOrchestrator().Register(containers.NewPgContainer(res.GetOrchestrator(), res.log)),
			// repository container
			res.GetOrchestrator().Register(containers.NewRepositoryContainer(res.GetOrchestrator(), res.log)),
			// use case container
			res.GetOrchestrator().Register(containers.NewUseCaseContainer(res.GetOrchestrator(), res.log)),
			// facade container
			res.GetOrchestrator().Register(containers.NewFacadeContainer(res.GetOrchestrator(), res.log)),
			// services container (inner kitchen)
			res.GetOrchestrator().Register(containers.NewServiceContainer(res.GetOrchestrator(), res.log)),
			// worker container
			res.GetOrchestrator().Register(containers.NewWorkerContainer(res.GetOrchestrator(), res.log)),
			// http container
			res.GetOrchestrator().Register(containers.NewHTTPContainer(res.GetOrchestrator(), res.log)),
			// gRPC container
			res.GetOrchestrator().Register(containers.NewGRPCContainer(res.GetOrchestrator(), res.log)),
		*/
	)
	if err != nil {
		return nil, errs.NewCommonError("application create failed", err)
	}

	return res, nil
}

func (app *Application) Init() error {
	appCnt, err := app.GetOrchestrator().GetContainer(containers.AppContainerName)
	if err != nil {
		return errs.NewCommonError("app init", err)
	}
	err = errors.Join(
		appCnt.RegisterInstance(containers.InstanceApplication, app),
		appCnt.RegisterInstance(containers.InstanceApplicationReady, app.IsReady),
	)
	if err != nil {
		return errs.NewCommonError("app init", err)
	}

	return app.BaseApplication.Init()
}
