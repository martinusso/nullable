package nullable

import (
	"encoding/json"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

const (
	name   = "John Doe"
	age    = 42
	height = 1.79
)

type Person struct {
	Name        String
	Nickname    String
	Age         Int64
	Married     Bool
	HasChildren Bool
	Children    Int64
	Height      Float64
	Weight      Float64
}

func TestMarshall(t *testing.T) {
	var person Person
	person.Name.Scan(name)
	person.Age.Scan(age)
	person.Married.Scan(true)
	person.Height.Scan(height)

	got, err := json.Marshal(person)
	if err != nil {
		t.Errorf("error marshalling data. '%s'", err)
	}
	expected := `{"Name":"John Doe","Nickname":null,"Age":42,"Married":true,"HasChildren":null,"Children":null,"Height":1.79,"Weight":null}`
	if string(got) != expected {
		t.Errorf("Expected '%s' got '%s'", expected, got)
	}
}

func TestUnmarshall(t *testing.T) {
	body := []byte(`{"Name":"John Doe","Nickname":null,"Age":42,"Married":true,"HasChildren":null,"Children":null,"Height":1.79,"Weight":null}`)
	var person Person
	err := json.Unmarshal(body, &person)
	if err != nil {
		t.Errorf("error unmarshalling data. '%s'", err)
	}
	if got := person.Name.String; got != name {
		t.Errorf("Expected '%s' got '%s'", name, got)
	}
	if got := person.Nickname.String; got != "" {
		t.Errorf("Expected no nickname got '%s'", got)
	}
	if person.Married.Bool == false {
		t.Errorf("Expected married")
	}
	if person.HasChildren.Bool == true {
		t.Errorf("Expected has not children")
	}
	if got := person.Age.Int64; got != age {
		t.Errorf("Expected '%d' got '%d'", age, got)
	}
	if got := person.Children.Int64; got != 0 {
		t.Errorf("Expected no children got '%d'", got)
	}
	if got := person.Height.Float64; got != height {
		t.Errorf("Expected '%f' got '%f'", height, got)
	}
	if got := person.Weight.Float64; got != 0 {
		t.Errorf("Expected no children got '%f'", got)
	}
}

type Persom struct {
	Name    String
	Age     Int64
	Married Bool
	Height  Float64
}

func TestUnmarshalla(t *testing.T) {
	var person Persom

	body := []byte(`{"Name":"John Doe","Age":42,"Married":true,"Height":1.79}`)
	json.Unmarshal(body, &person)
	spew.Dump(person)
	// fmt.Println(person)
}
