package dotenv

import (
	"os"
	"testing"
)

func TestGetEnvFile_Correctly(t *testing.T) {
	Configure("./examples/.env")

	envTest := os.Getenv("ENV_TEST")
	envWithSpace := os.Getenv("ENV_WITH_SPACE")
	envWithCommented := os.Getenv("ENV_WITH_COMMENTED")

	if envTest != "testing" {
		t.Error("env ENV_TEST must contain 'testing'")
	}

	if envWithSpace != "space" {
		t.Error("env ENV_WITH_SPACE must contain 'space' string")
	}

	if envWithCommented != "commented" {
		t.Error("env ENV_WITH_COMMENTED must contain 'commented' string")
	}
}
