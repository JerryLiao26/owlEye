package helper

import (
	"encoding/json"
	"os"
)

var confDir = GetHomePath() + "/.owl/"
var confName = "config.json"
var confPath = confDir + confName
var TextServerPath = "127.0.0.1"
var TextServerPort = "12580"
var ImageServerPath = "127.0.0.1"
var ImageServerPort = "12580"

type confStruct struct {
	TextServerPath  string `json:"text_server_path"`
	TextServerPort  string `json:"text_server_port"`
	ImageServerPath string `json:"image_server_path"`
	ImageServerPort string `json:"image_server_port"`
}

func LoadConf() {
	if checkFileExist(confPath) {
		var conf confStruct
		confStream := ReadLocalFile(confPath)

		// Extract JSON
		err := json.Unmarshal(confStream, &conf)

		if err != nil {
			panic(err)
		} else {
			TextServerPath = conf.TextServerPath
			TextServerPort = conf.TextServerPort
			ImageServerPath = conf.ImageServerPath
			ImageServerPort = conf.ImageServerPort
		}
	} else {
		initConf()
	}
}

func SaveConf() {
	conf := confStruct{
		TextServerPath:  TextServerPath,
		TextServerPort:  TextServerPort,
		ImageServerPath: ImageServerPath,
		ImageServerPort: ImageServerPort,
	}

	// Build JSON
	content, _ := json.Marshal(conf)

	// Write
	WriteLocalFile(confPath, []byte(content))
}

func initConf() {
	_ = os.Mkdir(confDir, 0755)

	SaveConf()
}
