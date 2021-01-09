// package dnd is a Go Client for the DnD 5e REST API
//
// Installation
//
//	go get github.com/brittonhayes/dnd
//
// Usage
//
//	func main() {
//		// Create a dnd client
//		c := dnd.NewClient()
//
//		// Fetch DnD rules about adventuring
//		r, err := c.Rules.FindRule("adventuring")
//		if err != nil {
//			panic(err)
//		}
//
//		// Print out the rule name
//		fmt.Println("Name", r.Name)
//
//		// Print out the rule description
//		fmt.Println("Description", r.Desc)
//	}
//
package dnd
