package main

//🔹 Задача: Гостевая книга
//Создай гостевую книгу, куда пользователи могут оставлять сообщения, а затем видеть список всех оставленных сообщений.
//
//📌 Функционал:
//Пользователь вводит имя и сообщение.
//Нажимает кнопку "Оставить сообщение".
//Сообщение сохраняется и отображается на странице.
//Все сообщения видны другим пользователям.
//Добавь кнопку "Очистить все сообщения", которая удаляет все записи.

import (
	"fmt"
	"net/http"
	"text/template"
)

var m = make(map[string]string)

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")

	if err != nil {
		fmt.Println("Loading error", err)
	}

	tmpl.Execute(w, m)
}

func addMesagge(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		name := r.Form.Get("name")
		message := r.Form.Get("message")

		if name != "" && message != "" {
			m[name] = message
			fmt.Println(m)
		} else {
			fmt.Fprintln(w, "Add message or name!")
		}
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func deleteMesagge(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		clear(m)
		fmt.Println(m)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func main() {

	http.HandleFunc("/", handler)
	http.HandleFunc("/send", addMesagge)
	http.HandleFunc("/clear", deleteMesagge)

	fmt.Println("Listening on port 8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Download error", err)
	}
}
