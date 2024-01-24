package omniServer

import ( 
	"database/sql"
	"strings"
)

func BuildParamaterizedPSQLString(querytemplate,input string) (string,error){
	builder := strings.Builder{}

	return result, nil
}

func BinarySearchTheDatabase(target string) {


}

func GetTotalRows(sqlLanguage, tablename string) (int, error) {
	var totalRows int
	var stmt *sql.Stmt
	var err error
	var query string

// use local builder not the function

	switch sqlLanguage {
	case "mysql":
		query = BuildParameterizedPSQLString("SHOW TABLE STATUS WHERE NAME = ", tablename)
		stmt, err = db.Prepare(query)
	case "postgres":
		
		query = BuildParameterizedPSQLString(`SELECT table_name, table_type, row_count FROM information_schema.tables WHERE table_type = 'BASE TABLE'`)

		stmt, err = db.Prepare(query)
	case "sqlite":
		
		query = BuildParameterizedPSQLString("PRAGMA table_info(table_name)")

		stmt, err = db.Prepare(query)
	case "oracle":
		query = BuildParameterizedPSQLString("SELECT table_name, num_rows FROM user_tables")

		stmt, err = db.Prepare(query)
	case "sqlserver":
		query = BuildParameterizedPSQLString(`SELECT table_name, table_type, row_count FROM information_schema.tables WHERE table_type = 'BASE TABLE'`)

		stmt, err = db.Prepare(query)
	default:
		log.Fatal("Unsupported database type")
	}

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Total rows:", totalRows)
	return totalRows
}


stmt, err := db.Prepare("SELECT * FROM users WHERE id = ?")
if err != nil {
   log.Fatal(err)
}

rows, err := stmt.Query(userId)
if err != nil {
   log.Fatal(err)
}
defer rows.Close()

for rows.Next() {
   var user User
   err = rows.Scan(&user.ID, &user.Name)
   if err != nil {
       log.Fatal(err)
   }
   fmt.Println(user)
}

err = rows.Err()
if err != nil {
   log.Fatal(err)
}

