package main

import (
    "github.com/cybersamx/to-do-go/app/models"
    "html/template"
    "net/http"
    "os"
    "time"
)

type notesTemplateData struct {
    Notes []*models.Note
}

// Read from an HTML template from the file system and stream the file's content to the
// HTTP response.

func outputHTML(w http.ResponseWriter, r *http.Request, filepath string) {
    file, err := os.Open(filepath)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer file.Close()

    http.ServeContent(w, r, file.Name(), time.Now(), file)
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

func createNoteHandler(app *App) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodGet {
            app.infoLog.Print("GET create note html page")

            files := []string{
                "../html/create-note-page.gohtml",
                "../html/base-layout.gohtml",
            }

            tpl, err := template.ParseFiles(files...)
            if err != nil {
                app.errLog.Print(err.Error())
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            err = tpl.Execute(w, nil)
            if err != nil {
                app.errLog.Print(err.Error())
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
        } else if r.Method == http.MethodPost {
            app.infoLog.Print("POST create note html page")

            title := r.FormValue("title")
            text := r.FormValue("text")

            note, err := app.noteModel.CreateNote(title, text)
            if err != nil {
                app.errLog.Print(err.Error())
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }

            app.infoLog.Print("Created note with ID ", note.ID)
            http.Redirect(w, r, "/", http.StatusSeeOther)
        }
    })
}
