package main
import (
	"net"
	"log"
)

func main(){
	s:= newServer()
	go s.run()
	listner, err := net.Listen("tcp",":8000")
    if err != nil {
		log.Fatalf("unable to connect server")
	}
	defer listner.Close()
	log.Println("listining on port 8000")

	for {
		conn, err := listner.Accept()
		if err != nil {
        log.Printf("unable to accept connection :", err.Error())
		continue	
		}
		go s.newClient(conn)
		
	}
    
}