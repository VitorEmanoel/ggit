package git

import "testing"

func TestOpen(t *testing.T) {
	repository := Open(".")
	t.Log(repository)
}

func TestInit(t *testing.T) {
	repository, err := Init(InitOptions{
		Path: "C:\\Users\\vitor\\Documents\\ggit-test",
	})
	if err != nil{
		t.Log(err)
	}
	t.Log(repository)
}
