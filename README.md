# conv
a golang util for convert data to the type you want


ToJSON
========
You can convert to json with many type of data source (`string`,`[]byte`,`io.Reader`,`conv.File`)

``` go
err = conv.ToJSON(conv.File("cfg.json"), &foo)
if err != nil {
	log.Println("EROR: ", err)
}
```
``` go
func Handle(w http.ResponseWriter, r *http.Request) {
	err := conv.ToJSON(r.Body, &foo)
	if err != nil {
		log.Println("EROR: ", err)
		w.Write([]byte("bad json"))
		return
	}
	w.Write([]byte("success"))
}
```

ToXML
========
similar to `ToJSON`