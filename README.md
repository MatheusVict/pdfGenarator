# PDF-Generator-Golang

This is a pdf generator made in golang, that uses html template to generate custom pdfs

### Used Technologies

* [Golang](https://github.com/golang/go)
## Dependencies and required versions

* wkHtmlToPdf - Version: 1.9.2
* uuid - Version: 1.4.0

### How to run

```shell
go mod tidy
```
and then

```shell
go run *.go
```

After that you will see on the ```tmp``` directory your pdfs

## How change pdf

you need to change html file in ```templates/example.html```

where you want to make a dynamic data you'll need to put ```{{.VariableName}}```, and change the *Data* struct in ```main.go``` or create your own

## Faced troubles

### Problem 1:
no command on start application
* Make sure you have wkHtmlToPdf installed or run it for your shell.


## Next steps

Make this project a go package

