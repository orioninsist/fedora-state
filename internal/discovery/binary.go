package discovery

import (
	"os"
	"os/user"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"syscall"

	"fedora-state/internal/model"
)

func DiscoverBinaries() []model.Binary {
	seen := map[string]bool{}
	result := []model.Binary{}

	for _, dir := range strings.Split(os.Getenv("PATH"), ":") {
		if dir == "" {
			continue
		}

		absoluteDir, err := filepath.Abs(dir)
		if err != nil {
			continue
		}

		entries, err := os.ReadDir(absoluteDir)
		if err != nil {
			continue
		}

		for _, entry := range entries {
			path := filepath.Join(absoluteDir, entry.Name())

			realPath, err := filepath.EvalSymlinks(path)
			if err != nil {
				continue
			}

			info, err := os.Stat(realPath)
			if err != nil || !info.Mode().IsRegular() || info.Mode()&0111 == 0 {
				continue
			}

			if seen[entry.Name()] {
				continue
			}
			seen[entry.Name()] = true

			result = append(
				result,
				model.Binary{
					Name:     entry.Name(),
					Path:     path,
					RealPath: realPath,
					Owner:    fileOwner(info),
				},
			)
		}
	}

	sort.Slice(result, func(i, j int) bool {
		if result[i].Path == result[j].Path {
			return result[i].RealPath < result[j].RealPath
		}

		return result[i].Path < result[j].Path
	})

	return result
}

func fileOwner(info os.FileInfo) string {
	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		return ""
	}

	uid := strconv.FormatUint(
		uint64(stat.Uid),
		10,
	)

	account, err := user.LookupId(uid)
	if err != nil {
		return uid
	}

	return account.Username
}
