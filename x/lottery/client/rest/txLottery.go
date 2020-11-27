package rest

import (
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/oweliie001/lottery/x/lottery/types"
)

// Used to not have an error if strconv is unused
var _ = strconv.Itoa(42)

type createLotteryRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Creator string `json:"creator"`
	Name string `json:"name"`
	Detail string `json:"detail"`
	Reward string `json:"reward"`
	Drawdate string `json:"drawdate"`
	Price string `json:"price"`
	
}

func createLotteryHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createLotteryRequest
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}
		creator, err := sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		
		parsedName := req.Name
		
		parsedDetail := req.Detail
		
		parsedReward, err2 := sdk.ParseCoins(req.Reward)
		if err2 != nil {
			return
		}
		
		parsedDrawdate := req.Drawdate
		
		parsedPrice, err2 := sdk.ParseCoins(req.Price)
		if err2 != nil {
			return
		}

		msg := types.NewMsgCreateLottery(
			creator,
			parsedName,
			parsedDetail,
			parsedReward,
			parsedDrawdate,
			parsedPrice,
			
		)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type setLotteryRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	ID 		string `json:"id"`
	Creator string `json:"creator"`
	Name string `json:"name"`
	Detail string `json:"detail"`
	Reward string `json:"reward"`
	Drawdate string `json:"drawdate"`
	Price string `json:"price"`
	
}

func setLotteryHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req setLotteryRequest
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}
		creator, err := sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		
		parsedName := req.Name
		
		parsedDetail := req.Detail
		
		parsedReward, err2 := sdk.ParseCoins(req.Reward)
		if err2 != nil {
			return
		}
		
		parsedDrawdate := req.Drawdate
		
		parsedPrice, err2 := sdk.ParseCoins(req.Price)
		if err2 != nil {
			return
		}
		

		msg := types.NewMsgSetLottery(
			creator,
			req.ID,
			parsedName,
			parsedDetail,
			parsedReward,
			parsedDrawdate,
			parsedPrice,
			
		)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type deleteLotteryRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Creator string `json:"creator"`
	ID 		string `json:"id"`
}

func deleteLotteryHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req deleteLotteryRequest
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}
		creator, err := sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		msg := types.NewMsgDeleteLottery(req.ID, creator)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}
