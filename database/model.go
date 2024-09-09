package database

import "time"

type User struct {
	ID        uint      `json:"id" param:"id"` // id
	Name      string    `json:"name"`          // name
	CreatedAt time.Time `json:"created_at"`    // created_at
	UpdatedAt time.Time `json:"updated_at"`    // updated_at
}

type Post struct {
	ID         uint      `json:"id" param:"id"` // id
	UserName   string    `json:"user_name"`     // user_name
	Task       string    `json:"task"`          // task
	Importance uint      `json:"importance"`    // importance
	Deadline   string    `json:"deadline"`      // deadline
	// DoneAt     time.Time `json:"done_at"`       // done_at
	CreatedAt  time.Time `json:"created_at"`    // created_at
	UpdatedAt  time.Time `json:"update_at"`     // update_at
	//DeletedAt  time.Time `json:"deleted_at"`    // deleted_at
}
