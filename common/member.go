/*
 *
 * In The Name of God
 *
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 01-11-2018
 * |
 * | File Name:     member.go
 * +===============================================
 */

package common

// Member represents a single contestant.
type Member struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	TShirt    string `json:"t-shirt_size"`
	StudentID string `json:"student_number"`
}
