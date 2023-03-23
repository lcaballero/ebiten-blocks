package main // "github.com/lcaballero/ebiten/gen-scripts/cli"

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"gopkg.in/yaml.v3"
)

func FlagType(t string) string {
	switch t {
	case "int":
		return "IntFlag"
	case "bool":
		return "BoolFlag"
	case "int64":
		return "Int64Flag"
	default:
		return "unknown"
	}
}

func ContextMethod(t string) string {
	switch t {
	case "int":
		return "Int"
	case "bool":
		return "Bool"
	case "int64":
		return "Int64"
	default:
		return "unknown"
	}
}

func TrimLeft(format string, args ...any) string {
	s := fmt.Sprintf(format, args...)
	s = strings.TrimLeft(s, "\n")
	return s
}

func Cap(s string) string {
	parts := strings.Split(s, "-")
	all := ""
	for _, p := range parts {
		cap := strings.ToUpper(p[0:1])
		all = all + cap + p[1:]
	}
	return all
}

func DefaultValue(ty string, t interface{}) string {
	s, isString := t.(string)
	if isString && strings.TrimSpace(s) == "" {
		return ""
	}
	switch ty {
	case "int", "int64":
		return fmt.Sprintf(`
						Value: %d,`, t)
	case "string":
		return fmt.Sprintf(`
						Value: %s,`, t)
	default:
		return ""
	}
}

type GenCLI struct {
	Package     string       `yaml:"package"`
	Name        string       `yaml:"name"`
	Usage       string       `yaml:"usage"`
	SubCommands []SubCommand `yaml:"sub-commands"`
}

type SubCommand struct {
	Name  string `yaml:"name"`
	Usage string `yaml:"usage"`
	Flags []Flag `yaml:"flags"`
}

type Flag struct {
	Name  string      `yaml:"name"`
	Type  string      `yaml:"type"`
	Usage string      `yaml:"usage"`
	Value interface{} `yaml:"value"`
}

func LoadNewCLI() *GenCLI {
	yamlFile, err := ioutil.ReadFile("cli.yaml")
	if err != nil {
		fmt.Errorf("yamlFile.Get err   #%v ", err)
	}
	c := &GenCLI{}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}

func main() {
	err := NewWriter(LoadNewCLI()).Generate().WriteFile("gen.cli.go")
	if err != nil {
		panic(err)
	}
}

type Writer struct {
	buf *bytes.Buffer
	cli *GenCLI
}

func NewWriter(cli *GenCLI) *Writer {
	return &Writer{
		buf: bytes.NewBufferString(""),
		cli: cli,
	}
}

func (w *Writer) Bytes() []byte {
	return w.buf.Bytes()
}

func (w *Writer) Write(format string, args ...interface{}) {
	fmt.Fprintf(w.buf, format, args...)
}

func (w *Writer) WriteFile(path string) error {
	err := ioutil.WriteFile(path, w.Bytes(), 0644)
	return err
}

func (w *Writer) Generate() *Writer {
	w.pkg()
	w.imports()
	w.vals()
	w.newApp()
	w.procs(w.cli.SubCommands)
	w.opts(w.cli.SubCommands)
	return w
}

func (w *Writer) pkg() {
	w.Write("package %s", w.cli.Package)
}

func (w *Writer) imports() {
	w.Write(`

import (
    "github.com/urfave/cli"
)
`)
}

func (w *Writer) procs(procs []SubCommand) {
	w.Write(`
type Procs struct {`)
	for _, p := range procs {
		w.Write(`
    %s ValProc`, Cap(p.Name))
	}
	w.Write(`
}
`)
}

func (w *Writer) newApp() {
	w.Write(`
func NewApp(procs Procs) *cli.App {
    app := &cli.App{
        Name:  "%s",
        Usage: "%s",
        Commands: []cli.Command{
			%s
        },
    }
    return app
}
`,
		w.cli.Name,
		w.cli.Usage,
		w.subs(w.cli.SubCommands),
	)
}

func (w *Writer) subs(subs []SubCommand) string {
	buf := `cli.Command{
`
	for _, cmd := range subs {
		s := TrimLeft(`
				Name: "%s",
				Usage: "%s",
				Action: procs.%s.ToProc(),
`,
			cmd.Name,
			cmd.Usage,
			Cap(cmd.Name),
		)
		buf += s
		buf += TrimLeft(`
				Flags: []cli.Flag{
`)
		for _, f := range cmd.Flags {
			buf += TrimLeft(`
					&cli.%s{
						Name: "%s",
						Usage: "%s",%s
					},
`,
				FlagType(f.Type),
				f.Name,
				f.Usage,
				DefaultValue(f.Type, f.Value),
			)
		}
		buf += `				},
`
	}
	s := `
			},`
	s = strings.TrimLeft(s, "\n")
	buf += s
	return buf
}

func (w *Writer) opts(subs []SubCommand) {
	for _, sub := range subs {
		w.Write(`
type %sOpts struct {
	vals Vals
`,
			Cap(sub.Name),
		)
		w.Write(`}
`)
		for _, f := range sub.Flags {
			w.Write(`
func (opt *%sOpts) %s() %s {
	return opt.vals.%s("%s")
}
`,
				Cap(sub.Name),
				Cap(f.Name),
				f.Type,
				ContextMethod(f.Type),
				f.Name,
			)
		}
		for _, f := range sub.Flags {
			w.Write(`
func (opt *%sOpts) Has%s() bool {
	return opt.vals.IsSet("%s")
}
`,
				Cap(sub.Name),
				Cap(f.Name),
				f.Name,
			)
		}
	}
}

func (w *Writer) vals() {
	w.Write(`
type ValProc func(Vals) error

func (e ValProc) ToProc() func(*cli.Context) error {
	return func(c *cli.Context) error {
		return e(c)
	}
}

type Vals interface {
	Int(string) int
	Int64(string) int64
	Bool(string) bool
	IsSet(string) bool
}
`)
}
