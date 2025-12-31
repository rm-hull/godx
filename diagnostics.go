package godx

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/user"
	"regexp"
	"strconv"
	"strings"

	"github.com/acarl005/stripansi"
	"github.com/earthboundkid/versioninfo/v2"
)

func Diagnostics(logger *slog.Logger) {
	currentUser, err := user.Current()
	if err != nil {
		logger.Warn("Error retrieving current user", "error", err)
	}

	groups, err := groups()
	if err != nil {
		logger.Warn("Error retrieving groups", "error", err)
	}

	logger.Info("Diagnostics",
		"git-version", versioninfo.Short(),
		"pid", os.Getpid(),
		"user", currentUser,
		"groups", groups,
		"environment", env(),
	)
}

func GitVersion() {
	log.Printf("Version: %s", versioninfo.Short())
}

func EnvironmentVars() {
	log.Println("Environment variables")
	for key, val := range env() {
		log.Printf("  %s: %s", key, val)
	}
}

func UserInfo() {
	log.Printf("PID: %d", os.Getpid())
	currentUser, err := user.Current()
	if err != nil {
		log.Printf("Error getting current user: %v", err)
	} else {
		log.Printf("User: uid=%s(%s) gid=%s", currentUser.Uid, currentUser.Username, currentUser.Gid)
	}
	groups, err := groups()
	if err != nil {
		log.Printf("Error getting groups: %v", err)
	} else {
		log.Printf("Groups: %v", groups)
	}
}

func groups() ([]string, error) {
	gids, err := os.Getgroups()
	if err != nil {
		return nil, err
	}
	groupNames := make([]string, 0, len(gids))
	for _, gid := range gids {
		group, err := user.LookupGroupId(strconv.Itoa(gid))
		if err != nil {
			groupNames = append(groupNames, strconv.Itoa(gid)) // Append ID if name lookup fails
		} else {
			groupNames = append(groupNames, fmt.Sprintf("%s(%s)", group.Name, group.Gid))
		}
	}
	return groupNames, nil
}

func env() map[string]string {
	sensitiveRegex := regexp.MustCompile(`(?i)(PASSWORD|API_KEY|ACCESS_KEY|SECRET|TOKEN)`)
	environ := os.Environ()
	envMap := make(map[string]string)
	for _, entry := range environ {
		kv := strings.SplitN(entry, "=", 2)
		if sensitiveRegex.MatchString(kv[0]) {
			envMap[kv[0]] = "********"
		} else {
			envMap[kv[0]] = stripansi.Strip(kv[1])
		}
	}
	return envMap
}
