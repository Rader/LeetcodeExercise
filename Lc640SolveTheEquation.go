package LeetcodeExercise

import "bytes"
import "container/list"
import "fmt"
import "strconv"
import "log"

/*
url:https://leetcode.com/problems/solve-the-equation

Solve a given equation and return the value of x in the form of string "x=#value". The equation contains only '+', '-' operation, the variable x and its coefficient.

If there is no solution for the equation, return "No solution".

If there are infinite solutions for the equation, return "Infinite solutions".

If there is exactly one solution for the equation, we ensure that the value of x is an integer.

Example 1:
Input: "x+5-3+x=6+x-2"
Output: "x=2"
Example 2:
Input: "x=x"
Output: "Infinite solutions"
Example 3:
Input: "2x=x"
Output: "x=0"
Example 4:
Input: "2x+3x-6x=x+2"
Output: "x=-1"
Example 5:
Input: "x=x+2"
Output: "No solution"

*/

func solveEquation(equation string) string {
	var i int
	ls := list.New()
	var buf bytes.Buffer
	for i = 0; i < len(equation); i++ {
		c := equation[i]
		if c == '=' {
			ls.PushBack(parseItem(buf))
			buf.Reset()
			i++
			break //end left of equation
		}

		if c == '+' || c == '-' {
			ls.PushBack(parseItem(buf))
			buf.Reset()
			buf.WriteByte(c) //buf for next item
		} else {
			buf.WriteByte(c)
		}

	}
	rs := list.New()
	for ; i < len(equation); i++ {
		c := equation[i]
		if c == '+' || c == '-' {
			rs.PushBack(parseItem(buf))
			buf.Reset()
			buf.WriteByte(c) //buf for next item
		} else {
			buf.WriteByte(c)
		}
	}

	if buf.Len() > 0 {
		rs.PushBack(parseItem(buf))
		buf.Reset()
	}

	if ls.Len() == 0 || rs.Len() == 0 {
		return "invalid equation"
	}
	//move to left or right of =
	var cur *list.Element
	cur = ls.Front()
	for cur != nil {
		next := cur.Next()
		if cur.Value.(*item).exponential == 0 {
			ls.Remove(cur)
			cur.Value.(*item).coefficient = 0 - cur.Value.(*item).coefficient
			rs.PushBack(cur.Value)
		}
		cur = next
	}
	cur = rs.Front()
	for cur != nil {
		next := cur.Next()
		if cur.Value.(*item).exponential != 0 {
			rs.Remove(cur)
			cur.Value.(*item).coefficient = 0 - cur.Value.(*item).coefficient
			ls.PushBack(cur.Value)
		}
		cur = next
	}
	//compute
	cur = ls.Front()
	for cur != nil {
		next := cur.Next()
		if next == nil {
			break //finish
		}
		next.Value.(*item).coefficient += cur.Value.(*item).coefficient
		ls.Remove(cur)
		cur = next
	}
	cur = rs.Front()
	for cur != nil {
		next := cur.Next()
		if next == nil {
			break //finish
		}
		next.Value.(*item).coefficient += cur.Value.(*item).coefficient
		rs.Remove(cur)
		cur = next
	}
	var result string
	if ls.Len() == 0 {
		ls.PushBack(&item{coefficient: 0, exponential: 0})
	}
	if rs.Len() == 0 {
		rs.PushBack(&item{coefficient: 0, exponential: 0})
	}
	litem := ls.Front().Value.(*item)
	ritem := rs.Front().Value.(*item)
	if litem.coefficient == 0 {
		if ritem.coefficient == 0 {
			result = "Infinite solutions"
		} else {
			result = "No solution"
		}
	} else {
		result = fmt.Sprintf("x=%d", ritem.coefficient/litem.coefficient)
	}
	return result
}

func parseItem(buf bytes.Buffer) *item {
	it := new(item)
	if buf.Bytes()[buf.Len()-1] == 'x' {
		it.exponential = 1
		buf.Truncate(buf.Len() - 1) //remove last char
	}
	if buf.Len() == 0 {
		it.coefficient = 1 //only "x"
	} else if buf.Len() == 1 {
		if buf.Bytes()[0] == '+' {
			it.coefficient = 1
		} else if buf.Bytes()[0] == '-' {
			it.coefficient = -1
		} else {
			it.coefficient, _ = strconv.Atoi(string(buf.Bytes()[0]))
		}
	} else {
		number, err := strconv.ParseInt(buf.String(), 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		it.coefficient = int(number)
	}

	return it
}

type item struct {
	coefficient int
	exponential int
}

func (i *item) ToString() string {
	var buf bytes.Buffer
	if i.coefficient != 0 {
		buf.WriteString(strconv.Itoa(i.coefficient))
	}
	if i.exponential > 0 {
		buf.WriteByte('x')
	}
	return buf.String()
}

func printList(l *list.List) {
	n := l.Front()
	for n != nil {
		fmt.Print(n.Value.(*item).ToString(), " ")
		n = n.Next()
	}
	fmt.Println()
}
