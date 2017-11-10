// Copyright (c) 2017 Mark K Mueller (github.com/mkmueller). All rights reserved.
// License: GNU GPL version 3 <http://gnu.org/licenses/gpl.html>

package golog_test

import (
	"io"
    "os"
    "fmt"
	"time"
	"bytes"
	"testing"
	"os/exec"
	"io/ioutil"
	log "github.com/mkmueller/golog"
    . "github.com/smartystreets/goconvey/convey"
)

const run_this = "GO_TEST_RUN_THIS"

func TestMain (m *testing.M) {
	time.Sleep( time.Second * 1 )	// slow things down a bit in case something goes wrong
	returnCode := 0
	switch os.Getenv(run_this) {
		case "testFatal_to_file":
			testFatal_to_file()
		case "testFatalf_to_file":
			testFatalf_to_file()
		case "":
			returnCode = m.Run()
		default:
			fmt.Println("Unknown test variable")
			returnCode = 1
	}
	os.Exit(returnCode)
}

func TestFatalf_to_file (t *testing.T) {

	temp_log_file := createTempFile()

	Convey( "Given a temporary log file name:"+temp_log_file, t, func(){

		Convey( "Log Fatal() message to file:"+temp_log_file, func(){
			cmd := exec.Command( os.Args[0], "test" )
			cmd.Env = []string{run_this+"=testFatal_to_file", "GO_TEMP_LOG="+temp_log_file}
			// Finally, we run the new process and wait for it to complete.
			err := cmd.Run()
			So( err, ShouldNotEqual, nil )
			s := readFile(temp_log_file)
			So( len(s), ShouldNotEqual, 0 )
			os.Remove(temp_log_file)
		})

		Convey( "Log Fatalf() message to file:"+temp_log_file, func(){
			cmd := exec.Command( os.Args[0], "test" )
			cmd.Env = []string{run_this+"=testFatalf_to_file", "GO_TEMP_LOG="+temp_log_file}
			// Finally, we run the new process and wait for it to complete.
			err := cmd.Run()
			So( err, ShouldNotEqual, nil )
			s := readFile(temp_log_file)
			So( len(s), ShouldNotEqual, 0 )
			os.Remove(temp_log_file)
		})

	})

}

func testFatal_to_file () {
	temp_log_file := os.Getenv("GO_TEMP_LOG")
	defer func(){
		if r := recover(); r == nil {

		}
	}()
	if temp_log_file != "" {
		mylog := log.New(temp_log_file)
		mylog.Fatal("This FATAL message will print to the file")
	}
}

func testFatalf_to_file () {
	temp_log_file := os.Getenv("GO_TEMP_LOG")
	if temp_log_file != "" {
		mylog := log.New(temp_log_file)
		mylog.Fatalf("This FATAL message will print to the file")
	}
}

func TestPrint_to_file (t *testing.T) {

	var temp_log_file string

	Convey( "Given a temporary log file name", t, func(){

		temp_log_file = createTempFile()

		Convey( "Print() ", func(){

			mylog := log.New(temp_log_file)
			mylog.Print("This message will print to the file")
			s := readFile(temp_log_file)

			So( len(s), ShouldNotEqual, 0 )

			os.Remove(temp_log_file)

		})

		Convey( "Printf() ", func(){

			mylog := log.New(temp_log_file)
			mylog.Printf("This message will print to the file")
			s := readFile(temp_log_file)

			So( len(s), ShouldNotEqual, 0 )

			os.Remove(temp_log_file)
		})

		Convey( "Println()", func(){

			mylog := log.New(temp_log_file)
			mylog.Println("This message will print to the file")
			s := readFile(temp_log_file)

			So( len(s), ShouldNotEqual, 0 )

			os.Remove(temp_log_file)
		})

	})

}

func TestPrintf_to_illegal_filename (t *testing.T) {

	Convey( "Printf() output all levels", t, func(){
		s := captureStderr( func(){
			mylog := log.New("./this/path/should/not/exist.log")
			mylog.Printf("This message will print to stderr", 1)
		})
		So( len(s), ShouldNotEqual, 0 )
	})

}

func TestPrintf_levels_all (t *testing.T) {

	Convey( "Printf() output all levels", t, func(){

		Convey( "Given Level 1 format:", func(){
			s := captureStderr( func(){
				mylog := log.New()
				mylog.Printf("[%LVL] This message will print", 1)
			})
			So( len(s), ShouldNotEqual, 0 )
		})

		Convey( "Given Level 99 format:", func(){
			s := captureStderr( func(){
				mylog := log.New()
				mylog.Printf("[%LVL] This message should not print", 99)
			})
			So( len(s), ShouldNotEqual, 0 )
		})

	})

}


func TestPrintf_level_1 (t *testing.T) {

	Convey( "Printf() Level 1 output only", t, func(){

		Convey( "Given Level 1 format:", func(){

			s := captureStderr( func(){
				mylog := log.New()
				mylog.SetLevel(1)
				mylog.Printf("[%LVL] This message will print", 1)
			})

			So( len(s), ShouldNotEqual, 0 )

		})

		Convey( "Given Level 2 format:", func(){

			s := captureStderr( func(){
				mylog := log.New()
				mylog.SetLevel(1)
				mylog.Printf("[%LVL] This message should not print", 2)
			})

			So( len(s), ShouldEqual, 0 )

		})

	})

}

