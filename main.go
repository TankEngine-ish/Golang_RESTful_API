package main

func main() {
	app := App{}
	app.Initialise()
	app.Run(":10000")
}
