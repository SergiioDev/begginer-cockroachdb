package main

import (
	"fmt"
	"github.com/SergiioDev/begginer-cockroachdb/authentication_service/custom_login"
	"github.com/SergiioDev/begginer-cockroachdb/authentication_service/google_login"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/auth/google", google_login.AuthGoogle)

	http.HandleFunc("/redirect", google_login.Redirect)

	http.HandleFunc("/auth/custom", custom_login.Login)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln("Unable to start the server", err)
	}

}
func index(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Login</title>
</head>
<body>
    <form action="/auth/google" method="post">
        <input type="submit" value="Login with Google">
    </form>

<form action="/auth/custom">
  <label for="email">Email:</label><br>
  <input type="text" id="email" name="email"><br>
  <label for="password">Password:</label><br>
  <input type="text" id="password" name="password" ><br><br>
  <input type="submit" value="Login">
</form> 

</body>
</html>`)

	hostname, _ := os.Hostname()
	hostname += ":8080"

	log.Println("You're talking to instance:", hostname)

	if err != nil {
		log.Fatal(err)
	}

}
