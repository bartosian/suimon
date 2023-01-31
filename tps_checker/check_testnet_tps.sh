#!/bin/bash
  
if [ $(dpkg-query -W -f='${Status}' jq 2>/dev/null | grep -c "ok installed") -eq 0 ];
then
  sudo apt update
  sudo apt install -y jq
fi

echo '-------------------------------------'
echo $(date)
echo 
SUISTART=$(curl --location --request POST https://fullnode.testnet.sui.io:443 \
--header 'Content-Type: application/json' \
--data-raw '{ "jsonrpc":"2.0", "method":"sui_getTotalTransactionNumber","id":1}' 2>/dev/null | jq .result)

NODESTART=$(curl --location --request POST 127.0.0.1:9000 \
--header 'Content-Type: application/json' \
--data-raw '{ "jsonrpc":"2.0", "method":"sui_getTotalTransactionNumber","id":1}' 2>/dev/null | jq .result)

for I in {1..10}; do
  sleep 1
  BAR="$(yes . | head -n ${I} | tr -d '\n')"
  printf "\rIN PROGRESS [%3d/100] %s" $((I * 10)) ${BAR}
done

printf "\n\n"

SUIEND=$(curl --location --request POST https://fullnode.testnet.sui.io:443 \
--header 'Content-Type: application/json' \
--data-raw '{ "jsonrpc":"2.0", "method":"sui_getTotalTransactionNumber","id":1}' 2>/dev/null | jq .result)

NODEEND=$(curl --location --request POST 127.0.0.1:9000 \
--header 'Content-Type: application/json' \
--data-raw '{ "jsonrpc":"2.0", "method":"sui_getTotalTransactionNumber","id":1}' 2>/dev/null | jq .result)

SUITPS=$((($SUIEND-$SUISTART)/10))
MYTPS=$((($NODEEND-$NODESTART)/10))

echo 'SUI TPS: '$SUITPS
echo 'NODE TPS: '$MYTPS
echo '-------------------------------------'
