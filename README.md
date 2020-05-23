# Ascii-art

## How to test
```sh
$ cd test
$ go build
$ ./test [string] [flags...]
```

## Flag rules
### fs
 - Only 1 font style can be used

### output
 - Flag --output can be used only with fs flag

### reverse
 - Flag --reverse must be used alone

### color
 - You can use --color without --slice if you want to colorize whole text to one color
 - Otherwise, their lengths must be equal
 - You must write colors and slices sequentially
 - In slices you dont need to consider empty spaces and new lines
 - RGB colors must be in the form "n.n.n"
 - Example: 
 ```sh
 $ ./test "hello world \nhello world" --color=cyan,gold,orange,lime --slice=0:5,5:10,10:15,15:20
 ```
 ![Output](https://i.imgur.com/TEqGKB7.png)
 
