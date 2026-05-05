
package main
import "fmt"
func main() {
	var nome string
	var idad int
	fmt.Scan(&nome, &idad)

	if idad < 12 {
		fmt.Printf("%s eh crianca\n", nome)
	} else if idad < 18 {
		fmt.Printf("%s eh jovem\n", nome)
	} else if idad < 65 {
		fmt.Printf("%s eh adulto\n", nome)
	} else if idad < 100 {
		fmt.Printf("%s eh idoso\n", nome)
	} else {
		fmt.Printf("%s eh mumia\n", nome)
	}
}
