package gotml
import (
	"fmt"
	"github.com/flosch/pongo2"
)



func M(string) {
    data := make(map[string]interface{})
tpl, err := pongo2.FromFile(dirsrc)
if err != nil {
	panic(err)
}
// Now you can render the template with the given 
// pongo2.Context how often you want to.
out, err := tpl.Execute(data)
if err != nil {
	panic(err)
}
fmt.Println(out) // Output: Hello Florian!

}