package main

import (
	"encoding/xml"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", IndexHanlder)
	router.Run()
}

type Person struct {
	XMLName   xml.Name `xml:"person"`
	FirstName string   `xml:"firstName,attr"`
	LastName  string   `xml:"lastName,attr"`
}

func IndexHanlder(c *gin.Context) {
	c.XML(200, Person{FirstName: "Bicky", LastName: "Tamang"})
}
