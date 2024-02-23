package git

import (
	"errors"
	"os/exec"
	"strings"
)

type FileDiff struct {
	File string
	Diff string
}

func Diff(target string) ([]FileDiff, error) {
	files, err := changedFiles(target)
	if err != nil {
		return nil, err
	}

	diffs := make([]FileDiff, len(files))
	for i, f := range files {
		d, err := fileDiff(target, f)
		if err != nil {
			return nil, err
		}

		diffs[i] = FileDiff{
			File: f,
			Diff: d,
		}
	}

	return diffs, nil
}

func changedFiles(target string) ([]string, error) {
	out, err := gitExec("diff", "--name-only", target)
	if err != nil {
		return nil, err
	}
	return strings.Split(strings.TrimSpace(out), "\n"), nil
}

func fileDiff(target, file string) (string, error) {
	return gitExec("diff", target, file)
}

func gitExec(args ...string) (string, error) {
	out, err := exec.Command("git", args...).Output()

	if exitErr, ok := err.(*exec.ExitError); ok {
		return string(out), errors.New(string(exitErr.Stderr))
	}

	return string(out), err
}
