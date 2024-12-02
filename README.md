# cobra-cli-samples

This is a CLI App example using the [Cobra](https://github.com/spf13/cobra) and [Viper](https://github.com/spf13/viper) libraries for Go. This application provides commands (per Cobra) to read, write, update, and delete configuration records from a configuration file using Viper.

## Using the example CLI App.

`<cli> config -h` provides the following documentation for using this CLI App.

```
Available Commands:  (See examples below for syntax.)
  add         The 'add' subcommand will add a passed in key value pair to the application configuration file. 
  delete      The 'delete' subcommand removes a key value pair from the configuration file. 
  update      The 'update' subcommand will update a passed in key value pair for an existing set of data to the application configuration file.
  view        The 'view' subcommand will provide a list of keys and a map of the values.

Flags:
  -h, --help           help for config
  -k, --key string     The key for the key value set to add to the configuration.
  -v, --value string   The value for the key value set to add to the configuration.
```

> Currently I got this to work but it isn't in the best shape. Check out the code [here](https://github.com/Adron/cobra-cli-samples/blob/master/cmd/delete.go).

## Get Involved, Add Samples, Make Requests

If you'd like to get involved and add samples or make request for additional samples please file an issue [here](https://github.com/Adron/cobra-cli-samples/issues/new?assignees=&labels=&template=feature_request.md&title=)!

If you've found any bugs or issues with the code please file a bug report [here](https://github.com/Adron/cobra-cli-samples/issues/new?assignees=&labels=&template=bug_report.md&title=)!

### Examples, The CRUD!

`./cli config add -k "blog" -v "https://compositecode.blog/"` example writes a record to the configuration file with a key of "blog" and a value of "https://compositecode.blog/".

`./cli config view` displays the contents of the configuration file and CLI specific environment variables. These are the configuration files located in the `.cobrae-cli-samples.yml` and environment variables prefaced with `COBRACLISAMPLES`. The list of keys is displaced first and then the keys and values are displayed below that.

`./cli config update -k "blog" -v "not found"` will update the blog entry in the configuration to read `not found` for the value.

`./cli config delete ...` will delete the key and value from the configuration file.

## Building the Project

Following a fairly standard clone, one can build this project with a single step using the `./build.sh` file. If you'd like to contribute the same for Windows, feel free I'd be happy to pull that PR in. Once the project is built use the CLI as defined above.

## Installing the CLI App

This application can be installed as a CLI app by referencing it's location in your bash (powershell? etc) startup script.
