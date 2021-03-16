# Bluefoxgo

Bluefoxgo is an api wrapper for bluefox api! Easy and beginner friendly!

> Bluefoxgo is currently under beta stage. Incase if you have found any kind of bugs kindly make an issue or make an pull request to improve us :)

## Quick example

```go
package main

import (
    "fmt"
    "github.com/Scientific-Guy/bluefoxgo"
)

func main(){
    client := bluefox.NewClient("TOKEN")
    server, err := client.GetServer("ID")

    if err != nil{
        fmt.Println("Failed fetching server details: " + err.Error())
        return
    }

    fmt.Println("Sucess!", server);
}
```

## Docs

You can view the full documentation on [`pkg.go.dev`](https://pkg.go.dev/github.com/Scientific-Guy/bluefoxgo)!
