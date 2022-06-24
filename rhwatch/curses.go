//This is in a seprate file because currently gopls does not work well with
// cgo
package main

//#cgo LDFLAGS: -lcurses
//#include <curses.h>
//#include <stdlib.h>
//#include <stdio.h>
//void curses_write(const char *str){
//printw("%s", str);
//}
import "C"
import (
	"fmt"
	"unsafe"
)

func cursesInit() {
	C.initscr()
}

func cursesWrite(str string) {
	cstr := C.CString(str)
	C.curses_write(cstr)
	C.free(unsafe.Pointer(cstr))
	C.refresh()
}

func cursesWritef(format string, args ...interface{}) {
	cursesWrite(fmt.Sprintf(format, args...))
}

func cursesClear() {
	C.clear()
	C.refresh()
}

func cursesClean() {
	C.endwin()
}

func cursesGetChar() rune {
	return rune(C.getch())
}
