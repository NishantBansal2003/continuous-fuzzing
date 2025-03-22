FROM gcr.io/oss-fuzz-base/base-builder-go:latest

COPY . $SRC/continuous-fuzzing

WORKDIR $SRC/continuous-fuzzing

COPY ./.clusterfuzzlite/build.sh $SRC/           