package db
import(
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var db *sql.DB

func ConnectDB() (*sql.DB, error) {
	connStr := "postgres://turcompanydb_user:wkFLT0jvemXY9QLneq198Bf2x81QKBes@dpg-d141nqndiees73f9eukg-a.oregon-postgres.render.com:5432/turcompanydb"
	
	db, err := sql.Open("postgres", connStr)
	if err!= nil {
		return nil,err
	}
	if err = db.Ping(); err != nil {
		return nil,err
	}
	fmt.Println("DB connected successfully")
	return db, nil
}