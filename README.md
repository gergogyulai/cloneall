# 🚀 CloneAll

CloneAll is a CLI tool written in Go that clones all repositories of a specified GitHub organization or user. It also generates a Markdown file listing all cloned repositories with their GitHub links and descriptions.

## ✨ Features

- 🗂️ Clones all repositories from a given GitHub user or organization.
- 📁 Organizes the cloned repositories into a folder named after the user or organization.
- 📝 Generates a `README.md` file containing a list of all cloned repositories, including their GitHub links and descriptions.

## 🛠️ Installation

### 1. Download Precompiled Binary

1. Go to the [Releases](https://github.com/gergogyulai/cloneall/releases) page.
2. Download the precompiled binary for your operating system.
3. Place the binary in a directory included in your PATH for easy access or just execute it where your heart desires.

### 2. Build from Source

1. Ensure you have Go installed on your machine. If not, download and install it from the [official Go website](https://golang.org/dl/).

2. Clone the repository:
   ```sh
   git clone https://github.com/gergogyulai/cloneall.git
   cd cloneall
   ```

3. Build the program by running:
   ```sh
   go build -o cloneall cloneall.go
   ```

4. Ensure you have `git` installed and available in your PATH.

## 🚀 Usage

To clone all repositories of a GitHub user or organization and generate a Markdown file, run:
```sh
cloneall <GitHub org or user URL>
```

Replace `<GitHub org or user URL>` with the actual URL of the GitHub user or organization. For example:
```sh
cloneall https://github.com/github
```

## 🌟 Example

After running the tool with the GitHub organization or user URL, a folder named after the user or organization will be created. Inside this folder, you will find all the cloned repositories and a `README.md` file.

The `README.md` file will look like this:

```markdown
# Repositories of github

## [repo1](https://github.com/github/repo1)

Description of repo1.

## [repo2](https://github.com/github/repo2)

Description of repo2.
```

## 📜 License 

This project is licensed under the the best license in the world, the [WTFPL License](https://en.wikipedia.org/wiki/WTFPL). See the [LICENSE](LICENSE) file for details.

## 🤝 Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.
