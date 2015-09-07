package main

import "github.com/yuin/gopher-lua"
import "github.com/nsf/termbox-go"
import "os"
import "flag"
import "path"
import "log"
import "strings"

/* This has to be a global - we write to it from a function called from Lua.
 * See loop.go, function qlobal_quit() */
var quit chan bool

func main() {
	events := make(chan lua.LValue, 256)
	quit = make(chan bool)

	/* either open the file given on the command line or main.lua */
	cfilename := "main.lua"
	if flag.Parse(); flag.Arg(0) != "" {
		cfilename = flag.Arg(0)
	}
	/* find any backslash separators and turn them into forward slashes */
	cfilename = strings.Replace(cfilename, "\\", "/", -1)

	/* process the filename - if it's in a directory, we'll need to
	 * chdir to it (so we can have relative imports) */
	directory := path.Dir(cfilename)
	filename := path.Base(cfilename)
	os.Chdir(directory)

	/* bare-minimum sanity check */
	if _, err := os.Stat(filename); err != nil {
		log.Fatal(err)
	}

	/* start up termbox
	 * see https://github.com/nsf/termbox-go/issues/80 */
	os.Setenv("TERM", "xterm")
	if err := termbox.Init(); err != nil {
		log.Fatal(err)
	}
	defer termbox_close()

	/* start the goroutines and wait */
	go termbox_listener(events)
	go timer(events)
	go mainloop(filename, events)
	<-quit
}
