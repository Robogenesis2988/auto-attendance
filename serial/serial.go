package serial

import (
	"log"
	"runtime"
	"strings"

	"go.bug.st/serial"
	"go.bug.st/serial/enumerator"
)

/*
If on linux, try:

sudo usermod -a -G dialout $USER

*/
func GetUSBPorts() []string {
	details, err := enumerator.GetDetailedPortsList()
	must(err, "error getting detailed ports list")
	var ports []string
	for _, port := range details {
		// log.Printf("Found port: %v\n", *port)
		if port.IsUSB {
			// log.Printf("   USB ID     %s:%s\n", port.VID, port.PID)
			// log.Printf("   USB serial %s\n", port.SerialNumber)
			ports = append(ports, port.Name)
		}
	}

	return ports
}

func Open(portName string, baud int) serial.Port {
	mode := &serial.Mode{BaudRate: baud}
	port, err := serial.Open(portName, mode)
	var help string
	if runtime.GOOS == "linux" {
		help = "try running: sudo usermod -a -G dialout $USER\n"
	}
	must(err, "there was a problem opening port "+portName+"\n"+help)
	return port
}

func ListenForData(port *serial.Port, c chan string) {
	buff := make([]byte, 10)
	for {
		n, err := (*port).Read(buff)
		if err != nil {
			log.Fatal(err)
			break
		}
		if n == 0 {
			// log.Println("\nEOF")
			break
		}
		// fmt.Printf("Data:%v", string(buff[:n]))
		// If we receive a newline stop reading
		if strings.Contains(string(buff[:n]), "\n") {
			break
		}
	}
	value := strings.TrimSpace(string(buff))
	// log.Printf("value: hi%vhi\n", value)
	c <- value
	close(c)
}

func should(err error, msg string) {
	if err := recover(); err != nil {
		log.Fatalf("%v:\n%v", msg, err)
	}
}
func must(err error, msg string) {
	if err != nil {
		log.Fatalf("%v:\n%v", msg, err)
	}
}
