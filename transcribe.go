/*
Problem:
An RNA string is a string formed from the alphabet containing 'A', 'C', 'G', and 'U'.
Given a DNA string t corresponding to a coding strand, its transcribed RNA string u is formed by replacing all occurrences of 'T' in t with 'U' in u.
Given: A DNA string t having length at most 1000 nt.
Return: The transcribed RNA string of t.
*/

package main

import (
    "fmt"
    "os"
    "strings"
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
        fmt.Println("./transcribe.go <input> <output>")
    } else {

        DNA := []string{"A","T","C","G"}

        in_file := os.Args[1]
        data, err := os.ReadFile(in_file)
        check(err)

        sequence := string(data)

        for _, x := range sequence {
            nuc := string(x)
            if !contains(DNA, nuc) {
                fmt.Println("Sequence contains invalid nucleotides")
                break
            }
        }
    
        sequence_transcribed := strings.ReplaceAll(sequence, "T", "U")

    
        out_file := os.Args[2]
        f, err := os.Create(out_file)
        data_transcribed, err := f.WriteString(sequence_transcribed)
        check(err)
        fmt.Println(data_transcribed)

    }
} // main
    