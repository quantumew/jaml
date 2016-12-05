//usr/bin/env go run $0 $@; exit;
package main

import (
	"github.com/docopt/docopt-go"
	"github.com/ghodss/yaml"
    "encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var logger = log.New(os.Stderr, "", 0)

func main() {
    doc := `Jamal

        Command line interface for converting JSON to YAML and YAML to JSON.
        Expects either a input file or data from stdin.

        Usage:
            jamal [<action>] [<input-file>]

        Options:
            -h --help       Show this message.

        Arguments:
        <action>        Conversion action. [yaml2json | json2yaml] [default: json2yaml]

            <input-file>    Path to data file.
    `
    arguments, _ := docopt.Parse(doc, nil, true, "Jamal 1.0.0", false)
    dataPath := arguments["<input-file>"]
    action := arguments["<action>"]

    var (
        err error
        data []byte
        decodedData []byte
    )

    if dataPath == nil {
        data, err = ioutil.ReadAll(os.Stdin)
    } else {
        path := dataPath.(string)
        data, err = ioutil.ReadFile(path)
    }

    if err != nil {
        logError("Error occurred loading data.", err)

        os.Exit(1)
    }

    if action == "yaml2json" {
        decodedData, err = yamlToJson(data)
    } else {
        decodedData, err = jsonToYaml(data)
    }

    if err != nil {
        logError("Error occurred converting data.", err)

        os.Exit(1)
    }

    os.Stdout.Write(decodedData)
}

// Converts YAML to JSON.
func yamlToJson(raw []byte) ([]byte, error) {
	var data interface{}

	err := yaml.Unmarshal(raw, &data)

    if err != nil {
        return nil, err
    }

    output, err := json.Marshal(data)

    return output, err
}

// Converts JSON to YAML.
func jsonToYaml(raw []byte) ([]byte, error) {
	var data interface{}

	err := json.Unmarshal(raw, &data)

    if err != nil {
        return nil, err
    }

    output, err := yaml.Marshal(data)

    return output, err
}

func logError(msg string, err error) {
    logger.Println(msg)
    logger.Println(err.Error())
}