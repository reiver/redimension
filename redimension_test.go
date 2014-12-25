package main


import "testing"
import "bytes"


func TestRedimension(t *testing.T) {

	tests := []struct{
		number_of_columns uint64
		stdin             string

		expected          string
	}{
		{
			number_of_columns: 0,
			stdin:    "1\n2\n3\n4\n5\n6\n7\n8\n9\n",
			expected: "\n",
		},
		{
			number_of_columns: 1,
			stdin:    "1\n2\n3\n4\n5\n6\n7\n8\n9\n",
			expected: "1\n2\n3\n4\n5\n6\n7\n8\n9\n",
		},
		{
			number_of_columns: 2,
			stdin:    "1\n2\n3\n4\n5\n6\n7\n8\n9\n",
			expected: "1\t2\n3\t4\n5\t6\n7\t8\n9\n",
		},
		{
			number_of_columns: 3,
			stdin:    "1\n2\n3\n4\n5\n6\n7\n8\n9\n",
			expected: "1\t2\t3\n4\t5\t6\n7\t8\t9\n",
		},
		{
			number_of_columns: 4,
			stdin:    "1\n2\n3\n4\n5\n6\n7\n8\n9\n",
			expected: "1\t2\t3\t4\n5\t6\t7\t8\n9\n",
		},
		{
			number_of_columns: 5,
			stdin:    "1\n2\n3\n4\n5\n6\n7\n8\n9\n",
			expected: "1\t2\t3\t4\t5\n6\t7\t8\t9\n",
		},
		{
			number_of_columns: 6,
			stdin:    "1\n2\n3\n4\n5\n6\n7\n8\n9\n",
			expected: "1\t2\t3\t4\t5\t6\n7\t8\t9\n",
		},
		{
			number_of_columns: 7,
			stdin:    "1\n2\n3\n4\n5\n6\n7\n8\n9\n",
			expected: "1\t2\t3\t4\t5\t6\t7\n8\t9\n",
		},
		{
			number_of_columns: 8,
			stdin:    "1\n2\n3\n4\n5\n6\n7\n8\n9\n",
			expected: "1\t2\t3\t4\t5\t6\t7\t8\n9\n",
		},
		{
			number_of_columns: 9,
			stdin:    "1\n2\n3\n4\n5\n6\n7\n8\n9\n",
			expected: "1\t2\t3\t4\t5\t6\t7\t8\t9\n",
		},
		{
			number_of_columns: 10,
			stdin:    "1\n2\n3\n4\n5\n6\n7\n8\n9\n",
			expected: "1\t2\t3\t4\t5\t6\t7\t8\t9\n",
		},
		{
			number_of_columns: 11,
			stdin:    "1\n2\n3\n4\n5\n6\n7\n8\n9\n",
			expected: "1\t2\t3\t4\t5\t6\t7\t8\t9\n",
		},
		{
			number_of_columns: 12,
			stdin:    "1\n2\n3\n4\n5\n6\n7\n8\n9\n",
			expected: "1\t2\t3\t4\t5\t6\t7\t8\t9\n",
		},
		{
			number_of_columns: 13,
			stdin:    "1\n2\n3\n4\n5\n6\n7\n8\n9\n",
			expected: "1\t2\t3\t4\t5\t6\t7\t8\t9\n",
		},
		{
			number_of_columns: 14,
			stdin:    "1\n2\n3\n4\n5\n6\n7\n8\n9\n",
			expected: "1\t2\t3\t4\t5\t6\t7\t8\t9\n",
		},
		{
			number_of_columns: 15,
			stdin:    "1\n2\n3\n4\n5\n6\n7\n8\n9\n",
			expected: "1\t2\t3\t4\t5\t6\t7\t8\t9\n",
		},



		{
			number_of_columns: 0,
			stdin:    "1\t2\n3\t4\n5\t6\n7\t8\n9\t10\n",
			expected: "\n",
		},
		{
			number_of_columns: 1,
			stdin:    "1\t2\n3\t4\n5\t6\n7\t8\n9\t10\n",
			expected: "1\n2\n3\n4\n5\n6\n7\n8\n9\n10\n",
		},
		{
			number_of_columns: 2,
			stdin:    "1\t2\n3\t4\n5\t6\n7\t8\n9\t10\n",
			expected: "1\t2\n3\t4\n5\t6\n7\t8\n9\t10\n",
		},
		{
			number_of_columns: 3,
			stdin:    "1\t2\n3\t4\n5\t6\n7\t8\n9\t10\n",
			expected: "1\t2\t3\n4\t5\t6\n7\t8\t9\n10\n",
		},
		{
			number_of_columns: 4,
			stdin:    "1\t2\n3\t4\n5\t6\n7\t8\n9\t10\n",
			expected: "1\t2\t3\t4\n5\t6\t7\t8\n9\t10\n",
		},
		{
			number_of_columns: 5,
			stdin:    "1\t2\n3\t4\n5\t6\n7\t8\n9\t10\n",
			expected: "1\t2\t3\t4\t5\n6\t7\t8\t9\t10\n",
		},
		{
			number_of_columns: 6,
			stdin:    "1\t2\n3\t4\n5\t6\n7\t8\n9\t10\n",
			expected: "1\t2\t3\t4\t5\t6\n7\t8\t9\t10\n",
		},
		{
			number_of_columns: 7,
			stdin:    "1\t2\n3\t4\n5\t6\n7\t8\n9\t10\n",
			expected: "1\t2\t3\t4\t5\t6\t7\n8\t9\t10\n",
		},
		{
			number_of_columns: 8,
			stdin:    "1\t2\n3\t4\n5\t6\n7\t8\n9\t10\n",
			expected: "1\t2\t3\t4\t5\t6\t7\t8\n9\t10\n",
		},
		{
			number_of_columns: 9,
			stdin:    "1\t2\n3\t4\n5\t6\n7\t8\n9\t10\n",
			expected: "1\t2\t3\t4\t5\t6\t7\t8\t9\n10\n",
		},
		{
			number_of_columns: 10,
			stdin:    "1\t2\n3\t4\n5\t6\n7\t8\n9\t10\n",
			expected: "1\t2\t3\t4\t5\t6\t7\t8\t9\t10\n",
		},
		{
			number_of_columns: 11,
			stdin:    "1\t2\n3\t4\n5\t6\n7\t8\n9\t10\n",
			expected: "1\t2\t3\t4\t5\t6\t7\t8\t9\t10\n",
		},
		{
			number_of_columns: 12,
			stdin:    "1\t2\n3\t4\n5\t6\n7\t8\n9\t10\n",
			expected: "1\t2\t3\t4\t5\t6\t7\t8\t9\t10\n",
		},
		{
			number_of_columns: 13,
			stdin:    "1\t2\n3\t4\n5\t6\n7\t8\n9\t10\n",
			expected: "1\t2\t3\t4\t5\t6\t7\t8\t9\t10\n",
		},
		{
			number_of_columns: 14,
			stdin:    "1\t2\n3\t4\n5\t6\n7\t8\n9\t10\n",
			expected: "1\t2\t3\t4\t5\t6\t7\t8\t9\t10\n",
		},
		{
			number_of_columns: 15,
			stdin:    "1\t2\n3\t4\n5\t6\n7\t8\n9\t10\n",
			expected: "1\t2\t3\t4\t5\t6\t7\t8\t9\t10\n",
		},
	}


	// Run each test.
	//
		for test_number, test := range tests {

			// Create the writer that we can use to capture the output.
				var w bytes.Buffer

			// Create the reader that we will use to send the input.
			//
			// And the put the input into it.
				var r bytes.Buffer

				_,_ = r.WriteString(test.stdin)

			// Invoke redimension().
				redimension(&w, &r, test.number_of_columns)

			// Check that redimension() outputted what we expected.
				if actual := w.String() ; test.expected != actual {
					t.Errorf("Did not get what was expected for test #%d.\nExpected: %q\nReceived: %q\n", test_number, test.expected, actual)
				}
		}
}
