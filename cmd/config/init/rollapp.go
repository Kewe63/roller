package init

import (
	"os"
	"os/exec"
	"path/filepath"

	toml "github.com/pelletier/go-toml"
)

func initializeRollappConfig(rollappExecutablePath string, chainId string, denom string) {
	initRollappCmd := exec.Command(rollappExecutablePath, "init", keyNames.HubSequencer, "--chain-id", chainId, "--home", filepath.Join(getRollerRootDir(), configDirName.Rollapp))
	err := initRollappCmd.Run()
	if err != nil {
		panic(err)
	}
	setRollappAppConfig(filepath.Join(getRollerRootDir(), configDirName.Rollapp, "config/app.toml"), denom)
}

func setRollappAppConfig(appConfigFilePath string, denom string) {
	config, _ := toml.LoadFile(appConfigFilePath)
	config.Set("minimum-gas-prices", "0"+denom)
	config.Set("api.enable", "true")
	file, _ := os.Create(appConfigFilePath)
	_, err := file.WriteString(config.String())
	if err != nil {
		panic(err)
	}
	file.Close()
}
