FROM golang as builder

RUN mkdir /build
ADD . /build/
WORKDIR /build

RUN apt-get update
RUN apt-get -y install curl
RUN apt-get -y install libfreetype6 -y --no-install-recommends apt-utils
RUN apt-get -y install libfontconfig -y --no-install-recommends apt-utils

RUN go get -v -u github.com/gorilla/mux

RUN go build -a -installsuffix cgo -ldflags '-extldflags "-s -L./indigo/lib -Wl,-rpath=./indigo/lib -lindigo"' -o service main.go
FROM busybox:glibc
COPY --from=builder /build/service /app/
COPY --from=builder /build/indigo/lib/libindigo.so /app/indigo/lib/
COPY --from=builder /build/indigo/lib/libindigo-renderer.so /app/indigo/lib/

COPY --from=builder /lib/x86_64-linux-gnu/libgcc_s.so.1 /lib/libgcc_s.so.1
COPY --from=builder /lib/x86_64-linux-gnu/libgcc_s.so.1 /lib/libgcc_s.so.1
COPY --from=builder /usr/lib/x86_64-linux-gnu/libfreetype.so.6 /lib/libfreetype.so.6
COPY --from=builder /usr/lib/x86_64-linux-gnu/libfontconfig.so.1 /lib/libfontconfig.so.1
COPY --from=builder /lib/x86_64-linux-gnu/libz.so.1 /lib/libz.so.1
COPY --from=builder /usr/lib/x86_64-linux-gnu/libpng16.so.16 /lib/libpng16.so.16
COPY --from=builder /lib/x86_64-linux-gnu/libexpat.so.1 /lib/libexpat.so.1

WORKDIR /app
CMD ["./service"]
