# Dotsync

[![GoDoc](https://godoc.org/github.com/mbaraa/dotsync?status.png)](https://godoc.org/github.com/mbaraa/dotsync)
[![build](https://github.com/mbaraa/dotsync/actions/workflows/build.yml/badge.svg)](https://github.com/mbaraa/dotsync/actions/workflows/build.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/mbaraa/dotsync)](https://goreportcard.com/report/github.com/mbaraa/dotsync)

A small, free, open-source, blazingly fast dotfiles synchronizer!

Dotsync's [server](https://github.com/mbaraa/dotsync_server) is the middleware between your connected computers, where all of your dotfiles stand there encrypted, and backed up!

## Features:

- Efficient
- Lightweight
- Blazingly Fast
- Cool Stack
- Open-source
- Free (of charge & evil data telemetry things)
- Self-hosting option, check [the server](https://github.com/mbaraa/dotsync_server)

## Dependencies:

*   [go](https://golang.org)
*   An internet connection
*   Linux or Unix-like system (I haven't tried it on Windows, a feedback is more than welcome)
*   A bunch of dotfiles to sync :)

## Installation:
<!---
### Using a Package Manager:

#### Gentoo

1. Add my [overlay](https://github.com/mbaraa/mbaraa-overlay)
2. Install `net-misc/dotsync` using emerge

#### Arch (AUR)

install [dotsync2](https://aur.archlinux.org/packages/dotsync2) -- I realized later on that an AUR package exists...
-->
### Using Go's installer

```bash
go install github.com/mbaraa/dotsync@latest

# or

make
sudo make install
```

## Usage:

### Create/Login using an email

And as mentioned above your email is encrypted, and won't be shared with anyone!

1. Login
```bash
dotsync -login someone@example.com
```
2. Enter the token which you recived as an email(it might arrive as a spam email)

3. Go nuts

### Sync

1. Add and upload a bunch of files

```bash
# add a file
dotsync -add ~/.bashrc
dotsync -add ~/.config/i3/config
dotsync -add ~/.config/nvim/

# upload your current files
dotsync -upload
```

2. Download your files on another computer to show the power of Dotsync
```bash
# you need to login first, so...
dotsync -download
```

**For a more detailed usage, visit the [Officical Docs](https://dotsync.org/docs)**!
