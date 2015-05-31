# cli
Command line interface framework
## Example
```
import (
	"fmt"
	"github.com/arvinkulagin/cli"
)

func main() {
	cli.Add("server {addr} {path}", func(r cli.Request) {
		fmt.Println(r.Raw())
		fmt.Printf("Arrd: %s\n", r.Vars()["addr"])
		fmt.Printf("Path: %s\n", r.Vars()["path"])
	})

	cli.Add("{url}", func(r cli.Request) {
		fmt.Println(r.Raw())
		fmt.Printf("URL: %s\n", r.Vars()["url"])
	})

	cli.Run()
}
```
## API
```
import "github.com/arvinkulagin/cli"
```
```
func Add(template string, handler Handler)

func Run() error
```
```
type App struct {
    // contains filtered or unexported fields
}

func NewApp() App

func (s *App) Add(template string, handler Handler)

func (s *App) Run() error
```
```
type Handler func(Request)
```
```
type Pattern struct {
    // contains filtered or unexported fields
}

func NewPattern(template string) Pattern

func (p Pattern) RE() string

func (p Pattern) Template() string
```
```
type Request struct {
    // contains filtered or unexported fields
}

func NewRequest(raw string, pattern Pattern) (Request, error)

func (r Request) Raw() string

func (r Request) Vars() map[string]string
```