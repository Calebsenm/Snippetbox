package main;

import (
    "log"
    "net/http"
    "flag"
    "os"
     _ "github.com/go-sql-driver/mysql"
    "database/sql"    
    "github.com/calebsenm/snippetbox/internal/models"

)

type application struct {
    errorLog    *log.Logger 
    infoLog     *log.Logger 
    snippets    *models.SnippetModel 
}


func main(){
    
    addr := flag.String("addr", ":4000", "HTTP network address");
    dns := flag.String("dns", "root:example@(127.0.0.1:3306)/snippetbox?parseTime=true", "MySQL data source name");

    flag.Parse()

    infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
    errorLog := log.New(os.Stderr , "ERROR\t", log.Ldate | log.Ltime | log.Lshortfile );


    db , err := openDB(*dns);
    if err != nil {
        errorLog.Fatal(err);
    }

    defer db.Close();

    //Dependencies 
    app := &application{
        errorLog: errorLog,
        infoLog: infoLog,
        snippets: &models.SnippetModel{BD:db},
    }


    srv := &http.Server{
        Addr:       *addr,
        ErrorLog:   errorLog,
        Handler:    app.routes(),
    }
    
    infoLog.Printf("Staring Server on %s", *addr );
    err = srv.ListenAndServe();
    errorLog.Fatal(err);
}

func openDB(dsn string) (*sql.DB, error) {
    db, err := sql.Open("mysql", dsn)
    
    if err != nil {
        return nil, err
    }
    
    if err = db.Ping(); err != nil {
        return nil, err
    }
    return db, nil
}

