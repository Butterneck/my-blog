set -x

cd ../src/blog-backend

GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o bootstrap 

# zip myFunction.zip bootstrap