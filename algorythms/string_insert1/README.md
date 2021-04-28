# string_insert1

Given 2 Strings: Check if you can make the second string by copy-pasting the whole
first string multiple times and inserting it in any place.

```
"XY" "XXXYYY" -> True (XY -> X[XY]Y -> XX[XY]YY)
```

## Trying it out

```
go run main.go -string XY -result XXXYYY
go run main.go -string XY -result XXXYYYA
```
