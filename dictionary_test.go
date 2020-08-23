package Dictionary

import (
	"log"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	var dict  = New()
	if reflect.ValueOf(dict).Kind() == reflect.Map {
		log.Println("Successful")
	}else {
		t.Errorf("%q wanted, got %q",reflect.TypeOf(Dictionary{}),reflect.TypeOf(dict))
	}
}

func TestDictionary_Add(t *testing.T) {
	var dict  = New()
	res , err := dict.Add("apple","a fruit")
	if res == "apple" && err == nil{
		val := dict["apple"]
		if val == "a fruit" {
			log.Println(val)
		}else {
			t.Errorf("%q wanted, got %q","a fruit",val)
		}

	}else if err != nil {
		t.Errorf("%q wanted, got %v","a fruit",err)
	}else {
		t.Errorf("%q wanted, got %q","apple",res)
	}
}

func TestDictionary_Add_InsertFailed(t *testing.T) {
	var dict  = New()
	dict["apple"] = "a fruit"
	_ , err := dict.Add("apple","a fruit")
	if err == nil{
		t.Errorf("%q wanted, got %q","Already exist",err)
	}
}

func TestDictionary_Update(t *testing.T) {
	var dict  = New()
	dict["apple"] = "a fruit"
	 _,err :=dict.Update("apple","updated")
	if dict["apple"] != "updated" || err != nil {
		t.Errorf("%q wanted, got %q","updated",dict["apple"])
	}
}

func TestDictionary_UpdateFailed(t *testing.T) {
	var dict  = New()
	_,err :=dict.Update("hello","will be failed")
	if err == nil{
		t.Errorf("%q wanted, got %q","updated",dict["apple"])
	}
}

func TestDictionary_Delete(t *testing.T) {
	var dict  = New()
	dict["apple"] = "a fruit"
	_,err :=dict.Delete("apple")
	if err != nil {
		t.Errorf("%v wanted, got %q",nil,err)
	}
}
func TestDictionary_Delete_Failed(t *testing.T) {
	var dict  = New()
	_,err :=dict.Delete("apple")
	if err == nil {
		t.Errorf("%v wanted, got %q","apple insertion failed",err)
	}
}

func TestDictionary_Search(t *testing.T) {
	var dict  = New()
	dict["apple"] = "a fruit"
	definition ,err :=dict.Search("apple")
	if err == nil {
		if definition !="a fruit"{
			t.Errorf("%v wanted, got %q","a fruit",definition)
		}
	}else {
		t.Errorf("%v wanted, got %v","a fruit",err)
	}
}

func TestDictionary_SearchFailed(t *testing.T) {
	var dict  = New()
	_ ,err :=dict.Search("apple")
	if err == nil {
			t.Errorf("%q wanted, got %v","Err",nil)
	}
}
