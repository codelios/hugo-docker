package api

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

const (
	// BranchPrefixNeedle represents the prefix in the name before the branch
	BranchPrefixNeedle = "origin/"
)

// GetDockerBranchName retrieves the local branch name for a given release
func GetDockerBranchName(release string) string {
	return release[1:]
}

func processReleases(releases []Release) error {
	remoteBranches, err := getDockerRemoteBranches()
	if err != nil {
		return err
	}
	branchMap := make(map[string]bool, 0)
	for k := 0; k < len(remoteBranches); k++ {
		branchMap[remoteBranches[k]] = true
	}
	fmt.Printf("output: %s\n", fmt.Sprint(remoteBranches))
	for i := 0; i < len(releases); i++ {
		if len(releases[i].Name) == 0 {
			continue
		}
		err := processRelease(releases[i].Name, branchMap)
		if err != nil {
			return err
		}
	}
	return nil
}

func runCommand(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
func createLocalBranch(branchName string) error {
	fmt.Printf("About to create local branch %s \n", branchName)
	err := runCommand("git", "checkout", "master")
	if err != nil {
		fmt.Printf("Failed to checkout master - Begin: %s\n", err)
		return err
	}
	err = runCommand("git", "checkout", "-b", branchName)
	if err != nil {
		fmt.Printf("Failed to create and checkout branch: %s\n", err)
		return err
	}
	err = runCommand("git", "push", "--set-upstream", "origin", branchName)
	if err != nil {
		fmt.Printf("Failed to set upstream for branch: %s\n", err)
		return err
	}
	err = runCommand("git", "checkout", "master")
	if err != nil {
		fmt.Printf("Failed to checkout master - End : %s\n", err)
		return err
	}
	return nil
}

func processRelease(releaseName string, branchMap map[string]bool) error {
	branchName := GetDockerBranchName(releaseName)
	_, ok := branchMap[branchName]
	fmt.Printf("%s -> %s -> %v \n", releaseName, branchName, ok)
	if ok {
		return nil
	}
	err := createLocalBranch(branchName)
	if err != nil {
		return err
	}
	branchMap[branchName] = true
	return nil
}

func getDockerRemoteBranches() ([]string, error) {
	branchRemoteCmd := exec.Command("git", "branch", "-r")
	cmdOutput := &bytes.Buffer{}
	branchRemoteCmd.Stdout = cmdOutput
	err := branchRemoteCmd.Run()
	if err != nil {
		return nil, err
	}
	allOutput := string(cmdOutput.Bytes())
	lines := strings.Split(allOutput, "\n")
	output := make([]string, 0)
	for i := 0; i < len(lines); i++ {
		trimmedLine := strings.Trim(lines[i], " ")
		if !strings.HasPrefix(trimmedLine, BranchPrefixNeedle) {
			continue
		}
		thisBranch := trimmedLine[len(BranchPrefixNeedle):]
		output = append(output, thisBranch)
	}
	return output, nil
}
