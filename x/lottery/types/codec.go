package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
  // this line is used by starport scaffolding # 1
		cdc.RegisterConcrete(MsgCreatePrizeAnnounce{}, "lottery/CreatePrizeAnnounce", nil)
		cdc.RegisterConcrete(MsgSetPrizeAnnounce{}, "lottery/SetPrizeAnnounce", nil)
		cdc.RegisterConcrete(MsgDeletePrizeAnnounce{}, "lottery/DeletePrizeAnnounce", nil)
		cdc.RegisterConcrete(MsgCreateTicket{}, "lottery/CreateTicket", nil)
		cdc.RegisterConcrete(MsgSetTicket{}, "lottery/SetTicket", nil)
		cdc.RegisterConcrete(MsgDeleteTicket{}, "lottery/DeleteTicket", nil)
		cdc.RegisterConcrete(MsgCreateLottery{}, "lottery/CreateLottery", nil)
		cdc.RegisterConcrete(MsgSetLottery{}, "lottery/SetLottery", nil)
		cdc.RegisterConcrete(MsgDeleteLottery{}, "lottery/DeleteLottery", nil)
}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
