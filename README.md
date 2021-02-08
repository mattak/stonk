# stonk

Stock market basic data tools.

## Install

```shell
$ make install
```

## Usage

List up symbols

```shell
$ stonk symbol
AACG
AACQ
AACQU
...
```

Price fetching

```shell
$ stonk price AAPL
date    open    close   high    low     volume
2000-01-01      0.936384        0.926339        1.084821        0.772321        12555177600
2000-02-01      0.928571        1.023438        1.070871        0.866071        7319782400
...
```
