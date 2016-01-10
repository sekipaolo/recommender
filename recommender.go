package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	main_user := stringToInt(os.Args[1])
	data_path := os.Args[2]

	// Open the data file
	file, err := os.Open(data_path)
	if err != nil {
    fmt.Printf("error opening file: %v\n",err)
    os.Exit(1)
	}
	defer file.Close()

	// read the data file
	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)

	// set global variables
	purchases := make([]string, 0)
	boughtTogether := make(map[string][]string)	
	currentUser := 0
	currentOrder := make([]string, 0)		
	// finally iterate rows
	for scanner.Scan() {
		line := scanner.Text()
		// break if end of file is reached: probably there is a more elegant way
		if line == "" {
			break
		}
		cols := strings.Split(line, ";")
		user 		:= stringToInt(cols[1])
		product := cols[2]

		if user == main_user { // item purchased by the user
			purchases = append(purchases, product)
		} else {
			if user != currentUser {
			// 	// save current order if contains more than one item			
				if len(currentOrder) > 1 { 
					for i := range currentOrder {
						currentItem := currentOrder[i]
			 			// exclude current item from the items to append
						itemsToAppend := deleteFromSlice(currentOrder, currentItem)
						boughtTogether[currentItem] = append(boughtTogether[currentItem], itemsToAppend...)						
					}
				}
			// 	// start a new order
				currentUser = user			
				currentOrder = nil
				currentOrder = make([]string, 0)
			}
			// // anyway .... add the product to the current cart
			currentOrder = append(currentOrder, product)
		}
	}	
	// we consider similar products the ones bought together with items
	// already bought by the user.
	similar := make([]string, 0)
	for i := range purchases {
		product := purchases[i]
		items := boughtTogether[product]
		similar = append(similar, items...)	
	}	
	// we keep duplicates so we can assign a weight to each product
	similar_map := freq_count(similar)
	similar_sorted := sort_map(similar_map)

	//now we print out the results	
	cnt := 0	
	for _, v := range similar_sorted {
		cnt++
		//fmt.Printf("%v %v \n", v.freq, v.product)	
		fmt.Printf("%v; %v\n", v.freq, v.product)	
		if cnt == 5 { 
		 	os.Exit(1)
		}
	}	
	// if we are here, we have less results as we should. Add results from boughtTogether
	// that is unsorted, so the default go access method garanties random results
	for k, _ := range boughtTogether {
		if sliceContains(similar, k) {
			continue
		}		
		cnt++
		fmt.Printf("%v\n", k)	
		if cnt == 5 { 
		 	os.Exit(1)
		}
	}

}
