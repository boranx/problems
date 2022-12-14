package main

import "fmt"

/*
Given a file system structure representing directories and files,
Return the total size of files in the structure.

home/
├── me/
│   ├── foo.txt : 416
│   ├── metrics.txt : 5892
│   └── src/
│       ├── site.html : 6051
│       ├── site.css : 5892
│       └── data.csv : 332789
└── you/
    └── dict.json : 4913364
bin/
├── bash: 618416
├── cat: 23648
└── ls: 38704
var/
└── log/
     ├── dmesg : 1783894
     ├── wifi.log : 924818
     └── httpd/
         ├── access.log : 17881
         └── access.log.0.gz : 4012

total_size(file_system) -> 8675777
*/

type Node struct {
	Name     string // Name of the path node
	Size     int    // Size of the file. Will be 0 for directories.
	Children []Node // Child nodes for directories.
}

var fs = []Node{
	{Name: "home/", Children: []Node{
		{Name: "me/", Children: []Node{
			{Name: "foo.txt", Size: 416},
			{Name: "metrics.txt", Size: 5892},
			{Name: "src/", Children: []Node{
				{Name: "site.html", Size: 6051},
				{Name: "site.css", Size: 5892},
				{Name: "data.csv", Size: 332789},
			}},
		}},
		{Name: "you/", Children: []Node{
			{Name: "dict.json", Size: 4913364},
		}},
	}},
	{Name: "bin/", Children: []Node{
		{Name: "bash", Size: 618416},
		{Name: "cat", Size: 23648},
		{Name: "ls", Size: 38704},
	}},
	{Name: "var/", Children: []Node{
		{Name: "log/", Children: []Node{
			{Name: "dmesg", Size: 1783894},
			{Name: "wifi.log", Size: 924818},
			{Name: "httpd/", Children: []Node{
				{Name: "access.log", Size: 17881},
				{Name: "access.log.0.gz", Size: 4012},
			}},
		}},
	}},
}

func calculateTotalSize(fs []Node) int {
	total := 0
	for _, node := range fs {
		total += calculateTotalSizeNode(node)
	}
	return total
}

func calculateTotalSizeNode(node Node) int {
	total := node.Size
	for _, child := range node.Children {
		total += calculateTotalSizeNode(child)
	}
	return total
}

func main() {
	total := calculateTotalSize(fs)
	fmt.Println(total) // 8675777
}
