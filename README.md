script.exe
==========

- Make typescript of terminal session like [that of Linux](https://www.man7.org/linux/man-pages/man1/script.1.html) on Windows10 or later
- Implemented with [aymanbagabas/go-pty](https://github.com/aymanbagabas/go-pty)

![demo](./demo.gif)

Install
-------

Download the binary package from [Releases](https://github.com/hymkor/script/releases) and extract the executable.

### go install

```
go install github.com/hymkor/script
```

### scoop-installer

```
scoop install https://raw.githubusercontent.com/hymkor/script/master/script.json
```

or

```
scoop bucket add hymkor https://github.com/hymkor/scoop-bucket
scoop install script
```

Usage
-----

```
script [options] [file]
```

- `-a` Append the output to specified file or **"typescript"**
- `-c command` Run the command rather than cmd.exe
