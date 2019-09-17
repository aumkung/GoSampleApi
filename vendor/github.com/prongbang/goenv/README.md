# Load environment variables from .env

## Installation:

```
go get github.com/prongbang/goenv
```

## Usage:

### `.env` file
```
DB_USER=root
DB_PASS=1234
```

### `main.go`
```golang
import (
	"fmt"
	"os"
	_ "github.com/prongbang/goenv"
)

func main() {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")

	fmt.Println("user:", user)
	fmt.Println("pass:", pass)
}
```