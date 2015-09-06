package main

import "github.com/yuin/gopher-lua"
import "github.com/nsf/termbox-go"
import "os"
import "flag"
import "path"

/* This has to be a global - we write to it from a function called from Lua.
 * See loop.go, function qlobal_quit() */
var quit chan bool

func main() {
	events := make(chan lua.LValue, 256)
	quit = make(chan bool)

	/* start up termbox
	 * see https://github.com/nsf/termbox-go/issues/80 */
	os.Setenv("TERM", "xterm")
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox_close()

	/* either open the file given on the command line or main.lua */
	cfilename := "main.lua"
	if flag.Parse(); flag.Arg(0) != "" {
		cfilename = flag.Arg(0)
	}

	/* process the filename - if it's in a directory, we'll need to
	 * chdir to it (so we can have relative imports) */
	directory := path.Dir(cfilename)
	filename := path.Base(cfilename)
	os.Chdir(directory)

	/* bare-minimum sanity check */
	if _, err := os.Stat(filename); err != nil {
		panic(err)
	}

	/* start the goroutines and wait */
	go termbox_listener(events)
	go timer(events)
	go mainloop(filename, events)
	<-quit
}
