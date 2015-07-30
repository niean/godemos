package g

import (
	"encoding/json"
	"github.com/toolkits/file"
	"log"
	"sync"
)

type HttpConfig struct {
	Enable bool   `json:"enable"`
	Listen string `json:"listen"`
}

type RrdConfig struct {
	Enable      bool   `json:"enable"`
	CallTimeout int64  `json:"callTimeout"`
	Protocol    string `json:"protocol"`
	Framed      bool   `json:"framed"`
	Buffered    bool   `json:"buffered"`
	Listen      string `json:"listen"`
}

type GlobalConfig struct {
	Debug bool        `json:"debug"`
	Http  *HttpConfig `json:"http"`
	Rrd   *RrdConfig  `json:"rrd"`
}

var (
	ConfigFile string
	config     *GlobalConfig
	configLock = new(sync.RWMutex)
)

func Config() *GlobalConfig {
	configLock.RLock()
	defer configLock.RUnlock()
	return config
}

func ParseConfig(cfg string) {
	if cfg == "" {
		log.Fatalln("use -c to specify one config file")
	}

	if !file.IsExist(cfg) {
		log.Fatalln("config file:", cfg, "not exist")
	}

	ConfigFile = cfg

	configContent, err := file.ToTrimString(cfg)
	if err != nil {
		log.Fatalln("read config file:", cfg, "failed:", err)
	}

	var c GlobalConfig
	err = json.Unmarshal([]byte(configContent), &c)
	if err != nil {
		log.Fatalln("parse config file:", cfg, "failed:", err)
	}

	// check
	if !checkConfig(c) {
		log.Fatalln("check config file:", cfg, "failed")
	}

	configLock.Lock()
	defer configLock.Unlock()
	config = &c

	log.Println("g.ParseConfig ok, file ", cfg)
}

// check
func checkConfig(c GlobalConfig) bool {
	return true
}
