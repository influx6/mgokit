Trail
---------

Trail brings web assets bundling and preprocessing made simple in Go. It provides set of simple processors that help
to manage and bundle your different assets and files, letting you bundle them along with your application binaries.


## Install

```
go get github.com/influx6/trail
```

## Usage

Simple call the `trail` CLI with a name for the giving package:

```bash
> trail
Usage: trail [options]
Trail creates a package for package of web assets using its internal bundlers.

COMMANDS:

	trail view [optional-name]	# Creates a generate.go file which bundles all assets in create directory.
	trail public [optional-name]	# Creates a complete package and content for asset bundling all static files

where:

	[optional-name] defines the name for the directory to be used for the assets if provided, else
	having files created within working directory.

EXAMPLES:

	trail view home			# Creates a generate.go file which bundles all assets in create directory.
	trail public static-data	# Creates a complete package and content for asset bundling all static files


FLAGS:
	-v      Print version.
	-f 	Force re-generation of all files
```


## Contribution

Please let all suggestions, thoughts and complaints come has issues and associated PRs.
