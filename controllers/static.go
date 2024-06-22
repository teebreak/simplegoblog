package controllers

import (
	"html/template"
	"net/http"
)

func StaticHandler(tpl Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, r, nil)
	}
}

func FAQ(tpl Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   template.HTML
	}{
		{
			Question: "How old are you?",
			Answer:   "I'm 24.",
		},
		{
			Question: "Do you enjoy your job?",
			Answer:   "Yeah, pretty much.",
		},
		{
			Question: "Can I contact you?",
			Answer:   `Yeah, please do so at <a href="/contact">Contact</a> page`,
		},
		{
			Question: "Do you go to the office?",
			Answer:   "I always worked remotely.",
		},
	}

	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, r, questions)
	}
}
