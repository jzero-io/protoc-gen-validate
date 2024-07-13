# gen validate.pb.go

```shell
protoc -I. validate.proto --go_out="."
mv github.com/envoyproxy/protoc-gen-validate/validate/validate.pb.go .
rm -rf github.com
```