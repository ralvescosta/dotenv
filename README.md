# DotEnv Package

This GoLang package was create to handle only .env files. 

The idea of ​​this package is to get the variables from the file and register them in the variable runtime environments.

### Installation

```bash
go get -u github.com/ralvescosta/dotenv
```

### Example

```golang
package main

import (
  "fmt"
  "os"
  
  "github.com/ralvescosta/dotenv"
)

func main() {
  dotenv.Configure(".env")

  fmt.Println(os.Getenv("SOME_ENV"))
}
```