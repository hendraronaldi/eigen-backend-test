package models

type Return struct {
	MemberID   string   `json:"member_id" schema:"member_id"`
	BookIDs    []string `json:"book_ids" schema:"book_ids"`
	ReturnedAt string   `json:"returned_at" schema:"returned_at"`
}
