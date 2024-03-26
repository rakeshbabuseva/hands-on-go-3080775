// challenges/testing/begin/main_test.go
package main

import "testing"

// write a test for letterCounter.count
func TestLetterCounter(t *testing.T){
  tests := map[string]struct{
		input string
		want int
	}{
		"empty": {input:"", want:0 },
		"one": {input:"a2 32^ &/", want:1 },
		"two": {input:"812 6ab//", want:2 },
	}
	l:=letterCounter{"LetterCounter"}
	for _, v := range tests{
		t.Run(v.input, func(t *testing.T) {
		got:= l.count(v.input)
			if got!=v.want{
				t.Errorf("got %v want %v", got, v.want)
			}
		})
	
}
}

func TestName(t *testing.T){
  tests := map[string]struct{
		input string
		want string
	}{
		"empty": {input:"", want:"" },
		"one": {input:"a2 32^ &/", want: "a2 32^ &/"},
		"two": {input:"812 6ab//", want:"812 6ab//" },
	}
	for _, v := range tests{
		t.Run(v.input, func(t *testing.T) {
			l:=numberCounter{v.input}
		got:= l.name()
		if got!=v.want{
			t.Errorf("got %v want %v", got, v.want)
		}
		})
	}
		for _, v := range tests{
		t.Run(v.input, func(t *testing.T) {
			l:=letterCounter{identifier: v.input}
		got:= l.name()
		if got!=v.want{
			t.Errorf("got %v want %v", got, v.want)
		}
		})
	}
		for _, v := range tests{
		t.Run(v.input, func(t *testing.T) {
			l:=symbolCounter{v.input}
		got:= l.name()
		if got!=v.want{
			t.Errorf("got %v want %v", got, v.want)
		}
		})
		}
}
// write a test for numberCounter.count
func TestNumberCounter(t *testing.T){
	tests := map[string]struct{
		input string
		want int
	}{
		"empty": {input:"", want:0 },
		"one": {input:"abc 1?", want:1 },
		"two": {input:"abc 12&8 ^", want:3 },
	}
	l:=numberCounter{"NumberCounter"}
	for _, v := range tests{
		t.Run(v.input, func(t *testing.T) {
		got:= l.count(v.input)
		if got!=v.want{
			t.Errorf("got %v want %v", got, v.want)
		}
		})
	
}
}

// write a test for symbolCounter.count
func TestSymbolCounter(t *testing.T){
	tests := map[string]struct{
		input string
		want int
	}{
		"empty": {input:"", want:0 },
		"one": {input:"abc 1,?/", want:4 },
		"two": {input:"abc 12&8 ^_", want:5 },
	}
	l:=symbolCounter{"SymbolCounter"}
	for _, v := range tests{
		t.Run(v.input, func(t *testing.T) {
		got:= l.count(v.input)
		if got!=v.want{
			t.Errorf("got %v want %v", got, v.want)
		}
		})
	
}
}

