TAG_PREFIX=wpenv__
TAG_WAIT_FOR_STATUS_CODE=${TAG_PREFIX}wait_for_status_code

WAIT_HOST?=localhost:32080
WAIT_STATUS?=404
WAIT_TIME?=30

all: build_wait_for_status_code \
	 run_wait_for_status_code

build_wait_for_status_code:
	docker build --progress=plain --tag=${TAG_WAIT_FOR_STATUS_CODE} -f ./wait_for_status_code.Dockerfile .

run_wait_for_status_code: build_wait_for_status_code
	docker run --network=host --rm ${TAG_WAIT_FOR_STATUS_CODE} ${WAIT_HOST} ${WAIT_STATUS} ${WAIT_TIME}

