package logic

import (
	"aidb_go/internal/aidb"
	"aidb_go/internal/svc"
	"aidb_go/internal/types"
	"context"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RigsterReq) (resp *types.StatusReply, err error) {
	// todo: add your logic here and delete this line
	fmt.Printf("~~~!! Register flow_uuid= %v, model=%v, backend=%s\n", req.FlowUUID, strings.Join(req.Model, ","), strings.Join(req.Backend, ","))
	resp = new(types.StatusReply)
	ret := aidb.AiDBRegister(req.FlowUUID, req.Model, req.Backend, req.Zoo)
	resp.Code = ret
	if ret != 0 {
		resp.Msg = "failed"
	} else {
		resp.Msg = "succeed"
	}
	return
}
