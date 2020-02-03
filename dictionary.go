package Dictionary

import "fmt"

type Dictionary struct {
	Entries map[string]string
}

func New() *Dictionary{
	newDictionary := new(Dictionary)
	newDictionary.Entries = make(map[string]string)
	return newDictionary
}

func(d Dictionary) Add(word string, definition string){
	if _, ok := d.Entries[word];ok{
		d.Entries[word] = definition
		fmt.Printf("%v is updated with the defintion %v !\n",word,definition)
	}else {
		d.Entries[word] = definition
		fmt.Printf("%v is inserted\n",word)
	}
}

func (d Dictionary) Delete(word string)  {
	if _, ok := d.Entries[word];ok{
		delete(d.Entries, word)
		fmt.Printf("%v is deleted\n",word)
	}else {
		fmt.Println("Sorry! Not Found. Deletion not successful.")
	}
}

func (d Dictionary) Search(word string){
	if definition, ok := d.Entries[word];ok{
		fmt.Printf("The defintion for the word %v is %v\n",word,definition)
	}else {
		fmt.Printf("Sorry %v not found",word)
	}
}


