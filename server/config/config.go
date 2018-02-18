package config

import (
	"encoding/json"
	"flag"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
)

// Variables used for command line parameters
var (
	Config configFile
	Flags  configFlags
)

type configFile struct {
	Token        string `json:"token"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectURI  string `json:"redirect_uri"`

	GuildID        string `json:"guild_id"`
	UploadPassword string `json:"upload_password"`

	BotPrefix  string `json:"bot_prefix"` //prefix to use for bot commands
	SoundsPath string `json:"sounds_path"`
	ClipsPath  string `json:"clips_path"`
	ServerAddr string `json:"server_addr"`

	Pubg struct {
		Enabled bool     `json:"enabled"`
		APIKey  string   `json:"api_key"`
		Players []string `json:"players"`
	} `json:"pubg"`
}

type configFlags struct {
	Prod bool
}

// Init -
func Init() {
	parseConfig()
	parseFlags()
}

func parseConfig() {

	log.Debug("Reading config file...")

	file, e := ioutil.ReadFile("./config.json")

	if e != nil {
		log.Fatal("File error: %v\n", e)
	}

	log.Debug("%s\n", string(file))

	err := json.Unmarshal(file, &Config)

	if err != nil {
		log.Error(err)
	}
}

func parseFlags() {

	Flags.Prod = false

	prod := flag.Bool("p", false, "Run in production")

	flag.Parse()

	Flags.Prod = *prod

	if Flags.Prod {
		log.Warn("Running in production mode")
	}

}
