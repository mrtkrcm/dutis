# Dutis

A command-line tool to select default applications. It is a wrapper around [duti](https://github.com/moretension/duti).

## Installation

### Using HomeBrew

```shell
brew tap mrtkrcm/dutis https://github.com/mrtkrcm/dutis
brew install dutis
```

### Using Go

```shell
go install github.com/mrtkrcm/dutis@latest
```

## Usage

```shell
dutis
```

## Screenshots

1. Waiting for environment checking

![env-check](./images/env-check.png)

1. Selecting suffix

![choose-suffix](./images/choose-suffix.png)

1. Checking recommended applications

![recommend](./images/recommend.png)

1. Selecting application UTI

![choose-uti](./images/choose-uti.png)

1. Finished

![finish](./images/finish.png)

## Stargazers over time

[![Stargazers over time](https://starchart.cc/mrtkrcm/dutis.svg?variant=adaptive)](https://starchart.cc/mrtkrcm/dutis)

## Release

To trigger a release using GitHub Actions, follow these steps:

1. Create a new release on GitHub.
2. The release workflow will automatically run and create a new release using GoReleaser.
