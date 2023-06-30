FROM golang:1.20-alpine AS build
WORKDIR /app
COPY . .
COPY vendor ./vendor
RUN go build -o sheet.report

FROM alpine AS runtime
RUN apk add --no-cache tzdata
ENV TZ=Asia/Ho_Chi_Minh
WORKDIR /app
COPY --from=build /app/sheet.report .
COPY sheet_connect ./sheet_connect
EXPOSE 8080
EXPOSE 8081
ENTRYPOINT ["/app/sheet.report"]
