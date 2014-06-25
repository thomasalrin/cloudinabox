package app

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/megamsys/cloudinabox/action"
	"github.com/megamsys/cloudinabox/exec"
//	"bitbucket.org/kardianos/osext"
	"log"
	"strings"
)

const (
	keyremote_repo = "remote_repo="
	keylocal_repo  = "local_repo="
	keyproject     = "project="
	ganetipreinstall = "bash conf/ganeti/mganeti_preinstall.sh"
	ganetiverify = "bash conf/ganeti/mganeti_verify.sh"
	ganetipostinstall = "bash conf/ganeti/mganeti_postinstall.sh"
	ganetiinstall = "bash conf/ganeti/mganeti_install.sh"
	opennebulapreinstall = "bash conf/opennebula/one_preinstall.sh"
	opennebulaverify = "bash conf/opennebula/one_verify.sh"
	opennebulapostinstall = "bash conf/opennebula/one_postinstall.sh"
	opennebulainstall = "bash conf/opennebula/one_install.sh"
	rootPath  = "/tmp"
	defaultEnvPath = "conf/env.sh"	
	drbd_mnt = "/drbd_mnt"
)

var ErrAppAlreadyExists = errors.New("there is already an app with this name.")

func CommandExecutor(app *App) (action.Result, error) {
	var e exec.OsExecutor
	var b bytes.Buffer
	var commandWords []string
	commandWords = strings.Fields(app.Command)
	
	fmt.Println(commandWords, len(commandWords))

	if len(commandWords) > 0 {
		if err := e.Execute(commandWords[0], commandWords[1:], nil, &b, &b); err != nil {
			log.Printf("stderr:%s", b)		
			return nil, err
		}
	}
   
	log.Printf("%s", b)
	return &app, nil
}

//Remove the installed packages..
var remove = action.Action{
	Name: "remove",
	Forward: func(ctx action.FWContext) (action.Result, error) {
		var app App
		switch ctx.Params[0].(type) {
		case App:
			app = ctx.Params[0].(App)
		case *App:
			app = *ctx.Params[0].(*App)
		default:
			return nil, errors.New("First parameter must be App or *App.")
		}
		
		return CommandExecutor(&app)
	},
	Backward: func(ctx action.BWContext) {
		app := ctx.FWResult.(*App)	
		log.Printf("[%s] Nothing to recover for %s", app.ClusterName)
	},
	MinParams: 1,
}

var ganetiVerify = action.Action{
	Name: "ganetiVerify",
	Forward: func(ctx action.FWContext) (action.Result, error) {
		var app App
		switch ctx.Params[0].(type) {
		case App:
			app = ctx.Params[0].(App)
		case *App:
			app = *ctx.Params[0].(*App)
		default:
			return nil, errors.New("First parameter must be App or *App.")
		}
		app.Command = ganetiverify
	   return CommandExecutor(&app)
	},
	Backward: func(ctx action.BWContext) {
	app := ctx.FWResult.(*App)
		log.Printf("[%s] Nothing to recover for %s", app.ClusterName)
	},
	MinParams: 1,
	}


var ganetiInstall = action.Action{
	Name: "ganetiInstall",
	Forward: func(ctx action.FWContext) (action.Result, error) {
		var app App
		switch ctx.Params[0].(type) {
		case App:
			app = ctx.Params[0].(App)
		case *App:
			app = *ctx.Params[0].(*App)
		default:
			return nil, errors.New("First parameter must be App or *App.")
		}
		
		log.Printf("Installation %s", app.ClusterName)
		app.Command = ganetiinstall
	   return CommandExecutor(&app)
	},
	Backward: func(ctx action.BWContext) {
		app := ctx.FWResult.(*App)		
		log.Printf("[%s] Nothing to recover for %s", app.ClusterName)		
	},
	MinParams: 1,
}

