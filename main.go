package main

import (
	"net/http"
	"web11/main/app"

	"github.com/urfave/negroni"
)

func main() {
	m := app.MakeHandler("./model/text.db")
	defer m.Close()
	n := negroni.Classic()
	n.UseHandler(m)

	http.ListenAndServe(":3000", n)
}
