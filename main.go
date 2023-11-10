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
	/* Message string
	   ViewCount int */
}

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

		/*if pageData.ViewCount%2 == 0 {
			pageData.Message = "Le nombre de vues est pair."
		} else {
			pageData.Message = "Le nombre de vues est impair."
		} */

	
		temp.ExecuteTemplate(w, "index", pageData)
	})

    fmt.Println("Serveur écoutant sur le port 8080")

    fileServer := http.FileServer(http.Dir("./asset"))
    http.Handle("/static/", http.StripPrefix("/static/", fileServer))

    http.ListenAndServe("localhost:8080", nil)
}


