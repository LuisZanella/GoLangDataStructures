package main

import "fmt"

/*
//haystack ="sadbutsad"
//needle = "sad"
//nc=3
//i=2
//found=2-3

	func strStr(haystack string, needle string) int {
	    if len(needle)>len(haystack){
	        return -1
	    }
	    nc:=0
	    found:=-1
	    for i:=0;i<len(haystack);i++{
	        if haystack[i]==needle[nc]{
	            nc++
	            if nc>=len(needle){
	                found=i-nc+1
	                break
	            }
	        }else{
	            i=i-nc
	            nc=0
	        }
	    }
	    return found
	}
*/
func findIndexInSubString(haystack string, needle string) int {
	subLength := len(needle)
	for i, _ := range haystack {
		if i+subLength > len(haystack) {
			return -1
		}
		if haystack[i:i+subLength] == needle {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(findIndexInSubString("HelloWorld", "HelloWOOORLD"))
	fmt.Println(findIndexInSubString("HelloWorld", "Hello"))
	fmt.Println(findIndexInSubString("HellHelloWorld", "Hello"))
	fmt.Println(findIndexInSubString("HellHelloWorld", ""))
	fmt.Println(findIndexInSubString("", "a"))
	fmt.Println(findIndexInSubString("a", "a"))
	fmt.Println(findIndexInSubString("", ""))
	fmt.Println(findIndexInSubString("sadbutsad", "sad"))
}