var ganetiPreInstall = action.Action{
	Name: "ganetiPreInstall",
	Forward: func(ctx action.FWContext) (action.Result, error) {
		var app App
		switch ctx.Params[0].(type) {
		case App:
			app = ctx.Params[0].(App)
		case *App:
			app = *ctx.Params[0].(*App)
		default:
			return nil, errors.New("First parameter must be App or *App.")
		}
		app.Command = ganetipreinstall
	   return CommandExecutor(&app)
	},
	Backward: func(ctx action.BWContext) {
		app := ctx.FWResult.(*App)		
		log.Printf("[%s] Nothing to recover for %s", app.ClusterName)
	},
	MinParams: 1,
	}

  var ganetiPostInstall = action.Action{
	Name: "ganetiPostInstall",
	Forward: func(ctx action.FWContext) (action.Result, error) {
		var app App
		switch ctx.Params[0].(type) {
		case App:
			app = ctx.Params[0].(App)
		case *App:
			app = *ctx.Params[0].(*App)
		default:
			return nil, errors.New("First parameter must be App or *App.")
		}
		app.Command = ganetipostinstall
	   return CommandExecutor(&app)
	},
	Backward: func(ctx action.BWContext) {
	app := ctx.FWResult.(*App)
		log.Printf("[%s] Nothing to recover for %s", app.ClusterName)
	},
	MinParams: 1,
	}
	
	var opennebulaVerify = action.Action{
	Name: "opennebulaVerify",
	Forward: func(ctx action.FWContext) (action.Result, error) {
		var app App
		switch ctx.Params[0].(type) {
		case App:
			app = ctx.Params[0].(App)
		case *App:
			app = *ctx.Params[0].(*App)
		default:
			return nil, errors.New("First parameter must be App or *App.")
		}
		//filename, _ := osext.Executable()
		app.Command = opennebulaverify
	   return CommandExecutor(&app)
	},
	Backward: func(ctx action.BWContext) {
	app := ctx.FWResult.(*App)
		log.Printf("[%s] Nothing to recover for %s", app.ClusterName)
	},
	MinParams: 1,
	}

var opennebulaInstall = action.Action{
	Name: "opennebulaInstall",
	Forward: func(ctx action.FWContext) (action.Result, error) {
		var app App
		switch ctx.Params[0].(type) {
		case App:
			app = ctx.Params[0].(App)
		case *App:
			app = *ctx.Params[0].(*App)
		default:
			return nil, errors.New("First parameter must be App or *App.")
		}
		
		log.Printf("Installation %s", app.ClusterName)
		app.Command = opennebulainstall
	   return CommandExecutor(&app)
	},
	Backward: func(ctx action.BWContext) {
		app := ctx.FWResult.(*App)		
		log.Printf("[%s] Nothing to recover for %s", app.ClusterName)		
	},
	MinParams: 1,
}

var opennebulaPreInstall = action.Action{
	Name: "opennebulaPreInstall",
	Forward: func(ctx action.FWContext) (action.Result, error) {
		var app App
		switch ctx.Params[0].(type) {
		case App:
			app = ctx.Params[0].(App)
		case *App:
			app = *ctx.Params[0].(*App)
		default:
			return nil, errors.New("First parameter must be App or *App.")
		}
		app.Command = opennebulapreinstall
	   return CommandExecutor(&app)
	},
	Backward: func(ctx action.BWContext) {
		app := ctx.FWResult.(*App)		
		log.Printf("[%s] Nothing to recover for %s", app.ClusterName)
	},
	MinParams: 1,
	}

  var opennebulaPostInstall = action.Action{
	Name: "opennebulaPostInstall",
	Forward: func(ctx action.FWContext) (action.Result, error) {
		var app App
		switch ctx.Params[0].(type) {
		case App:
			app = ctx.Params[0].(App)
		case *App:
			app = *ctx.Params[0].(*App)
		default:
			return nil, errors.New("First parameter must be App or *App.")
		}
		app.Command = opennebulapostinstall
	   return CommandExecutor(&app)
	},
	Backward: func(ctx action.BWContext) {
	app := ctx.FWResult.(*App)
		log.Printf("[%s] Nothing to recover for %s", app.ClusterName)
	},
	MinParams: 1,
	}



