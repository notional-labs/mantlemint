package main

import (
	"fmt"
	dbm "github.com/tendermint/tm-db"
)

func main() {
	fmt.Println("Hello")

	dbType := dbm.BackendType(dbm.GoLevelDBBackend)
	// dir := "/Users/tuanpa/.mantlemint"
	dir := "/root/.mantlemint"

	stateDB, err := dbm.NewDB("state", dbType, dir)
	if err != nil {
		panic(err)
	}
	blockstoreDB, err := dbm.NewDB("blockstore", dbType, dir)
	if err != nil {
		panic(err)
	}
	mmDB, err := dbm.NewDB("mantlemint", dbType, dir)
	if err != nil {
		panic(err)
	}

	defer func() {
		stateDB.Close()
		blockstoreDB.Close()
		mmDB.Close()
	}()

	itr, err := stateDB.Iterator(nil, nil)
	if err != nil {
		panic(err)
	}

	for ; itr.Valid(); itr.Next() {
		k := cp(itr.Key())
		v := cp(itr.Value())
		mmDB.Set(k, v)
	}

	itr, err = blockstoreDB.Iterator(nil, nil)
	if err != nil {
		panic(err)
	}

	for ; itr.Valid(); itr.Next() {
		k := cp(itr.Key())
		v := cp(itr.Value())
		mmDB.Set(k, v)
	}

	////dbType := dbm.BackendType(dbm.GoLevelDBBackend)
	////stateDB, err := dbm.NewDB("state", dbType, "/Users/tuanpa/.mantlemint")
	////if err != nil {
	////	panic(err)
	////}
	////
	////defer func() {
	////	stateDB.Close()
	////}()
	////
	////stateStore := sm.NewStore(stateDB)
	////lastState, err := stateStore.Load()
	////if err != nil {
	////	panic(err)
	////}
	////fmt.Printf("LastBlockHeight=%d\n",lastState.LastBlockHeight)
	////
	////
	////valset, err := stateStore.LoadValidators(11800315)
	////if err != nil {
	////	panic(err)
	////}
	////fmt.Printf("LastValidators=%v\n", valset)
	//
	//
	//
	//
	//dbType := dbm.BackendType(dbm.GoLevelDBBackend)
	//blockstoreDB, err := dbm.NewDB("blockstore", dbType, "/Users/tuanpa/.mantlemint")
	//if err != nil {
	//	panic(err)
	//}
	//
	//defer func() {
	//	blockstoreDB.Close()
	//}()
	//
	//blockStore := bm.NewBlockStore(blockstoreDB)
	////blockStore.
	//blk := blockStore.LoadBlock(11800315)
	//fmt.Printf("blk=%v\n", blk)
	////blk
	////blockStore.SaveBlock()
}

func cp(bz []byte) (ret []byte) {
	ret = make([]byte, len(bz))
	copy(ret, bz)
	return ret
}
