# AWS CodeDeploy Agent in GO

## WORK IN PROGRESS

### Introduction

This is an attempt to rewrite AWS CodeDeploy Agent in Golang, since the current Ruby implementation is not supported on all [ruby/Ubuntu versions](https://github.com/aws/aws-codedeploy-agent/issues/61), has [high memory consumption](https://github.com/aws/aws-codedeploy-agent/issues/32) and add a tons of [system dependencies](https://github.com/aws/aws-codedeploy-agent/blob/master/bin/install).

This project is a result of some frustration I had with the current solution. Time permitting I will continue on this attempt. Stay tuned :)

### Blockers

* It appears that [aws-sdk-go](https://github.com/aws/aws-sdk-go) doesn't support the [codedeploy-commands API](https://github.com/aws/aws-codedeploy-agent/tree/master/vendor/gems/codedeploy-commands-1.0.0)