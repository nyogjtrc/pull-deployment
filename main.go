package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	yaml "gopkg.in/yaml.v2"
)

// Project of git repo
type Project struct {
	Name     string
	URL      string
	WorkTree string `yaml:"work_tree"`
	Version  string
}

// Config of work env
type Config struct {
	RepoPath string `yaml:"repo_path"`
	WorkPath string `yaml:"work_path"`
	Projects []Project
}

// LoadConfig in yaml format
func LoadConfig(filename string) *Config {
	conf := new(Config)
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		log.Fatal(err)
	}

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(bytes, &conf)
	if err != nil {
		log.Fatal(err)
	}

	return conf
}

func execPrinting(name string, args ...string) error {
	cmd := exec.Command(name, args...)

	var stdOut, errOut bytes.Buffer
	cmd.Stdout = &stdOut
	cmd.Stderr = &errOut

	err := cmd.Run()
	fmt.Println(stdOut.String(), errOut.String())
	return err
}

// mirror execute git clone --mirror
func mirror(url string, path string) error {
	return execPrinting("git", "clone", "--mirror", url, path)
}

// pull execute git --git-dir=<name> remote update --prune
// to update git mirror repo with delete branch
// reference:
// https://stackoverflow.com/questions/7068541/
func pull(name string) error {
	return execPrinting(
		"git",
		"--git-dir="+name,
		"remote",
		"update",
		"--prune",
	)
}

// checkout file
// git --git-dir=<name> --work-tree=<work-tree> checkout <commit>
func checkout(name string, worktree string, version string) error {
	return execPrinting(
		"git",
		"--git-dir="+name,
		"--work-tree="+worktree,
		"checkout",
		"-f",
		version,
	)
}

func findAndCreateDir(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		er := os.Mkdir(path, os.ModePerm)
		if er != nil {
			return er
		}
		return nil
	}

	return err
}

func main() {
	conf := LoadConfig("config.yaml")

	err := findAndCreateDir(conf.RepoPath)
	PanicWhenError(err)
	err = findAndCreateDir(conf.WorkPath)
	PanicWhenError(err)

	for _, c := range conf.Projects {
		// update
		repoPath := fmt.Sprintf("%s/%s", conf.RepoPath, c.Name)
		fmt.Println(repoPath)
		_, err := os.Stat(repoPath)
		if os.IsNotExist(err) {
			fmt.Println(c.Name, "not existed")
			mirror(c.URL, repoPath)
		} else {
			fmt.Println(c.Name, "existed")
			pull(repoPath)
		}

		// checkout
		workPath := fmt.Sprintf("%s/%s/", conf.WorkPath, c.WorkTree)
		fmt.Println(workPath)
		err = findAndCreateDir(workPath)
		PanicWhenError(err)
		checkout(repoPath, workPath, c.Version)
	}
}

// PanicWhenError execute panic when error is not nil
func PanicWhenError(err error) {
	if err != nil {
		panic(err)
	}
}
