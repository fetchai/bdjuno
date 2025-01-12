package modules

import (
	"github.com/desmos-labs/juno/v2/modules"
	"github.com/desmos-labs/juno/v2/types/config"

	"github.com/forbole/bdjuno/v2/database"
)

var (
	_ modules.Module                     = &Module{}
	_ modules.AdditionalOperationsModule = &Module{}
)

type Module struct {
	cfg config.ChainConfig
	db  *database.Db
}

// NewModule returns a new Module instance
func NewModule(cfg config.ChainConfig, db *database.Db) *Module {
	return &Module{
		cfg: cfg,
		db:  db,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "modules"
}
