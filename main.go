package main

import (
    "os"
    "fmt"
    "log"
    "encoding/json"
    "io/ioutil"

    "github.com/urfave/cli"

    types "github.com/silmin/odc-combiner/typefile"
    "github.com/silmin/odc-combiner/combine"
)

func isExistFile(filename string) bool {
    _, err := os.Stat(filename)
    return err == nil
}

func file2Figure(filename string) (types.Figure, error) {
    bytes, err := ioutil.ReadFile(filename)
    if err != nil {
        log.Fatal(err)
        return types.Figure{}, err
    }

    var figure types.Figure
    if err := json.Unmarshal(bytes, &figure); err != nil {
        log.Fatal(err)
        return types.Figure{}, err
    }

    return figure, nil
}

func isExistAreas(figure types.Figure, names []string) bool {
    for _, name := range names {
        flg := false
        for _, area := range figure.Areas {
            if name == area.Name {
                flg = true
                break
            }
        }
        if !flg {
            return false
        }
    }
    return true
}

func main() {
    app := &cli.App {

        Name: "ODC-combiner",
        Usage: "combine json of opendatacam",
        Version: "0.0.1",

        Commands: []*cli.Command {
            {
                Name: "areas",
                Usage: "combine any areas",

                Flags: []cli.Flag {
                    &cli.StringFlag {
                        Name: "src",
                        Aliases: []string{"s"},
                        Usage: "input json file",
                        Value: "input.json",
                    },
                    &cli.StringFlag {
                        Name: "dst",
                        Aliases: []string{"d"},
                        Usage: "output json file",
                        Value: "output.json",
                    },
                },
                Action: func (context *cli.Context) error {
                    filename := context.String("file")
                    if !isExistFile(filename) {
                        fmt.Println(filename, "not exist.")
                        return nil
                    }
                    args := context.Args().Slice()
                    fmt.Println("filename:", filename)
                    fmt.Println("args:", args)

                    figure, err := file2Figure(filename)
                    if err != nil {
                        log.Fatal(err)
                        return err
                    }

                    if !isExistAreas(figure, args) {
                        fmt.Println(args, "not exist.")
                        return nil
                    }

                    //CombineAreas(filename, context.Args().Slice())

                    return nil
                },
            },
        },
    }

    err := app.Run(os.Args)
    if err != nil {
        log.Fatal(err)
    }
}

