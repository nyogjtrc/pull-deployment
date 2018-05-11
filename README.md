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
    - name: f2e-jacket
      url: git@github.com:nyogjtrc/f2e-jacket.git
      version: master
    - name: awesome
      url: git@github.com:nyogjtrc/awesome.git
      version: master
```

- repo_path: base dir to put git repo
