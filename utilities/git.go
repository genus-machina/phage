package utilities

import (
	"log"
	"os/exec"
	"path/filepath"
)

func attachCommandLogger(cmd *exec.Cmd, logger *log.Logger) *exec.Cmd {
	cmd.Stderr = logger.Writer()
	cmd.Stdout = logger.Writer()
	return cmd
}

func CloneGitRepository(logger *log.Logger, repo, destination string) error {
	cmd := exec.Command("git", "clone", repo, destination)
	attachCommandLogger(cmd, logger)
	return cmd.Run()
}

func UpdateGitRepository(logger *log.Logger, repo string) error {
	path, err := filepath.Abs(repo)
	if err != nil {
		return err
	}

	cmd := exec.Command("git", "pull", "origin")
	cmd.Dir = path
	attachCommandLogger(cmd, logger)
	return cmd.Run()
}
