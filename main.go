package main

import (

	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)



var topic = "test"
const broker = "localhost"

// {"status":"success","data":[{"id":"1","employee_name":"Tiger Nixon","employee_salary":"320800","employee_age":"61","profile_image":""},{"id":"2","employee_name".......},]}

type employee struct {
	Status string `json:"status"`
	Data []data   `json:"data"`
}

type data struct {

	Id string   `json:"id"`
	EmployeeName string `json:"employee_name"`
	EmployeeSalary string `json:"employee_salary"`
	EmployeeAge string `json:"employee_age"`
	ProfileImage string `json:"profile_image"`
}

func main()  {
	fmt.Println("Sending an HTTP GET request to http://dummy.restapiexample.com/api/v1/employees")
	resp,err := http.Get("http://dummy.restapiexample.com/api/v1/employees")
	if err!=nil {
		panic(err)
	}

	defer resp.Body.Close()
	b,err:=ioutil.ReadAll(resp.Body)
	if err!=nil {
		panic(err)
	}
	Produce(b)
	time.Sleep(5 * time.Second)

}