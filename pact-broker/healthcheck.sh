#!/bin/sh

set -e

# print error and exit
die () {
  echoerr "ERROR: $0: $1"
  # if $2 is defined AND NOT EMPTY, use $2; otherwise, set to "150"
  errnum=${2-115}
  exit $errnum
}

echo ""
echo "Checking that the Pact Broker container is still up and running"
docker inspect -f "{{ .State.Running }}" ${PACT_CONT_NAME} | grep true || die \
  "The Pact Broker container is not running!"
if [ -z "${TEST_IP}" ]; then
  TEST_IP=`docker inspect -f='{{ .NetworkSettings.IPAddress }}' ${PACT_CONT_NAME}`
fi
TEST_URL="http://${TEST_IP}:${PACT_BROKER_PORT}"
echo "TEST_URL is '${TEST_URL}'"
echo ""
echo "Checking that server can be connected from outside the Docker container"
PACT_BROKER_HOST=${TEST_IP} $(dirname "$0")/wait_pact.sh ${PACT_WAIT_TIMEOUT} ${PACT_BROKER_BASIC_AUTH_USERNAME} ${PACT_BROKER_BASIC_AUTH_PASSWORD}
echo ""
echo "Checking that server accepts and return HTML from outside"
curl -H "Accept:text/html" --user ${PACT_BROKER_BASIC_AUTH_USERNAME}:${PACT_BROKER_BASIC_AUTH_PASSWORD} -s "${TEST_URL}"
echo ""
echo "Checking for specific HTML content from outside: 'Pacts'"
curl -H "Accept:text/html" --user ${PACT_BROKER_BASIC_AUTH_USERNAME}:${PACT_BROKER_BASIC_AUTH_PASSWORD} -s "${TEST_URL}" | grep "Pacts"
echo "Checking that server accepts and responds with status 200"
response_code=$(curl -s -o /dev/null -w "%{http_code}" --user ${PACT_BROKER_BASIC_AUTH_USERNAME}:${PACT_BROKER_BASIC_AUTH_PASSWORD} ${TEST_URL})
if [[ "${response_code}" -ne '200' ]]; then
  die "Expected response code to be 200, but was ${response_code}"
fi
if [[ ! -z "${PACT_BROKER_BASIC_AUTH_USERNAME}" ]]; then
  echo ""
  echo "Checking that basic auth is configured"
  response_code=$(curl -s -o /dev/null -w "%{http_code}" ${TEST_URL})
  if [[ "${response_code}" -ne '401' ]]; then
    die "Expected response code to be 401, but was ${response_code}"
  else
    echo "Correctly received a 401 for an unauthorised request"
  fi
fi

BODY=$(ruby -e "require 'json'; j = JSON.parse(File.read('script/foo-bar.json')); j['interactions'][0]['providerState'] = 'it is ' + Time.now.to_s; puts j.to_json")
echo ${BODY} >> tmp.json
curl -v -XPUT -u foo:bar \-H "Content-Type: application/json" \
-d@tmp.json \
${TEST_URL}/pacts/provider/Bar/consumer/Foo/version/1.1.0
rm tmp.json
echo ""

echo ""
echo "Checking that badges can be accessed without basic auth"
response_code=$(curl -s -o /dev/null -w "%{http_code}" ${TEST_URL}/pacts/provider/Bar/consumer/Foo/latest/badge.svg)
if [[ "${response_code}" -ne '200' ]]; then
  die "Expected response code to be 200, but was ${response_code}"
fi
echo "Checking that the heartbeat URL can be accessed without basic auth"
response_code=$(curl -s -o /dev/null -w "%{http_code}" ${TEST_URL}/diagnostic/status/heartbeat)
if [[ "${response_code}" -ne '200' ]]; then
  die "Expected response code to be 200, but was ${response_code}"
fi
echo "SUCCESS: All tests passed!"