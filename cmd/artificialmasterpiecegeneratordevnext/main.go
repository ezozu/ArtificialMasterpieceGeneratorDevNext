// cmd/artificialmasterpiecegeneratordevnext/main.go
package main

import (
"flag"
"log"
"os"

"artificialmasterpiecegeneratordevnext/internal/artificialmasterpiecegeneratordevnext"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := artificialmasterpiecegeneratordevnext.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
