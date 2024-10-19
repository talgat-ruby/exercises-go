package constants

import "os"

var NAME = os.Getenv("NAME")
var PORT = os.Getenv("PORT")
var URL = "localhost:" + PORT
var FullUrl = "http://" + URL
