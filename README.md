# tut

A command-line tool to ease the management of multiple GitHub accounts from the same workstation.

## Features

*   **Effortless Account Switching:** Quickly switch between different GitHub accounts within your local Git configuration.
*   **Simple Configuration:** Intuitive configuration management for your accounts.
*   **Cross-Platform Support:** Works seamlessly on macOS, Linux, and Windows.

## Installation

**Download Pre-built Binaries**

1.  Visit the latest release page: [https://github.com/ZiadMansourM/tut/releases](https://github.com/ZiadMansourM/tut/releases)
2.  Download the appropriate `.tgz` archive for your operating system and architecture (e.g., `tut_linux_amd64.tgz`).
3.  Extract the archive: `tar -xzvf tut_linux_amd64.tgz`
4.  Move the `tut` binary into a directory in your PATH (e.g., `/usr/local/bin`).
5.  Add inside `$HOME/.tut/.ini` file, your account credentials just as the following:
```.ini
[account.a]
name = ZiadMansourM
email = personal@gmail.com
sshCommand = ssh -i ~/.ssh/ziadmansourm
[account.b]
name = ziadmmh
email = ziadh@work.com
```
> [!IMPORTANT]
> Make sure to prefex each section with `account.`

6. Edit your `$HOME/.gitconfig` to include this:
```.ini
[include]
path = ~/.tut/.gitconfig
```

## Usage
If you use one workstation to contribute to projects for more than one account on GitHub.com, you can modify your Git configuration to simplify the contribution process. There are many ways you can achieve that. [Read here for more details](https://docs.github.com/en/account-and-profile/setting-up-and-managing-your-personal-account-on-github/managing-your-personal-account/managing-multiple-accounts).

> [!WARNING]
> Please, be aware that `tut` supports only using two differnet protocols
> for each account. E.g. use ssh for your personal account and https for work account.

The protocol you use to clone a repository determines which credentials your workstation will use to authenticate when you access the repository. That is exacly why we choose this approach.

1. Configure your personal account with ssh access.
2. Configure your work account with https access.
3. Create the `$HOME/.tut/.ini` file as described above.
4. Edit `$HOME/.gitconfig` as described above.
5. Clone a private repo from both of your accounts, you should have no problem on doing that.
6. Before you make any changes run `tut` and choose correct account. Other than that the wrong github username and email will appear in the commit messages.

## Contribution Guidelines

**I welcome your contributions!  Here's how you can help:**

- `Open Issues`: Share your ideas, feature suggestions, or any bugs you discover.
- `Submit Pull Requests`: Feel free to directly propose code improvements or new features.

## License

This project is licensed under the MIT License – see the LICENSE file for details.
