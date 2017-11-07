FROM iron/go:dev

# Create Env Vars and Working Directory
ARG POSTMATES_ENV=dev
ARG POSTMATES_KEY
WORKDIR /app

# Build API
ENV SRC_DIR=/go/src/github.com/effortless-technologies/elt-delivery-api
ADD . $SRC_DIR
ENV POSTMATES_ENV ${POSTMATES_ENV}
RUN cd $SRC_DIR; go build -o api -tags ${POSTMATES_ENV}; cp api /app/

# Run API
ENV POSTMATES_KEY ${POSTMATES_KEY}
ENTRYPOINT POSTMATES_PROD_KEY=${POSTMATES_KEY} ./api
