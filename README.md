cast 
===

[![Build Status](https://travis-ci.org/goraz/cast.svg?branch=master)](https://travis-ci.org/goraz/cast)
[![Coverage Status](https://coveralls.io/repos/github/goraz/cast/badge.svg?branch=master)](https://coveralls.io/github/goraz/cast?branch=master)
[![GoDoc](https://godoc.org/github.com/goraz/cast?status.svg)](https://godoc.org/github.com/goraz/cast)

Cast is a package that can cast different go types to other types.
You don't need to be worried about performance because this package doesn't use reflect package so it's fast enough




## Installiton

```bash
go get -u github.com/goraz/cast
```

## Usage
This package has String, Int, Float, Duration, Bool and Uint functions than can cast different input type to specified type.
Example:

```go
castVal, _ := String(124) // "124"
castVal, _ := String(true) // "true"
castVal, _ := String(false) // "false"
castVal, _ := String(nil) // "false"
```


Also, there are equivalent MustString, MustInt, MustFloat, MustDuration, MustBool and MustUint functions which instead of returning error, make a panic error

## Example:

```go
castVal := MustString(124) // "124"
castVal := MustString(true) // "true"
castVal := MustString(false) // "false"
castVal := MustString(nil) // "false"
```

For more examples read tests


Contributing
------------
If you are interested in contributing to this project please create a pull request. I appreciate your help!
