# Contributing Guide

We happily welcome contributions to this package. We use [GitHub Issues](https://github.com/databricks/databricks-sql-go/issues) to track community reported issues and [GitHub Pull Requests](https://github.com/databricks/databricks-sql-go/pulls) for accepting changes.

Contributions are licensed on a license-in/license-out basis.

## Communication

Before starting work on a major feature, please reach out to us via GitHub, Slack, email, etc. We will make sure no one else is already working on it and ask you to open a GitHub issue.
A "major feature" is defined as any change that is > 100 LOC altered (not including tests), or changes any user-facing behavior.
We will use the GitHub issue to discuss the feature and come to agreement.
This is to prevent your time being wasted, as well as ours.
The GitHub review process for major features is also important so that organizations with commit access can come to agreement on design.
If it is appropriate to write a design document, the document must be hosted either in the GitHub tracking issue, or linked to from the issue and hosted in a world-readable location.
Specifically, if the goal is to add a new extension, please read the extension policy.
Small patches and bug fixes don't need prior communication.

## Developer Certificate of Origin

To contribute to this repository, you must sign off your commits to certify that you have the right to contribute the code and that it complies with the open source license. The rules are simple: if you can certify the content of [DCO](./DCO), then add a "Signed-off-by" line to your commit message to certify your compliance. Please use your real name, as pseudonymous/anonymous contributions are not accepted.

```
Signed-off-by: Joe Smith <joe.smith@email.com>
```

If you set your `user.name` and `user.email` git configs, you can sign your commit automatically with `git commit -s`:

```
git commit -s -m "Your commit message"
```

## Pull Request Process

1. Update the [CHANGELOG](CHANGELOG.md) with details of your changes, if applicable.
2. Add any appropriate tests.
3. Make your code or other changes.
4. Review guidelines such as
   [How to write the perfect pull request][github-perfect-pr], thanks!

[github-perfect-pr]: https://blog.github.com/2015-01-21-how-to-write-the-perfect-pull-request/
