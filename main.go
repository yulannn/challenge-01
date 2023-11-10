package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)


type Data struct {
	Nom string
	Filiere string
	Niveau int
	NbEtudiant int
}

type User struct {
    FirstName string
    LastName  string
	Age int
	Sexe string
}

type Profil struct {
	Users []User
}

type PageData struct {
    Data   Data
    Profil Profil
}

type DataFormUser struct{
	Nom string
	Prenom string
	Date string
	Sexe string
}

var DataForm DataFormUser = DataFormUser{}

func main() {

    temp, errTemp := template.ParseGlob("./*.html")
    if errTemp != nil {
        fmt.Println(errTemp)
        os.Exit(1)
    }


	http.HandleFunc("/promo", func(w http.ResponseWriter, r *http.Request) {
		data := Data{
			Nom:        "Mentor'ac",
			Filiere:    "Informatique",
			Niveau:     5,
			NbEtudiant: 3,
		}
	
		profil := Profil{
			Users: []User{
				{"Corentin", "DEPREZ", 19, "male"},
				{"Romain", "GOUD", 17, "femme"},
				{"Yulan", "NGUYEN",  18 ,"male"},
			},
		}
	
		pageData := PageData{
			Data:   data,
			Profil: profil,
		}


	
		temp.ExecuteTemplate(w, "index", pageData)
	})

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "form", nil)
	})

	http.HandleFunc("/user/data", func(w http.ResponseWriter, r *http.Request) {
		DataForm = DataFormUser{r.FormValue("Nom"),r.FormValue("Prenom"),r.FormValue("DateDeNaissance"),r.FormValue("Sexe")}
		fmt.Println(DataForm)
		http.Redirect(w,r,"/user/display",301)
	})
		

    fmt.Println("Serveur écoutant sur le port 8080")

    fileServer := http.FileServer(http.Dir("./asset"))
    http.Handle("/static/", http.StripPrefix("/static/", fileServer))

    http.ListenAndServe("localhost:8080", nil)
}


