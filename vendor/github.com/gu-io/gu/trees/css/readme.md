# CSS
CSS provides a library which greatly simplify how we write css styles in a more flexible way by using the power of Go templates.


## Install

```bash
go get -u github.com/gu-io/gu/css
```

## Example

- Create a new css style with properties fed in

```go
csr := css.New(`

    $:hover {
      color: red;
    }

    $::before {
      content: "bugger";
    }

    $ div a {
      color: black;
      font-family: {{ .Font }}
    }

    @media (max-width: 400px){

      $:hover {
        color: blue;
        font-family: {{ .Font }}
      }

    }
`, nil)

  sheet, err := csr.Stylesheet(struct {
    Font string
  }{Font: "Helvetica"}, "#galatica")

  sheet.String() // => "#galatica:hover {\n  color: red;\n}\n#galatica::before {\n  content: \"bugger\";\n}\n#galatica div a {\n  color: black;\n  font-family: Helvetica;\n}\n@media (max-width: 400px) {\n  #galatica:hover {\n    color: blue;\n    font-family: Helvetica;\n  }\n}"

```

- Extend parts of another css rule into a giving style selector

```go
	csr := css.New(`
    block {
      font-family: {{ .Font }};
      color: {{ .Color }};
    }
  `, nil)

	csx := css.New(`

    ::before {
      content: "bugger";
    }

    div a {
			{{ extend "block" }}
			border: 1px solid #000;
    }

    @media (max-width: 400px){

      :hover {
        color: blue;
        font-family: {{ .Font }};
      }

    }
`, csr)

	sheet, err := csx.Stylesheet(struct {
		Font  string
		Color string
	}{
		Font:  "Helvetica",
		Color: "Pink",
	}, "#galatica")

  sheet.String() /*=>

#galatica::before {
  content: "bugger";
}
div a {
  font-family: Helvetica;
  color: Pink;
  border: 1px solid #000;
}
@media (max-width: 400px) {
  #galatica:hover {
    color: blue;
    font-family: Helvetica;
  }
}

*/
```

## Gratitude
Thanks to the awesome work of the [CSS tokenizer by the Gorilla team](https://github.com/gorilla/css)  
and [Aymerick's css parser](https://github.com/aymerick/douceur) through all whom by God's grace made this library possible.
