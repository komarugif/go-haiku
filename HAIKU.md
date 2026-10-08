# Go for Haiku

Go 1.27.1 for `haiku/amd64`, kept for [KomaruGram Go](https://github.com/komarugif/komarugram-go),
a Telegram client that runs on Haiku R1/beta6. The branch is
`golang-1.27-haiku`:

1. the tag `go1.27.1` of Go;
2. the Haiku port of Jérôme Duval, [korli/go](https://github.com/korli/go),
   branch `golang-1.26-haiku` at `244072a4`, as its difference from the
   `release-branch.go1.26` it last merged (`ee49ded5`), applied onto it,
   with seven conflicts resolved in go1.27's form;
3. fixes found running a real program, a Gio window with many threads,
   cgo, SQLite in wasm and a network client, each a commit of its own:
   - the runtime's `_SS_DISABLE` is Haiku's 2, not Solaris's 4: with cgo,
     signal handlers ran on goroutines' stacks, and programs died at random
     (`unexpected return pc`, `morestack on gsignal`);
   - the process ends with `_exit`, as on Solaris: `exit` ran every
     library's C++ destructors while other threads ran, and programs
     crashed on exit;
   - go1.27's `libinit.go` and shared `gcc_unix.c` taken in;
   - `syscall.sysvicall6` marked for `golang.org/x/sys/unix`, which
     failed to link without it;
   - sockets made nonblocking with `fcntl`: on Haiku a socket made with
     `SOCK_NONBLOCK` in `socket()` still blocks in `accept`, and closing a
     listener waited forever;
   - `AT_FDCWD` is `-100`, as `fcntl.h` has it.

## Building

From Linux (or any system with Go 1.24 or later):

```sh
cd src && ./make.bash
```

Then build for Haiku with `GOOS=haiku GOARCH=amd64`. cgo works when asked
for (`CGO_ENABLED=1`, with a C compiler for Haiku such as clang
`--target=x86_64-unknown-haiku` and a sysroot copied from Haiku), with Go's
internal linker only, `-ldflags=-linkmode=internal`: Haiku links programs
as shared objects, and the external linker fails on Go's local-exec TLS.

How KomaruGram is built with it, and what was checked on Haiku, is in
KomaruGram's `docs/PLATFORMS.md`, "Haiku".

`golang.org/x/sys/unix` has Haiku's files only in this toolchain's own
copy, `src/cmd/vendor/golang.org/x/sys/unix`; a program using `x/sys`
needs a copy of it with them.
