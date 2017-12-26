package publisher

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// Repo ...
type Repo struct {
	Name      string
	Bucket    string
	Directory string
	UrlPrefix string
}

// Path ...
func (r *Repo) Path(filename string) string {
	path := filename

	if r.Directory != "" {
		path = r.Directory + "/" + filename
	}

	return path
}

// ProxyUrl
func (r *Repo) ProxyUrl(filename string) string {
	path := filename

	if r.UrlPrefix != "" && r.Directory != "" {
		path = r.UrlPrefix + "/" + r.Directory + "/" + filename
	} else if r.UrlPrefix != "" {
		path = r.UrlPrefix + "/" + filename
	}

	return path
}

// Repos ...
type Repos []Repo

// Get ...
func (repos *Repos) Get(name string) (*Repo, error) {
	for _, r := range *repos {
		if r.Name == name {
			return &r, nil
		}
	}

	return nil, ResourceNotFoundErr("repository not defined")
}

func decodeRepos(c interface{}) (Repos, error) {
	repos := Repos{}
	if err := mapstructure.Decode(c, &repos); err != nil {
		return nil, err
	}

	for _, r := range repos {
		if r.Bucket == "" {
			return nil, errors.New("invalid config: bucket is required")
		}
	}

	return repos, nil
}
