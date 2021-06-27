package cron

import (
	"testing"

	"github.com/axiaoxin-com/x-stock/services"
	"github.com/stretchr/testify/require"
)

func TestSyncFundAllList(t *testing.T) {
	SyncFundAllList()
}

func TestUpdate4433(t *testing.T) {
	services.FundAllListFilename = "../fund_all_list.json"
	err := services.InitFundAllList()
	require.Nil(t, err)
	Update4433()
}