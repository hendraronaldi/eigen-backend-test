package services

import (
	"app-library/internal/models"
	"app-library/internal/repositories"
	"context"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LibraryInterface interface {
	GetAllBooks(ctx context.Context) ([]*models.Book, error)
	GetAllMembers(ctx context.Context) ([]*models.Member, error)
	BorrowBook(ctx context.Context, p *models.Borrow) error
	ReturnBook(ctx context.Context, p *models.Return) error
}

type Library struct {
	libraryRepo repositories.Interface
}

func NewLibrary(l repositories.Interface) *Library {
	return &Library{libraryRepo: l}
}

func (lb *Library) GetAllMembers(ctx context.Context) ([]*models.Member, error) {
	rs, err := lb.libraryRepo.FindAllMembers(ctx)
	for _, m := range rs {
		m.MemberID = m.ID.Hex()
		m.TotalBorrowedBooks = len(m.BoorowedBooks)
	}
	return rs, err
}

func (lb *Library) GetAllBooks(ctx context.Context) ([]*models.Book, error) {
	r, err := lb.libraryRepo.FindAllBooks(ctx)
	return r, err
}

func (lb *Library) BorrowBook(ctx context.Context, p *models.Borrow) error {
	oid, err := primitive.ObjectIDFromHex(p.MemberID)
	if err != nil {
		return errors.Wrap(err, "invalid member id")
	}

	m, err := lb.libraryRepo.FindOneMember(ctx, oid)
	if err != nil {
		return errors.Wrap(err, "find member")
	}

	if len(m.BoorowedBooks) >= 2 {
		return errors.New("member already borrowed max number of books")
	}

	borrowedAt, err := time.Parse("2006-01-02 15:04:05", p.BorrowedAt)
	if err != nil {
		return errors.Wrap(err, "invalid borrowed at date")
	}
	if m.PenalizedAt != nil {
		days := int(borrowedAt.Sub(*m.PenalizedAt).Hours() / 24)
		if days <= 3 {
			return errors.New("member currently penalized")
		}
	}

	var bs []*models.Book
	for _, bookID := range p.BookIDs {
		bid, err := primitive.ObjectIDFromHex(bookID)
		if err != nil {
			return errors.Wrap(err, "invalid book id")
		}

		b, err := lb.libraryRepo.FindOneBook(ctx, bid)
		if err != nil {
			return errors.Wrap(err, "find book")
		}

		b.BorrowedAt = &borrowedAt
		bs = append(bs, b)
	}
	m.BoorowedBooks = bs

	if err := lb.libraryRepo.UpdateMember(ctx, m); err != nil {
		return errors.Wrap(err, "update member")
	}

	for _, b := range bs {
		b.IsBorrowed = true
		if err := lb.libraryRepo.UpdateBook(ctx, b); err != nil {
			return errors.Wrap(err, "update book status borrowed")
		}
	}
	return nil
}

func (lb *Library) ReturnBook(ctx context.Context, p *models.Return) error {
	oid, err := primitive.ObjectIDFromHex(p.MemberID)
	if err != nil {
		return errors.Wrap(err, "invalid member id")
	}

	m, err := lb.libraryRepo.FindOneMember(ctx, oid)
	if err != nil {
		return errors.Wrap(err, "find member")
	}

	if len(m.BoorowedBooks) == 0 {
		return errors.New("member already return all the borrowed books")
	}

	mb := make(map[string]*models.Book)
	for _, b := range m.BoorowedBooks {
		mb[b.ID.Hex()] = b
	}

	for _, bid := range p.BookIDs {
		if _, ok := mb[bid]; !ok {
			return errors.New("member cannot return not own borrowed books")
		}
	}

	rb := make(map[string]bool)
	for _, bid := range p.BookIDs {
		rb[bid] = true
	}

	returnedAt, err := time.Parse("2006-01-02 15:04:05", p.ReturnedAt)
	if err != nil {
		return errors.Wrap(err, "invalid returned at date")
	}

	stillBorrowedBooks := []*models.Book{}
	for _, b := range m.BoorowedBooks {
		if _, ok := rb[b.ID.Hex()]; !ok {
			stillBorrowedBooks = append(stillBorrowedBooks, b)
			continue
		}

		b.IsBorrowed = false
		days := int(returnedAt.Sub(*b.BorrowedAt).Hours() / 24)
		if days > 7 {
			m.PenalizedAt = &returnedAt
		}
		b.BorrowedAt = nil
		if err := lb.libraryRepo.UpdateBook(ctx, b); err != nil {
			return errors.Wrap(err, "update book")
		}

	}

	m.BoorowedBooks = stillBorrowedBooks
	if err := lb.libraryRepo.UpdateMember(ctx, m); err != nil {
		return errors.Wrap(err, "update member")
	}
	return nil
}
