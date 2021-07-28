# Simple Go DotEnv Package

### Get Started

```bash
go get -u github.com/ralvescosta/dotenv
```

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