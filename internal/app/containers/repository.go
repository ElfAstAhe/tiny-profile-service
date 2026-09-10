package containers

import (
	"context"
	"errors"

	"github.com/ElfAstAhe/go-service-template/pkg/container"
	"github.com/ElfAstAhe/go-service-template/pkg/errs"
	"github.com/ElfAstAhe/go-service-template/pkg/logger"
)

const (
	InstancePersonRepo        string = "person-repository"
	InstancePersonMetricsRepo string = "person-metrics-repository"
	InstancePersonTraceRepo   string = "person-trace-repository"
	InstancePersonAuditRepo   string = "person-audit-repository"

	InstanceProfileRepo        string = "profile-repository"
	InstanceProfileMetricsRepo string = "profile-metrics-repository"
	InstanceProfileTraceRepo   string = "profile-trace-repository"
	InstanceProfileAuditRepo   string = "profile-audit-repository"
)

type RepositoryContainer struct {
	*container.BaseLazyContainer
}

var _ container.Container = (*RepositoryContainer)(nil)
var _ container.LazyContainer = (*RepositoryContainer)(nil)

func NewRepositoryContainer(
	orchestrator container.Orchestrator,
	log logger.Logger,
) *RepositoryContainer {
	return &RepositoryContainer{
		BaseLazyContainer: container.NewBaseLazyContainer(
			container.WithLazyName(RepositoryContainerName),
			container.WithLazyOrchestrator(orchestrator),
			container.WithLazyLogger(log),
		),
	}
}

func (rc *RepositoryContainer) Init(ctx context.Context) error {
	err := errors.Join(
		rc.RegisterProvider(InstanceProfileRepo, rc.providerProfileRepo),
		rc.RegisterProvider(InstanceProfileMetricsRepo, rc.providerProfileMetricsRepo),
		rc.RegisterProvider(InstanceProfileTraceRepo, rc.providerProfileTraceRepo),
		rc.RegisterProvider(InstanceProfileAuditRepo, rc.providerProfileAuditRepo),

		rc.RegisterProvider(InstancePersonRepo, rc.providerPersonRepo),
		rc.RegisterProvider(InstancePersonMetricsRepo, rc.providerPersonMetricsRepo),
		rc.RegisterProvider(InstancePersonTraceRepo, rc.providerPersonTraceRepo),
		rc.RegisterProvider(InstancePersonAuditRepo, rc.providerPersonAuditRepo),
	)
	if err != nil {
		return errs.NewContainerError(rc.GetName(), "container init: register providers failed", err)
	}

	return nil
}
