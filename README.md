# Cross-platform `Exec` for Go

⚡️ \*nix [`CmdExt.Exec`](https://pkg.go.dev/github.com/jcbhmr/go-exec#CmdExt.Exec) but works on Windows

<table align=center><td>

```go
path, err := exec.LookPath("go")
if err != nil {
    panic(err)
}
log.Fatal(crossexec.CrossExec(path, os.Args, os.Environ()))
```

</table>

## Installation

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=Go&logoColor=FFFFFF)

```sh
go get go.jcbhmr.com/crossexec
```

## Usage

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=Go&logoColor=FFFFFF)

This module exports only a single function: [`crossexec.Exec`](https://pkg.go.dev/go.jcbhmr.com/crossexec#CrossExec). It takes in the same arguments and has the same behaviour as [`syscall.Exec`](https://pkg.go.dev/syscall#Exec) (available on `unix || plan9`; returns `EWINDOWS` on `windows`) except `crossexec.Exec` **works on Windows too**.

[📚 See the docs](https://pkg.go.dev/go.jcbhmr.com/crossexec#CrossExec)

## Development

![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=Go&logoColor=FFFFFF)

Inspired by [`exec_replace`](https://docs.rs/cargo-util/latest/cargo_util/struct.ProcessBuilder.html#method.exec_replace) from [`cargo-util`](https://docs.rs/cargo-util/latest/cargo_util/).
