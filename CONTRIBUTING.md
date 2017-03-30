# How to contribute

FuildCurrency is MIT licensed and accepts contributions via GitHub pull requests. This document outlines some of the conventions on commit message formatting, contact points for developers, and other resources to help get contributions into FuildCurrency.

## Getting started

- Fork the repository on GitHub
- Read the README.md for build instructions

## Contribution flow

This is a rough outline of what a contributor's workflow looks like:

- Fork the project.
- Make sure commit messages are verbose enough.
- Push changes in a feature branch.
- Submit a pull request to MichaelTSS/FuildCurrency.
- If you don't get an answer within 2 days, please ping me again!

Thanks for contributing!

### Code style

The coding style suggested by the Golang community is used in FuildCurrency. See the [style doc](https://github.com/golang/go/wiki/CodeReviewComments) for details.

Please follow this style to make FuildCurrency easy to review, maintain and develop.

### Format of the commit message

We follow a rough convention for commit messages that is designed to answer two
questions: what changed and why. The subject line should feature the what and
the body of the commit should describe the why.

```
This add a new crypto-currency provider that has > 99% uptime.

Fixes #38
```
