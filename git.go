package git

import (
	"log"
	"os/exec"
	"reflect"
)

func Open(path string) *Repository{
	return NewRepository(path)
}

type InitOptions struct {
	Path 	string
	Bare	bool	`args:"--bare"`
}

func Init(init InitOptions) (*Repository, error){
	command := exec.Command("git", processOptions(init, "init")...)
	command.Dir = init.Path
	log.Println(command.String())
	outputBytes, err := command.Output()
	if err != nil{
		return nil, err
	}
	log.Println(string(outputBytes))
	return NewRepository(init.Path), nil

}

func processOptions(option interface{}, initalArgs ...string) []string {
	var outputArgs []string
	outputArgs = append(outputArgs, initalArgs...)
	optionType := reflect.TypeOf(option)
	for i := 0; i < optionType.NumField(); i++{
		field := optionType.Field(i)
		if arg, ok := field.Tag.Lookup("args"); ok {
			fieldValue := reflect.ValueOf(option).Field(i)
			if field.Type.Kind() == reflect.Bool{
				if fieldValue.Bool() == true{
					outputArgs = append(outputArgs, arg)
				}
			}else if field.Type.Kind() == reflect.String{
				if fieldValue.IsValid() && fieldValue.String() != ""{
					outputArgs = append(outputArgs, arg)
					outputArgs = append(outputArgs, fieldValue.String())
				}
			}
		}
	}
	return outputArgs
}