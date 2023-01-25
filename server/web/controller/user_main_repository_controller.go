package controller

// // UserMainRepositoryController 仓库控制器
// type UserMainRepositoryController struct {
// 	markup.RestController `class:"rest-controller"`

// 	RepoService service.UserMainRepositoryService `inject:"#UserMainRepositoryService"`
// 	Responder   glass.MainResponder               `inject:"#glass-main-responder"`
// }

// func (inst *UserMainRepositoryController) _Impl() glass.Controller {
// 	return inst
// }

// // Init 初始化
// func (inst *UserMainRepositoryController) Init(ec glass.EngineConnection) error {

// 	ec = ec.RequestMapping("user-main-repositories")

// 	ec.Handle(http.MethodGet, "", inst.handleGetList)
// 	ec.Handle(http.MethodPost, "", inst.handlePost)

// 	ec.Handle(http.MethodGet, ":id", inst.handleGetOne)
// 	ec.Handle(http.MethodPut, ":id", inst.handlePut)
// 	ec.Handle(http.MethodDelete, ":id", inst.handleDelete)

// 	return nil
// }

// func (inst *UserMainRepositoryController) handleGetList(c *gin.Context) {
// 	req := &myUserMainRepositoryRequest{
// 		gc:              c,
// 		controller:      inst,
// 		wantRequestID:   false,
// 		wantRequestBody: false,
// 	}
// 	err := req.open()
// 	if err == nil {
// 		err = req.doGetList()
// 	}
// 	req.send(err)
// }

// func (inst *UserMainRepositoryController) handleGetOne(c *gin.Context) {
// 	req := &myUserMainRepositoryRequest{
// 		gc:              c,
// 		controller:      inst,
// 		wantRequestID:   true,
// 		wantRequestBody: false,
// 	}
// 	err := req.open()
// 	if err == nil {
// 		err = req.doGetOne()
// 	}
// 	req.send(err)
// }

// func (inst *UserMainRepositoryController) handlePost(c *gin.Context) {
// 	req := &myUserMainRepositoryRequest{
// 		gc:              c,
// 		controller:      inst,
// 		wantRequestID:   false,
// 		wantRequestBody: true,
// 	}
// 	err := req.open()
// 	if err == nil {
// 		err = req.doPost()
// 	}
// 	req.send(err)
// }

// func (inst *UserMainRepositoryController) handlePut(c *gin.Context) {
// 	req := &myUserMainRepositoryRequest{
// 		gc:              c,
// 		controller:      inst,
// 		wantRequestID:   true,
// 		wantRequestBody: true,
// 	}
// 	err := req.open()
// 	if err == nil {
// 		err = req.doPut()
// 	}
// 	req.send(err)
// }

// func (inst *UserMainRepositoryController) handleDelete(c *gin.Context) {
// 	req := &myUserMainRepositoryRequest{
// 		gc:              c,
// 		controller:      inst,
// 		wantRequestID:   true,
// 		wantRequestBody: false,
// 	}
// 	err := req.open()
// 	if err == nil {
// 		err = req.doDelete()
// 	}
// 	req.send(err)
// }

// ////////////////////////////////////////////////////////////////////////////////

// type myUserMainRepositoryRequest struct {
// 	gc         *gin.Context
// 	controller *UserMainRepositoryController

// 	wantRequestID   bool
// 	wantRequestBody bool

// 	id    dxo.UserMainRepositoryID
// 	body1 vo.UserMainRepository
// 	body2 vo.UserMainRepository
// }

// func (inst *myUserMainRepositoryRequest) open() error {

// 	c := inst.gc

// 	if inst.wantRequestID {
// 		id := c.Param("id")
// 		n, err := strconv.Atoi(id)
// 		if err != nil {
// 			return err
// 		}
// 		inst.id = dxo.UserMainRepositoryID(n)
// 	}

// 	if inst.wantRequestBody {
// 		err := c.BindJSON(&inst.body1)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

// func (inst *myUserMainRepositoryRequest) send(err error) {
// 	ctx := inst.gc
// 	data := &inst.body2
// 	status := data.Status
// 	resp := &glass.Response{
// 		Context: ctx,
// 		Data:    data,
// 		Error:   err,
// 		Status:  status,
// 	}
// 	inst.controller.Responder.Send(resp)
// }

// func (inst *myUserMainRepositoryRequest) doGetOne() error {
// 	id := inst.id
// 	ctx := inst.gc
// 	ser := inst.controller.RepoService
// 	o, err := ser.Find(ctx, id)
// 	if err != nil {
// 		return err
// 	}
// 	inst.body2.Repositories = []*dto.UserMainRepository{o}
// 	return nil
// }

// func (inst *myUserMainRepositoryRequest) doGetList() error {
// 	ctx := inst.gc
// 	ser := inst.controller.RepoService
// 	list, err := ser.ListAll(ctx)
// 	if err != nil {
// 		return err
// 	}
// 	inst.body2.Repositories = list
// 	return nil
// }

// func (inst *myUserMainRepositoryRequest) doPost() error {

// 	return nil
// }

// func (inst *myUserMainRepositoryRequest) doPut() error {
// 	return nil
// }

// func (inst *myUserMainRepositoryRequest) doDelete() error {

// 	return nil
// }
