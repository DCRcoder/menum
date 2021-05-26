# Menum
 An enum generator for go
enum 类型数据生成器
## 安装
```shell
// clone
git clone git@github.com:DCRcoder/menum.git
// cd 到目录下
go install
```
## 命令用例
```shell
NAME:
   enumer - An enum generator for go

USAGE:
   menum [global options] [arguments...]

GLOBAL OPTIONS:
   --file value, -f value    The file(s) to generate enums.  Use more than one flag for more files.
   --output value, -o value  output file name; default srcdir/<type>_enum.go
   --help, -h                show help (default: false)
```

## go generate 用例
```
//go:generate menum -output=gen_enums.go
```

### 生成代码用例 查看 [examples](./examples/macro.go)

## TODO
