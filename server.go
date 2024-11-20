package main

import (
	"net/http"
	"net/smtp"
)

type homeHandler struct{}
type mailHandler struct{}




func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//Writes the response to the server
	w.Write([]byte("This is my home page"))

}

func (h *mailHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//Writes the response to the server
	response := sendMail()
	w.Write(response)

}

func sendMail() []byte {
	from := "your_mail_id"
	password := "my_secret_pass"

	toEmailAddress := "recipent's_mail_id"
	to := []string{toEmailAddress}

	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port

	subject := "this is the subject of the mail"
	body := "This is the body of the mail"
	message := []byte(subject + body)

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
	panic(err)
} else {
	return ([]byte("Email successfully sent!"))
}
}

func main() {
	//Take incoming requests and dispatch them to the matching handlers
	mux := http.NewServeMux()
	
	//Register the routes and handlers
	//& passes the address of the homeHandler struct as a argument ti the h parameter of type pointer
	mux.Handle("/", &homeHandler{})
	mux.Handle("/mailer", &mailHandler{})	

	//Run the server
	http.ListenAndServe(":8080", mux)
}
