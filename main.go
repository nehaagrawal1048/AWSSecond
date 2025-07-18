package main

import "net/http"

//fmt.Println("hello")

// // // // another way to declare
// // // var (
// // // 	name   string
// // // 	age    int
// // // 	block  string
// // // 	flatNo int
// // // )

// // // const pi = 3.14

// // //func main() {
// // //var name = "Neha" //Type inferred

// // // 	//var name string = "Neha" //long hand syntax

// // // 	//age := 30
// // // 	//earnings := 3000.098

// // // 	fmt.Print("Hello ", name)

// // // 	isGraduate := true

// // // 	isGraduate = false
// // // 	fmt.Println(isGraduate)

// // // 	//arrays, slices, maps, struct

// // // 	arr := [3]int{1, 2, 3}
// // // 	fmt.Println(arr)
// // // 	fmt.Println(arr, arr[1], arr[2])

// // // 	strs := []string{"name", "firstname"}
// // // 	fmt.Println(strs)

// // // 	strs = append(strs, "lastname") //appending the values using append function
// // // 	fmt.Println(strs)

// // // 	// slices

// // // 	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
// // // 	fmt.Println(num, num[1])

// // // 	fmt.Println("slicing", num[:5]) //excluding 5rd index
// // // 	//fmt.Println("slicing", num[:])
// // // 	//fmt.Println("slicing", num[3:])

// // // 	//maps keys values storage in go lang

// // // 	userInfo := map[string]string{
// // // 		"name": "neha",
// // // 		"age":  "12"}

// // // 	fmt.Println(userInfo, userInfo["name"], userInfo["age"])

// // // 	type clientInfo struct {
// // // 		Name       string
// // // 		Age        int
// // // 		isGraduate bool
// // // 	}

// // // 	neha := clientInfo{
// // // 		Name:       "neha",
// // // 		Age:        12,
// // // 		isGraduate: true}

// // // 	ABC := clientInfo{
// // // 		Name:       "abc",
// // // 		Age:        2,
// // // 		isGraduate: true}

// // // 	data := []clientInfo{
// // // 		ABC,
// // // 		neha}
// // // 	fmt.Println("data", data)

// // // 	// employee struct
// // // 	//salary, manager, name, designation

// // // 	type employee struct {
// // // 		Name        string
// // // 		Designation string
// // // 		Manager     string
// // // 		Salary      int
// // // 	}

// // // 	Neha := employee{
// // // 		Name:        "Neha",
// // // 		Designation: "Engineer",
// // // 		Manager:     "ABC",
// // // 		Salary:      10000}

// // // 	Akash := employee{
// // // 		Name:        "Akash",
// // // 		Designation: "Manager",
// // // 		Manager:     "ABC",
// // // 		Salary:      50000}

// // // 	Details := []employee{
// // // 		Neha,
// // // 		Akash}

// // // 	fmt.Println("Details:", Details, Akash.Designation)

// // // 	desig := Akash.Designation
// // // 	fmt.Println(desig)

// // // 	addressOfDesig := &Akash.Designation
// // // 	fmt.Println(Akash.Designation, addressOfDesig)

// // // 	designation(addressOfDesig)
// // // 	fmt.Println(Akash.Designation)
// // // }

// // // func designation(desig *string) {
// // // 	*desig = "PO"
// // // }

// // // 	//loops

// // // 	nums := []int{1, 2, 3, 4}
// // // 	for i := 0; i < len(nums); i++ {
// // // 		fmt.Println(i, "index")
// // // 		fmt.Println(nums[i], "value")
// // // 	}

// // // 	// for i, v := range nums {
// // // 	// } //range = built in function to iterate over the values
// // // 	// fmt.Println((i))

// // // 	// for _, v := range nums {
// // // 	// } //range = built in function to iterate over the values
// // // 	// //fmt.Println(("index", i ))

// // // 	// //func greetUser() {
// // // 	// //fmt.Println("Hello Neha")
// // // 	// //}

// // // 	// //func greetUser1(name string) {
// // // 	// fmt.Println("Hello", name)

// // // 	// greetUser1("Neha") // named functions

// // // 	// }

