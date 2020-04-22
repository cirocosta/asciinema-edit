package main

import (
	"io"
	"os"
	"strings"

	"github.com/Masterminds/sprig"
	"github.com/cirocosta/asciinema-edit/commands"
	"gopkg.in/urfave/cli.v1"
)

var (
	version = "dev"
	commit  = "HEAD"
)

var CmdHelpOld = `COMMANDS:{{range .VisibleCategories}}{{if .Name}}
   {{.Name}}:{{end}}{{range .VisibleCommands}}
     {{join .Names ", "}}{{"\t"}}{{.Usage}}{{end}}{{end}}`

var CmdHelpNew = `COMMANDS:{{range .VisibleCommands}}

{{.Name}}

  {{.Usage | indent 2}}{{end}}
`

func main() {
	cli.HelpPrinter = func(w io.Writer, templ string, data interface{}) {
		cli.HelpPrinterCustom(w, templ, data, sprig.FuncMap())
	}
	cli.AppHelpTemplate = strings.Replace(cli.AppHelpTemplate, CmdHelpOld, CmdHelpNew, 1)

	app := cli.NewApp()

	app.Version = version + " - " + commit
	app.Usage = "edit recorded asciinema casts"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Ciro da Silva da Costa",
			Email: "ciro.costa@liferay.com",
		},
	}
	app.Description = `asciinema-edit provides missing features from the "asciinema" tool
   when it comes to editing a cast that has already been recorded.`
	app.Commands = []cli.Command{
		commands.Cut,
		commands.Quantize,
		commands.Speed,
	}

	app.Run(os.Args)
}
