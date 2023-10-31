set -x

cd ../src

GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o bootstrap 

# zip myFunction.zip bootstrap