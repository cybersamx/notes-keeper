package main

import (
	"github.com/cybersamx/to-do-go/app/models"
	"html/template"
	"net/http"
)

type notesTemplateData struct {
	Notes []*models.Note
}

type noteTemplateData struct {
	Note *models.Note
}

func notesHandler(app *App) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Mux supports both fix path and subtree path patterns ie. path ending w/o `/` and
		// path that ends with `/` respectively. So a single `/` will match anything not
		// handled by any handler. So we use a special guard treatment.
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		// Handle the notes page where user can create a note or view a list of notes.
		if r.Method == http.MethodGet {
			app.infoLog.Print("GET notes html page")

			notes, err := app.noteModel.GetNotes()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			data := &notesTemplateData{
				Notes: notes,
			}
			files := []string{
				"../html/notes-page.gohtml",
				"../html/notes-component.gohtml",
				"../html/base-layout.gohtml",
			}
			tpl, err := template.ParseFiles(files...)
			if err != nil {
				app.errLog.Print(err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			err = tpl.Execute(w, data)
			if err != nil {
				app.errLog.Print(err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	})
}

func parseForNote(app *App, r *http.Request) *models.Note {
	nilNote := models.Note{
		ID:    "",
		Title: "",
		Text:  "",
	}

	queryStr, ok := r.URL.Query()["noteID"]
	if ok {
		note := app.noteModel.GetNote(queryStr[0])
		if note != nil {
			return note
		}
	}

	return &nilNote
}

func editNoteHandler(app *App) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			app.infoLog.Print("GET create note html page")

			// Get param from the URL
			note := parseForNote(app, r)

			data := noteTemplateData{
				Note: note,
			}

			files := []string{
				"../html/edit-note-page.gohtml",
				"../html/base-layout.gohtml",
			}

			tpl, err := template.ParseFiles(files...)
			if err != nil {
				app.errLog.Print(err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			err = tpl.Execute(w, data)
			if err != nil {
				app.errLog.Print(err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else if r.Method == http.MethodPost {
			app.infoLog.Print("POST create note html page")

			title := r.FormValue("title")
			text := r.FormValue("text")
			noteID := r.FormValue("noteID")

			note, err := app.noteModel.Upsert(noteID, title, text)
			if err != nil {
				app.errLog.Print(err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if noteID == "" {
				app.infoLog.Print("created note with ID ", note.ID)
			} else {
				app.infoLog.Print("updated note with ID ", note.ID)
			}
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	})
}

func removeNoteHandler(app *App) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			app.infoLog.Print("POST remove note")

			noteID := r.FormValue("noteID")
			if noteID == "" {
				http.Error(w, "No noteID", http.StatusBadRequest)
				return
			}

			err := app.noteModel.RemoveNote(noteID)
			if err != nil {
				app.errLog.Print(err.Error())
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			app.infoLog.Print("deleted note with ID ", noteID)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	})
}
