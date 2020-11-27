package types

// lottery module event types
const (
	// TODO: Create your event types
	// EventType<Action>    		= "action"
	EventTypeCreateLottery = "CreateLotttery"
	EventTypeCreateTicket  = "CreateTicket"
	EventTypeCreatePrizeAnnounce = "PrizeAnnounce"
	// TODO: Create keys fo your events, the values will be derivided from the msg
	// AttributeKeyAddress  		= "address"
	AttributeDetail  = "detail"
	AttributeReward  = "reward"
	AttributePrice   = "price"
	AttributeLottery = "lottery"
	AttributeNumber  = "number"
	// TODO: Some events may not have values for that reason you want to emit that something happened.
	// AttributeValueDoubleSign = "double_sign"

	AttributeValueCategory = ModuleName
)
