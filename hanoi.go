package main

import (
	`fmt`
	`strconv`
)

var count int
var hanoiList Hanois

type Hanoi struct {
	Arr []int
	Index string
}

type Hanois = []Hanoi

func hanoi(a, b, c *Hanoi, n int){
	if n > 1 {
		hanoi(a, c, b, n-1)
		a, c = setings(a,c,n)
		printHanoi(a, b, c, n,count)
		hanoi(b, a, c, n-1)
	} else {
		a, c = setings(a,c,n)
		printHanoi(a, b, c, n,count)
	}
}

func setings(a, c *Hanoi, n int) (*Hanoi, *Hanoi) {
	count += 1
	return unset(a, n), set(c,n)
}

func unset(hanoi *Hanoi, n int) *Hanoi {
	tmp := []int{}
	for i := 0; i < len(hanoi.Arr); i++ {
		if hanoi.Arr[i] == n {
			continue
		}else {
			tmp = append(tmp, hanoi.Arr[i])
		}
	}
	hanoi.Arr = tmp
	return hanoi
}

func set(hanoi *Hanoi, n int) *Hanoi {
	hanoi.Arr = append(hanoi.Arr, n)
	return hanoi
}

func printHanoi(a, b, c *Hanoi, n, count int)  {
	fmt.Printf("Step %d: move %d %s to %s\n",count, n, a.Index, c.Index)
	fmt.Printf("%s: %v\n",a.Index, a.Arr)
	fmt.Printf("%s: %v\n",b.Index, b.Arr)
	fmt.Printf("%s: %v\n",c.Index, c.Arr)
}

func printHanois(hanois *Hanois, n, count int)  {
	fmt.Printf("Step %d: move %d %s to %s\n",count, n, (*hanois)[0].Index, (*hanois)[2].Index)
	fmt.Printf("%s: %v\n",(*hanois)[0].Index, (*hanois)[0].Arr)
	fmt.Printf("%s: %v\n",(*hanois)[1].Index, (*hanois)[1].Arr)
	fmt.Printf("%s: %v\n",(*hanois)[2].Index, (*hanois)[2].Arr)
}

func hanoiL(a, b, c *Hanoi, n int){
	if n > 1 {
		hanoiL(a, c, b, n-1)
		setings(a,c,n)
		printHanois(&hanoiList, n,count)
		hanoiL(b, a, c, n-1)
	} else {
		setings(a,c,n)
		printHanois(&hanoiList, n,count)
	}
}

func main()  {
	n := 5
	aUnicode := 41
	hanoiList = make(Hanois, n)
	for i := 0; i < 3; i++ {
		uni := strconv.Itoa(aUnicode)
		hanoiList[i].Index,_ = strconv.Unquote(`"\u00` + uni + `"`)
		aUnicode+=1
		if i == 0 {
			for j := 0; j < n; j++{
				hanoiList[i].Arr = append(hanoiList[i].Arr, n-j)
			}
		}
	}
	//hanoi(a, b, c, n)

	hanoiL(&hanoiList[0],&hanoiList[1],&hanoiList[2],n)
}
