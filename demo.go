package main

import (
	"fmt"
	"os"
	"time"

	"github.com/kardianos/service"
)

var logger = service.ConsoleLogger

type program struct{}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) run() {
	// 运行逻辑
	i := 1
	for {
		i = i + 1
		a, err := os.OpenFile("/Users/Zerpro/Downloads/testestetest.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0)
		if err != nil {
			fmt.Println(err)
		}
		_, err = a.WriteString("test123")
		if err != nil {
			fmt.Println(err)
		}
		a.Close()
		fmt.Println(i)
		time.Sleep(time.Second * 1)
	}
}

func (p *program) Stop(s service.Service) error {
	return nil
}

func main() {
	svcConfig := &service.Config{
		Name:        "DemoService",      //服务显示名称
		DisplayName: "demodemodemodemo", //服务名称
		Description: "sadsadadsaddas",   //服务描述
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		logger.Error(err)
	}

	if err != nil {
		logger.Error(err)
	}

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "install":
			s.Install()
			logger.Info("服务安装成功!")
			s.Start()
			logger.Info("服务启动成功!")
			break
		case "start":
			s.Start()
			logger.Info("服务启动成功!")
			break
		case "stop":
			s.Stop()
			logger.Info("服务关闭成功!")
			break
		case "restart":
			s.Stop()
			logger.Info("服务关闭成功!")
			s.Start()
			logger.Info("服务启动成功!")
			break
		case "remove":
			s.Stop()
			logger.Info("服务关闭成功!")
			s.Uninstall()
			logger.Info("服务卸载成功!")
			break
		}
		return
	}
	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}
