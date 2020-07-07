module github.com/aliensero/go-lotus-interaction

go 1.14

replace github.com/filecoin-project/filecoin-ffi => ./extern/filecoin-ffi

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/filecoin-project/go-address v0.0.2-0.20200504173055-8b6f2fb2b3ef
	github.com/filecoin-project/go-data-transfer v0.3.0
	github.com/filecoin-project/go-fil-markets v0.3.2
	github.com/filecoin-project/go-jsonrpc v0.1.1-0.20200602181149-522144ab4e24
	github.com/filecoin-project/go-statestore v0.1.0
	github.com/filecoin-project/sector-storage v0.0.0-20200630180318-4c1968f62a8f
	github.com/filecoin-project/specs-actors v0.7.1
	github.com/filecoin-project/specs-storage v0.1.1-0.20200622113353-88a9704877ea
	github.com/gbrlsnchs/jwt/v3 v3.0.0-rc.2
	github.com/ipfs/go-block-format v0.0.2
	github.com/ipfs/go-blockservice v0.1.4-0.20200624145336-a978cec6e834
	github.com/ipfs/go-cid v0.0.6
	github.com/ipfs/go-cidutil v0.0.2
	github.com/ipfs/go-datastore v0.4.4
	github.com/ipfs/go-ds-badger2 v0.1.0
	github.com/ipfs/go-ds-leveldb v0.4.2
	github.com/ipfs/go-ds-measure v0.1.0
	github.com/ipfs/go-filestore v1.0.0
	github.com/ipfs/go-fs-lock v0.0.5
	github.com/ipfs/go-graphsync v0.0.6-0.20200504202014-9d5f2c26a103
	github.com/ipfs/go-ipfs-blockstore v1.0.0
	github.com/ipfs/go-ipfs-exchange-interface v0.0.1
	github.com/ipfs/go-ipld-cbor v0.0.5-0.20200428170625-a0bd04d3cbdf
	github.com/ipfs/go-ipld-format v0.2.0
	github.com/ipfs/go-log/v2 v2.1.2-0.20200626104915-0016c0b4b3e4
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/libp2p/go-libp2p-core v0.6.0
	github.com/minio/blake2b-simd v0.0.0-20160723061019-3f5f724cb5b1
	github.com/mitchellh/go-homedir v1.1.0
	github.com/multiformats/go-base32 v0.0.3
	github.com/multiformats/go-multiaddr v0.2.2
	github.com/multiformats/go-multiaddr-net v0.1.2
	github.com/multiformats/go-multibase v0.0.3
	github.com/multiformats/go-multihash v0.0.14
	github.com/prometheus/common v0.10.0
	github.com/syndtr/goleveldb v1.0.0
	github.com/urfave/cli/v2 v2.2.0
	github.com/whyrusleeping/cbor-gen v0.0.0-20200706173030-3bb387cdd4d1
	go.uber.org/zap v1.15.0
	golang.org/x/xerrors v0.0.0-20191204190536-9bdfabe68543
)
