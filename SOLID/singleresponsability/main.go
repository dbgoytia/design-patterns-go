package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

/*
	Single responsability.

	The single responsability behind the Journal struct is to manage the entries of the journal
	The single responsiability of the Persistent struct is to manage the persistency of data

	Below we introduce a couple of methods that break this pattern, save, load and loadfromweb, that mostly deal
	with persistency of the data. So it seems that this could be part of a different struct and we will show how
	to deal with this situation (sepparation of concerns).
*/

var entryCount = 0

type Journal struct {
	entries []string
}

func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s",
		entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

func (j *Journal) RemoveEntry(index int) {
	// ...
	j.entries = append(j.entries[:index], j.entries[index+1:]...)
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

// Sepparation of concerns
/*
// We could have placed this under the same package, but it's generally a good idea to have things sepparated
// for a single responsability

func (j *Journal) Save(filename string) {
	_ = ioutil.WriteFile(filename, []byte(j.String()), 0644)
}

func (j *Journal) Load(filename string) {
	// ...
}

func (j *Journal) LoadFromWeb(url *url.URL) {
	// ...
}
*/

// OPTION 1, STANDALONE METHOD:
// var LineSeparator = "\n"

// func SaveToFile(j *Journal, filename string) {
// 	_ = ioutil.WriteFile(filename, []byte(strings.Join(j.entries, LineSeparator)), 0644)
// }

type Persistence struct {
	lineSeparator string
}

func (p *Persistence) SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename, []byte(strings.Join(j.entries, p.lineSeparator)), 0644)
}

func (p *Persistence) LoadFile(filepath string) []string {
	dat, _ := ioutil.ReadFile(filepath)
	fmt.Println(string(dat))
	return strings.Split(string(dat), p.lineSeparator)
}

func main() {
	j := Journal{}
	fmt.Println("Inserting some data:")
	j.AddEntry("I cried today.")
	j.AddEntry("I ate a bug.")
	j.AddEntry("Still a noob.")
	fmt.Println(j.String())
	fmt.Println("Deleting some data:")
	j.RemoveEntry(0)
	fmt.Println(j.String())

	p := Persistence{"\r\n"}

	// Saving to file
	p.SaveToFile(&j, "journal.txt")

	// Loading from file
	fmt.Println("Loading saved journal.")
	j2 := Journal{}
	j2.entries = p.LoadFile("journal.txt")

}
