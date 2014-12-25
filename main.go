package main


import "flag"
import "io"
import "os"


func main() {

	// Declare the command line flags.
	//
	// NOTE that we need to run "flag.Parse()" before we will actually get the data from the command line.
	//
		columns_ptr := flag.Uint64("columns", 1, "Determines how many column should be outputted. (default: 1)")

	// Parse the command line.
	//
	// This results in getting the command line flags we previously declared.
	// And also lets us get the (remaining) command line arguments (after the flags have been
	// removed) with "flag.Args()".
	//
		flag.Parse()

	// Decide what is going to get done.
	//
	// There are three (3) things that are part of the decision.
	// * writer
	// * reader
	// * columns
	//
	// #1: writer
	//
	//     The "writer" variable decides where this program will put its output onto.
	//
	//     This will be STDOUT. So that "writer" will be os.Stdout.
	//     (This is hardcoded.)
	//
	//     Of course, a shell could redirect the STDOUT from this program. So this seemingly
	//     fixed choice ends up having a lot of flexibility.
	//
	// #2: reader
	//
	//     The "reader" variable decides where this program will take its input from.
	//
	//     It can come from multiple places. STDIN if something it piped to it or
	//     (if nothing is piped to it from STDIN) from a file or named pipe if a
	//     file name or pipe name is specified on the command line. (The latter option
	//     can be used for process substitution, for shells that support it, such as
	//     bash and zsh.)
	//
	//     If STDIN has input on it, such as would be the case if the user piped something
	//     to this program, then "reader" will be os.Stdin.
	//
	//     So, for example, this corresponds to the case where the user does something like
	//     the following:
	//
	//         cat TABLE.tsv | redimension
	//
	//     However, if STDIN does NOT have input on it, then the command line arguments will
	//     be checked to see if any argument is a file name, pipe name, etc. If one is found,
	//     then only the first one found will be used as the input stream (by opening it with
	//     os.Open()).
	//
	//     That means that all other file names, pipe names, etc will be ignored! in this
	//     situation.
	//
	//     So, for example, this corresponds to the case where the user does something like
	//     the following:
	//
	//         redimension FILE1.tsv FILE2.tsv FILES3.tsv
	//
	//     In this case, only FILE1.tsv will be opened and used as input.
	//     Both FILE2.tsv and FILE3.tsv will be ignored.
	//
	//     Note that this allows for support process substitution, with shells that support
	//     process substitution (such as bash and zsh).
	//
	//     I.e. something like the following:
	//
	//         redimension <(cat TABLE.tsv)
	//
	// #3: columns
	//
	//     The "columns" variable decides how "wide" the columns of the output will be.
	//     I.e., how many columns there are going to be.
	//
	//     This is probably best described with an example.
	//
	//     Imagine that this program receives an input stream like the following:
	//
	//		1.01	0.301
	//		2.02	0.402
	//		3.03	0.503
	//		4.04	0.604
	//		5.05	0.705
	//		6.06	0.806
	//		7.07	0.907
	//		8.08	1.008
	//
	//     (Just a note that there are tab characters -- "\t" -- between each column.)
	//
	//     The "width" of columns of this input stream is: 2.
	//
	//     But let's says this command has "columns" with a value of 4. Then this program
	//     would transform that stream into the following (which it would send over the output
	//     stream).
	//
	//		1.01	0.301	2.02	0.402
	//		3.03	0.503	4.04	0.604
	//		5.05	0.705	6.06	0.806
	//		7.07	0.907	8.08	1.008
	//
	//     (Again, just a note that there are tab characters -- "\t" -- between each column.)
	//
	//     Here is another example.
	//
	//     Now consider the following input stream (that we might generate using the command "seq 1 9"):
	//
	//		1
	//		2
	//		3
	//		4
	//		5
	//		6
	//		7
	//		8
	//		9
	//
	//     If "columns" has a value of 3, then this would be transformed to the following on the output stream:
	//
	//		1	2	3
	//		4	5	6
	//		7	8	9
	//
	//     (And yet again, just a note that there are tab characters -- "\t" -- between each column.)
	//
		writer := os.Stdout

		var reader io.Reader
		if haveInputOnStdin() {
			reader = os.Stdin
		} else {
			if 1 > flag.NArg() {
				flag.PrintDefaults()
				os.Exit(1)
			}

			arg0 := flag.Arg(0)
			if ! isOpenable(arg0) {
				flag.PrintDefaults()
				os.Exit(1)
			}

			var err error
			reader, err = os.Open(arg0)
			if nil != err {
				flag.PrintDefaults()
				os.Exit(1)
			}
		}

		columns := *columns_ptr

	// Actually do what needs to get done.
	//
	// I.e., redimension.
	//
		redimension(writer, reader, columns)
}


// haveInputOnStdin returns true if input has been piped onto STDIN; returns false otherwise.
func haveInputOnStdin() bool {

	// Check to see if this program is receiving input from STDIN.
	//
	// I.e., check to see if data is being piped to this program.
	//
		have_input_on_stdin := false

		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) == 0 {
			have_input_on_stdin = true
		} else {
			have_input_on_stdin = false
		}

	// Return.
		return have_input_on_stdin
}


// isOpenable return true if the path parameter is openable and false otherwise.
// This is used to detect if path is a file name or a named pipe.
func isOpenable(path string) bool {

	// Check to see if this is an input file name or pipe name.
		is_openable := true
		if _, err := os.Stat(path); os.IsNotExist(err) {
			is_openable = false
		}

	// Return.
		return is_openable
}

