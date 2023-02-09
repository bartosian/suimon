#!/bin/bash

if [ $(dpkg-query -W -f='${Status}' jq 2>/dev/null | grep -c "ok installed") -eq 0 ];
then
  sudo apt update
  sudo apt install -y jq
fi

echo '------------------DEVNET TPS 1.0.0-------------------'
echo $(date)
echo

REMOTE_RPC="https://fullnode.devnet.sui.io:443"
LOCAL_RPC="127.0.0.1:9000"

function get_transactions {
  result=$(curl -m 2 --location --request POST $1 \
  --header 'Content-Type: application/json' \
  --data-raw '{ "jsonrpc":"2.0", "method":"sui_getTotalTransactionNumber","id":1}' 2>/dev/null | jq .result)

  echo "$result"
}

SUISTART=$(get_transactions $REMOTE_RPC)
if [ -z "$SUISTART" ]; then
    echo "Failed to calculate TPS: check if remote devnet RPC is up and running"

    exit 1
fi

NODESTART=$(get_transactions $LOCAL_RPC)
if [ -z "$NODESTART" ]; then
  echo "Failed to calculate TPS: check if your node is up and running on port 9000"

  exit 1
fi

for I in {1..10}; do
  sleep 1
  BAR="$(yes . | head -n ${I} | tr -d '\n')"
  printf "\rIN PROGRESS [%3d/100] %s" $((I * 10)) ${BAR}
done

printf "\n\n"

SUIEND=$(get_transactions $REMOTE_RPC)
NODEEND=$(get_transactions $LOCAL_RPC)

SUITPS=$(((SUIEND-SUISTART)/10))
MYTPS=$(((NODEEND-NODESTART)/10))

echo 'SUI TPS: '$SUITPS
echo 'NODE TPS: '$MYTPS
echo '-----------------------------------------------------'