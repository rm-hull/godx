package godx

import (
	"bytes"
	"log"
	"log/slog"
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

func TestEnvironmentVarsStrippingANSI(t *testing.T) {
	ansiValue := "\x1b[31mred\x1b[0mtext"
	_ = os.Setenv("TEST_ANSI_VAR", ansiValue)
	defer func() {
		_ = os.Unsetenv("TEST_ANSI_VAR")
	}()

	output := captureLog(EnvironmentVars)

	assert.Contains(t, output, "TEST_ANSI_VAR: redtext", "EnvironmentVars() should strip ANSI codes")
	assert.NotContains(t, output, ansiValue, "EnvironmentVars() should not contain raw ANSI codes")
}

func TestDiagnostics(t *testing.T) {
	var buf bytes.Buffer
	logger := slog.New(slog.NewTextHandler(&buf, nil))

	Diagnostics(logger)

	logOutput := buf.String()
	assert.Contains(t, logOutput, "Diagnostics", "Diagnostics() should log diagnostics information")
	assert.Contains(t, logOutput, "git-version=", "Diagnostics() should log git version")
	assert.Contains(t, logOutput, "pid=", "Diagnostics() should log PID")
	assert.Contains(t, logOutput, "user=", "Diagnostics() should log user information")
	assert.Contains(t, logOutput, "groups=", "Diagnostics() should log group information")
	assert.Contains(t, logOutput, "environment=", "Diagnostics() should log environment variables")
}
