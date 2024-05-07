package main

import (
	"github.com/TrueBlocks/trueblocks-core/sdk"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/logger"
	"github.com/TrueBlocks/trueblocks-core/src/apps/chifra/pkg/types"
)

// DoMonitors tests the monitors sdk function
func DoMonitors() {
	opts := sdk.MonitorsOptions{
		Addrs: []string{testAddrs[0]},
	}
	ShowHeader("DoMonitors", &opts)

	if monitors, _, err := opts.MonitorsList(); err != nil {
		logger.Error(err)
	} else {
		if err := SaveAndClean[types.Monitor]("sdkFuzzer/monitors.json", monitors, &opts, func() error {
			_, _, err := opts.MonitorsList()
			return err
		}); err != nil {
			logger.Error(err)
		}
	}

	// opts.List = false
	// if monitors, _, err := opts.Monitors(); err != nil {
	// 	logger.Error(err)
	// } else {
	// 	if err := SaveAndClean[bool]("sdkFuzzer/monitors.json", monitors, &opts, func() error {
	// 		_, _, err := opts.Monitors()
	// 		return err
	// 	}); err != nil {
	// 		logger.Error(err)
	// 	}
	// }

	// if monitors, _, err := opts.MonitorsClean(); err != nil {
	// 	logger.Error(err)
	// } else {
	// 	if err := SaveAndClean[types.MonitorClean]("sdkFuzzer/monitors.json", monitors, &opts, func() error {
	// 		_, _, err := opts.MonitorsClean()
	// 		return err
	// 	}); err != nil {
	// 		logger.Error(err)
	// 	}
	// }
}