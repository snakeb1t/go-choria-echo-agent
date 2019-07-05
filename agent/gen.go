// +build ignore

package main

import (
"encoding/base64"
"fmt"
"io/ioutil"
"os"
"time"

"text/template"
)

const ftempl = `package provision
// auto generated {{.Now}}
import (
	"encoding/base64"
	"encoding/json"
	addl "github.com/choria-io/mcorpc-agent-provider/mcorpc/ddl/agent"
)
var echoddl = ` + "`{{.EchoDDL}}`" + `
// DDL is the agent DDL
var DDL = make(map[string]*addl.DDL)
func init() {
	DDL["echo"] = &addl.DDL{}
	ddl, _ := base64.StdEncoding.DecodeString(echoddl)
	json.Unmarshal(ddl, DDL["echo"])
}
`

type dat struct {
EchoDDL    string
}

func (d dat) Now() string {
return fmt.Sprintf("%s", time.Now())
}

func main() {
echoj, err := ioutil.ReadFile("agent/echo.json")
if err != nil {
panic(fmt.Sprintf("Could not read agents spec file agent/echo.json: %s", err))
}

templ := template.Must(template.New("ddl").Parse(ftempl))

f, err := os.Create("agent/ddl.go")
if err != nil {
panic(fmt.Sprintf("cannot create file agent/ddl.go: %s", err))
}
defer f.Close()

input := dat{
EchoDDL:    base64.StdEncoding.EncodeToString(echoj),
}

err = templ.Execute(f, input)
if err != nil {
panic(fmt.Sprintf("executing template:", err))
}
}


