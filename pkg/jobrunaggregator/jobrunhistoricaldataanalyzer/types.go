package jobrunhistoricaldataanalyzer

import (
	_ "embed"
	"time"

	"github.com/openshift/ci-tools/pkg/jobrunaggregator/jobrunaggregatorapi"
)

//go:embed pr_message_template.md
var prTemplate string

type parsedJobData struct {
	PercentTimeDiff                    float64       `json:"-"`
	TimeDiff                           time.Duration `json:"-"`
	PrevP99                            time.Duration `json:"-"`
	DurationP95                        time.Duration `json:"-"`
	DurationP99                        time.Duration `json:"-"`
	JobResults                         int           `json:"-"`
	jobrunaggregatorapi.HistoricalData `json:",inline"`
}

type compareResults struct {
	newReleaseEvent bool
	increaseCount   int
	decreaseCount   int
	addedJobs       []string
	jobs            []parsedJobData
	missingJobs     []parsedJobData
}
