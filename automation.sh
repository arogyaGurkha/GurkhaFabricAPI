#!/bin/bash

## Package the CC "asset-transfer-basic"
PACKAGE=$(curl --location --request POST '15.164.97.167:8080/fabric/lifecycle/package' \
--header 'Content-Type: application/json' \
--data-raw '{
    "cc_source_name": "asset-transfer-basic",
    "label": "basic_1.0",
    "language": "go",
    "package_name": "basic.tar.gz"
}')
echo $PACKAGE

## Set Org1 as the admin
ADMIN1=$(curl --location --request POST '15.164.97.167:8080/fabric/lifecycle/admin/Org1')
echo $ADMIN1

## Install for Org1
INSTALLCC=$(curl --location --request POST '15.164.97.167:8080/fabric/lifecycle/install/basic.tar.gz')
echo $INSTALLCC

## Set Org2 as the admin
ADMIN2=$(curl --location --request POST '15.164.97.167:8080/fabric/lifecycle/admin/Org2')
echo $ADMIN2

## Install for Org2
INSTALLCC=$(curl --location --request POST '15.164.97.167:8080/fabric/lifecycle/install/basic.tar.gz')
echo $INSTALLCC

## Check if installed succesfully
INSTALLID=$(curl --location --request GET '15.164.97.167:8080/fabric/lifecycle/install')
echo $INSTALLID

ADMIN1=$(curl --location --request POST '15.164.97.167:8080/fabric/lifecycle/admin/Org1')
echo $ADMIN1

## Approve for Org1
APPROVEORG1=$(curl --location --request POST '15.164.97.167:8080/fabric/lifecycle/approve' \
--header 'Content-Type: application/json' \
--data-raw '{
    "cc_name": "basic",
    "cc_sequence": 1,
    "cc_version": "1.0",
    "channel_name": "mychannel",
    "package_ID": "basic_1.0:ff2e5d5f8ff054f8e4f951c592068e863c80b62cb6a9a355d4b51de886273c96"
}')
echo $APPROVEORG1

## Check commit readiness from both Orgs
QUERYCOMMITREADY=$(curl --location --request GET '15.164.97.167:8080/fabric/lifecycle/commit/organizations' \
--header 'Content-Type: application/json' \
--data-raw '{
  "cc_name": "basic",
  "cc_sequence": 1,
  "cc_version": "1.0",
  "channel_name": "mychannel"
}')
echo $QUERYCOMMITREADY

ADMIN2=$(curl --location --request POST '15.164.97.167:8080/fabric/lifecycle/admin/Org2')
echo $ADMIN2

QUERYCOMMITREADY=$(curl --location --request GET '15.164.97.167:8080/fabric/lifecycle/commit/organizations' \
--header 'Content-Type: application/json' \
--data-raw '{
  "cc_name": "basic",
  "cc_sequence": 1,
  "cc_version": "1.0",
  "channel_name": "mychannel"
}')
echo $QUERYCOMMITREADY

## Approve for Org2
APPROVEORG2=$(curl --location --request POST '15.164.97.167:8080/fabric/lifecycle/approve' \
--header 'Content-Type: application/json' \
--data-raw '{
    "cc_name": "basic",
    "cc_sequence": 1,
    "cc_version": "1.0",
    "channel_name": "mychannel",
    "package_ID": "basic_1.0:ff2e5d5f8ff054f8e4f951c592068e863c80b62cb6a9a355d4b51de886273c96"
}')
echo $APPROVEORG2

## Check commit readiness from both Orgs
ADMIN1=$(curl --location --request POST '15.164.97.167:8080/fabric/lifecycle/admin/Org1')
echo $ADMIN1

QUERYCOMMITREADY=$(curl --location --request GET '15.164.97.167:8080/fabric/lifecycle/commit/organizations' \
--header 'Content-Type: application/json' \
--data-raw '{
  "cc_name": "basic",
  "cc_sequence": 1,
  "cc_version": "1.0",
  "channel_name": "mychannel"
}')
echo $QUERYCOMMITREADY

ADMIN2=$(curl --location --request POST '15.164.97.167:8080/fabric/lifecycle/admin/Org2')
echo $ADMIN2

QUERYCOMMITREADY=$(curl --location --request GET '15.164.97.167:8080/fabric/lifecycle/commit/organizations' \
--header 'Content-Type: application/json' \
--data-raw '{
  "cc_name": "basic",
  "cc_sequence": 1,
  "cc_version": "1.0",
  "channel_name": "mychannel"
}')
echo $QUERYCOMMITREADY

## Commit the definition
COMMITCC=$(curl --location --request POST '15.164.97.167:8080/fabric/lifecycle/commit' \
--header 'Content-Type: application/json' \
--data-raw '{
    "cc_name": "basic",
    "cc_sequence": 1,
    "cc_version": "1.0",
    "channel_name": "mychannel"
}')
echo $COMMITCC