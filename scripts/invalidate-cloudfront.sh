#!/bin/bash

invalidationid=$(aws cloudfront create-invalidation --distribution-id "$DISTRIBUTIONID" --paths "/*" | jq .Invalidation.Id)
invalidationstatus=$(aws cloudfront get-invalidation --id "$invalidationid" --distribution-id "$DISTRIBUTIONID" | jq .Invalidation.Status)
echo "invalidating cloudfront cache"
while [ "$invalidationstatus" == "InProgress" ]
do
  echo "checking cloudfront invalidation status"
  if [ "$invalidationstatus" == "InProgress" ]; then
    echo "cloudfront invalidation status still in progress checking again in 30 seconds"
    sleep 30
    invalidationstatus=$(aws cloudfront get-invalidation --id "$invalidationid" --distribution-id "$DISTRIBUTIONID" | jq .Invalidation.Status)
  fi
done
if [ "$invalidationstatus" == "Completed" ]; then
  echo "cloudfront invalidation complete!"
fi
