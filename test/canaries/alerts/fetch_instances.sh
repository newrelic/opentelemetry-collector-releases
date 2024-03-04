#!/usr/bin/env bash

# This helps in debugging your scripts. TRACE=1 ./script.sh
if [[ "${TRACE-0}" == "1" ]]; then set -o xtrace; fi

usage() {
 echo 'Usage: ./fetch_instances.sh otel-canary:nr-otel-collector-0.5.0'
}

CURRENT_VERSION=$1
if [ -z "${CURRENT_VERSION}" ];then
  { usage; echo "CURRENT_VERSION is not provided"; exit 1; }
fi

INSTANCES=$( aws ec2 describe-instances --no-paginate \
  --filters 'Name=instance-state-code,Values=16' \
  --query 'Reservations[*].Instances[*].{PrivateIpAddress:PrivateIpAddress,Name:[Tags[?Key==`Name`].Value][0][0],InstanceId:InstanceId}' \
  | jq -c '.[]' \
  | tr -d "[]{}\""
)

cat << EOF

linux_display_names  = [
EOF
for ins in $INSTANCES;do
  NAME=$( echo $ins | awk -F "," '{print $2}' | sed "s/^Name://g" )

  if [[ ! "${NAME}" =~ $CURRENT_VERSION ]]; then
    continue
  fi

  echo "    { previous = \"$NAME-docker-previous\", current = \"$NAME-docker-current\"},"

done

cat << EOF
]
EOF
