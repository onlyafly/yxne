https://json5.org/
https://dhall-lang.org/#

## What is a good data format have?

- comments
- no end of line delimiter

## How to document the grammar

https://www.crockford.com/mckeeman.html

## WIP

### Heterogenous List/Vector

(1 42 "foo")

1
24
"foo"

### Map

{x:1 y:42 z:"foo"}

x: 1
y: 42
z: "foo"

### Struct/Named List

coord(1 2)

coord..
  1
  2

OR

place{city:"Stockholm" lat:59 long:18}

place..
  city: "Stockholm"
  lat: 59
  long: 18

### Homogenous List/Vector/Array

[1 2 3]

--no expanded syntax for arrays--

## Attempts to see if this is a good idea

+(,
  1
  +(,
    2
    3

def(double [x] +(1 2))

def(,
  double
  [x]
  +(1 2)

def(double [x],
  if(isnum(x)
    return(*(x 2))
    panic(x "is not a number")

def double x
  if isnum(x)
    return *(x 2)
    panic x "is not a number"
