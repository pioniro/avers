AutoVERSion
===========
tool for automatic versioning

get
---
```
$ go get github.com/pioniro/avers
```

Examples:
---------

```
$avers next --current="1.2.3-alpha+commit.2e02c0e"
1.2.4
```

```
$avers next --current="1.2.3-alpha+commit.2e02c0e" --keep-pre
1.2.4-alpha
```

```
$avers next --current="1.2.3-alpha+commit.2e02c0e" --keep-pre --change=M
2.0.0-alpha
```

```
$avers next --current="1.2.3-alpha+commit.2e02c0e" --keep-pre --change=m
1.3.0-alpha
```

```
# --change=p - default
$avers next --current="1.2.3-alpha+commit.2e02c0e" --keep-pre --change=p
1.2.4-alpha
```

```
# --change=n - none
$avers next --current="1.2.3-alpha+commit.2e02c0e" --keep-pre --change=n
1.2.3-alpha
```

```
$avers next --current="1.2.3-alpha+commit.2e02c0e" --keep-pre --build="commit.2997e81"
1.2.4-alpha+commit.2997e81
```