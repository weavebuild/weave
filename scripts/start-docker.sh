#!/bin/bash

KEY="mykey"
CHAINID="tomahawk_9000-1"
MONIKER="mymoniker"
DATA_DIR=$(mktemp -d -t tomahawk-datadir.XXXXX)

echo "create and add new keys"
./tomad keys add $KEY --home $DATA_DIR --no-backup --chain-id $CHAINID --algo "eth_secp256k1" --keyring-backend test
echo "init tomahawk with moniker=$MONIKER and chain-id=$CHAINID"
./tomad init $MONIKER --chain-id $CHAINID --home $DATA_DIR
echo "prepare genesis: Allocate genesis accounts"
./tomad add-genesis-account \
"$(./tomad keys show $KEY -a --home $DATA_DIR --keyring-backend test)" 1000000000000000000atoma,1000000000000000000stake \
--home $DATA_DIR --keyring-backend test
echo "prepare genesis: Sign genesis transaction"
./tomad gentx $KEY 1000000000000000000stake --keyring-backend test --home $DATA_DIR --keyring-backend test --chain-id $CHAINID
echo "prepare genesis: Collect genesis tx"
./tomad collect-gentxs --home $DATA_DIR
echo "prepare genesis: Run validate-genesis to ensure everything worked and that the genesis file is setup correctly"
./tomad validate-genesis --home $DATA_DIR

echo "starting tomahawk node $i in background ..."
./tomad start --pruning=nothing --rpc.unsafe \
--keyring-backend test --home $DATA_DIR \
>$DATA_DIR/node.log 2>&1 & disown

echo "started tomahawk node"
tail -f /dev/null