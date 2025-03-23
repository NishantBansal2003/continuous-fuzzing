#!/bin/bash -eu

function compile_fuzzer {
  path=$1
  function=$2
  fuzzer=$3

   # Instrument all Go files relevant to this fuzzer
  go-fuzz -tags gofuzz -func $function -o $fuzzer.a $path

   # Instrumented, compiled Go ($fuzzer.a) + fuzzing engine = fuzzer binary
  $CXX $CXXFLAGS $LIB_FUZZING_ENGINE $fuzzer.a -lpthread -o $OUT/$fuzzer
}

compile_fuzzer ./fuzz Fuzz fuzz_parse_string