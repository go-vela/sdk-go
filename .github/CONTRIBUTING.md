# Contributing

## Getting Started

We'd love to accept your contributions to this project! If you are a first time contributor, please review our [Contributing Guidelines](https://go-vela.github.io/docs/community/contributing_guidelines/) before proceeding.

### Prerequisites

* [Review the commit guide we follow](https://chris.beams.io/posts/git-commit/#seven-rules) - ensure your commits follow our standards
* Review our [style guide](https://go-vela.github.io/docs/community/contributing_guidelines/#style-guide) to ensure your code is clean and consistent.

### Setup

* [Fork](/fork) this repository

* Clone this repository to your workstation:

```bash
# Clone the project
git clone git@github.com:go-vela/sdk-go.git $HOME/go-vela/sdk-go
```

* Navigate to the repository code:

```bash
# Change into the project directory
cd $HOME/go-vela/sdk-go
```

* Point the original code at your fork:

```bash
# Add a remote branch pointing to your fork
git remote add fork https://github.com/your_fork/sdk-go
```

### Development

* Navigate to the repository code:

```bash
# Change into the project directory
cd $HOME/go-vela/sdk-go
```

* Write your code and tests to implement the changes you desire.

* Run the repository tests (ensures your changes perform as you desire):

```bash
# Test the code with `go`
go test ./...
```

* Ensure your code meets the project standards:

```bash
# Clean the code with `go`
go mod tidy
go fmt ./...
go vet ./...
```

* Push to your fork:

```bash
# Push your code up to your fork
git push fork main
```
* Make sure to follow our [PR process](https://go-vela.github.io/docs/community/contributing_guidelines/#development-workflow) when opening a pull request

Thank you for your contribution!
