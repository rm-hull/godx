package godx

import (
	"bytes"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Helper to capture log output
func captureLog(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr) // Reset to default stderr
	}()
	f()
	return buf.String()
}

func TestGitVersion(t *testing.T) {
	output := captureLog(GitVersion)
	assert.Contains(t, output, "Version:", "GitVersion() should log version information")
}

func TestEnvironmentVars(t *testing.T) {
	// Set some dummy environment variables for testing
	_ = os.Setenv("TEST_VAR_1", "value1")
	_ = os.Setenv("TEST_PASSWORD", "secret")
	defer func() {
		_ = os.Unsetenv("TEST_VAR_1")
		_ = os.Unsetenv("TEST_PASSWORD")
	}()

	output := captureLog(EnvironmentVars)

	assert.Contains(t, output, "Environment variables", "EnvironmentVars() should log 'Environment variables'")
	assert.Contains(t, output, "TEST_VAR_1: value1", "EnvironmentVars() should log TEST_VAR_1 correctly")
	assert.Contains(t, output, "TEST_PASSWORD: ********", "EnvironmentVars() should mask TEST_PASSWORD")
}

func TestUserInfo(t *testing.T) {
	output := captureLog(UserInfo)

	assert.Contains(t, output, "PID:", "UserInfo() should log PID")
	assert.Contains(t, output, "User:", "UserInfo() should log user information")
	// We can't reliably test group information without mocking os.Getgroups and user.LookupGroupId,
	// which is beyond the scope of a simple diagnostic test.
	// Just check if the "Groups:" string is present, indicating the function attempted to get groups.
	assert.Contains(t, output, "Groups:", "UserInfo() should log group information")
}
