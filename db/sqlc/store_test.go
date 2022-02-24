package db

import (
	"database/sql"
	"testing"

	"github.com/harissucipto/xendit-task/util"
)

func TestNewStore (t *testing.T) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		t.Errorf("cannot load config: %v", err)
	}

	db, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		t.Errorf("cannot connect to db: %v", err)
	}

	store := NewStore(db)
	t.Logf("store: %v", store)
}