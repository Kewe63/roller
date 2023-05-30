package init

import (
	"os"
	"os/exec"
	"path/filepath"
)

func initializeLightNodeConfig() error {
	initLightNodeCmd := exec.Command(celestiaExecutablePath, "light", "init", "--p2p.network", "arabica", "--node.store", filepath.Join(os.Getenv("HOME"), configDirName.DALightNode))
	err := initLightNodeCmd.Run()
	if err != nil {
		return err
	}
	return nil
}
