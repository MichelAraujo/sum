# Just sum and return the total based in the lines in a file

## To compile for linux
```
$ go build -o sum main.go
$ sudo cp sum /usr/local/bin/
```

### Example of the file format

```
# 10
# 20
# 30
```
Run:
```
$ sum example
60
```