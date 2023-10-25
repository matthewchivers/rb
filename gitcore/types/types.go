package types

import "strings"

// Repo is a struct containing information about a git repository
type Repo struct {
	Owner string
	Name  string
	Host  string
}

// Validate checks a repo for valid host, owner and repo name
func (r *Repo) Validate() bool {
	return isValidHost(r.Host) && isValidOwner(r.Owner) && isValidRepo(r.Name)
}

func isValidHost(host string) bool {
	return host != "" && strings.Contains(host, ".")
}

func isValidOwner(owner string) bool {
	return owner != ""
}

func isValidRepo(repo string) bool {
	return repo != ""
}
