#!/bin/bash

if [ 2 -gt $# ]
then
	echo 'need params,$1:miner $2:file';
	exit 400;
fi

importVar="dataCid="`lotus client import $2`;

eval $importVar;

if [ ! "$dataCid" ]
then
	echo "client import data fialed"
	exit 400;
fi

dealVar="lotus client deal $dataCid $1 0 10";

eval $dealVar;

sleep 3;
watch -n 5 lotus client list-deals;

