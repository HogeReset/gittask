// New cal project main.go
package main

import (
	"bufio"
	"data-structure/stack"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	print("请输入需要计算的四则运算式: ")
	var stat string
	reader := bufio.NewReader(os.Stdin)
	stat, _ = reader.ReadString('\n')
	stat = strings.TrimSpace(stat)
	postfix := infix2ToPostfix(stat)
	fmt.Printf("后缀表达式：%s\n", postfix)
	fmt.Printf("计算结果: %d\n", calculate(postfix))
}

func calculate(postfix string) int 
{
	stack := stack.ItemStack{}
	fixLen := len(postfix)
	for i := 0; i < fixLen; i++ {
		nextChar := string(postfix[i])
		if unicode.IsDigit(rune(postfix[i])) 
		{
			stack.Push(nextChar)
		} else 
		{
			num1, _ := strconv.Atoi(stack.Pop())
			num2, _ := strconv.Atoi(stack.Pop())
			switch nextChar 
			{
			case "+":
				stack.Push(strconv.Itoa(num1 + num2))
			case "-":
				stack.Push(strconv.Itoa(num1 - num2))
			case "*":
				stack.Push(strconv.Itoa(num1 * num2))
			case "/":
				stack.Push(strconv.Itoa(num1 / num2))
			}
		}
	}
	result, _ := strconv.Atoi(stack.Top())
	return result
}

func infix2ToPostfix(exp string) string 
{
	stack := stack.ItemStack{}
	postfix := ""
	expLen := len(exp)
	for i := 0; i < expLen; i++
	 {

		char := string(exp[i])

		switch char
		 {
		case " ":
			continue
		case "(":
			stack.Push("(")
		case ")":
			for !stack.IsEmpty()
			 {
				preChar := stack.Top()
				if preChar == "(" 
				{
					stack.Pop()
					break
				}
				postfix += preChar
				stack.Pop()
			}

		case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
			j := i
			digit := ""
			for ; j < expLen && unicode.IsDigit(rune(exp[j])); j++ 
			{
				digit += string(exp[j])
			}
			postfix += digit
			i = j - 1 

		default:
			for !stack.IsEmpty() 
			{
				top := stack.Top()
				if top == "(" || isLower(top, char) 
				{
					break
				}
				postfix += top
				stack.Pop()
			}
			stack.Push(char)
		}
	}

	for !stack.IsEmpty() 
	{
		postfix += stack.Pop()
	}

	return postfix
}

func isLower(top string, newTop string) bool
{
	
	switch top 
	{
	case "+", "-":
		if newTop == "*" || newTop == "/" 
		{
			return true
		}
	case "(":
		return true
	}
	return false
}


package stack

import
(
	"sync"
)

type Item string

type ItemStack struct 
{
	items []string
	lock  sync.RWMutex
}

func (s *ItemStack) New() *ItemStack 
{
	s.items = []string{}
	return s
}


func (s *ItemStack) Push(t string) 
{
	s.lock.Lock()
	s.items = append(s.items, t)
	s.lock.Unlock()
}

func (s *ItemStack) Pop() string
{
	s.lock.Lock()
	item := s.items[len(s.items)-1]
	s.items = s.items[0 : len(s.items)-1]
	s.lock.Unlock()
	return item
}

func (s *ItemStack) Top() string
{
	return s.items[len(s.items)-1]
}

func (s *ItemStack) IsEmpty() bool 
{
	return len(s.items) == 0
}


package stack

import "testing"

var stack ItemStack

func initStack() *ItemStack 
{
	if stack.items == nil
    {
		stack = ItemStack{}
		stack.New()
	}
	return &stack
}

func TestPush(t *testing.T) 
{
	stack := initStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	if size := len(stack.items); size != 3
    {
		t.Errorf("Wrong stack size, expected 3 and got %d", size)
	}
}

func TestPop(t *testing.T) 
    {
	stack.Pop()
	if size := len(stack.items); size != 2 
	{
		t.Errorf("Wrong stack size, expected 2 and got %d", size)
	}
	stack.Pop()
	stack.Pop()
	if size := len(stack.items); size != 0 
	{
		t.Errorf("Wrong stack size, expected 0 and got %d", size)
	}
}
