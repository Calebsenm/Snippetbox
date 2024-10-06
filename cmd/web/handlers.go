package main; 

import (
    "fmt"
    "net/http"
    "strconv"
    "html/template"
    "github.com/calebsenm/snippetbox/internal/models"
    "errors"
)

func (app *application) home(w http.ResponseWriter, r *http.Request ){
    if r.URL.Path != "/" {
        http.NotFound(w , r );
        return 
    }

    snippets , err := app.snippets.Latest();
    if err != nil {
        app.serveError(w, err)
        return 
    }

    for _ , snippet := range snippets {
        fmt.Fprintf(w, "%+v\n", snippet);
    }

/*
    files := []string{
        "./ui/html/base.tmpl",
        "./ui/html/pages/home.tmpl",
        "./ui/html/partials/nav.tmpl",
    }

    ts , err := template.ParseFiles(files...);
    
    if err != nil {
        app.serveError(w , err )     
        return 
    }
    
    err = ts.ExecuteTemplate(w , "base", nil )
    if err != nil {
        app.serveError(w , err )  
    }
    */
}

func (app *application) snippetView(w http.ResponseWriter , r *http.Request ){
    id , err := strconv.Atoi(r.URL.Query().Get("id"));
    
    if err != nil || id < 1 {
        app.notFound(w)
        return 
    }
    
    snippet , err := app.snippets.Get(id);
    
    if err != nil {
         if errors.Is(err, models.ErrNoRecord) {
            app.notFound(w)
        } else {
            app.serveError(w, err)
        }
        return
    }

    // Initialize a slice containing the paths to the view.tmpl file,
    // plus the base layout and navigation partial that we made earlier.
    files := []string{
        "./ui/html/base.tmpl",
        "./ui/html/partials/nav.tmpl",
        "./ui/html/pages/view.tmpl",
    }
    
    ts , err := template.ParseFiles(files...)
    if err != nil {
        app.serveError(w , err)    
        return 
    }
    // And then execute them. Notice how we are passing in the snippet
    // data (a models.Snippet struct) as the final parameter?

    err = ts.ExecuteTemplate(w, "base", snippet);
    
    if err != nil {
         app.serveError(w, err)
    }
}

func (app *application ) snippetCreate(w http.ResponseWriter , r *http.Request){
    if r.Method != http.MethodPost{
        w.Header().Set("Allow", http.MethodPost );
        app.clientError(w , http.StatusMethodNotAllowed)
        return 
    }
   
    title := "LOL"
    content := "O LOL oooo LOL"
    expires := 7

    id , err := app.snippets.Insert(title , content , expires );
    if err != nil {
        app.serveError(w , err );
        return 
    }    
    
    http.Redirect(w , r ,fmt.Sprintf("/snippet/view?id=%d" , id),http.StatusSeeOther)
}

