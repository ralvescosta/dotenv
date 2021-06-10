# Simple Go DotEnv Package

### How we can use

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