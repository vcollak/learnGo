package main

import (
	"fmt"
	"time"
)

//site structure
type Site struct {
	Name    string
	Address string
	City    string
	Zip     string
}

//some function to get site info maybe via network call
//returns a site or an error
//notice it returns an error object
func GetSite() (error, Site) {

	//error
	var site Site
	return &SiteError{time.Now(), "it didn't work"}, site

	//no error
	//site := Site{"Bobby Site", "123 Address", "Houston", "77077"}
	//return nil, site

}

//------ Error --------
//error structure
type SiteError struct {
	Date         time.Time
	ErrorMessage string
}

//error method
//returns a string with the error nessage
func (se *SiteError) Error() string {
	return fmt.Sprintf("%v %s", se.Date, se.ErrorMessage)
}

func main() {

	fmt.Println("### Errors1 ###")

	//if GetSite returns an error (of type SiteError)
	//then we know we had an error
	if err, site := GetSite(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(site.Name)
	}

}
