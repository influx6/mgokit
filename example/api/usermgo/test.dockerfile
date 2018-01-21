FROM influx6/mongrel-0.0.1
MAINTAINER GOKIT(gitbub.com/gokit) <trinoxf@gmail.com>

# Set script to run at startup
ENV MONGO_INIT /mnt/db/mongodb/db.js

# This is a test image, dont do this for any production
# system please, am begging please. Instead load secrets
# through the --env-file flag for docker run .
ENV API_MONGO_TEST_HOST 0.0.0.0:27017
ENV API_MONGO_TEST_DB test_db
ENV API_MONGO_TEST_AUTHDB test_db

CMD [/bin/sh -c "/bin/bootmgo --no-auth --fork && go test -v ./..."]