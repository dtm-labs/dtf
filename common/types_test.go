package common

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

type testConfig struct {
	DB map[string]string `yaml:"DB"`
}

var config = testConfig{}

func init() {
	InitConfig(GetProjectDir(), &config)
	config.DB["database"] = ""
}

func TestDb(t *testing.T) {
	db := DbGet(config.DB)
	err := func() (rerr error) {
		defer P2E(&rerr)
		dbr := db.NoMust().Exec("select a")
		assert.NotEqual(t, nil, dbr.Error)
		db.Must().Exec("select a")
		return nil
	}()
	assert.NotEqual(t, nil, err)
}

func TestDbAlone(t *testing.T) {
	db := SdbAlone(config.DB)
	_, err := SdbExec(db, "select 1")
	assert.Equal(t, nil, err)
	db.Close()
	_, err = SdbExec(db, "select 1")
	assert.NotEqual(t, nil, err)
}
