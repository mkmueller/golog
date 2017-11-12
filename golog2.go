// Copyright 2017 Mark K Mueller. (github.com/mkmueller)
// This package is a fork of golang/src/log/

/*
The golog package is a fork of Go's log package. It is functionally identical
to Go's log package which makes it a drop-in replacement with a few added
features. The New method is variadic so it can be called using zero to four
arguments; the optional fourth argument will set the logging level. The New
method accepts a filename as the first argument. Also, the Printf method
accepts a logging level parameter.
*/
package golog

import (
	"io"
	"os"
	"fmt"
	"time"
	"sync"
	"bytes"
	"strings"
	"runtime"
)

type Logger struct {
	mu     		sync.Mutex 			// ensures atomic writes; protects the following fields
	prefix 		string     			// prefix to write at beginning of each line
	flag   		int        			// properties
	out    		io.Writer  			// destination for output
	buf    		[]byte     			// for accumulating text to write
    isfile      bool                // true if file was opened successfully
	level		int					// The level of messages that will be output using Printf
	Err			error				// Set if an error was encountered during file opening
    Filename    string              // Output filename
    FileMode    os.FileMode         // Permissions of output file. Default is 0660
    Options     int                 // Output file flags. Default is os.O_WRONLY|os.O_CREATE|os.O_APPEND
}

/*
New creates a new Logger.  Similar in usage to the New function in Go's log
package, it may be called in a similar way using out, prefix and flag arguments.
However, you may also call it using no arguments at all in which case the
default output will be stderr, no prefix, and the flag set to 0. Being a
variadic function, New will allow one, two, three or four arguments. The first
argument may be a string filename or an io.Writer like *os.File or
*bytes.Buffer.  An optional fourth argument will set the logging level for the
Printf method.

The following argument combinations may be used with the New method:

  file   := "logfile.log"
  out    := os.Stdout
  prefix := "PREFIX "
  flags  := golog.LstdFlags|golog.LUTC
  level  := 2

  golog.New()
  golog.New(flags)
  golog.New(flags, prefix)
  golog.New(out, prefix, flags)
  golog.New(file)
  golog.New(file, flags)
  golog.New(file, prefix, flags)
  golog.New(file, prefix, flags, level)
*/
func New (v ...interface{}) *Logger {
	l 			:= new(Logger)
	l.mu.Lock()
	defer l.mu.Unlock()
	l.out		= os.Stderr
	l.level		= -1
    l.Options 	= os.O_WRONLY|os.O_CREATE|os.O_APPEND
	l.FileMode	= 0660

	if len(v) > 0 {
		switch v[0].(type) {
			case string:
				l.Filename = v[0].(string)
				l.flag = LstdFlags
				l.openFile()
				l.closeFile()
			case *os.File:
				l.out = v[0].(io.Writer)
			case *bytes.Buffer:
				l.out = v[0].(io.Writer)
			case int:
				l.flag = v[0].(int)
			default:
				panic("expecting a filename or an io.Writer in the first argument")
		}
		if len(v) > 1 {
			switch v[1].(type) {
				case string:
					l.prefix = v[1].(string)
				case int:
					l.flag = v[1].(int)
				default:
					panic("expecting prefix or flags in the second argument")
			}
		}
		if len(v) > 2 {
			switch v[2].(type) {
				case int:
					l.flag = v[2].(int)
				default:
					panic("expecting flags in the third argument")
			}
		}
		if len(v) > 3 {
			switch v[3].(type) {
				case int:
					l.level = v[3].(int)
				default:
					panic("expecting level in the fourth argument")
			}
		}
	}
	return l
}

// SetOutput sets the output destination for the logger.  The argument may be an
// io.Writer or a filename string. If a string is supplied, the file is opened then
// immediately closed which will create an empty file if it does not already exist.
// If an error is encountered during file opening, the output will be set to stderr
// and the logger Err value will be set.
//
func (l *Logger) SetOutput(v interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	switch v.(type) {
		case string:
			l.Filename = v.(string)
			l.openFile()
			l.closeFile()
		case *os.File:
			l.out = v.(io.Writer)
		case *bytes.Buffer:
			l.out = v.(io.Writer)
		default:
			panic("expecting a filename or an io.Writer in the first argument")
	}
}

// Printf calls l.Output to print to the logger. Arguments are handled in the
// manner of fmt.Printf. If the format string "%LVL" is used in combination with
// an integer as the first value, the output level will be checked.  Output will
// proceed only if the supplied level is less than or equal to the configured
// logging level.
//
func (l *Logger) Printf (format string, v ...interface{}) {
	if !l.checkLevel(&format, v...) { return }
	l.Output(2, fmt.Sprintf(format, v...))
}

// SetLevel will set the output logging level for the Printf method. The default
// logging level is -1 causing all levels to be output. Setting the level to any
// value greater than -1 will allow all levels up to and including that level to
// be output. For example, setting the level to 9 will allow Printf levels 0 - 9
// to be output.
//
func  (l *Logger) SetLevel (n int) {
	l.level = n
}


// Output writes the output for a logging event. The string s contains the text to
// print after the prefix specified by the flags of the Logger. A newline is
// appended if the last character of s is not already a newline. Calldepth is used
// to recover the PC and is provided for generality, although at the moment on all
// pre-defined paths it will be 2. Output will be written to a file if a filename
// was supplied in the New or SetOutput methods. If the file cannot be opened,
// output will be sent to stderr instead and logger.Err value will be set.
//
func (l *Logger) Output(calldepth int, s string) error {
	// Get time early if we need it.
	var now time.Time
	if l.flag&(Ldate|Ltime|Lmicroseconds) != 0 {
		now = time.Now()
	}
	var file string
	var line int
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.flag&(Lshortfile|Llongfile) != 0 {
		// Release lock while getting caller info - it's expensive.
		l.mu.Unlock()
		var ok bool
		_, file, line, ok = runtime.Caller(calldepth)
		if !ok {
			file = "???"
			line = 0
		}
		l.mu.Lock()
	}
	l.buf = l.buf[:0]
	l.formatHeader(&l.buf, now, file, line)
	l.buf = append(l.buf, s...)
	if len(s) == 0 || s[len(s)-1] != '\n' {
		l.buf = append(l.buf, '\n')
	}
	l.openFile()
	_, err := l.out.Write(l.buf)
	l.closeFile()
	return err
}

// Open a file for append. If the file cannot be opened, logger.out will be
// set to os.Stderr and logger.Err will be set.
func (l *Logger) openFile () {
	if l.Filename == "" { return }
	l.out, l.Err = os.OpenFile(l.Filename, l.Options, l.FileMode)
	if l.Err == nil {
		l.isfile = true
	} else {
		// Cannot open the output file. Using stderr instead.
		l.out = os.Stderr
		l.isfile = false
	}
}

// Close the file if it was opened.
func (l *Logger) closeFile () {
	if l.isfile {
		l.out.(*os.File).Close()
	}
}

// Check output level. Return true if it's okay to proceed, false if not.
// Will also replace %LVL with %v.
func (l *Logger) checkLevel (formatptr *string, v ...interface{}) bool {
	if strings.Index(*formatptr, "%LVL") > -1 && len(v) > 0 {
		*formatptr = strings.Replace(*formatptr, "%LVL", "%v", 1)
		switch v[0].(type) { case int:
			if l.level != -1 && l.level < v[0].(int) {
				return false
			}
		}
	}
	return true
}

