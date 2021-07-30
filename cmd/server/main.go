package main

import "fmt"

// App - the struct contains stuffs like pointers
// to db connections
type App struct{}

func (app *App) Run() error {
  fmt.Println("Setting up Our App")
  return nil
}
func main() {
  fmt.Println("Project kicking off")
  app := App{}
  if err := app.Run(); err != nil {
    fmt.Println("Error starting up our Rest API")
    fmt.Println(err)
  }
}
