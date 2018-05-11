package gaia

import (
    "errors"
    "fmt"
    "os"
    "strconv"
)

const INVALID_PORT = -1
const MIN_PORT = 1
const MAX_PORT = 65535

func GetPortFromEnv(envStr string) (int, error) {
  portStr, found := os.LookupEnv(envStr)
    if !found  {
        msg := fmt.Sprintf("Porter ... unset enviroment variable %s", portStr)
        return INVALID_PORT, errors.New(msg)
    }

    port, err := strconv.Atoi(portStr)
    if err != nil {
        msg := fmt.Sprintf("Porter ... non-integer port %s for environment %s", portStr, envStr)
        return INVALID_PORT, errors.New(msg)
    }

    if port < MIN_PORT {
        msg := fmt.Sprintf("Port value %d is below min port %d", port, MIN_PORT)
        return INVALID_PORT, errors.New(msg)
    }

    if port > MAX_PORT {
        msg := fmt.Sprintf("Port value %d is above max port %d", port, MAX_PORT)
        return INVALID_PORT, errors.New(msg)
    }
    return port, nil
}
