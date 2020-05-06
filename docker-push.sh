#!/bin/bash
# inspiration: https://github.com/ShelterTechSF/askdarcel-web/blob/master/tools/docker-build.sh
set -ex

SERVICE_NAME=restaurants
REPO="docker.pkg.github.com/rohan-luthra/service-$SERVICE_NAME-docker/service-$SERVICE_NAME"

COMMIT=$GIT_COMMIT
if [[ -z "$COMMIT" ]]; then
  COMMIT=$(git log -1 --format=%H)
fi
COMMIT=${COMMIT::8}

DOCKER_HOST="docker.pkg.github.com"

ACCOUNT_SVC_PROFILE="development"

if [[ "$ACCOUNT_SVC_PROFILE" == "development" ]]; then
  TAG="dev"
elif [[ "$ACCOUNT_SVC_PROFILE" == "testing" ]]; then
  TAG="test"
elif [[ "$ACCOUNT_SVC_PROFILE" == "production" ]]; then
  TAG="prod"
elif [[ "$ACCOUNT_SVC_PROFILE" == "staging" ]]; then
  TAG="stag"
fi

echo $TAG
echo $COMMIT

echo "Creating version.json..."
echo "{
  \"commit\": \"$COMMIT\",
  \"image\": \"$TAG\"
}" > ./scripts/version.json

echo "Building docker image..."
docker build -f scripts/docker/Dockerfile -t $REPO:$COMMIT .
docker tag $REPO:$COMMIT $REPO:$TAG

$token | docker login docker.pkg.github.com -u "rohan-luthra" --password-stdin
echo "Pushing docker image..."
docker push $REPO:$TAG
echo "Successfully pushed docker image"