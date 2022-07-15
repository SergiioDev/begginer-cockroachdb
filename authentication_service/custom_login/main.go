package custom_login

import (
	"fmt"
	"log"
	"net/http"
)

type User struct {
}

func Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	email := r.FormValue("email")

	_, err = fmt.Fprint(w, "Welcome,", email)
	if err != nil {
		log.Println(err)
	}

}
