[![Build Status](https://travis-ci.org/bpmason1/gaia.svg?branch=master)](https://travis-ci.org/bpmason1/gaia)

# gaia
Simplified environment management for golang apps.

The gaia library is intended as a way to minimize writing boilerplate code when reading and parsing environment variables in an application.  It provides methods that can either apply a sane default or abort the applicatino (fail-fast) when an environment variable is missing or malformed.

## functions
* GetEnvOrDie(envStr string) string
* GetEnvWithDefault(envStr, defaultStr string) string
* GetDirectoryNameOrDie(envStr string) string   ... (will exit if adirectory ${envStr} is not exist)
* GetInteger(envStr string) (int, error)
* GetIntegerWithDefault(envStr string, defaultInt int) int
* GetIntegerInRange(envStr string, minInt, maxInt int) (int, error) ... (the valid range is inclusive of both maxInt and minInt)
* GetPort(envStr string) (int, error) ... (an error will be raised if the port number is outside the range 1 to 65535)
* GetPortWithDefault(envStr string, defaultPort int) int
* GetPortOrDie(envStr string) int

In all cases, the function parameter `envStr` is the name of the environmnt variable.

## Example
from the command line (assumes you are using bash)  
`export YourEnvVar=25`

then create a main.go with the following content  
```
package main

import (
  "fmt"
  "github.com/bpmason1/gaia"
)

func main() {
  fmt.Println("Hello")

  minInt := 20
  maxInt := 30
  i, err := GetIntegerInRange("YourEnvVar", minInt, maxInt)

  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println("Success ... %d is in the range %d to %d", i, minInt, maxInt)
  }
}
```  
This simple example shows gaia's ability to verify not only that the integer in the environment variable is within bounds but to also perform the initial type casting to ensure the input string can be turned into in integer at all.