// // // 	func(name string) {
// // // 		fmt.Println("Hi How r u?")
// // // 	}("Agrawal") //anonymous functions calling itself with no name

// // // 	greetUse("Neha", 30)

// // // }

// // // func greetUse(name string, age int) {
// // // 	fmt.Println("Hello", name, age)

// // // 	//you cant call named functions from within the main function only anonymous function u can call from

// // // 	utility.SayHiToUser("Neha")

// // // name:= "Neha"
// // // addressOfName := &name

// // // //greetUser2(name string){
// // // 	address := &name
// // // //}

// // // func greetUser(name *string){
// // // 	*name = "neha"
// // // }

// // // total: = sum(1,2)

// // // func sum(num1 int, num2 int) int {
// // // 	return num1 + num2
// // // }

// // //goROutines are light weight parallel threads
// // //when we put go in front of any function it becomes go routine and runs parallely and not synchronously

// // // 	var wg sync.WaitGroup

// // // 	go func(){
// // // 		defer wg.Done()
// // // 		wg.Add(1)
// // // 		fmt.Println(("all go routines finished"))
// // // time.sleep(2 * time.Second)

// // // 	}(&wg)

// // // 		wg.Wait()

// // // 		fmt.Println("all programs finished")
// // // 	}
// // //Add
// // //Done
// // //Wait

// // // 	wg.Add(5) //start with 5 go routines

// // // 	go greetUser("A", &wg) //4
// // // 	//wg.Done()

// // // 	go greetUser("B", &wg)
// // // 	go greetUser("C", &wg)
// // // 	go greetUser("D", &wg)
// // // 	go greetUser("E", &wg)

// // // 	wg.Wait() //blocking //it will wait till the counter reaches to 0
// // // 	time.Sleep(4 * time.Second)

// // // 	fmt.Println("Finished Exeution")
// // // }

// // // func greetUser(name string, wg *sync.WaitGroup) {
// // // 	fmt.Println("hello", name)
// // // 	wg.Done()
// // // }

// // //channels

// // // greeting := make(chan string)

// // // go func(){
// // // 	time.sleep(2 * time.Second)
// // // 	greeting <- "pratik" //push the data in greeting chanel

// // // }()

// // // <-greeting // recieving the data that is blocking in nature

// // // 	greeting := make(chan string)

// // // 	go func() {
// // // 		time.Sleep(4 * time.Second)

// // // 		greeting <- "Neha"
// // // 	}()

// // // 	fmt.Println("hello finished", <-greeting)
// // // 	fmt.Println("all finished")
// // // }

// // // // type struct user{
// // // // 	Name string
// // // // 	Age integer

// // // type user struct {
// // // 	Name string
// // // 	Age int
// // // }

// // // func main(){
// // // 	data := []byte()
// // // 	employee := user{

// // // 		json.Unmarshal(data, &employee)
// // // 	}
// // // }

// // // rest api -> POST, name -> username, email
// // // 	data := []byte(`{"user_name": "Neha", "user_age": 12, "office_location": "delhi"}`) //mocking api

// // // 	employee := user{}

// // // 	err := json.Unmarshal(data, &employee)
// // // 	if err != nil {
// // // 		fmt.Println("Error in Decoding", err)
// // // 		panic("were in error")
// // // 	}

// // // 	fmt.Println("employee", employee)

// // //initialize the server

// // // type user struct {
// // // 	Name           string `json:"user_name"`
// // // 	Age            int    `json:"user_age"`
// // // 	OfficeLocation string `json:"office_location"`

// // // type UserInfo struct {
// // // Name string `json:"name"`
// // // Email string `json:"email"`
// // // Age int `json:"age"`
// // // }
// // // var allUsers []UserInfo
// // // func main() {
// // // // initialise a server
// // // routeHandler := http.NewServeMux()
// // // // create a route handler
// // // routeHandler.HandleFunc("POST /user", addUser)
// // // routeHandler.HandleFunc("GET /user", getUser)
// // // // start the server on particular port
// // // http.ListenAndServe(":8050", routeHandler)
// // // }
// // // func addUser(w http.ResponseWriter, r *http.Request) {
// // //  // w.WriteHeader(http.StatusAccepted) // we are modifying the response data set
// // // currentUser := UserInfo{}
// // // // decode the req body into currentuser
// // // json.NewDecoder(r.Body).Decode(&currentUser)
// // // if currentUser.Name == "" {
// // // w.WriteHeader(http.StatusBadRequest)
// // // w.Write([]byte("Please insert username"))
// // // return
// // //  }
// // // allUsers = append(allUsers, currentUser)
// // // w.Write([]byte("Data inserted successfully"))
// // // }
// // // func getUser(w http.ResponseWriter, r *http.Request) {
// // // // w.Header().Set("Content-Type", "application/json")
// // // json.NewEncoder(w).Encode(allUsers)
// // // }

