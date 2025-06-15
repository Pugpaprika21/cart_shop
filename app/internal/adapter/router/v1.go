package router

func (r *router) v1() {
	server := r.server
	handler := r.handler

	api := server.Group("api")

	v1 := api.Group("/v1")
	{
		user := v1.Group("/user", r.jwtx.Validate())
		user.POST("/getUser", handler.User.GetUsers)
		user.POST("/findUser", handler.User.FindUser)
		user.POST("/creUsers", handler.User.CreUsers)
		user.POST("/updUser", handler.User.UpdUser)
		user.POST("/delUser", handler.User.DelUser)

		products := v1.Group("/products", r.jwtx.Validate())
		products.POST("/create", nil)
	}
}