func TestNew_toStdout (t *testing.T) {

	Convey( "New() given os.Stdout argument:", t, func(){

		s := captureStdout( func(){
			mylog := log.New(os.Stdout)
			mylog.Printf("test")
		})

		So( len(s), ShouldNotEqual, 0 )

	})

}

func TestNew_toStderr (t *testing.T) {

	Convey( "New() given no arguments:", t, func(){

		s := captureStderr( func(){
			mylog := log.New()
			mylog.Printf("test")
		})

		So( len(s), ShouldNotEqual, 0 )

	})

}


func TestNew_flags_only (t *testing.T) {

	Convey( "New() given flags only:", t, func(){
		s := captureStderr( func(){
			mylog := log.New(log.LstdFlags|log.LUTC)
			mylog.Printf("test")
		})
		So( len(s), ShouldNotEqual, 0 )
	})

}

func TestNew_Stdout_and_flags (t *testing.T) {

	Convey( "New() given Stdout and flags:", t, func(){
		s := captureStdout( func(){
			mylog := log.New(os.Stdout, log.LstdFlags|log.LUTC)
			mylog.Printf("test")
		})
		So( len(s), ShouldNotEqual, 0 )
	})

}

func TestNew_flags_and_prefix (t *testing.T) {

	Convey( "New() given only flags and prefix", t, func(){
		s := captureStderr( func(){
			mylog := log.New(log.LstdFlags|log.LUTC, "[PREFIX]")
			mylog.Printf("test")
		})
		So( len(s), ShouldNotEqual, 0 )
	})

}

func TestNew_stdout_flags_and_prefix (t *testing.T) {

	Convey( "New() given traditional arguments", t, func(){
		s := captureStdout( func(){
			mylog := log.New(os.Stdout, "[PREFIX]", log.LstdFlags|log.LUTC)
			mylog.Printf("test")
		})
		So( s[:8], ShouldEqual, "[PREFIX]" )
	})

}

func TestPrefix (t *testing.T) {

	Convey( "Prefix() method", t, func(){
		var pref string
		s := captureStdout( func(){
			mylog := log.New(os.Stdout, "[PREFIX]")
			pref = mylog.Prefix()
			mylog.Printf("test")
		})
		So( s[:8], ShouldEqual, "[PREFIX]" )
		So( pref, ShouldEqual, "[PREFIX]" )
	})

}


func TestSetPrefix (t *testing.T) {

	Convey( "SetPrefix() method", t, func(){
		s := captureStderr( func(){
			mylog := log.New()
			mylog.SetPrefix("[PREFIX]")
			mylog.Printf("test")
		})
		So( s[:8], ShouldEqual, "[PREFIX]" )
	})

}

func TestSetFlags (t *testing.T) {

	Convey( "SetFlags() method", t, func(){
		s := captureStderr( func(){
			mylog := log.New()
			mylog.SetFlags(log.LstdFlags|log.LUTC)
			mylog.Printf("test")
		})
		So( len(s), ShouldNotEqual, 0 )
	})

}

func TestSetOutput (t *testing.T) {

	Convey( "SetOutput() method", t, func(){
		s := captureStdout( func(){
			mylog := log.New()
			mylog.SetOutput(os.Stdout)
			mylog.Printf("test")
		})
		So( len(s), ShouldNotEqual, 0 )
	})

}

func TestOutput_2 (t *testing.T) {

	var err error
	Convey( "Output() method", t, func(){
		s := captureStderr( func(){
			mylog := log.New()
			err = mylog.Output(1, "Output something")
		})
		So( err, ShouldEqual, nil )
		So( s, ShouldNotEqual, "" )
	})

}


//  //  //  //  //  //  //  //  //  //  //  //  //  //  //  //  //  //  //  //  //  //  //

func captureStdout ( fn func() ) string {
    prevStdout := os.Stdout
    r, w, _ := os.Pipe()
    os.Stdout = w
	fn()
    ch := make(chan string)
    go func() {
        var buff bytes.Buffer
        io.Copy(&buff, r)
        ch <- buff.String()
    }()
    w.Close()
    os.Stdout = prevStdout
    str := <-ch
    return str
}

func captureStderr ( fn func() ) string {
    prevStderr := os.Stderr
    r, w, _ := os.Pipe()
    os.Stderr = w
	fn()
    ch := make(chan string)
    go func() {
        var buff bytes.Buffer
        io.Copy(&buff, r)
        ch <- buff.String()
    }()
    w.Close()
    os.Stderr = prevStderr
    str := <-ch
    return str
}

// Create a temp file for output
func createTempFile () string {
	fl, err := ioutil.TempFile("/tmp", "GOTEST_GOLOG_")
	dieIfError(err)
	fl.Close()
	return fl.Name()
}

func readFile ( file string ) []byte {
	b,err := ioutil.ReadFile(file)
	dieIfError(err)
	return b
}

func dieIfError ( err error ) {
	if err != nil {
		//fmt.Println(err.Error())
		panic(err.Error())
		//os.Exit(1)
	}
}

func openFile (file string) *os.File {
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0660)
	dieIfError(err)
	return f
}


//func Convey (msg string, t *testing.T, fn func() ) {
//	fmt.Println("\n"+msg)
//	fn()
//}
