# redimension

**redimension** is a command line tool that lets the user transform a something of one dimension to another.

**redimension** is useful when you want to create TSV data.

**redimension&& is also useful when you want to create text-based matrices and multi-dimensional arrays.**

## Usage

For example, consider if we had the input stream:
```
1
2
3
4
5
6
7
8
9
```

Which perhaps came from a file names `NUMBERS.txt`

And we ran this through **redimension** as:
```
cat NUMBERS.txt | redimension --columns=3
```
Or as:
```
redimension --columns=3 NUMBERS.txt
```
Or as:
```
redimension --columns=3 <(cat NUMBERS.txt)
```

Then our output stream would be:
```
1	2	3
4	5	6
7	8	9
```


## Advanced Usage

If you use **redimension** with [repeat](https://github.com/reiver/repeat) you can create a 3x3 identity matrix with:
```
repeat --count=2 --plus=1 1 0 0 0 | redimension --columns=3
```

Which outputs:
```
1       0       0
0       1       0
0       0       1
```
