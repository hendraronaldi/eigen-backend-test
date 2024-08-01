package models

type Borrow struct {
	MemberID   string   `json:"member_id" schema:"member_id"`
	BookIDs    []string `json:"book_ids" schema:"book_ids"`
	BorrowedAt string   `json:"borrowed_at" schema:"borrowed_at"`
}
