package service

import (
	"context"
	"guanli/pkg"
	"guanli/srv/basic/config"
	"guanli/srv/handler/model"

	__ "guanli/proto"
)

// server is used to implement helloworld.GreeterServer.
type Server struct {
	__.UnimplementedStreamGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *Server) UserAdd(_ context.Context, in *__.UserAddReq) (*__.UserAddResp, error) {

	var user model.User

	if in.Name == "" || in.Password == "" {
		return &__.UserAddResp{
			Msg:  "不能为空",
			Code: 400,
		}, nil
	}

	err := user.FindUser(config.DB, in.Name)
	if err != nil {
		return &__.UserAddResp{
			Msg:  "用户不存在",
			Code: 400,
		}, nil
	}
	m := model.User{
		Name:     in.Name,
		Password: pkg.Md5(in.Password),
	}
	err = m.UserAdd(config.DB)
	if err != nil {
		return &__.UserAddResp{
			Msg:  "添加失败",
			Code: 400,
		}, nil
	}
	return &__.UserAddResp{
		Msg:  "添加成功",
		Code: 200,
	}, nil
}

// SayHello implements helloworld.GreeterServer
func (s *Server) RoleAdd(_ context.Context, in *__.RoleAddReq) (*__.RoleAddResp, error) {

	var Role model.Role
	err := Role.FindRole(config.DB, in.Name)
	if err != nil {
		return &__.RoleAddResp{
			Msg:  "角色不存在",
			Code: 200,
		}, nil
	}
	role := model.Role{
		Uid:  int(in.Uid),
		Name: in.Name,
	}
	err = role.RoleAdd(config.DB)
	if err != nil {
		return &__.RoleAddResp{
			Msg:  "角色添加失败",
			Code: 200,
		}, nil
	}
	return &__.RoleAddResp{
		Msg:  "角色添加成功",
		Code: 200,
	}, nil
}

// SayHello implements helloworld.GreeterServer
func (s *Server) QuanXianAdd(_ context.Context, in *__.QuanXianAddReq) (*__.QuanXianAddResp, error) {

	var quan model.QuanXian
	err := quan.FindQuan(config.DB, in.Name)
	if err != nil {
		return &__.QuanXianAddResp{
			Msg:  "权限不存在",
			Code: 200,
		}, nil
	}
	xian := model.QuanXian{
		Rid:  int(in.Rid),
		Uid:  int(in.Uid),
		Name: in.Name,
	}
	err = xian.QuanAdd(config.DB)
	if err != nil {
		return &__.QuanXianAddResp{
			Msg:  "权限添加成功",
			Code: 200,
		}, nil
	}

	return &__.QuanXianAddResp{
		Msg:  "权限添加成功",
		Code: 200,
	}, nil
}

// SayHello implements helloworld.GreeterServer
func (s *Server) RoleList(_ context.Context, in *__.RoleListReq) (*__.RoleListResp, error) {

	var role model.Role
	err, _ := role.RoleList(config.DB, in.Id)
	if err != nil {
		return &__.RoleListResp{
			Msg:  "列表查询失败",
			Code: 400,
		}, nil
	}
	return &__.RoleListResp{
		Msg:  "列表查询成功",
		Code: 200,
	}, nil
}

// SayHello implements helloworld.GreeterServer
func (s *Server) RoleDelete(_ context.Context, in *__.RoleDeleteReq) (*__.RoleDeleteResp, error) {

	var role model.Role
	err := role.FindName(config.DB, in.Id)
	if err != nil {
		return &__.RoleDeleteResp{
			Msg:  "角色不存在",
			Code: 400,
		}, nil
	}
	err = role.DeleteRole(config.DB, in.Id)
	if err != nil {
		return &__.RoleDeleteResp{
			Msg:  "删除失败",
			Code: 400,
		}, nil
	}
	return &__.RoleDeleteResp{
		Msg:  "删除成功",
		Code: 200,
	}, nil
}
