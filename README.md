# Fielder

*Fielder* is a tool to generate Go code that extracts fields from a struct and transforms them into ENUM.
Also, it adds useful methods and functions. 

## Motivation

When using ORM-s like `gorm`, `go-pg`, `bun` you have to pass column names as arguments of different methods.
This is a pain to use raw strings for that, and it also might be a security risk.
Much better to rely on ENUM which represents columns for a specific table.
Also, the generated field names can be used in combination with Golang reflection (`FieldByName`) for different purposes.

## Features

The library provides the following features:

  * Struct fields representation with ENUM 
  * Different functions and methods to work with ENUM (validation, listing, conversion to string, etc)
  * Tag-based field names extraction
  * Embedded fields extraction
  * Fields excluding  
  * Different formatting (camel, pascal, snake)
  * Template overriding

## Installation

    $ go install github.com/kpeu3i/fielder@v1.1.0

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
    * Method `<Type><Suffix>List::Add` adds the value to the collection.
    * Method `<Type><Suffix>List::AddIfNotContains` adds the value to the collection if it is not already in the collection.
    * Method `<Type><Suffix>List::Remove` removes the value from the collection.
    * Method `<Type><Suffix>List::Clear` clears the collection.
    * Method `<Type><Suffix>List::Strings` returns a slice with all the strings of the collection items.
* Functions:
    * `<Type><Suffix>Values` returns a slice with all the values of the ENUM.
    * `<Type><Suffix>Strings` returns a slice with all the strings of the ENUM.
    * `New<Type><Suffix>` returns a new collection with all the values of the ENUM.
