package gaia

import (
  "math"
  "math/rand"
  "os"
  "strconv"
  "testing"
)

//------------------------------------------------------------------------------
// helper to create random strngs
//------------------------------------------------------------------------------
var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandString(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}

//------------------------------------------------------------------------------
//------------------------------------------------------------------------------
func GetUnusedEnv() string {
    var randEnvKey string
    for {
      randEnvKey = RandString(60)
      if _, found := os.LookupEnv(randEnvKey); !found {
        return randEnvKey
      }
    }
}

//------------------------------------------------------------------------------
// The tests
//------------------------------------------------------------------------------
func TestMinPort(t *testing.T) {
    if MIN_PORT != 1 {
       t.Errorf("MIN_PORT is %d ... should be 1", MIN_PORT)
    }
}

func TestMaxPort(t *testing.T) {
    expected_max_port := int(math.Pow(2, 16)) - 1
    if MAX_PORT != expected_max_port {
       t.Errorf("MAX_PORT is %d ... should be %d", MAX_PORT, expected_max_port)
    }
}

func TestGetEnvWithDefault(t *testing.T) {
    envName := GetUnusedEnv()
    defaultValue := "mock-default"
    value := GetEnvWithDefault(envName, defaultValue)
    if value != defaultValue {
        t.Errorf("Environment is %s ... should be default %s", value, defaultValue)
    }

    envValue := RandString(60)
    os.Setenv(envName, envValue)
    defer os.Unsetenv(envName)

    value = GetEnvWithDefault(envName, "mock default")
    if value != envValue {
        t.Errorf("Environment is %s ... should be %s", value, envValue)
    }
}

func TestGetInteger(t *testing.T) {
    var value int
    var err error
    envValues := []int{42389, 32469, -44353, 0, -1}
    for _, envValue := range envValues {
        envValueStr := strconv.Itoa(envValue)
        envName := GetUnusedEnv()
        defer os.Unsetenv(envName)  // call is safe even if the env isn't set

        // return a error since the environmenet is unset
        value, err = GetInteger(envName)
        if err == nil {
          t.Errorf("Integer Environment %s was unset but returned value %d", envName, value)
        }

        // set the environment
        os.Setenv(envName, envValueStr)

        // make sure a value (any value) is returned when the environment is
        value, err = GetInteger(envName)
        if err != nil {
          t.Errorf("Integer Environment %s was set but nothing was returned", envName)
        }

        // make sure the returned value matches the set value
        if value != envValue {
            t.Errorf("Integer Environment %s was %d ... should be %d", envName, value, envValue)
        }
    }
}

func TestGetIntegerWithDefault(t *testing.T) {
    var value int
    defaultValues := []int{9, 0, 490, -8787}
    envValues := []int{42389, 32469, -44353, 0, -1}
    for _, envValue := range envValues {
        for _, defaultValue := range defaultValues {
            envValueStr := strconv.Itoa(envValue)
            envName := GetUnusedEnv()
            defer os.Unsetenv(envName)  // call is safe even if the env isn't set

            // return the default since the environmenet is unset
            value = GetIntegerWithDefault(envName, defaultValue)
            if value != defaultValue {
              t.Errorf("Unset integer environment %s was returned %d instead of default %d", envName, value, defaultValue)
            }

            // set the environment
            os.Setenv(envName, envValueStr)

            // make sure a value (any value) is returned when the environment is
            value = GetIntegerWithDefault(envName, defaultValue)

            // make sure the returned value matches the set value
            if value != envValue {
                t.Errorf("Integer Environment %s was %d ... should be %d", envName, value, envValue)
            }
        }
    }
}
