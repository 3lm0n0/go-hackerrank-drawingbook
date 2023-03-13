package main

import "fmt"

func main() {
	fmt.Println(pageCount(int32(6), int32(2)))
}

func pageCount(n int32, p int32) int32 {
	if n == 1 || p == 1 {
		return 0
	}
	if n == p {
		return 0
	}
	pageNumber := int32(0)
	isEven := bool(false)
	x := int32(0)
	y := int32(1)
	// is even? beacuse evens will have at last page only one number
	if n%2 == 0 {
		isEven = true
	}
	fmt.Println("isEven = ", isEven)
	middlePage := int32(n/2) + 1
	fmt.Println("middlePage = ", middlePage)
	// starts from last page
	if p >= middlePage {
		fmt.Println("starting from last page, n = ", n)
		// p is even
		if isEven {
			x = int32(0)
			y = int32(n)
			for i := int32(0); i < n; i++ {
				if i == 1 {
					x = n - 1
				}
				fmt.Println("x = ", x)
				fmt.Println("y = ", y)
				if p == x || p == y {
					pageNumber = i
					break
				}
				x -= 2
				y -= 2
			}
			return pageNumber
		}
		// p is odd
		x = int32(n - 1)
		y = int32(n)
		for i := int32(0); i < n; i++ {
			if p == x || p == y {
				pageNumber = i
				break
			}
			x -= 2
			y -= 2
		}
		return pageNumber
	}
	// starts from first page
	fmt.Println("starting from first page, n = ", n)
	// starts from first page
	for i := int32(0); i < n; i++ {
		if p == x || p == y {
			pageNumber = i
			break
		}
		x += 2
		y += 2
	}

	return pageNumber
	}
	// starts from first page
	fmt.Println("starting from first page, n = ", n)
	// starts from first page
	for i := int32(0); i < n; i++ {
		if p == x || p == y {
			pageNumber = i
			break
		}
		x += 2
		y += 2
	}

	return pageNumber
}

// int n: the number of pages in the book
// int p: the page number to turn to
