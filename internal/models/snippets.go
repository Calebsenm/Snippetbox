package models 

import (
    "database/sql"
    "time"
    "errors"
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
   stmt := `SELECT id, title, content, created, expires FROM snippets
    WHERE expires > UTC_TIMESTAMP() AND id = ?`

    row := m.BD.QueryRow(stmt , id );
    s := &Snippet{}

    err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
    if err != nil {
        if errors.Is(err , sql.ErrNoRows ) {
            return nil , ErrNoRecord 
        }   else {
            return nil , err 
        }
    }
    return s , nil 

}

// return the 10 most recently created snippets.
func (m *SnippetModel) Latest() ([]*Snippet , error ) {

    // Write the SQL statement we want to execute.
    stmt :=  `SELECT id, title, content, created, expires FROM snippets
              WHERE expires > UTC_TIMESTAMP() ORDER BY id DESC LIMIT 10`
    
    // Use the Query() method on the connection pool to execute our
    // SQL statement. This returns a sql.Rows resultset containing the result of
    // our query.

    rows  , err := m.BD.Query(stmt);
    if err != nil {
        return nil , err 
    }

    // We defer rows.Close() to ensure the sql.Rows resultset is
    // always properly closed before the Latest() method returns. This defer
    // statement should come *after* you check for an error from the Query()
    // method. Otherwise, if Query() returns an error, you'll get a panic
    // trying to close a nil resultset

    defer rows.Close();
    
    // Initialize an empty slice to hold the snippet struct.
//    snippet := []*Snippet{}
    
    // Use rows.Next to iterate through the rows in the resultset. This
    // prepares the first (and then each subsequent) row to be acted on by the
    // rows.Scan() method. If iteration over all the rows completes then the
    // resultset automatically closes itself and frees-up the underlying
    // database connection.

    //  for rows.Next() {
    // Create a pointer to a new zeroed Snippet struct.
     //   s := &Snippet{}
//117

    //}
     return nil  , nil 
}


