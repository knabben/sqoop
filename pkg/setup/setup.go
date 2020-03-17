package setup

import (
	"time"

	"github.com/solo-io/gloo/pkg/utils/setuputils"

	check "github.com/solo-io/go-checkpoint"

	"github.com/solo-io/sqoop/pkg/syncer"
	"github.com/solo-io/sqoop/version"
)

func Main() error {
	start := time.Now()
	check.CallCheck("sqoop", version.Version, start)
	return setuputils.Main(setuputils.SetupOpts{
		SetupFunc:         syncer.Setup,
		ExitOnError:       true,
		LoggingPrefixVals: []interface{}{"sqoop"},
	})
}
