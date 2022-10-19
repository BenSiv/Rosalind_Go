
/* 
Generates a random DNA sequence and saves it to a given file
*/

package main

import (
    "fmt"
    "time"
    "math/rand"
    "strconv"
    "os"
    "bufio"
)

	
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    if os.Args[1] == "help" {
        fmt.Println("./DNA_sequence_generator.go <length of sequence> <file name>")
    } else {

        DNA := [4]string{"A","T","C","G"}
        
        n, err1 := strconv.Atoi(os.Args[1])
        out_file := os.Args[2]

        if err1 != nil {
            fmt.Println("Given invalid number of nucleotides")
        } else {

            // set random seed for sampling
            rand.Seed(time.Now().UnixNano())

            // generate a sequence of random nucleotides
            sequence := make([]string, n)
            for i := range sequence {
                sequence[i] = DNA[rand.Intn(len(DNA))]
            }

            // create a file
            f, err2 := os.Create(out_file)
            check(err2)
            
            // create new writer
            w := bufio.NewWriter(f)

            // write each nucleotide to file
            for _, nuc := range sequence {
                w.WriteString(nuc)
            }

            // flush writer and close file
            w.Flush()
            f.Close()
        } 
    } // if help
} // main

