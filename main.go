/*
Oh is a Unix shell written in Go. Like the rc shell, oh is similar in spirit
but different in detail from other Unix shells.

Oh extends the shell's programming language features without sacrificing the
shell's interactive features. The following commands behave as expected:

    date
    cat /usr/share/dict/words
    who >user.names
    who >>user.names
    wc <file
    echo [a-f]*.c
    who | wc
    who; date
    cc *.c &
    mkdir junk && cd junk
    cd ..
    rm -r junk || echo "rm failed!"

Oh has objects but no classes. Objects can be created from scratch using the
'object' command. Private members are defined using the 'define' command and
public members are defined using the 'public' command:

    define point: object {
        define x: integer 0
        define y: integer 0

        public move: method a b {
            set $self::x: add $self::x a
            set $self::y: add $self::y b
        }

        public show: method {
            echo $self::x $self::y
        }
    }

Objects can also be created by cloning an existing object:

    define o: point::clone

Modules are objects. The command below creates an object called 'm'. Public,
top-level definitions in 'file' can be accessed using the object 'm'.

    define m: import file

Pipes are objects. Oh exposes pipes, which are implicit in other shells, as
first-class values. Pipes can be created with the 'pipe' command:

    define p: pipe

Oh incorporates many features, including first-class functions, from the
Scheme dialect of Lisp. Like Lisp, oh uses the same syntax for code and data.
When data is sent across a channel it is converted to text so that it can be
sent to (or even through) external Unix programs.

To compile and run oh you will need to install (./configure; make install)
the C library libtecla:

    http://www.astro.caltech.edu/~mcs/tecla/libtecla.tar.gz

On Ubuntu you may also need to install libncurses5-dev.

Then go get oh:

    go get github.com/michaelmacinnis/oh

Oh is released under an MIT-style license. See LICENSE.
*/
package main

import (
	"github.com/michaelmacinnis/tecla"
	"os"
	"runtime"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	Start(len(os.Args) <= 1)

	if len(os.Args) <= 1 {
		Parse(tecla.New("> "), Evaluate)
	} else {
		f, err := os.OpenFile(os.Args[1], os.O_RDONLY, 0666)
		if err == nil {
			ParseFile(f, Evaluate)
		}
	}

	os.Exit(ExitStatus())
}
