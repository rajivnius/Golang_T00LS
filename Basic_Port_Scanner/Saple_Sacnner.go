package main
import(
	"fmt"
	"net"
	"sync"
)
func main(){
	var wg sync.WaitGroup // initialize wait Group 
	for i := 1; i <=10000; i++ {
		wg.Add(1) // increment wait group counter
		go func(p int){
			defer wg.Done() // decrement wait group counter when the goroutine completes
			addr := fmt.Sprintf("scanme.nmap.org:%d", p)
			conn,err := net.Dial("tcp", addr) // net.Dial use three way handshake to check if port is open deeply scan 
			if err == nil {
				fmt.Printf("Port %d is open\n", p)
				conn.Close()
			}
		}(i) // pass port number as argument to avoid closure capture issue
	}
	wg.Wait() // wait for all goroutines to finish
	fmt.Println("Scan Finished")
}
