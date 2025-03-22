FROM gcr.io/oss-fuzz-base/base-builder-go:v1

COPY . $SRC/continuous-fuzzing

WORKDIR continuous-fuzzing

COPY .clusterfuzzlite/build.sh $SRC/           