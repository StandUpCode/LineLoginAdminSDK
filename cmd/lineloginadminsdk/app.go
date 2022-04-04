package cmd
import (

	"flag"


)

var (
	// Options

	flagKey     = flag.String("key", "", "path to key file or '-' to read from stdin")

	// Modes - exactly one of these is required

	flagVerify = flag.String("verify", "", "path to JWT token to verify or '-' to read from stdin")

)

func main (){

}