package cron

import (
	"errors"
	"fmt"
	"github.com/KubeOperator/KubeOperator/pkg/cron/job"
	"github.com/robfig/cron/v3"
)

var Cron *cron.Cron

const phaseName = "cron"

type InitCronPhase struct {
	Enable bool
}

func (c *InitCronPhase) Init() error {
	if c.Enable {
		Cron := cron.New()
		_, err := Cron.AddJob("@every 5m", job.NewRefreshHostInfo())
		if err != nil {
			return errors.New(fmt.Sprintf("can not add corn job: %s", err.Error()))
		}
		Cron.Start()
	}
	return nil
}

func (c *InitCronPhase) PhaseName() string {
	return phaseName
}
