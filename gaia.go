package gaia

import (
    "errors"
    "fmt"
    "log"
    "os"
    "strconv"
)

const INVALID_PORT = -1
const MIN_PORT = 1
const MAX_PORT = 65535

func GetEnvOrDie(envStr string) string {
    value, found := os.LookupEnv(envStr)
    if !found {
        msg := fmt.Sprintf("Missing required enviroment variable %s", envStr)
        log.Fatal(msg)
        os.Exit(-1)
    }
    return value
}

func GetEnvWithDefault(envStr, defaultStr string) string {
    value, found := os.LookupEnv(envStr)
    if !found {
        msg := fmt.Sprintf("Missing enviroment variable %s ... using default %s\n", envStr, defaultStr)
        log.Println(msg)
        return defaultStr
    }
    return value
}

func GetDirectoryNameOrDie(envStr string) string {
  dirname := GetEnvOrDie(envStr)
  fileInfo, err := os.Stat(dirname)
  if err != nil {
    log.Println(err)
    os.Exit(-1)
  } else if(fileInfo.IsDir() == false) {
    log.Println(dirname, "exists but is not a directory")
    os.Exit(-1)
  }
  return dirname
}

func GetInteger(envStr string) (int, error) {
    intStr, found := os.LookupEnv(envStr)
    if !found {
        msg := fmt.Sprintf("Unset enviroment variable %s", envStr)
        return 0, errors.New(msg)
    }

    num, err := strconv.Atoi(intStr)
    if err != nil {
        msg := fmt.Sprintf("Non-integer value %s for environment %s", intStr, envStr)
        return 0, errors.New(msg)
    }
    return num, nil
}

func GetIntegerWithDefault(envStr string, defaultInt int) int {
    val, err := GetInteger(envStr)
    if err != nil {
        return defaultInt
    }
    return val
}

func GetIntegerInRange(envStr string, minInt, maxInt int) (int, error) {
  num, err := GetInteger(envStr)
  if err != nil {
      return 0, err
  }

  if num < minInt {
      msg := fmt.Sprintf("Value %d is below min port %d", num, minInt)
      return 0, errors.New(msg)
  }

  if num > maxInt {
      msg := fmt.Sprintf("Value %d is above max port %d", num, maxInt)
      return 0, errors.New(msg)
  }
  return num, nil
}

func GetPort(envStr string) (int, error) {
    return GetIntegerInRange(envStr, MIN_PORT, MAX_PORT)
}

func GetPortWithDefault(envStr string, defaultPort int) int {
    port, err := GetPort(envStr)
    if err != nil {
        log.Println(err)
        log.Println("Warning - using default port")
        return defaultPort
    }
    return port
}

func GetPortOrDie(envStr string) int {
    port, err := GetPort(envStr)
    if err != nil {
        log.Fatal(err)
        os.Exit(-1)
    }
    return port
}
