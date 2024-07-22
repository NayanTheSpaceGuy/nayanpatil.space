# nayanpatil.space

NayanTheSpaceGuy's personal website.

[![License](https://img.shields.io/github/license/NayanTheSpaceGuy/nayanpatil.space)](https://mozilla.org/MPL/2.0/)

### Tech Stack

[![Go](https://img.shields.io/badge/Go-v1.22.5-blue.svg?logo=go&logoColor=white&labelColor=grey)](https://github.com/golang/go)
[![Echo](https://img.shields.io/badge/Echo-v4.11.3-blue.svg?logo=echo&logoColor=white&labelColor=grey)](https://github.com/labstack/echo)
[![Templ](https://img.shields.io/badge/Templ-v0.2-blue.svg?logo=templ&logoColor=white&labelColor=grey)](https://github.com/a-h/templ)
[![Htmx](https://img.shields.io/badge/Htmx-v2.0.1-blue.svg?logo=htmx&logoColor=white&labelColor=grey)](https://github.com/bigskysoftware/htmx)
[![Tailwind](https://img.shields.io/badge/Tailwind-v3.4-blue.svg?logo=tailwindcss&logoColor=white&labelColor=grey)](https://github.com/tailwindlabs/tailwindcss)

and more...

## Table of Contents
- [Introduction](#nayanpatilspace)
- [Tech Stack](#tech-stack)
- [Important Documents](#important-documents)
- [Building the source](#building-the-source)
- [Running](#running)
- [Miscellaneous](#miscellaneous)
- [Connect & Support](#connect--support)

## Important Documents

Please refer to the following documents for additional information:
- [LEGAL.md](LEGAL.md): Detailed licensing information and copyright notices for the project and its components.
- [CREDITS.md](CREDITS.md): Comprehensive list of contributors, third-party resources, and acknowledgments for the project.

Contributors and users are encouraged to review these documents thoroughly.

## Building the source

### Dependencies

[Arch-based](https://archlinux.org/packages/extra/x86_64/go/) (Manjaro / EndeavourOS / ... ):

```shell
pacman -S go
```

### Installation

```shell
git clone https://github.com/NayanTheSpaceGuy/nayanpatil.space.git
cd nayanpatil.space
```

### Building

```shell
make build
```

## Running

```shell
make run
```
This uses the default port 42069. To specify a custom port:
```shell
SERVER_PORT=<VALUE> make run
```

For live reloading during development:
```shell
make watch
```
Or with a custom port:
```shell
SERVER_PORT=<VALUE> make watch
```

Options:
- `SERVER_PORT` — Custom port number (default: 42069)

## Miscellaneous

### Other makefile usage instructions

Run all make commands with clean tests
```shell
make all build
```

Run the test suite
```shell
make test
```

Clean up binary from the last build
```shell
make clean
```

## Connect & Support

[![GitHub](https://img.shields.io/badge/GitHub-NayanTheSpaceGuy-181717?style=for-the-badge&logo=github)](https://github.com/NayanTheSpaceGuy)

### Star this repo if you find it useful! 🌟

[![License](https://img.shields.io/github/license/NayanTheSpaceGuy/nayanpatil.space)](https://mozilla.org/MPL/2.0/)
![Last Commit](https://img.shields.io/github/last-commit/NayanTheSpaceGuy/nayanpatil.space)
![Repo Size](https://img.shields.io/github/repo-size/NayanTheSpaceGuy/nayanpatil.space)
