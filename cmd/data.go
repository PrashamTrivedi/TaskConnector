package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
)

var commandFile string
var commandMapping map[string]string

func init() {
	if len(commandMapping) == 0 {
		readCommandMapping()
	}
}
func SetcommandFile(commandFilePath string) {
	commandFile = commandFilePath
	if len(commandMapping) == 0 {
		readCommandMapping()
	}
}

func readCommandMapping() {
	data, _ := readConfigFile()

	if len(data) > 0 {
		json.Unmarshal(data, &commandMapping)
	} else {
		commandMapping = make(map[string]string)
	}
}

func readConfigFile() ([]byte, error) {
	configFile := getConfigFilePath()
	fileBytes, err := ioutil.ReadFile(configFile)
	return fileBytes, err

}

func getConfigFilePath() string {
	configFile := commandFile
	return configFile
}

func writeCommandMappingFile(bytes []byte) error {
	configFilePath := getConfigFilePath()
	configFile, errorData := os.OpenFile(configFilePath, os.O_RDWR|os.O_CREATE, 0600)
	if errorData != nil {
		fmt.Println("Error in processing file:", errorData.Error())
		os.Exit(1)
	}

	err := ioutil.WriteFile(configFile.Name(), bytes, 0600)
	return err
}

func writeCommandMapping() {

	commands, err := json.Marshal(commandMapping)
	if err != nil {
		fmt.Println("Error in unmarshalling json:", err.Error())
		os.Exit(1)
	}
	err = writeCommandMappingFile(commands)
	if err != nil {
		fmt.Println("Error in writing to file:", err.Error())
		os.Exit(1)
	}

}

func GetCommandMapping() map[string]string {
	if len(commandMapping) == 0 {
		readCommandMapping()
	}
	return commandMapping
}

func AddConfig(key, value string) {

	commandMapping[key] = value
	writeCommandMapping()
}

func getCommandRunners() (string, string) {
	mainCommand := "bash"
	commandFlag := "-c"
	if runtime.GOOS == "windows" {
		mainCommand = "cmd"
		commandFlag = "/c"
	}
	return mainCommand, commandFlag
}

func RunCommand(url string) string {

	if len(commandMapping) == 0 {
		readCommandMapping()
	}

	command := commandMapping[url]
	if command == "" {
		return fmt.Sprintf("Command not found for url %s", url)
	}
	mainCommand, commandFlag := getCommandRunners()
	output, errorData := exec.Command(mainCommand, commandFlag, command).CombinedOutput()
	if errorData != nil {
		return fmt.Sprintf("Error in running command %s, error: %s", command, errorData.Error())
	}
	return string(output)

}
