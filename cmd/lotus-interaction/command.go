package main

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	// "reflect"
	"strconv"
	"syscall"
	"text/tabwriter"
	"time"

	lapi "github.com/aliensero/go-lotus-interaction/api"
	lcli "github.com/aliensero/go-lotus-interaction/cli"
	"github.com/aliensero/go-lotus-interaction/types"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-fil-markets/storagemarket"
	"github.com/filecoin-project/specs-actors/actors/abi"
	"github.com/filecoin-project/specs-actors/actors/builtin/market"
	"github.com/ipfs/go-cid"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
)

var ClientDealCmd = &cli.Command{
	Name:  "deal",
	Usage: "deal",
	Subcommands: []*cli.Command{
		ClientDealSubMakeCmd,
		ClientDealSubListCmd,
		ClientDealSubInterCmd,
	},
}

var ClientDealSubMakeCmd = &cli.Command{
	Name:      "make",
	Usage:     "deal make",
	ArgsUsage: "[inputfile minerActor]",
	Action: func(cctx *cli.Context) error {
		api, closer, err := lcli.GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()
		if cctx.NArg() != 2 {
			return xerrors.New("Usage: deal <inputfile> <minerActor>")
		}
		path := cctx.Args().Get(0)
		actor := cctx.Args().Get(1)
		cid, err := importFucn(cctx, api, path)
		if err != nil {
			fmt.Println(err)
			return err
		}
		err = dealFunc(cctx, api, cid, actor)
		if err != nil {
			fmt.Println(err)
			return err
		}

		return nil
	},
}

var ClientDealSubListCmd = &cli.Command{
	Name:  "list",
	Usage: "deal list",
	Action: func(cctx *cli.Context) error {
		api, closer, err := lcli.GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()
		return listDealFunc(cctx, api)
	},
}

var ClientDealSubInterCmd = &cli.Command{
	Name:  "inter",
	Usage: "deal inter",
	Action: func(cctx *cli.Context) error {
		api, closer, err := lcli.GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()
		for {
			helpFunc()
			fmt.Print("select:")
			var op string
			fmt.Scanln(&op)
			fmt.Println("")
			switch op {
			case "1":
				var path, actor string

				fmt.Print("inputfile:")
				fmt.Scanln(&path)
				fmt.Print("minerActor:")
				fmt.Scanln(&actor)
				cid, err := importFucn(cctx, api, path)
				if err != nil {
					fmt.Println(err)
					continue
				}
				err = dealFunc(cctx, api, cid, actor)
				if err != nil {
					fmt.Println(err)
				}
			case "2":
				err := listDealFunc(cctx, api)
				if err != nil {
					fmt.Println(err)
				}
				// IntervalExec(3, listDealFunc, cctx, api)
			case "99":
				fmt.Println("QUIT")
				return nil
			default:
				fmt.Println("you have no choice.")
			}
		}
	},
}

func importFucn(cctx *cli.Context, api lapi.FullNode, path string) (string, error) {
	cid := ""
	ctx := lcli.ReqContext(cctx)

	absPath, err := filepath.Abs(path)
	if err != nil {
		return cid, errors.New("Get asbPath error:" + err.Error())
	}

	ref := lapi.FileRef{
		Path:  absPath,
		IsCAR: false,
	}
	c, err := api.ClientImport(ctx, ref)
	if err != nil {
		return cid, errors.New("Import error:" + err.Error())
	}
	encoder, err := lcli.GetCidEncoder(cctx)
	if err != nil {
		return cid, errors.New("Get cidEncoder error:" + err.Error())
	}
	cid = encoder.Encode(c)
	fmt.Printf("cid: %s\n", cid)

	return cid, nil
}

