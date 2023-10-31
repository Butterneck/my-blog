FROM golang:1.20 as build
WORKDIR /build
# Copy dependencies list
COPY go.mod go.sum ./
# Download dependencies
RUN go mod download
# Build with optional lambda.norpc tag
COPY . .
RUN CGO_ENABLED=0 go build -tags lambda.norpc -o main

# Copy artifacts to a clean image
FROM public.ecr.aws/lambda/provided:al2
COPY --from=build /build/main ./main
ENTRYPOINT [ "./main" ]