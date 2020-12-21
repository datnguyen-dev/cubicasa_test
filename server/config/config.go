package config

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/caarlos0/env"
	"github.com/hashicorp/hcl"
)

// Config is a configuration struct.
type Config struct {
	Postgres struct {
		Host      string `hcl:"host"`
		Port      int    `hcl:"port"`
		UserName  string `hcl:"user_name"`
		Password  string `hcl:"password"`
		DBName    string `hcl:"db_name"`
		SSL       string `hcl:"ssl"`
		SetupFile string `hcl:"setup_file"`
	} `hcl:"postgres"`
}

//Config for Docker Enviroment
type envConfig struct {
	Host      string `env:"HOST"`
	Port      int    `env:"PORT"`
	UserName  string `env:"USER_NAME"`
	Password  string `env:"PASS_WORD"`
	DBName    string `env:"DB_NAME"`
	SSL       string `env:"SSL"`
	SetupFile string `env:"SETUP_FILE"`
}

func readFile(filename string) (*Config, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %v", err)
	}
	defer f.Close()

	cfg, err := read(f)
	if err == nil {
		readEnvConfig(cfg)
	}
	return cfg, err
}

func read(r io.Reader) (*Config, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read config: %v", err)
	}

	cfg := &Config{}
	err = hcl.Unmarshal(data, cfg)
	if err != nil {
		return nil, fmt.Errorf("failed unmarshal hcl: %v", err)
	}

	return cfg, nil
}

//InitConfig - init config file
func InitConfig(configFile string) (*Config, error) {
	if _, err := os.Stat(configFile); !os.IsNotExist(err) {
		res, err := readFile(configFile)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
	cfg, err := initfile()
	if err != nil {
		return nil, err
	}

	err = ioutil.WriteFile(configFile, []byte(cfg), 0666)
	if err != nil {
		return nil, err
	}

	res, err := readFile(configFile)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func initfile() (string, error) {
	buf := new(bytes.Buffer)
	err := tpl.Execute(buf, map[string]interface{}{
		"host":       "localhost",
		"port":       5432,
		"user_name":  "postgres",
		"password":   "1234567890",
		"db_name":    "cubicasa_db",
		"ssl":        "disable",
		"setup_file": "./setup/script.sql",
	})
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func readEnvConfig(cfg *Config) {
	envCfg := envConfig{}
	err := env.Parse(&envCfg)
	if err != nil {
		return
	}

	// Postgres
	{
		if envCfg.Host != "" {
			cfg.Postgres.Host = envCfg.Host
		}
		if envCfg.Port > 0 {
			cfg.Postgres.Port = envCfg.Port
		}
		if envCfg.UserName != "" {
			cfg.Postgres.UserName = envCfg.UserName
		}
		if envCfg.Password != "" {
			cfg.Postgres.Password = envCfg.Password
		}
		if envCfg.DBName != "" {
			cfg.Postgres.DBName = envCfg.DBName
		}
		if envCfg.SSL != "" {
			cfg.Postgres.SSL = envCfg.SSL
		}
		if envCfg.SetupFile != "" {
			cfg.Postgres.SetupFile = envCfg.SetupFile
		}
	}
}

var tpl = template.Must(template.New("initial-config").Parse(strings.TrimSpace(`
postgres {
  host    = "{{.host}}"
	port = "{{.port}}"
	user_name = "{{.user_name}}"
	password = "{{.password}}"
	db_name = "{{.db_name}}"
	ssl = "{{.ssl}}"
	setup_file = "{{.setup_file}}"
}
`)))
