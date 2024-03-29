// package dnd is a Go Client for the DnD 5e REST API
//
//
// Installation
//
// Install with the go get command
//
//	go get github.com/brittonhayes/dnd
//
// Documentation
//
// View the full docs on pkg.go.dev https://pkg.go.dev/github.com/brittonhayes/dnd
//
// View the API here https://www.dnd5eapi.co/
//
// Usage
//
// Using the package is as easy as create client, pick the endpoint, and run the method.
// This applies across every data type so it is consistent across the board.
// Here's a simple _example of how to fetch a rule from the DnD 5e ruleset.
//
//	func main() {
//		// Create a dnd client
//		c := dnd.NewClient()
//
//		// Fetch DnD rules about adventuring
//		r, err := c.Rules.Find("adventuring")
//		if err != nil {
//			// handle error
//		}
//
//		// Print out the rule name
//		fmt.Println("Name", r.Name)
//
//		// Print out the rule description
//		fmt.Println("Description", r.Desc)
//	}
//
// Social image by Ashley Mcnamara https://twitter.com/ashleymcnamara 💖
//
// Development
//
// If you'd like to contribute to DnD, make sure you've got mage installed: https://magefile.org
//
//	# Download dependencies and run tests
//	mage download
//	mage test
//
package dnd
