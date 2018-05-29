# Pull Deployment

[![Go Report Card](https://goreportcard.com/badge/github.com/nyogjtrc/pull-repo)](https://goreportcard.com/report/github.com/nyogjtrc/pull-repo)

## Install

```
$ go get -u github.com/nyogjtrc/pull-repo
```

## Configuration File

```yaml
repo_path: ./repo
work_path: ./work
projects:
    - name: f2e-jacket.git
      url: git@github.com:nyogjtrc/f2e-jacket.git
      work_tree: f2e-jacket-master
      version: master
    - name: f2e-jacket.git
      url: git@github.com:nyogjtrc/f2e-jacket.git
      work_tree: f2e-jacket-release
      version: v0.1.3
```

- repo_path: base dir to put git repo
- work_path: base dir to put work tree
