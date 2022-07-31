# Fielder

*Fielder* is a tool to generate Go code that extracts fields from a struct and transforms them into ENUM.
Also, it adds useful methods and functions. 

## Motivation

When using ORM-s like `gorm`, `go-pg`, `bun` you have to pass column names as arguments of different methods.
This is a pain to use raw strings for that, and it also might be a security risk.
Much better to rely on ENUM which represents columns for a specific table.

## Features

The library provides the following features:

  * Generating ENUM for the struct fields
  * Generating method to validate ENUM values
  * Generating function to list all ENUM values
  * Tag-based field names extraction
  * Extraction of embedded fields
  * Excluding fields 
  * Different formatting (camel, pascal, snake)
  * Template overriding

## Installation

    $ go install github.com/kpeu3i/fielder@v1.0.1

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


For more details, check example in the `example` folder. 

The following CLI flags are allowed:

```
Usage of fielder:
  -embedded
        Extract embedded fields (default false)
  -excluded string
        Comma separated list of excluded fields (default "")
  -format string
        Format of the generated type values (default "as_is")
  -output string
        Set output filename (default "<src_dir>/<type>_fielder.go")
  -pkg string
        Package name to extract type from (default ".")
  -suffix string
        Suffix for the generated struct (default "Field")
  -tag string
        Use tag to extract field values from (default "")
  -tpl string
        Set template filename (default "")
  -type string
        Type to extract fields from
```
