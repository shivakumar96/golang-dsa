package main

import (
	"container/list"
	"fmt"
	"sort"
	"strings"
)

type studnet struct {
	name string
	age  int
}

type studnetSlice []studnet

func (s studnetSlice) Len() int           { return len(s) }
func (s studnetSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s studnetSlice) Less(i, j int) bool { return strings.Compare(s[i].name, s[j].name) < 0 }

func main() {
	s1 := studnet{name: "abc", age: 14}
	s2 := studnet{name: "xyz", age: 13}
	s3 := studnet{name: "efg", age: 15}

	studnets := studnetSlice{s1, s2, s3}
	fmt.Println(studnets)

	sort.Sort(studnets)
	fmt.Println(studnets)

	sMap := make(map[studnet]*list.Element)

	studnetLists := list.New()
	studnetLists.Init()
	sMap[s1] = studnetLists.PushFront(s1)
	sMap[s2] = studnetLists.PushFront(s2)
	sMap[s3] = studnetLists.PushFront(s3)

	for ele := studnetLists.Front(); ele != nil; ele = ele.Next() {
		fmt.Println(ele)
	}
	val, _ := sMap[s2]
	studnetLists.Remove(val)
	fmt.Println("\n\n")
	for ele := studnetLists.Front(); ele != nil; ele = ele.Next() {
		fmt.Println(ele)
	}
}
