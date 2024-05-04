package types

import (
	"github.com/furysport/streampay/v2/x/streampay/exported"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	authzcodec "github.com/cosmos/cosmos-sdk/x/authz/codec"
	govcodec "github.com/cosmos/cosmos-sdk/x/gov/codec"
)

const (
	AminoTypeStreamSendMsg          = "furya/streampay/MsgStreamSend"
	AminoTypeStopStreamMsg          = "furya/streampay/MsgStopStream"
	AminoTypeClaimStreamedAmountMsg = "furya/streampay/MsgClaimStream"
	AminoTypeStreamPayment          = "furya/streampay/StreamPayment"
	AminoTypeUpdateParamsMsg        = "furya/streampay/MsgUpdateParams"
	AminoTypeParams                 = "furya/streampay/Params"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	legacy.RegisterAminoMsg(cdc, &MsgStreamSend{}, AminoTypeStreamSendMsg)
	legacy.RegisterAminoMsg(cdc, &MsgStopStream{}, AminoTypeStopStreamMsg)
	legacy.RegisterAminoMsg(cdc, &MsgClaimStreamedAmount{}, AminoTypeClaimStreamedAmountMsg)
	legacy.RegisterAminoMsg(cdc, &MsgUpdateParams{}, AminoTypeUpdateParamsMsg)

	cdc.RegisterInterface((*exported.StreamPaymentI)(nil), nil)
	cdc.RegisterConcrete(&StreamPayment{}, AminoTypeStreamPayment, nil)
	cdc.RegisterConcrete(&Params{}, AminoTypeParams, nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgStreamSend{},
		&MsgStopStream{},
		&MsgClaimStreamedAmount{},
		&MsgUpdateParams{},
	)

	registry.RegisterImplementations((*exported.StreamPaymentI)(nil),
		&StreamPayment{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)

	// Register all Amino interfaces and concrete types on the authz Amino codec
	// so that this can later be used to properly serialize MsgGrant and MsgExec
	// instances.
	RegisterLegacyAminoCodec(authzcodec.Amino)
	RegisterLegacyAminoCodec(govcodec.Amino)
}
