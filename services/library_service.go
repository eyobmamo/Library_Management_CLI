package services

import (
	"LMS/models"
	"fmt"
)

type LibraryManager interface {
	AddBook(book models.Book) string
	RemoveBook(bookID int) error
	BorrowBook(bookID int ,memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) ([]models.Book,error)
	RegisterMember(newMember models.Member) error
}
type library_service struct {
	Library_Members  map[int]models.Member
	Library_Books   map[int]models.Book
}

func NewLibraryService() *library_service {
	return &library_service{
		Library_Members  : make(map[int]models.Member),
		Library_Books   : make(map[int]models.Book),
}
}



func (ls *library_service)AddBook(book models.Book) string{

	ls.Library_Books[book.ID]=book
	return "Book add successfully "
} 

func (ls *library_service)RemoveBook(bookID int) error {
	if _, error := ls.Library_Books[bookID]; !error{
		return fmt.Errorf("element with given id :%d not found:",bookID)
	}
	return nil
}

func (ls *library_service)BorrowBook(bookID int,memberID int) error {
	if _,ok := ls.Library_Members[memberID]; !ok {
		return fmt.Errorf("Member with given id Not Registered")

	}
	if ls.Library_Books[bookID].Status != "Available"{
		return fmt.Errorf("book is not currently available ")
	}

	member := ls.Library_Members[memberID]
	book := ls.Library_Books[bookID]

	member.BorrowedBooks = append(member.BorrowedBooks,book )
	book.Status ="Borrowed"
	fmt.Println("book borrowed successfully")

	return nil
}

func (ls *library_service)ReturnBook(bookID int,memberID int) error {
	if _,ok := ls.Library_Members[memberID]; !ok {
		return fmt.Errorf("member not register")
	}
	 
	if _,ok := ls.Library_Books[bookID]; !ok {
		return fmt.Errorf("book not register in library ")
		}
	
	userborrowed := ls.Library_Members[memberID].BorrowedBooks
	for i,book := range userborrowed {
		if book.ID == bookID {
			userborrowed = append(userborrowed[:i],userborrowed[i+1:]... )
		} 
	book := ls.Library_Books[bookID]
	book.Status ="Available"

	}
	return nil
}

func (ls *library_service)ListAvailableBooks() []models.Book {
	var AvailableBooks []models.Book
	for _, book := range ls.Library_Books {
		if book.Status == "Available" {
			AvailableBooks = append(AvailableBooks, book)
		}
	}

	return AvailableBooks
}

func (ls *library_service)ListBorrowedBooks(memberID int) ([]models.Book,error) {
	var BorrowBooks []models.Book
	if _,exist := ls.Library_Members[memberID]; !exist {
		return nil,fmt.Errorf("member not exist")
	}
	for _, book := range ls.Library_Books {
		if book.Status == "Available" {
			BorrowBooks = append(BorrowBooks, book)
		}
	}
	return BorrowBooks,nil
}

func (ls *library_service)RegisterMember(newMember models.Member) error {
	if _,ok := ls.Library_Members[newMember.ID]; ok {
		return fmt.Errorf("user register with same ID")
	} 

	ls.Library_Members[newMember.ID]=newMember
	return nil 
}
