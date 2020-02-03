package Dictionary

import "fmt"

type Dictionary struct {
	Entries map[string]string
}

func New() *Dictionary{
	return new(Dictionary)
}

func(d Dictionary) Add(word string, definition string){
	if _, ok := d.Entries[word];ok{
		d.Entries[word] = definition
		fmt.Printf("%v is updated with the defintion %v !",word,definition)
	}else {
		d.Entries[word] = definition
		fmt.Printf("%v is inserted",word)
	}
}

func (d Dictionary) Delete(word string)  {
	if _, ok := d.Entries[word];ok{
		delete(d.Entries, word)
		fmt.Printf("%v is deleted",word)
	}else {
		fmt.Println("Sorry! Not Found. Deletion not successful.")
	}
}

func (d Dictionary) Search(word string){
	if definition, ok := d.Entries[word];ok{
		fmt.Println(definition)
	}else {
		fmt.Println("Sorry Not Found")
	}
}


