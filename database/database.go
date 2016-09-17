package database

import(
  "database/sql"
  "fmt"
  "../config"
)

import _ "github.com/go-sql-driver/mysql"


// func Connect(){
//   db, err := sql.Open("mysql", "root:pass@/my_database")
// }

func GetContacts(){
  conf := config.Get().Db
  db, err := sql.Open("mysql", 
                      fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
                                  conf.User,
                                  conf.Password,
                                  conf.Host,
                                  conf.Port,
                                  conf.Schema))

  rows, err := db.Query("SELECT * FROM users")
  if err != nil {
      fmt.Println(err)
  }

  for rows.Next() {
      var id string
      var name string
      var email string
      var reg_date string
      if err := rows.Scan(&id, &name, &email, &reg_date); err != nil {
          fmt.Println(err)
      }
      fmt.Printf("%s is", name)
  }

  if err := rows.Err(); err != nil {
      fmt.Println(err)
  }
}