#!/bin/bash
  
echo '-------------------------------------'
echo $(date)
SUIVER=$(curl --location --request POST https://fullnode.devnet.sui.io:443 \
--header 'Content-Type: application/json' \
--data-raw '{ "jsonrpc":"2.0", "method":"sui_getTotalTransactionNumber","id":1}' 2>/dev/null | jq .result)

MYVER=$(curl --location --request POST 127.0.0.1:9000 \
--header 'Content-Type: application/json' \
--data-raw '{ "jsonrpc":"2.0", "method":"sui_getTotalTransactionNumber","id":1}' 2>/dev/null | jq .result)

echo 'Sui ledger version : '$SUIVER
echo 'My ledger version : '$MYVER

for I in {1..10}; do
  sleep 1
  BAR="$(yes . | head -n ${I} | tr -d '\n')"
  printf "\r[%3d/100] %s" $((I * 10)) ${BAR}
done
printf "\n"

echo '-------------------------------------'
echo $(date)
NOWSUIVER=$(curl --location --request POST https://fullnode.devnet.sui.io:443 \
--header 'Content-Type: application/json' \
--data-raw '{ "jsonrpc":"2.0", "method":"sui_getTotalTransactionNumber","id":1}' 2>/dev/null | jq .result)

NOWMYVER=$(curl --location --request POST 127.0.0.1:9000 \
--header 'Content-Type: application/json' \
--data-raw '{ "jsonrpc":"2.0", "method":"sui_getTotalTransactionNumber","id":1}' 2>/dev/null | jq .result)

SUITPS=$((($NOWSUIVER-$SUIVER)/10))
MYTPS=$((($NOWMYVER-$MYVER)/10))
echo 'Sui ledger version : '$NOWSUIVER
echo 'My ledger version : '$NOWMYVER
echo '-------------------------------------'
echo 'Sui TPS : '$SUITPS
echo 'My TPS : '$MYTPS
