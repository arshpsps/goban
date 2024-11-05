### goban

> Simple kanban board app written in Go.

This app uses [bubbletea](https://github.com/charmbracelet/bubbletea) for TUI. (bubbletea is sick (and annoying))

#### Usage

##### 1. GNU Make

`make` to run without building

`make build` to compile in builds/out

##### 2. Docker

> Docker tty has some trouble, input doesn't go through.
> A GitHub Issue regarding how to fix it would be much appreciated.

`docker build --tag 'goban' .`

`docker run -t 'goban'`

---

~Might wanna look into [cobra](https://pkg.go.dev/github.com/spf13/cobra)~

A cobra based approach wouldn't really work unless the app runs as a service or the like.
