#!/bin/bash

# Split the REGISTRY_IMAGE variable into repository and tag parts
# e.g. kabanero/kabanero-operator-registry:0.3.0 -> kabanero/kabanero-operator-registry
IFS=’:’ read -ra REPOSITORY <<< "$REGISTRY_IMAGE"

# Set the tag for the kabanero CatalogSource
sed -i.bak -e "s,kabanero/kabanero-operator-registry:latest,kabanero/kabanero-operator-registry:$TRAVIS_TAG," deploy/kabanero-subscriptions.yaml

# Set the tag for the install script
sed -i.bak -e 's/RELEASE=.*/RELEASE="${RELEASE:-'$TRAVIS_TAG'}"/' deploy/install.sh

# Set the tag for the uninstall script
sed -i.bak -e 's/RELEASE=.*/RELEASE="${RELEASE:-'$TRAVIS_TAG'}"/' deploy/uninstall.sh