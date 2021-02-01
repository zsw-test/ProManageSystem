package middleware

// func Ownerauth(c *gin.Context) {
// 	fmt.Println("ownerauth中间件启动")
// 	session, err := cache.SessionStore.Get(c.Request, "sessionid")
// 	if err != nil {
// 		fmt.Println("没有获取到session")
// 		c.Redirect(http.StatusFound, "/login")
// 		return
// 	}
// 	username, ok := session.Values["username"]
// 	if !ok {
// 		fmt.Println("没有获取到username")
// 		c.Redirect(http.StatusFound, "/login")
// 		return
// 	}
// 	c.Set("username", username)

// }
