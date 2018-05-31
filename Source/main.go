package main;

import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "log"
//    "syscall"
    "time"
    "strconv"
)
var port_r = make(chan int)
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "<h1>Las World!</h1>\n")
}
func Server(){
    var arr = [13][13]int{}
    if(<-port_r == -1){
        port_r<-8080
    }

    router := httprouter.New()
    router.GET("/", Index)
    router.GET("/A",func (w http.ResponseWriter, r *http.Request, _ httprouter.Params){
        fmt.Fprint(w, "<h1>Las World!</h1>\n")
        fmt.Fprint(w, "",arr)
    })
    router.GET("/A/:X",func (w http.ResponseWriter, r *http.Request, ps httprouter.Params){
        I := ps.ByName("X")
        X, _ := strconv.Atoi(I[0:2])
        Y, _ := strconv.Atoi(I[2:4])
        V, _ := strconv.Atoi(I[4:])
        arr[X][Y]=V
        fmt.Fprint(w, "X :", X, "\nY :", Y, "\nV :", V)
    })
    fmt.Println("port is",<-port_r)
    log.Fatal(http.ListenAndServe(":"+strconv.Itoa(<-port_r), router))
}
func Start(){
    fmt.Println("Start Successfully")
    fmt.Println("Las Server Controller Start~!")
}
func Method(command string, s chan int){
    if command == "end"{    
        s<-0
    }else if command == "state"{
        fmt.Println("State Ok....")
    }else if command == "version" || command == "ver"{
        fmt.Println("Server Version is 0.3\r\nAuther is Wonho Ha")
    }else if command == "start"{
        fmt.Println("Server Start")
        go Server()
    }else if command == "help"{
        help()
    }else if command[:4] == "port"{
        port_r<-2
    }else {
        fmt.Println("-->wrong command")
        fmt.Println("-->help command will help you")
    }
}
func help(){
    fmt.Println("This Program is Wonho's Server Controller")
    fmt.Println("Used Golang, Source is https://github.com/Las-Wonho/GoToServer")
    fmt.Println("Server Port was 8080")
    fmt.Println("Command list")
    fmt.Println("    ->help")
    
    fmt.Println("    ->state")
    
    fmt.Println("    ->start")
    
    fmt.Println("    ->version or ver")
    
    fmt.Println("    ->end")
    
    fmt.Println("    ->port")
}
func CLI_io(state chan int){
    var command string
    for {
        fmt.Print(">")
        fmt.Scanln(&command)
        Method(command, state)
    }
}
func Init(){
}
func main() {

    go Init()
    Start()
    state := make(chan int,0)

    go CLI_io(state)
    for <-state != 0{
        time.Sleep(time.Second)
    }
    fmt.Println("Sysyem End")
}
