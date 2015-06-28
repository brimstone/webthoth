#!/bin/bash
set -ueo pipefail

BASE=http://localhost:8080

echo "==> Checking for existing offer"
curl -s -m 1 -i $BASE/room/test
# this should return nothing
echo

echo "==> Sending offer"
offerid="$(curl -s -m 1 -d "sdp bits" $BASE/room/test/offer)"
echo "==> Got offer id: '$offerid'"
echo

echo "==> Checking for existing offer"
offerid2="$(curl -s -m 1 $BASE/room/test)"
echo

echo "==> Got offer id: '$offerid2'"
echo


echo "==> Answering offer"
curl -s -m 1 -d "offer bits" "$BASE/room/test/$offerid"