func dealFunc(cctx *cli.Context, api lapi.FullNode, dataCid, actor string) error {

	ctx := lcli.ReqContext(cctx)
	data, err := cid.Parse(dataCid)
	if err != nil {
		return err
	}

	miner, err := address.NewFromString(actor)
	if err != nil {
		return err
	}

	price, err := types.ParseFIL("0")
	if err != nil {
		return err
	}

	dur, err := strconv.ParseInt("10", 10, 32)
	if err != nil {
		return err
	}

	var a address.Address
	def, err := api.WalletDefaultAddress(ctx)
	if err != nil {
		return err
	}
	a = def

	ref := &storagemarket.DataRef{
		TransferType: storagemarket.TTGraphsync,
		Root:         data,
	}

	proposal, err := api.ClientStartDeal(ctx, &lapi.StartDealParams{
		Data:              ref,
		Wallet:            a,
		Miner:             miner,
		EpochPrice:        types.BigInt(price),
		MinBlocksDuration: uint64(dur),
		DealStartEpoch:    abi.ChainEpoch(cctx.Int64("start-epoch")),
	})
	if err != nil {
		return err
	}

	encoder, err := lcli.GetCidEncoder(cctx)
	if err != nil {
		return err
	}

	fmt.Println("message:" + encoder.Encode(*proposal))
	return nil
}

func listDealFunc(cctx *cli.Context, api lapi.FullNode) error {

	ctx := lcli.ReqContext(cctx)

	head, err := api.ChainHead(ctx)
	if err != nil {
		return err
	}

	localDeals, err := api.ClientListDeals(ctx)
	if err != nil {
		return err
	}

	var deals []deal
	for _, v := range localDeals {
		if v.DealID == 0 {
			deals = append(deals, deal{
				LocalDeal: v,
				OnChainDealState: market.DealState{
					SectorStartEpoch: -1,
					LastUpdatedEpoch: -1,
					SlashEpoch:       -1,
				},
			})
		} else {
			onChain, err := api.StateMarketStorageDeal(ctx, v.DealID, head.Key())
			if err != nil {
				return err
			}

			deals = append(deals, deal{
				LocalDeal:        v,
				OnChainDealState: onChain.State,
			})
		}
	}

	w := tabwriter.NewWriter(os.Stdout, 2, 4, 2, ' ', 0)
	fmt.Fprintf(w, "DealCid\tDealId\tProvider\tState\tOn Chain?\tSlashed?\tPieceCID\tSize\tPrice\tDuration\tMessage\n")
	for _, d := range deals {
		onChain := "N"
		if d.OnChainDealState.SectorStartEpoch != -1 {
			onChain = fmt.Sprintf("Y (epoch %d)", d.OnChainDealState.SectorStartEpoch)
		}

		slashed := "N"
		if d.OnChainDealState.SlashEpoch != -1 {
			slashed = fmt.Sprintf("Y (epoch %d)", d.OnChainDealState.SlashEpoch)
		}

		fmt.Fprintf(w, "%s\t%d\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%d\t%s\n", d.LocalDeal.ProposalCid, d.LocalDeal.DealID, d.LocalDeal.Provider, storagemarket.DealStates[d.LocalDeal.State], onChain, slashed, d.LocalDeal.PieceCID, types.SizeStr(types.NewInt(d.LocalDeal.Size)), d.LocalDeal.PricePerEpoch, d.LocalDeal.Duration, d.LocalDeal.Message)
	}
	return w.Flush()
}

func helpFunc() {
	fmt.Println(`---------------------------------------
|  1:make deal 2:list deals 99:exit   |
---------------------------------------`)
}

type deal struct {
	LocalDeal        lapi.DealInfo
	OnChainDealState market.DealState
}

func IntervalExec(i int64, cb interface{}, params ...interface{}) {
	sch := make(chan os.Signal, 1)
	signal.Notify(sch, syscall.SIGINT)
	for {
		select {
		case <-sch:
			return
		default:
			time.Sleep(time.Duration(i) * time.Second)
			// mFunc := reflect.ValueOf(cb).Interface()

			// mFunc(reflect.ValueOf(params[0]).Interface())
		}
	}
}
