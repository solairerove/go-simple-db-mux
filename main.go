package main

func main() {
	a := App{}
	a.Initialize("griotrard_user", "griotrard_pass", "griotrard")

	a.Run(":8000")
}
