package setting

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/astaxie/beego/config"
)

var (
	//@Global Config

	//AppName should be "dockyard"
	AppName string
	//Usage is short description
	Usage   string
	Version string
	Author  string
	Email   string

	//@Basic Runtime Config

	RunMode        string
	ListenMode     string
	HTTPSCertFile  string
	HTTPSKeyFile   string
	LogPath        string
	LogLevel       string
	DatabaseDriver string
	DatabaseURI    string
	Domains        string
)

func init() {
	conf, err := getConfig()
	if err == nil {
		err = setGlobalConfig(conf)
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getConfig() (conf config.Configer, err error) {
	home := os.Getenv("HOME")
	if home != "" {
		homePath := filepath.Join(home, ".dockyard", "containerops.conf")
		conf, err = config.NewConfig("ini", homePath)
	}

	if err != nil {
		conf, err = config.NewConfig("ini", "conf/containerops.conf")
	}

	return
}

func setGlobalConfig(conf config.Configer) error {
	//config globals
	if appname := conf.String("appname"); appname != "" {
		AppName = appname
	} else if appname == "" {
		return fmt.Errorf("AppName config value is null")
	}

	if usage := conf.String("usage"); usage != "" {
		Usage = usage
	} else if usage == "" {
		return fmt.Errorf("Usage config value is null")
	}

	if version := conf.String("version"); version != "" {
		Version = version
	} else if version == "" {
		return fmt.Errorf("Version config value is null")
	}

	if author := conf.String("author"); author != "" {
		Author = author
	} else if author == "" {
		return fmt.Errorf("Author config value is null")
	}

	if email := conf.String("email"); email != "" {
		Email = email
	} else if email == "" {
		return fmt.Errorf("Email config value is null")
	}

	return nil
}

func setServerConfig(conf config.Configer) error {
	//config runtime
	if runmode := conf.String("runmode"); runmode != "" {
		RunMode = runmode
	} else if runmode == "" {
		return fmt.Errorf("RunMode config value is null")
	}

	if listenmode := conf.String("listenmode"); listenmode != "" {
		ListenMode = listenmode
	} else if listenmode == "" {
		return fmt.Errorf("ListenMode config value is null")
	}

	if httpscertfile := conf.String("httpscertfile"); httpscertfile != "" {
		HTTPSCertFile = httpscertfile
	} else if httpscertfile == "" {
		return fmt.Errorf("HttpsCertFile config value is null")
	}

	if httpskeyfile := conf.String("httpskeyfile"); httpskeyfile != "" {
		HTTPSKeyFile = httpskeyfile
	} else if httpskeyfile == "" {
		return fmt.Errorf("HttpsKeyFile config value is null")
	}

	if logpath := conf.String("log::filepath"); logpath != "" {
		LogPath = logpath
	} else if logpath == "" {
		return fmt.Errorf("LogPath config value is null")
	}

	if loglevel := conf.String("log::level"); loglevel != "" {
		LogLevel = loglevel
	} else if loglevel == "" {
		return fmt.Errorf("LogLevel config value is null")
	}

	if databasedriver := conf.String("database::driver"); databasedriver != "" {
		DatabaseDriver = databasedriver
	} else if databasedriver == "" {
		return fmt.Errorf("Database Driver config value is null")
	}

	if databaseuri := conf.String("database::uri"); databaseuri != "" {
		DatabaseURI = databaseuri
	} else if databaseuri == "" {
		return fmt.Errorf("Database URI config vaule is null")
	}

	// Deployment domain could be empty
	if domains := conf.String("deployment::domains"); domains != "" {
		Domains = domains
	} else if domains == "" {
		return fmt.Errorf("Deployment domains config vaule is null")
	}

	return nil
}
