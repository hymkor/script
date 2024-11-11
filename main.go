package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"runtime"

	"golang.org/x/term"

	"github.com/aymanbagabas/go-pty"

	"github.com/hymkor/go-windows1x-virtualterminal"
)

var (
	flagAppend  = flag.Bool("a", false, "append")
	flagCommand = flag.String("c", "cmd.exe", "execute command instead of interactive shell")
)

func mains(args []string) error {
	disableStdout, err := virtualterminal.EnableStdout()
	if err != nil {
		return err
	}
	defer disableStdout()

	disableStdin, err := virtualterminal.EnableStdin()
	if err != nil {
		return err
	}
	defer disableStdin()

	fn := "typescript"
	if len(args) > 0 {
		fn, args = args[0], args[1:]
	}
	var typeScript *os.File
	if *flagAppend {
		typeScript, err = os.OpenFile(fn, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	} else {
		typeScript, err = os.Create(fn)
	}
	if err != nil {
		return err
	}
	defer typeScript.Close()

	ptmx, err := pty.New()
	if err != nil {
		return err
	}
	defer ptmx.Close()

	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return err
	}
	ptmx.Resize(width, height)

	fields := splitField(*flagCommand)
	sh := ptmx.Command(fields[0], fields[1:]...)
	if err := sh.Start(); err != nil {
		return err
	}

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return err
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	go io.Copy(ptmx, os.Stdin)
	go io.Copy(io.MultiWriter(os.Stdout, typeScript), ptmx)

	return sh.Wait()
}

var version = "snapshot"

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s %s-%s-%s:\n",
			os.Args[0], version, runtime.GOOS, runtime.GOARCH)
		flag.PrintDefaults()
	}
	flag.Parse()
	if err := mains(flag.Args()); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
