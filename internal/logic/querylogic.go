package logic

import (
	"aidb_go/internal/aidb"
	"aidb_go/internal/svc"
	"aidb_go/internal/types"
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryLogic {
	return &QueryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryLogic) Query(req *types.QueryReq) (resp *types.QueryReply, err error) {
	// todo: add your logic here and delete this line
	fmt.Printf("~~~!! Query flow_uuid= %v\n", req.FlowUUID)
	resp = new(types.QueryReply)
	aidb_output := aidb.AiDBOutput{}
	ret := aidb.AiDBForward(req.FlowUUID, req.ImageBase64, &aidb_output)
	resp.Code = ret
	if ret != 0 {
		resp.Msg = "failed"
	} else {
		resp.Msg = "succeed"
		if len(aidb_output.Anime) > 0 {
			resp.Anime = aidb_output.Anime
		}
		if len(aidb_output.Tddfa) > 0 {
			resp.Tddfa = aidb_output.Tddfa
		}
		if len(aidb_output.Face) > 0 {
			for _, value := range aidb_output.Face {
				cur_face_result := types.FaceResult{}
				cur_face_result.Conf = value.Conf
				cur_face_result.BBox = value.BBox
				cur_face_result.LandMark = value.LandMark
				if len(value.Parsing) > 0 {
					cur_face_result.Parsing = value.Parsing
				}

				resp.Face = append(resp.Face, cur_face_result)
			}
		}
		if len(aidb_output.Object) > 0 {
			for _, value := range aidb_output.Object {
				cur_obj_result := types.ObjectResult{}
				cur_obj_result.BBox = value.BBox
				cur_obj_result.Conf = value.Conf
				cur_obj_result.Label = value.Label
				resp.Object = append(resp.Object, cur_obj_result)
			}
		}
		if len(aidb_output.Ocr) > 0 {
			for _, value := range aidb_output.Ocr {
				cur_ocr_result := types.OCRResult{}
				cur_ocr_result.Box = value.Box
				cur_ocr_result.Conf = value.Conf
				cur_ocr_result.Label = value.Label
				cur_ocr_result.ConfRotate = value.ConfRotate
				resp.Ocr = append(resp.Ocr, cur_ocr_result)
			}
		}
		if len(aidb_output.Cls) > 0 {
			for _, value := range aidb_output.Cls {
				cur_cls_result := types.ClsResult{}
				cur_cls_result.Conf = value.Conf
				cur_cls_result.Label = value.Label
				resp.Cls = append(resp.Cls, cur_cls_result)
			}
		}
		if len(aidb_output.KeyPoints) > 0 {
			resp.KeyPoints = aidb_output.KeyPoints
		}

	}
	return
}
