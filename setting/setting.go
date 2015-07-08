package setting

import (
	"fmt"
	"os"

	"github.com/astaxie/beego/config"
)

var (
	conf          config.ConfigContainer
	AppName       string
	Usage         string
	Version       string
	Author        string
	Email         string
	RunMode       string
	ListenMode    string
	HttpsCertFile string
	HttpsKeyFile  string
	LogPath       string
	DBURI         string
	DBPasswd      string
	DBDB          int64
)

func init() {
	if err := SetSettings("conf/containerops.conf"); err != nil {
		fmt.Println("Read configuration file error: %s", err.Error())
		os.Exit(1)
	}
}

func SetSettings(path string) error {
	var err error

	conf, err = config.NewConfig("ini", path)
	if err != nil {
		fmt.Errorf("Read conf/crew.conf error: %v", err.Error())
	}

	if appname := conf.String("appname"); appname != "" {
		AppName = appname
	} else if appname == "" {
		err = fmt.Errorf("AppName config value is null")
	}

	if usage := conf.String("usage"); usage != "" {
		Usage = usage
	} else if usage == "" {
		err = fmt.Errorf("Usage config value is null")
	}

	if version := conf.String("version"); version != "" {
		Version = version
	} else if version == "" {
		err = fmt.Errorf("Version config value is null")
	}

	if author := conf.String("author"); author != "" {
		Author = author
	} else if author == "" {
		err = fmt.Errorf("Author config value is null")
	}

	if email := conf.String("email"); email != "" {
		Email = email
	} else if email == "" {
		err = fmt.Errorf("Email config value is null")
	}

	if runmode := conf.String("runmode"); runmode != "" {
		RunMode = runmode
	} else if runmode == "" {
		err = fmt.Errorf("RunMode config value is null")
	}

	if listenmode := conf.String("listenmode"); listenmode != "" {
		ListenMode = listenmode
	} else if listenmode == "" {
		err = fmt.Errorf("ListenMode config value is null")
	}

	if httpscertfile := conf.String("httpscertfile"); httpscertfile != "" {
		HttpsCertFile = httpscertfile
	} else if httpscertfile == "" {
		err = fmt.Errorf("HttpsCertFile config value is null")
	}

	if httpskeyfile := conf.String("httpskeyfile"); httpskeyfile != "" {
		HttpsKeyFile = httpskeyfile
	} else if httpskeyfile == "" {
		err = fmt.Errorf("HttpsKeyFile config value is null")
	}

	if logpath := conf.String("log::filepath"); logpath != "" {
		LogPath = logpath
	} else if logpath == "" {
		err = fmt.Errorf("LogPath config value is null")
	}

	if dburi := conf.String("db::uri"); dburi != "" {
		DBURI = dburi
	} else if dburi == "" {
		err = fmt.Errorf("DBURI config value is null")
	}

	if dbpass := conf.String("db::passwd"); dbpass != "" {
		DBPasswd = dbpass
	}

	DBDB, err = conf.Int64("db::db")

	return err
}
