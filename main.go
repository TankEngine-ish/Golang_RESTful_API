package main

func main() {
	app := App{}
	app.Initialise(DBUser, DBPassword, DBName)
	app.Run(":10000")
}
