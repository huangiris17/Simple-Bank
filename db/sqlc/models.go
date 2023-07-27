// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package db

import (
	"time"
)

type Accounts struct {
	ID        int64     `json:"id"`
	Owner     string    `json:"owner"`
	Balance   int64     `json:"balance"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"createdAt"`
}

type Entries struct {
	ID        int64 `json:"id"`
	AccountID int64 `json:"accountId"`
	// can be negative or positive
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"createdAt"`
}

type Transfers struct {
	ID            int64 `json:"id"`
	FromAccountID int64 `json:"fromAccountId"`
	ToAccountID   int64 `json:"toAccountId"`
	// must be positive
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"createdAt"`
}

type Users struct {
	Username          string    `json:"username"`
	HashedPassword    string    `json:"hashedPassword"`
	FullName          string    `json:"fullName"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"passwordChangedAt"`
	CreatedAt         time.Time `json:"createdAt"`
}
