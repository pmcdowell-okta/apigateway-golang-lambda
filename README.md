# GOLANG Lambda

Quick example of how to user Serverless Framework to deploy an Amazon AWS API Gateway that uses a GOLANG Lambda function on the back end

* Checkout Code
* Compile like this: `GOOS=linux go build -o bin/main`
* Deploy  like this: `sls deploy`
* Test like this: 

`curl -X POST -d "car" \
                     https://tmgc8to3dh.execute-api.us-east-1.amazonaws.com/dev/echo \
                     -H 'accept: application/json' \
                     -H 'cache-control: no-cache' \
                     -H 'content-type: application/x-www-form-urlencoded'`
                     
* **Use your URL, your URL will be different from mine..**

The result should look like this:

`{"Dude where is my":"car"}`              