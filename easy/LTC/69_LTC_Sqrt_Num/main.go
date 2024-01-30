package main

/*
Given a non-negative integer x, return the square root of x rounded down to the nearest integer. The returned integer should be non-negative as well.

You must not use any built-in exponent function or operator.

For example, do not use pow(x, 0.5) in c++ or x ** 0.5 in python.


Example 1:

Input: x = 4
Output: 2
Explanation: The square root of 4 is 2, so we return 2.
Example 2:

Input: x = 8
Output: 2
Explanation: The square root of 8 is 2.82842..., and since we round it down to the nearest integer, 2 is returned.


Constraints:

0 <= x <= 231 - 1

*/
/*
*

	func mySqrt(x int) int {
	    if x < 2 {
	        return x;
	    }

	    var l, h, res int;
	    l = 1;
	    h = x / 2;

	    for l <= h {
	        mid := l + (h - l) / 2;

	        if mid * mid <= x {
	            res = mid;
	            l = mid + 1;
	        } else {
	            h = mid - 1;
	        }
	    }

	    return res;
	}
*/
func mySqrt(x int) int {
	if x <= 0 {
		return 0
	}
	var ans = 0
	p1 := 1
	p2 := x
	for p1 <= p2 {
		mid := p1 + ((p2 - p1) / 2)
		if mid*mid == x {
			return mid
		}
		if mid*mid < x {
			p1 = mid + 1
			ans = mid
		} else {
			p2 = mid - 1
		}
	}
	return ans
}
