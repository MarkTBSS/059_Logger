package main

import (
	"os"

	_pkgConfig "github.com/MarkTBSS/059_Logger/config"
	_pkgModulesServers "github.com/MarkTBSS/059_Logger/modules/servers"
)

func envPath() string {
	if len(os.Args) == 1 {
		return ".env"
	} else {
		return os.Args[1]
	}
}

func main() {
	cfg := _pkgConfig.LoadConfig(envPath())
	_pkgModulesServers.NewServer(cfg).Start()
}