// // import (
// // 	"encoding/json"
// // 	"net/http"
// // )

// // // CRUD -> Memory,
// // // Create User -> POST
// // // EC2 instance -> public ip address

// // type UserInfo struct {
// // 	Name  string `json:"name"`
// // 	Email string `json:"email"`
// // 	Age   int    `json:"age"`
// // }

// // var allUsers []UserInfo

// // func main() {
// // 	// initialise a server
// // 	routeHandler := http.NewServeMux()

// // 	// create a route handler
// // 	routeHandler.HandleFunc("POST /user", addUser)
// // 	routeHandler.HandleFunc("GET /user", getUser)
// // 	routeHandler.HandleFunc("DELETE /user", deleteUser)

// // 	// start the server on particular port
// // 	http.ListenAndServe(":8050", routeHandler)
// // }

// // func deleteUser(w http.ResponseWriter, r *http.Request) {
// // 	allUsers = make([]UserInfo, 0)
// // 	w.Header().Set("Content-Type", "application/json")
// // 	json.NewEncoder(w).Encode(allUsers)
// // }

// // func addUser(w http.ResponseWriter, r *http.Request) {
// // 	// w.WriteHeader(http.StatusAccepted) // we are modifying the response data set
// // 	currentUser := UserInfo{}

// // 	// decode the req body into currentuser
// // 	json.NewDecoder(r.Body).Decode(&currentUser)

// // 	if currentUser.Name == "" {
// // 		w.WriteHeader(http.StatusBadRequest)
// // 		w.Write([]byte("Please insert username"))
// // 		return
// // 	}
// // 	allUsers = append(allUsers, currentUser)
// // 	w.Write([]byte("Data inserted successfully"))
// // }

// // func getUser(w http.ResponseWriter, r *http.Request) {
// // 	// w.Header().Set("Content-Type", "application/json")
// // 	json.NewEncoder(w).Encode(allUsers)
// // }

// package main

// import (
// 	"fmt"
// 	"nehaagrawal148/First/firstcontainer"
// )

// type Rectangle struct {
// 	Length int
// 	Width  int
// }

// func main() {
// 	rec := Rectangle{Length: 10, Width: 10}
// 	firstcontainer.Area(rec)
// 	fmt.Printf("the Length of the rectangle %d\n", rec.Length)
// }

func main() {

	// type NehaStruct struct {
	// 	Fname string
	// 	Lname string
	// }

	// MyChannel := make(chan NehaStruct)

	// go func() {

	// 	Neha := NehaStruct{Fname: "PING", Lname: "PONG"}

	// 	MyChannel <- Neha
	// }()

	// fmt.Println(<-MyChannel)

	// 	router := http.NewServeMux()
	// 	router.HandleFunc("GET /greet", greetUser)

	// 	http.ListenAndServe(":7080", router)
	// }

	// func greetUser(w http.ResponseWriter, req *http.Request) {
	// 	w.Write([]byte("Welcome user"))

	// myChannel := make(chan string)

	// go func() {
	// 	for {
	// 		myChannel <- "ping"
	// 		fmt.Println("hello", "ping")
	// 		time.Sleep(2 * time.Second)
	// 		<-myChannel
	// 		fmt.Println("hello", "pong")
	// 		time.Sleep(2 * time.Second)

	// 	}

	// }()
	// time.Sleep(5 * time.Second)

	router := http.NewServeMux()
	router.HandleFunc("GET /greet", greetUser)
	http.ListenAndServe(":7080", router)
}
func greetUser(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Welcome user!!"))

}
