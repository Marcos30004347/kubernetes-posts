#!/bin/bash -e

CURRENT_DIR=$(echo "$(pwd)/$line")
REPO_DIR="$CURRENT_DIR"

PROJECT_MODULE="github.com/Marcos30004347/kubernetes-posts"
IMAGE_NAME="kubernetes-codegen:latest"

CUSTOM_RESOURCE_NAME="baz"
CUSTOM_RESOURCE_VERSION="v1alpha1"

echo "Building codegen Docker image..."
docker build -f "${CURRENT_DIR}/hack/docker/codegen.dockerfile" \
             -t "${IMAGE_NAME}" \
             "${REPO_DIR}"

cmd0="./generate-groups.sh all \
    "$PROJECT_MODULE/pkg/generated" \
    "$PROJECT_MODULE/pkg/apis" \
    $CUSTOM_RESOURCE_NAME:$CUSTOM_RESOURCE_VERSION"

cmd1="./generate-internal-groups.sh "deepcopy,defaulter,conversion,informer,listers,client" \
    "$PROJECT_MODULE/pkg/generated" \
    "$PROJECT_MODULE/pkg/apis" \
    "$PROJECT_MODULE/pkg/apis" \
    $CUSTOM_RESOURCE_NAME:$CUSTOM_RESOURCE_VERSION"

echo "Generating client codes..."

docker run --rm \
           -v "${REPO_DIR}:/go/src/${PROJECT_MODULE}" \
           "${IMAGE_NAME}" $cmd0

docker run --rm \
           -v "${REPO_DIR}:/go/src/${PROJECT_MODULE}" \
           "${IMAGE_NAME}" $cmd1
