# tut

A command-line tool to streamline the management of multiple GitHub accounts.

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
5.  Add inside `$HOME/.tut/.ini` file, your account credentials jsut as the following:
```.ini
[account.a]
name = ZiadMansourM
email = personal@gmail.com
sshCommand = ssh -i ~/.ssh/ziadmansourm
[account.b]
name = ziadmmh
email = ziadh@work.com
sshCommand = ssh -i ~/.ssh/ziadmmh
```
> [!IMPORTANT]
> Make sure to prefex each section with `account.`

## Usage
1.  **List Accounts:** `tut list`.
2.  **Switch Accounts:** `tut`.

## Contribution Guidelines

**I welcome your contributions!  Here's how you can help:**

- `Open Issues`: Share your ideas, feature suggestions, or any bugs you discover.
- `Submit Pull Requests`: Feel free to directly propose code improvements or new features.

## License

This project is licensed under the MIT License – see the LICENSE file for details.
