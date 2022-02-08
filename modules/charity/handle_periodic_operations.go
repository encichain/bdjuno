package charity

import (
	"github.com/forbole/bdjuno/v2/modules/utils"

	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"
)

// RegisterPeriodicOperations implements modules.PeriodicOperationsModule
func (m *Module) RegisterPeriodicOperations(scheduler *gocron.Scheduler) error {
	log.Debug().Str("module", "charity").Msg("setting up periodic tasks")

	// Setup a cron job to update params periodically to ensure no sync errors
	// Run once every other day at midnight
	if _, err := scheduler.Every(2).Day().At("0:00").Do(func() {
		utils.WatchMethod(m.PeriodicUpdateParams)
	}); err != nil {
		return err
	}

	return nil
}
