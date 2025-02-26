package controllers

import (
	"LMS/models"
	"LMS/services"
	"strings"

	// "LMS/services"
	"LMS/utils"
	"fmt"
)



type library_service_control struct{
	library_service services.LibraryManager
}

func NewLibraryServiceControl (library_service services.LibraryManager) *library_service_control {
	return &library_service_control{
		library_service: library_service,
	}
}

func (LSC *library_service_control)Library_controller() {
	for {fmt.Println("Library services")

	fmt.Println(" 1. Register Member")
	fmt.Println(" 2. Add a New Book")
	fmt.Println(" 3. Borrow a Book")
	fmt.Println(" 4. Return a Book")
	fmt.Println(" 5. List All Available Books")
	fmt.Println(" 6. List all Borrowed books by a Member")
	fmt.Println(" 7. Remove an Exixting book")

    fmt.Println("Enter your choice:")
	choice := utils.Readline()
	choice = strings.TrimSpace(choice)

	fmt.Println(choice)

	switch choice {
		case "1":
			LSC.Register_member_controller()
		case "2":
			LSC.Add_new_book_controller()
		case "3":
			LSC.Borrow_a_Book_controller()
		case "4":
			LSC.Retern_Book_controller()
		case "5":
			LSC.List_all_the_available_Book_controller()
		case "6":
			LSC.Borrowed_book_controller()
		case "7":
			LSC.Remove_Book_controller()	
		
		default :
			fmt.Println("Invalid choice choose correctly")
			

	}
	utils.Pause()
	utils.ClearScreen()
	LSC.Library_controller()
}
}

func (LSC *library_service_control) Add_new_book_controller(){
	fmt.Println("Add new book")
	var newBook models.Book

	fmt.Println("Enter of the Book: ")
	newBook.Title = utils.Readline()

	fmt.Println("Enter of Author: ")
	newBook.Author = utils.Readline()

	fmt.Println("Enter of ID: ")
	newBook.ID = utils.Readint()

	newBook.Status = "Available" 

	LSC.library_service.AddBook(newBook)
}

func (LSC *library_service_control) Register_member_controller(){
	fmt.Println("Register to as New member:")
	var newMember models.Member

	fmt.Println("Enter Name: ")
	newMember.Name = utils.Readline()

	fmt.Println("Enter Member ID: ")
	newMember.ID = utils.Readint()

	LSC.library_service.RegisterMember(newMember)
}

func  (LSC *library_service_control) Borrow_a_Book_controller(){
	fmt.Println("borrow_book_service")

	fmt.Println("Enter the ID of the book")
	var searched_bookID int
	fmt.Println("Enter the user ID")
	var userID int
	userID = utils.Readint()

	searched_bookID = utils.Readint()

	LSC.library_service.BorrowBook(searched_bookID,userID)
}
func (LSC *library_service_control) Retern_Book_controller(){
	fmt.Println("Return a book:")
	var bookID int
	fmt.Println("Inter the ID of Returnig Book: ")
	bookID = utils.Readint()

	var userID int
	userID = utils.Readint()

	LSC.library_service.ReturnBook(bookID,userID)
}

func (LSC *library_service_control) Remove_Book_controller(){
	fmt.Println("Remove Book Service")
	var removeBookID int
	removeBookID = utils.Readint()
	LSC.library_service.RemoveBook(removeBookID)

}

func (LSC *library_service_control) Borrowed_book_controller(){
	fmt.Println("borowed book ")
	var userID int

	fmt.Println("Inter user ID: ")
	userID = utils.Readint()

	books,error :=LSC.library_service.ListBorrowedBooks(userID)
	if error == nil && books != nil {
		for _,book := range books  {
			fmt.Printf("book: %s   author : %s",book.Title,book.Author)
		}
	}
	
}

func (LCS *library_service_control) List_all_the_available_Book_controller(){
	fmt.Println("list all available books service ")
	for _,v := range LCS.library_service.ListAvailableBooks(){
		fmt.Println(v)
	}

}