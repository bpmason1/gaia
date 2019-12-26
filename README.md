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
