
import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/wybiral/torgo"
)

func HiddenService() {
	// Setup handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world!")
	})
	// Create local listener on next available port
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	// Get listener port
	port := listener.Addr().(*net.TCPAddr).Port
	addr := fmt.Sprintf("127.0.0.1:%d", port)
	// Create controller
	controller, err := torgo.NewController(ControllerAddr)
	if err != nil {
		log.Fatal(err)
	}
	// Authenticate to controller using filesystem cookie
	// You may need to change this depending on your torrc configuration
	err = controller.AuthenticateCookie()
	if err != nil {
		log.Fatal(err)
	}
	// Configure onion to route hidden service port 80 to server address
	onion := &torgo.Onion{
		Ports:          map[int]string{80: addr},
	}
	// Start the hidden service
	err = controller.AddOnion(onion)
	if err != nil {
		log.Fatal(err)
	}
	// Print newly created onion address
	fmt.Println("Local port is", port)
	fmt.Println("Serving at http://" + onion.ServiceID + ".onion")
	// Start serving
	http.Serve(listener, nil)
}