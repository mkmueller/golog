

# golog
`import "github.com/mkmueller/golog"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Examples](#pkg-examples)

## <a name="pkg-overview">Overview</a>
The golog package is a fork of Go's log package. It is functionally identical
to Go's log package which makes it a drop-in replacement with a few added
features. The New method is variadic so it can be called using zero to four
arguments; the optional fourth argument will set the logging level. The New
method accepts a filename as the first argument. Also, the Printf method
accepts a logging level parameter.




## <a name="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [func Fatal(v ...interface{})](#Fatal)
* [func Fatalf(format string, v ...interface{})](#Fatalf)
* [func Fatalln(v ...interface{})](#Fatalln)
* [func Flags() int](#Flags)
* [func Output(calldepth int, s string) error](#Output)
* [func Panic(v ...interface{})](#Panic)
* [func Panicf(format string, v ...interface{})](#Panicf)
* [func Panicln(v ...interface{})](#Panicln)
* [func Prefix() string](#Prefix)
* [func Print(v ...interface{})](#Print)
* [func Printf(format string, v ...interface{})](#Printf)
* [func Println(v ...interface{})](#Println)
* [func SetFlags(flag int)](#SetFlags)
* [func SetOutput(w io.Writer)](#SetOutput)
* [func SetPrefix(prefix string)](#SetPrefix)
* [type Logger](#Logger)
  * [func New(v ...interface{}) *Logger](#New)
  * [func (l *Logger) Fatal(v ...interface{})](#Logger.Fatal)
  * [func (l *Logger) Fatalf(format string, v ...interface{})](#Logger.Fatalf)
  * [func (l *Logger) Fatalln(v ...interface{})](#Logger.Fatalln)
  * [func (l *Logger) Flags() int](#Logger.Flags)
  * [func (l *Logger) Output(calldepth int, s string) error](#Logger.Output)
  * [func (l *Logger) Panic(v ...interface{})](#Logger.Panic)
  * [func (l *Logger) Panicf(format string, v ...interface{})](#Logger.Panicf)
  * [func (l *Logger) Panicln(v ...interface{})](#Logger.Panicln)
  * [func (l *Logger) Prefix() string](#Logger.Prefix)
  * [func (l *Logger) Print(v ...interface{})](#Logger.Print)
  * [func (l *Logger) Printf(format string, v ...interface{})](#Logger.Printf)
  * [func (l *Logger) Println(v ...interface{})](#Logger.Println)
  * [func (l *Logger) SetFlags(flag int)](#Logger.SetFlags)
  * [func (l *Logger) SetLevel(n int)](#Logger.SetLevel)
  * [func (l *Logger) SetOutput(v interface{})](#Logger.SetOutput)
  * [func (l *Logger) SetPrefix(prefix string)](#Logger.SetPrefix)

#### <a name="pkg-examples">Examples</a>
* [Logger](#example_Logger)
* [Logger.Output](#example_Logger_Output)
* [Logger.Printf](#example_Logger_Printf)
* [New](#example_New)

#### <a name="pkg-files">Package files</a>
[golog.go](/src/github.com/mkmueller/golog/golog.go) [golog2.go](/src/github.com/mkmueller/golog/golog2.go) 


## <a name="pkg-constants">Constants</a>
``` go
const (
    // Bits or'ed together to control what's printed.
    // There is no control over the order they appear (the order listed
    // here) or the format they present (as described in the comments).
    // The prefix is followed by a colon only when Llongfile or Lshortfile
    // is specified.
    // For example, flags Ldate | Ltime (or LstdFlags) produce,
    //	2009/01/23 01:23:23 message
    // while flags Ldate | Ltime | Lmicroseconds | Llongfile produce,
    //	2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
    Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
    Ltime                         // the time in the local time zone: 01:23:23
    Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
    Llongfile                     // full file name and line number: /a/b/c/d.go:23
    Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
    LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
    LstdFlags     = Ldate | Ltime // initial values for the standard logger
)
```
These flags define which text to prefix to each log entry generated by the Logger.




## <a name="Fatal">func</a> [Fatal](/src/target/golog.go?s=9516:9544#L310)
``` go
func Fatal(v ...interface{})
```
Fatal is equivalent to Print() followed by a call to os.Exit(1).



## <a name="Fatalf">func</a> [Fatalf](/src/target/golog.go?s=9665:9709#L316)
``` go
func Fatalf(format string, v ...interface{})
```
Fatalf is equivalent to Printf() followed by a call to os.Exit(1).



## <a name="Fatalln">func</a> [Fatalln](/src/target/golog.go?s=9841:9871#L322)
``` go
func Fatalln(v ...interface{})
```
Fatalln is equivalent to Println() followed by a call to os.Exit(1).



## <a name="Flags">func</a> [Flags](/src/target/golog.go?s=8461:8477#L270)
``` go
func Flags() int
```
Flags returns the output flags for the standard logger.



## <a name="Output">func</a> [Output](/src/target/golog.go?s=10843:10885#L355)
``` go
func Output(calldepth int, s string) error
```
Output writes the output for a logging event. The string s contains
the text to print after the prefix specified by the flags of the
Logger. A newline is appended if the last character of s is not
already a newline. Calldepth is the count of the number of
frames to skip when computing the file name and line number
if Llongfile or Lshortfile is set; a value of 1 will print the details
for the caller of Output.



## <a name="Panic">func</a> [Panic](/src/target/golog.go?s=9989:10017#L328)
``` go
func Panic(v ...interface{})
```
Panic is equivalent to Print() followed by a call to panic().



## <a name="Panicf">func</a> [Panicf](/src/target/golog.go?s=10141:10185#L335)
``` go
func Panicf(format string, v ...interface{})
```
Panicf is equivalent to Printf() followed by a call to panic().



## <a name="Panicln">func</a> [Panicln](/src/target/golog.go?s=10320:10350#L342)
``` go
func Panicln(v ...interface{})
```
Panicln is equivalent to Println() followed by a call to panic().



## <a name="Prefix">func</a> [Prefix](/src/target/golog.go?s=8672:8692#L280)
``` go
func Prefix() string
```
Prefix returns the output prefix for the standard logger.



## <a name="Print">func</a> [Print](/src/target/golog.go?s=8996:9024#L293)
``` go
func Print(v ...interface{})
```
Print calls Output to print to the standard logger.
Arguments are handled in the manner of fmt.Print.



## <a name="Printf">func</a> [Printf](/src/target/golog.go?s=9173:9217#L299)
``` go
func Printf(format string, v ...interface{})
```
Printf calls Output to print to the standard logger.
Arguments are handled in the manner of fmt.Printf.



## <a name="Println">func</a> [Println](/src/target/golog.go?s=9377:9407#L305)
``` go
func Println(v ...interface{})
```
Println calls Output to print to the standard logger.
Arguments are handled in the manner of fmt.Println.



## <a name="SetFlags">func</a> [SetFlags](/src/target/golog.go?s=8562:8585#L275)
``` go
func SetFlags(flag int)
```
SetFlags sets the output flags for the standard logger.



## <a name="SetOutput">func</a> [SetOutput](/src/target/golog.go?s=8318:8345#L263)
``` go
func SetOutput(w io.Writer)
```
SetOutput sets the output destination for the standard logger.



## <a name="SetPrefix">func</a> [SetPrefix](/src/target/golog.go?s=8780:8809#L285)
``` go
func SetPrefix(prefix string)
```
SetPrefix sets the output prefix for the standard logger.




## <a name="Logger">type</a> [Logger](/src/target/golog2.go?s=638:1431#L25)
``` go
type Logger struct {
    Err      error       // Set if an error was encountered during file opening
    Filename string      // Output filename
    FileMode os.FileMode // Permissions of output file. Default is 0660
    Options  int         // Output file flags. Default is os.O_WRONLY|os.O_CREATE|os.O_APPEND
    // contains filtered or unexported fields
}
```






### <a name="New">func</a> [New](/src/target/golog2.go?s=2415:2450#L66)
``` go
func New(v ...interface{}) *Logger
```
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





### <a name="Logger.Fatal">func</a> (\*Logger) [Fatal](/src/target/golog.go?s=6723:6763#L196)
``` go
func (l *Logger) Fatal(v ...interface{})
```
Fatal is equivalent to l.Print() followed by a call to os.Exit(1).




### <a name="Logger.Fatalf">func</a> (\*Logger) [Fatalf](/src/target/golog.go?s=6884:6940#L202)
``` go
func (l *Logger) Fatalf(format string, v ...interface{})
```
Fatalf is equivalent to l.Printf() followed by a call to os.Exit(1).




### <a name="Logger.Fatalln">func</a> (\*Logger) [Fatalln](/src/target/golog.go?s=7072:7114#L208)
``` go
func (l *Logger) Fatalln(v ...interface{})
```
Fatalln is equivalent to l.Println() followed by a call to os.Exit(1).




### <a name="Logger.Flags">func</a> (\*Logger) [Flags](/src/target/golog.go?s=7736:7764#L235)
``` go
func (l *Logger) Flags() int
```
Flags returns the output flags for the logger.




### <a name="Logger.Output">func</a> (\*Logger) [Output](/src/target/golog2.go?s=5785:5839#L173)
``` go
func (l *Logger) Output(calldepth int, s string) error
```
Output writes the output for a logging event. The string s contains the text to
print after the prefix specified by the flags of the Logger. A newline is
appended if the last character of s is not already a newline. Calldepth is used
to recover the PC and is provided for generality, although at the moment on all
pre-defined paths it will be 2. Output will be written to a file if a filename
was supplied in the New or SetOutput methods. If the file cannot be opened,
output will be sent to stderr instead and logger.Err value will be set.




### <a name="Logger.Panic">func</a> (\*Logger) [Panic](/src/target/golog.go?s=7232:7272#L214)
``` go
func (l *Logger) Panic(v ...interface{})
```
Panic is equivalent to l.Print() followed by a call to panic().




### <a name="Logger.Panicf">func</a> (\*Logger) [Panicf](/src/target/golog.go?s=7396:7452#L221)
``` go
func (l *Logger) Panicf(format string, v ...interface{})
```
Panicf is equivalent to l.Printf() followed by a call to panic().




### <a name="Logger.Panicln">func</a> (\*Logger) [Panicln](/src/target/golog.go?s=7587:7629#L228)
``` go
func (l *Logger) Panicln(v ...interface{})
```
Panicln is equivalent to l.Println() followed by a call to panic().




### <a name="Logger.Prefix">func</a> (\*Logger) [Prefix](/src/target/golog.go?s=8011:8043#L249)
``` go
func (l *Logger) Prefix() string
```
Prefix returns the output prefix for the logger.




### <a name="Logger.Print">func</a> (\*Logger) [Print](/src/target/golog.go?s=6392:6432#L189)
``` go
func (l *Logger) Print(v ...interface{})
```
Print calls l.Output to print to the logger.
Arguments are handled in the manner of fmt.Print.




### <a name="Logger.Printf">func</a> (\*Logger) [Printf](/src/target/golog2.go?s=4651:4708#L149)
``` go
func (l *Logger) Printf(format string, v ...interface{})
```
Printf calls l.Output to print to the logger. Arguments are handled in the
manner of fmt.Printf. If the format string "%LVL" is used in combination with
an integer as the first value, the output level will be checked.  Output will
proceed only if the supplied level is less than or equal to the configured
logging level.




### <a name="Logger.Println">func</a> (\*Logger) [Println](/src/target/golog.go?s=6573:6615#L193)
``` go
func (l *Logger) Println(v ...interface{})
```
Println calls l.Output to print to the logger.
Arguments are handled in the manner of fmt.Println.




### <a name="Logger.SetFlags">func</a> (\*Logger) [SetFlags](/src/target/golog.go?s=7869:7904#L242)
``` go
func (l *Logger) SetFlags(flag int)
```
SetFlags sets the output flags for the logger.




### <a name="Logger.SetLevel">func</a> (\*Logger) [SetLevel](/src/target/golog2.go?s=5153:5187#L160)
``` go
func (l *Logger) SetLevel(n int)
```
SetLevel will set the output logging level for the Printf method. The default
logging level is -1 causing all levels to be output. Setting the level to any
value greater than -1 will allow all levels up to and including that level to
be output. For example, setting the level to 9 will allow Printf levels 0 - 9
to be output.




### <a name="Logger.SetOutput">func</a> (\*Logger) [SetOutput](/src/target/golog2.go?s=3940:3981#L126)
``` go
func (l *Logger) SetOutput(v interface{})
```
SetOutput sets the output destination for the logger.  The argument may be an
io.Writer or a filename string. If a string is supplied, the file is opened then
immediately closed which will create an empty file if it does not already exist.
If an error is encountered during file opening, the output will be set to stderr
and the logger Err value will be set.




### <a name="Logger.SetPrefix">func</a> (\*Logger) [SetPrefix](/src/target/golog.go?s=8152:8193#L256)
``` go
func (l *Logger) SetPrefix(prefix string)
```
SetPrefix sets the output prefix for the logger.








- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)