package models 

import (
    "database/sql"
    "time" 
)

type Snippet struct {
    ID      int 
    Title   string 
    Content string 
    Created time.Time 
    Expires time.Time 
}


type SnippetModel struct {
    BD *sql.DB 
}

// Insert a new Snippet into the database 
func (m *SnippetModel) Insert(title string , content string , expires int ) ( int , error ){
    stmt := `INSERT INTO snippets (title, content, created, expires)
             VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`
 

    result , err := m.BD.Exec(stmt , title , content, expires);
    if err  != nil{
        return 0 , err 
    }
    
    id , err := result.LastInsertId();
    if err != nil {
        return 0 , err  
    }
    return int(id), nil   
}

// return a specific snippet based on its id 
func (m *SnippetModel ) Get(id int) (*Snippet , error){
   return nil , nil  
}

// return the 10 most recently created snippets.
func (m *SnippetModel) Latest() ([]*Snippet , error ) {
    return nil , nil 
}