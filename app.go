package main

import (
	"log"
	"os"
	"net/http"
	"encoding/json"
	"net/smtp"
)

type resp struct{
	Message string
}
type data struct{
	Nombre string `json:"nombre"`
	Email string `json:"email"`
	Nkids string `json:"nkids"`
	Nrooms string `json:"nrooms"`
	Descripcion string `json:"descripcion"`
}

func send_resp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	out := &resp{
		Message : "Hemos recibido sus datos.<br/>Pronto nos pondremos en contacto con vosotros."}
	w.WriteHeader(http.StatusOK)
	js, err := json.Marshal(out)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(js)
}

func send_email(body data) error{
	from := os.Getenv("EMAIL_FROM")
	pass := os.Getenv("EMAIL_PASS")
	to := os.Getenv("EMAIL_TO")

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Contacto desde todosairlanda.com\n\n" +
		"Nombre: " + body.Nombre +"\n\n" +
		"Email: " + body.Email+"\n\n" +
		"Numero de niños: " + body.Nkids+"\n\n" +
		"Numero de habitaciones: " + body.Nrooms+"\n\n" +
		"Descripcion: " + body.Descripcion 


	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return err
	}

	return nil
}

func Contacta(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST"{
		if r.Body == nil {
			http.Error(w, "Necesitamos a request body", http.StatusBadRequest)
			return	
		}
		var d data
		err := json.NewDecoder(r.Body).Decode(&d)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = send_email(d)
		if err != nil {
			http.Error(w, "Error enviando datos del formulario.", http.StatusBadRequest)
			return	
		} else {
			send_resp(w,r)
		}
		
	}
}

func main() {
	fs := http.FileServer(http.Dir("dist"))
	http.Handle("/", fs)
	http.HandleFunc("/contacta", Contacta)

	log.Println("Listening...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Err: ", err)
	}
}
