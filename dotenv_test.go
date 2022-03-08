package dotenv

import (
	"os"
	"testing"
)

func Test_Configure_GetEnvFile_Correctly(t *testing.T) {
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

func Test_Configure_ReturnErr_If_File_Dont_Exist(t *testing.T) {
	err := Configure("/examples/.wrongfile")

	if err == nil {
		t.Error("env should return error when file do not exist")
	}
}
