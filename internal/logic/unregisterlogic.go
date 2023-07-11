package logic

import (
	"aidb_go/internal/aidb"
	"aidb_go/internal/svc"
	"aidb_go/internal/types"
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnregisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUnregisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnregisterLogic {
	return &UnregisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UnregisterLogic) Unregister(req *types.UnRigsterReq) (resp *types.StatusReply, err error) {
	// todo: add your logic here and delete this line
	fmt.Printf("~~~!! Register flow_uuid= %v\n", req.FlowUUID)
	ret := aidb.AiDBUnRegister(req.FlowUUID)
	resp = new(types.StatusReply)
	resp.Code = ret
	if ret != 0 {
		resp.Msg = "failed"
	} else {
		resp.Msg = "succeed"
	}
	return
}
