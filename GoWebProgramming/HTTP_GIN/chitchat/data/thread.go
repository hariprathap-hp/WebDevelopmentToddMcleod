package data

import (
	"fmt"
	"time"
)

type Thread struct {
	Id        int
	Uuid      string
	Topic     string
	UserId    int
	CreatedAt time.Time
}

type Post struct {
	Id        int
	Uuid      string
	Body      string
	UserId    int
	ThreadId  int
	CreatedAt time.Time
}

// format the CreatedAt date to display nicely on the screen
func (thread *Thread) CreatedAtDate() string {
	return thread.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

func Threads() (threads []Thread, err error) {
	query := "select id, uuid, topic, user_id, created_at from threads order by created_at desc"
	rows, err := Db.Query(query)
	if err != nil {
		return
	}
	for rows.Next() {
		thread := Thread{}
		if err = rows.Scan(&thread.Id, &thread.Uuid, &thread.Topic, &thread.UserId, &thread.CreatedAt); err != nil {
			return
		}
		threads = append(threads, thread)
	}
	fmt.Println("Threads are -- ", threads)
	rows.Close()
	return
}

func (thread *Thread) Create() {

}

func ThreadByUUID(uuid string) (t Thread, err error) {
	fmt.Println("Querying thread -- ", uuid)
	t = Thread{}
	statement := "select id,uuid, topic,user_id,created_at from threads where uuid=$1"
	stmt, _ := Db.Prepare(statement)
	fmt.Println("statement is -- ", stmt)
	err = stmt.QueryRow(uuid).Scan(&t.Id, &t.Uuid, &t.Topic, &t.UserId, &t.CreatedAt)
	return
}
