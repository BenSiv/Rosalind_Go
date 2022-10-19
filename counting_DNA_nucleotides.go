
/* 
Problem:
A string is simply an ordered collection of symbols selected from some alphabet and formed into a word; the length of a string is the number of symbols that it contains.
An example of a length 21 DNA string (whose alphabet contains the symbols 'A', 'C', 'G', and 'T') is "ATGCTTCAGAAAGGTCTTACG."
Given: A DNA string s of length at most 1000 nt.
Return: Four integers (separated by spaces) counting the respective number of times that the symbols 'A', 'C', 'G', and 'T' occur in s.
*/

package main

import (
    "fmt"
    "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func main() {

    if os.Args[1] == "help" {
        fmt.Println("./counting_DNA_nucleotides.go <file name>")
    } else {
        
        DNA := []string{"A","T","C","G"}

        in_file := os.Args[1]
        data, err := os.ReadFile(in_file)
        check(err)

        sequence := string(data)
        
        //Create a dictionary of values for each element
        DNA_dict := make(map[string]int)

        for _, x := range sequence {
            nuc := string(x)
            if !contains(DNA, nuc) {
                fmt.Println("Sequence contains invalid nucleotides")
                break
            } else {
                DNA_dict[nuc] = DNA_dict[nuc]+1
            }
        }

        fmt.Println(DNA_dict["A"], DNA_dict["C"], DNA_dict["G"], DNA_dict["T"])
    }
} // main
