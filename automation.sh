#!/bin/bash

DELAY=${10:-"3"}

approveOrg(){
    APPROVEORG=$(curl --location --request POST '15.164.97.167:8080/fabric/lifecycle/approve' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "cc_name": "basic",
        "cc_sequence": 1,
        "cc_version": "1.0",
        "channel_name": "mychannel",
        "package_ID": "basic_1.0:ff2e5d5f8ff054f8e4f951c592068e863c80b62cb6a9a355d4b51de886273c96"
    }')
    echo $APPROVEORG
}

queryCommitReady(){
     sleep $DELAY
    QUERYCOMMITREADY=$(curl --location --request GET '15.164.97.167:8080/fabric/lifecycle/commit/organizations' \
    --header 'Content-Type: application/json' \
    --data-raw '{
    "cc_name": "basic",
    "cc_sequence": 1,
    "cc_version": "1.0",
    "channel_name": "mychannel"
    }')
    echo $QUERYCOMMITREADY
}

queryCommitted(){
     sleep $DELAY
    QUERYCOMMIT=$(curl --location --request GET '15.164.97.167:8080/fabric/lifecycle/commit' \
    --header 'Content-Type: application/json' \
    --data-raw '{
    "cc_name": "basic",
    "channel_name": "mychannel"
    }')
    echo $QUERYCOMMIT
}

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

## Set Org1 as the admin
ADMIN1=$(curl --location --request POST '15.164.97.167:8080/fabric/lifecycle/admin/Org1')
echo $ADMIN1

## Check if installed succesfully
INSTALLID=$(curl --location --request GET '15.164.97.167:8080/fabric/lifecycle/install')
echo $INSTALLID

## Approve for Org1
ADMIN1=$(curl --location --request POST '15.164.97.167:8080/fabric/lifecycle/admin/Org1')
echo $ADMIN1
approveOrg

## Check commit readiness from both Orgs
ADMIN1=$(curl --location --request POST '15.164.97.167:8080/fabric/lifecycle/admin/Org1')
echo $ADMIN1
queryCommitReady

ADMIN2=$(curl --location --request POST '15.164.97.167:8080/fabric/lifecycle/admin/Org2')
echo $ADMIN2
queryCommitReady

## Approve for Org2
ADMIN2=$(curl --location --request POST '15.164.97.167:8080/fabric/lifecycle/admin/Org2')
echo $ADMIN2
approveOrg

## Check commit readiness from both Orgs
ADMIN1=$(curl --location --request POST '15.164.97.167:8080/fabric/lifecycle/admin/Org1')
echo $ADMIN1
queryCommitReady

ADMIN2=$(curl --location --request POST '15.164.97.167:8080/fabric/lifecycle/admin/Org2')
echo $ADMIN2
queryCommitReady

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

## Check committed CC from both Orgs
ADMIN1=$(curl --location --request POST '15.164.97.167:8080/fabric/lifecycle/admin/Org1')
echo $ADMIN1
queryCommitted

## Check committed CC from both Orgs
ADMIN2=$(curl --location --request POST '15.164.97.167:8080/fabric/lifecycle/admin/Org1')
echo $ADMIN2
queryCommitted
