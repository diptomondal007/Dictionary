package Dictionary

import "fmt"

type  Dictionary map[string]string

func New() Dictionary{
	return Dictionary{}
}

func(d Dictionary) Add(word string, definition string) (string , error){
	if _, ok := d[word];!ok{
		d[word] = definition
		return word , nil
	}else {
		return word ,fmt.Errorf("%v already exists",word)
	}
}

func (d Dictionary) Update(word string, definition string)(string, error){
	if _, ok := d[word];ok{
		d[word] = definition
		return fmt.Sprintf("%v is updated",word),nil
	}else{
		return word,fmt.Errorf("sorry %v not found can not be updated",word)
	}
}

func (d Dictionary) Delete(word string) (string ,error) {
	if _, ok := d[word];ok{
		delete(d, word)
		return fmt.Sprintf("%v deleted",word),nil
	}else {
		return word,fmt.Errorf("sorry %v not found",word)
	}
}

func (d Dictionary) Search(word string) (string , error){
	if definition, ok := d[word];ok{
		return definition,nil
	}else {
		return word,fmt.Errorf("sorry %v not found",word)
	}
}


