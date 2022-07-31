# Fielder

Fielder is a tool to generate Go code that extracts fields from a struct and transforms them into ENUM.
Also, it adds useful methods and functions. 

## Motivation

When using ORM-s like `gorm`, `go-pg`, `bun` you have to pass column names as arguments of different methods.
This is a pain to use raw strings for that, and it also might be a security risk.
Much better to rely on ENUM which represents columns for a specific table.

## Features

The library provides the following features:

  * Generates ENUM for the struct fields
  * Generates method to validate ENUM values
  * Generates function to list all fields
  * A tag can be used when extracting field names from a struct
  * Extracting embedded struct fields are supported
  * Fields excluding is supported
  * Different formats of generated field values are supported (camel, pascal, snake)
  * Template which is used to generate the code can be overridden

## Installation

    $ go install github.com/kpeu3i/fielder@v1.0.0

## Usage

Put the `go:generate` directive in the same package as the struct you want to generate.
For example:

```go
//go:generate fielder -type=UserAccount

package domain
 
type UserAccount struct {
    FirstName string
    LastName  string
    Email     string
    Password  string
}
```

Then, run command bellow to generate the code:

    $ go generate ./...

You can check for example in the `example` folder.

The following CLI flags are allowed:

```
Usage of fielder:
  -embedded
    	Extract embedded fields
  -excluded string
    	Comma separated list of excluded fields
  -format string
    	Format of the generated type values (default "as_is")
  -output string
    	Output file name (default <src_dir>/<type>_fielder.go)
  -pkg string
    	Package name to extract type from (default ".")
  -suffix string
    	Generated type name suffix (default "Field")
  -tag string
    	Extract field names from struct tag
  -tpl string
    	Override template file
  -type string
    	Type name to generate fields for
```
