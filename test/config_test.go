package test

import (
	"testing"

	"github.com/Muntaha369/Go-CRUD-Mongo/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	cfg, err := config.LoadConfig()

	assert.NoError(t, err)
	assert.NotNil(t, cfg)

	assert.Equal(t, 8080, cfg.Server.Port)
	assert.Equal(t, "mongodb://localhost:27017/?directConnection=true", cfg.Database.URI)
}
