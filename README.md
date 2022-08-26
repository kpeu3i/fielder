# Fielder

*Fielder* is a tool to generate Go code that extracts fields from a struct and transforms them into ENUM.
Also, it adds useful types, methods and functions. 

## Motivation

When using ORM-s like `gorm`, `go-pg`, `bun` you have to pass column names as arguments of different methods.
This is a pain to use raw strings for that, and it also might be a security risk.
Much better to rely on ENUM which represents columns for a specific table.  
Also, the generated field names can be used in combination with Golang reflection (`FieldByName`) for different purposes.

## Features

The library provides the following features:

  * Struct fields representation with ENUM.
  * Different functions and methods to work with ENUM (validation, listing, conversion to string, etc).
  * Tag-based field names extraction (regex can be used to extract a value from a tag).
  * Embedded fields extraction.
  * Fields excluding.
  * Different formatting (camel, pascal, snake).
  * Template overriding.

## Installation

    go install github.com/kpeu3i/fielder@v1.5.0

## Usage

Put the `go:generate` directive in the same package as the struct you want to generate.
For example:

```go
//go:generate fielder -type=UserAccount

package models
 
type UserAccount struct {
    FirstName string
    LastName  string
    Email     string
    Password  string
}
```

Then, run command bellow to generate the code:

    $ go generate ./...

The following formatting strategies can be applied for the extracted field names (see `format` and `tag_format` flag):
 * `snake_case` (e.g `first_name`)
 * `camel_case` (e.g `firstName`)
 * `pascal_case` (e.g `FirstName`)

For more details, check the [examples](examples) folder in the project root. 

The following CLI flags are allowed:

```
Usage of fielder:
  -embedded
    	Extract embedded fields (default false)
  -excluded string
    	Comma separated list of excluded fields (default "")
  -format string
    	Format of the generated type values extracted from the struct field (default "as_is")
  -output string
    	Set output filename (default "<src_dir>/<type>_fielder.go")
  -pkg string
    	Package name to extract type from (default ".")
  -suffix string
    	Suffix for the generated struct (default "Field")
  -tag string
    	Tag to extract field values from (default "")
  -tag_format string
    	Format of the generated type values extracted from the tag (default "as_is")
  -tag_regexp string
    	Regular expression to parse field value from a tag (default "")
  -tag_strict
    	Strict mode for tag parsing (returns error if tag is not found)
  -tpl string
    	Set template filename (default "")
  -type string
    	Type to extract fields from
```

## Generated types, functions and methods

When *Fielder* is applied to a type, it will generate public functions/methods and types:

* Type `<Type><Suffix>` represents fields ENUM:
    * Method `<Type><Suffix>::IsValid` returns true if the value is a valid ENUM.
    * Method `<Type><Suffix>::String` returns the string representation of the value.
* Type `<Type><Suffix>List` represents collection of ENUM fields:
    * Method `<Type><Suffix>List::Len` returns the number of values in the collection.
    * Method `<Type><Suffix>List::Contains` returns true if the collection contains the value.
    * Method `<Type><Suffix>List::Equals` returns true if the two collections are equal.
    * Method `<Type><Suffix>List::Similar` returns true if the two collections contain the same values.
    * Method `<Type><Suffix>List::Add` adds the values to the collection.
    * Method `<Type><Suffix>List::AddIfNotContains` adds the values to the collection if they are not already present.
    * Method `<Type><Suffix>List::Remove` removes the values from the collection.
    * Method `<Type><Suffix>List::Clear` clears the collection.
    * Method `<Type><Suffix>List::Clone` returns a pointer to a copy of the collection.
    * Method `<Type><Suffix>List::Strings` returns a slice with all the strings of the collection items.
* Functions:
    * `<Type><Suffix>Values` returns a slice with all the values of the ENUM.
    * `<Type><Suffix>Strings` returns a slice with all the strings of the ENUM.
    * `New<Type><Suffix>` returns a new collection with all the values of the ENUM.
    * `New<Type><Suffix>With` returns a new collection with the given values of the ENUM.

## Inspiring projects
  * [enumer](https://github.com/dmarkham/enumer)
  * [stringer](https://pkg.go.dev/golang.org/x/tools/cmd/stringer)
