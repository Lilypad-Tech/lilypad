# Lilypad Contribution Guide

# Contributing to Lilypad

Thanks for your interest in contributing to the Lilypad Network! Welcome, we are happy to have you. 🐸 💚

This guide will help you get started with contributing to the Lilypad network. There are many ways to get contribute in the network, including contacting us about bugs on Discord, opening or triaging issues, or submitting a pull request to address an issue.

**No contribution is too small and all contributions are valued.**

## Code of conduct

We expect contributors to abide by our Code of Conduct. Violations of the code of conduct can be reported to the Lilypad team at contact@lilypad.tech.

## Contact on Discord

We accept bug reports in the Lilypad Discord [#i-need-help channel](https://discord.com/channels/1212897693450641498/1230231823674642513). For Resource Providers and Job creators, please use the [ticketing format](https://docs.lilypad.tech/lilypad/hardware-providers/troubleshooting#dont-see-your-issue-below) to ensure a smooth exchange of info. If you have a well-defined, reproducible issue we invite you to open an issue as described below.

## Issues

Issues are how we consider new features or address bugs in our code base.

### Open an issue

When you open a issue, you will be presented with an option to open a feature request or report a bug. Select the appropriate option, and fill in the provided template. Please do your best to fill out the details, but feel free to skip parts if you're not sure how to answer.

For bug reports, the most important information is a description of the behavior you are seeing and a test case to help us reproduce the behavior. If we are unable to reproduce the bug, we may ask for additional details to help us reproduce.

Feature requests should include a detailed description of the proposed feature and an explanation of how it will benefit the project. We will consider feature requests in the broader context of the project to determine whether we would like to add them.

Some changes like updating docs or typo fixes do not require an issue. Feel free to submit a pull request directly for these changes.

### Claim an issue

Leave a comment on the issue like “I would like to work on this issue” if you would like to claim it.

If an issue is claimed, but work on the issue in not progressing, a Lilypad team member will ping the contributor to ask them if they want to continue work on the issue. If they do not wish to continue or we do not get a timely response, we will comment and open the issue up to another contributor.

### Triage a bug report

Once an issue has been opened, there will often be discussion about the issue internally and in the community. Some contributors may have differing opinions about the issue, including whether the behavior being seen is a bug or a feature. This is a part of the process and should be kept focused, helpful, and professional.

Contributors are encouraged to help one another make forward progress as much as possible, empowering one another to solve issues collaboratively. If you choose to comment on an issue that you feel either is not a problem that needs to be fixed, or if you encounter information in an issue that you feel is incorrect, explain why you feel that way with additional supporting context, and be willing to be convinced that you may be wrong.

### Resolve an issue

In the majority of cases, pull requests resolve issues. Pull requests require review and approval to ensure that proposed changes meet minimum quality and functional requirements to be accepted into the Lilypad code base.

## Pull requests

Pull requests are how changes are made to code and docs in the Lilypad Network.

Pull requests of any size are greatly appreciated. To stay aligned with our team, we request that contributors first open an issue describing proposed, non-trivial change to get feedback and a check whether the work is likely to be merged.

We encourage the community to review any PR on any repo in the [Lilypad GitHub organization](https://github.com/Lilypad-Tech)!

See our [local development guide](LOCAL_DEVELOPMENT.md) for instructions on running the Lilypad stack.

*Note:* any pull request created for an issue that already has someone else assigned **will be closed without review**.

### Commits

We require [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/) for all changes submitted in a pull request. Conventional commits help us to assess the impact of changes we merge into the code base.

Please use one of the following conventional commit types in your commit messages:

```
fix      - code changes linked to a known issue
feat     - new feature
refactor - code refactoring
perf     - performance improvements
hotfix   - code changes fixing code changes that introduced an issue
revert   - backwards remove of previous changes
chore    - changes not related to a fix or feature that don't modify source code
docs     - documentation only changes
test     - add missing tests or correct existing tests
build    - changes that affect the build system or external dependencies
```

### Signed commits

We require commits to be signed before they can be merged to `main`. Signing commits proves their authenticity.

#### Create a new SSH key pair

Configuring signed commits has gotten quite a bit easier now that GitHub allows you to sign with SSH keys.

```yaml
ssh-keygen -t ed25519 -C "<yourname>@example.com"
```

The last part is a comment. You could add something like “git signing key” if you want extra information to identify the key’s purpose.

You’ll be asked where you want to store the key and if you want to protect it with a passphrase. Adding a passphrase is more secure, but you will need to enter it each time you commit or set up an agent to help with that.

You will also need to configure `git` to use the signing key. See this [guide](https://vanjacosic.com/posts/sign-your-git-commits-using-ssh-keys/#prepare-git-for-signing) for `git` configuration options.

#### Upload public key to GitHub

After creating the key pair, upload the public key to GitHub. See this [guide](https://docs.github.com/en/authentication/connecting-to-github-with-ssh/adding-a-new-ssh-key-to-your-github-account#adding-a-new-ssh-key-to-your-account) and make sure to select “Signing Key” for the key type.

#### Verify configuration

This [guide](https://vanjacosic.com/posts/sign-your-git-commits-using-ssh-keys/#testing) on signing your git commits using SSH keys shows how to test signed commits locally. If you push the commits to GitHub, you will see a “Verified” badge next to your commit.


### Contributor license agreement

We require all code contributors to sign a contributor license agreement (CLA) that grants license on contributions to Lilypad and Lilypad users. Agreement to the contributor license agreement does not restrict usage of contributions outside of this repository. See [CLA.md](CLA.md) for more details.

To sign the CLA, open a pull request that adds a license accpetance entry to [NOTICES.md](NOTICES.md) and adds your GitHub username to the contributors list in [.clabot](.clabot). All commits on the pull request must be signed as described in [Signed commits](CONTRIBUTING.md#signed-commits) section of this document.

We will review the pull request and add you as a contributor. Agreement to the CLA is required before we merge other contributions.

### Opening the Pull Request

Opening a pull request on GitHub will present you with a pull request template. Please do your best to fill out the details, but feel free to skip parts if you're not sure how to answer.

Pull requests should document how to test new features or bug fixes. If we are unable to test the pull request, we may request additional context.

Pull request titles should follow the conventional commit format. We automatically create releases based on the pull request title. Pull requests with `feat:`, `fix:`, or `refactor:` titles will create a release.

### Tests

When a Pull Request is submitted, integration tests will run. The pull request must pass integration tests before we will merge it. We may merge some pull requests, such as changes to the docs without the integration test requirement.

### Commit Squashing

In most cases, do not squash commits while your pull request is under review. We use a squash and merge strategy when merging pull requests, and we may edit the individual commit messages squashed into the final commit message. The commit history of your initial pull request will remain intact on the pull request page.

### Reviewing contributions

PR review from members of the community are greatly encouraged and will assist the team in the final assessment of the PR. Reviews and feedback must be helpful, insightful, and geared towards improving the contribution as opposed to simply blocking it. If there are reasons why you feel the PR should not land, explain what those are.

Reviews that are dismissive or disrespectful of the contributor or any other reviewers are strictly counter to the Code of Conduct.

### Abandoned or stalled pull requests

If a Pull Request appears to be abandoned or has not seen any activity, please first check with the contributor to see if they intend to continue the work before asking if they would like you to take over the work. In this scenario, it is courteous to give the original contributor credit for work that has been already completed (use the `Author:`  or `Co-authored-by:` metadata tags in the commit to attribute with name and email)

## Rewards

We tag some issues with a `community` tag to invite the community to work on them. Work on these issues is eligible rewards once they have been resolved with an associated pull request.

Community issues can be claimed by commenting that you would like to work on the issue. Please be respectful of other community members if they have already claimed an issue. If an issue is claimed, but the work is not progressing, a Lilypad team member will comment to indicate the issue is open to claim again.

If you submit a new issue requesting a feature or reporting a bug, please indicate if you would like to work on the issue. At our discretion, we may determine the issue is eligible for rewards, and we will tag it with the `community` tag and ping you to let you know. The issue must contain a detailed description and an explanation of how it will benefit the project.

## License

The Lilypad Network is an open source project licensed under an [Apache 2.0 license](https://github.com/Lilypad-Tech/lilypad/blob/main/LICENSE).

## Attribution

This contribution guide was adapted from the [Tokio contributing guide](https://github.com/tokio-rs/tokio/blob/master/CONTRIBUTING.md)
