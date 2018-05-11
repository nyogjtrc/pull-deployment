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
	Name    string
	URL     string
	Version string
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

func mirror(url string) error {
	cmd := exec.Command("git", "clone", "--mirror", url)

	var stdOut, errOut bytes.Buffer
	cmd.Stdout = &stdOut
	cmd.Stderr = &errOut

	err := cmd.Run()
	fmt.Println(stdOut.String(), errOut.String())
	return err
}

// pull update git mirror repo with delete branch
// reference:
// https://stackoverflow.com/questions/7068541/
func pull(name string) error {
	cmd := exec.Command("git", "--git-dir="+name+".git", "remote", "update", "--prune")

	var stdOut, errOut bytes.Buffer
	cmd.Stdout = &stdOut
	cmd.Stderr = &errOut

	err := cmd.Run()
	fmt.Println(stdOut.String(), errOut.String())
	return err
}

func findRepoDir(path string) error {
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

	err := findRepoDir(conf.RepoPath)
	if err != nil {
		panic(err)
	}
	err = os.Chdir(conf.RepoPath)
	if err != nil {
		panic(err)
	}

	for _, c := range conf.Projects {
		path := fmt.Sprintf("%s.git", c.Name)
		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			fmt.Println(c.Name, "not existed")
			mirror(c.URL)
		} else {
			fmt.Println(c.Name, "existed")
			pull(c.Name)
		}
	}
}
