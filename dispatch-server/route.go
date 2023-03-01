package main

func (app *Config) Routes() {
	app.router.GET("/")
	app.router.POST("/", app.submitHandler)
}
