# Cross-platform `Exec` for Go

⚡️ [`syscall.Exec`](https://pkg.go.dev/syscall#Exec) that works on Windows

<table align=center><td>

```go
p, err := exec.LookPath("go")
if err != nil {
    panic(err)
}
log.Fatal(crossexec.Exec(p, os.Args, os.Environ()))
```

</table>

## Installation

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=Go&logoColor=FFFFFF)

```sh
go get go.jcbhmr.com/crossexec
```

## Usage

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=Go&logoColor=FFFFFF)

This module exports only a single function: **[`crossexec.Exec`](https://pkg.go.dev/go.jcbhmr.com/crossexec#Exec)**. It takes in the same arguments and has the same behaviour as [`syscall.Exec`](https://pkg.go.dev/syscall#Exec) (available on `unix || plan9`; returns `EWINDOWS` on `windows`) except `crossexec.Exec` **works on Windows too** using a normal subprocess instead of replacing the current process.

[📚 See the docs](https://pkg.go.dev/go.jcbhmr.com/crossexec#Exec)

## Development

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=Go&logoColor=FFFFFF)

Inspired by [`exec_replace`](https://docs.rs/cargo-util/latest/cargo_util/struct.ProcessBuilder.html#method.exec_replace) from [`cargo-util`](https://docs.rs/cargo-util/latest/cargo_util/).
