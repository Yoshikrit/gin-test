package configs_test

import (
    "os"
    "testing"
	"fmt"

    "github.com/stretchr/testify/assert"

    "gin-test/configs"
)

func TestLoadEnv(t *testing.T) {
    testCases := []struct {
        name       string
        env        string
        envName    string
        expected   string
    }{
        {name: "test case : development environment", env: "development", envName: "dev", expected: ".env.dev"},
        {name: "test case : production environment", env: "production", envName: "prod", expected: ".env.prod"},
        {name: "test case : invalid environment", env: "", envName: "", expected: ".env."},
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
			t.Setenv("APP_ENV", tc.env)

            configs.LoadEnv()

			res := os.Getenv("APP_ENV")
            assert.Equal(t, tc.env, res)

			filename := fmt.Sprintf(".env.%s", tc.envName)
            assert.Equal(t, tc.expected, filename)
        })
    }
}
