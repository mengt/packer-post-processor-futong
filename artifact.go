package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const BuilderId = "packer.post-processor.futong"

type Artifact struct {
	files []string
	Url string
}

func NewArtifact(files []string) (*Artifact, error) {
	artifact := &Artifact{}
	for _, f := range files {
		globfiles, err := filepath.Glob(f)
		if err != nil {
			return nil, err
		}
		for _, gf := range globfiles {
			if _, err := os.Stat(gf); err != nil {
				return nil, err
			}
			artifact.files = append(artifact.files, gf)
		}
	}
	return artifact, nil
}

func (a *Artifact) BuilderId() string {
	return BuilderId
}

func (a *Artifact) Files() []string {
	return a.files
}

func (a *Artifact) Id() string {
	return ""
}

func (a *Artifact) String() string {
	files := strings.Join(a.files, ", ")
	return fmt.Sprintf("URL of the template : %s", a.Url)
}

func (a *Artifact) State(name string) interface{} {
	return nil
}

func (a *Artifact) Destroy() error {
        // TODO
	// Delete the template from Futong
	return nil
}
