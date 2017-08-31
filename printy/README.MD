# Printy

Intended for development only!

Quickly print request/response or any arbitrary interface to stdout.

Methods never return anything.

Errors are just printed to stdout.

## Usage

```
r := somepackage.DoSomeStuff("args",1512)

// wanna look at the object
printy.AsJson(r, true)
```

more:
```
var h http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
	printy.RequestBody(r,true)
}
http.ListenAndServe(":3000", h)
```

## Am I pretty?

```
{
  "answer":"yes",
  "really?":"yes"
}
```

```
{"answer":"no","really not?":"no"}
```
