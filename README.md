# Continuous Fuzzing

To Run the Fuzz test use:

```sh
go test -fuzz=FuzzParseString ./go-native-fuzzing -test.fuzzcachedir="../"
```

## Fuzzing using ClutserFuzzLite

Testing ClusterFuzzLite Locally:

```sh
python ../oss-fuzz/infra/helper.py build_image --external $PATH_TO_PROJECT
python ../oss-fuzz/infra/helper.py build_fuzzers --external $PATH_TO_PROJECT --sanitizer address
```
