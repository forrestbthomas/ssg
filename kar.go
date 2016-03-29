// +build kar

package main

import (
  "github.com/omeid/gonzo/context"
  "github.com/omeid/kar"
  "github.com/omeid/kargar"
  "github.com/go-gonzo/fs"
  "github.com/go-gonzo/util"
  "github.com/go-gonzo/css"
)


// You must put your task definations inside an init.
// Read more about the order of init functions and Package Initialization
// at https://golang.org/ref/spec#Package_initialization
func init() {

  kar.Run(func(build *kargar.Build) error {

    return build.Add(
      kargar.Task{
        Name:  "minify-css",
        Usage: "Concat frontend CSS to app.min.css",
        Action: func(ctx context.Context) error {
          return fs.Src(ctx, "./assets/stylesheets/*.css").Then(
            util.Concat(ctx, "app.css"),
            css.Minify(),
            fs.Dest("./dist/"),
          )
        },
      })
  })
}
