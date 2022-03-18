package main

import (
	"fmt"
	"io"
	"os"
	"time"

	"gitee.com/yjhi/golib/jsql"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

type Worker struct {
	Cfg     *Config
	StrCon  string
	LogFile *os.File
}

func (w *Worker) Run() {
	var use float64 = 0

	t1 := time.Now()

	sqlUtil := jsql.BuildSqlServer()

	b := sqlUtil.ConnectSqlServer(w.StrCon)

	ret := ""

	if b {

		err := sqlUtil.Db.Ping()

		if err == nil {
			ret = "Success"

			count, isok := sqlUtil.GetInt(w.Cfg.Sql)

			if isok {
				ret = "Query OK," + fmt.Sprintf("%d", count)
			} else {
				ret = "Query Fail"
			}

			sqlUtil.CloseSqlServer()

		} else {
			ret = "Connect Fail"
		}

	} else {
		ret = "Fail"
	}

	t2 := time.Now()

	use = t2.Sub(t1).Seconds()

	str := fmt.Sprintf("Use Time:%vs\t%s\n", use, ret)

	logrus.Info(str)
}

func StartWork(cfg *Config) {

	strIp := cfg.IP

	if len(cfg.Port) > 0 {
		strIp += "," + cfg.Port
	}

	job := &Worker{
		Cfg:     cfg,
		StrCon:  jsql.CreateSqlServerString(strIp, cfg.User, cfg.Pass, cfg.Db),
		LogFile: nil,
	}

	var err error
	job.LogFile, err = os.OpenFile("worker.log", os.O_WRONLY|os.O_CREATE, 0755)

	if err != nil {
		fmt.Sprintln("Create Log Fail.")
		job.LogFile = nil
	}

	if job.LogFile != nil {
		logrus.SetOutput(io.MultiWriter(os.Stdout, job.LogFile))
	}
	logrus.Info("-------------------------------------------------------")

	logrus.Info(cfg.Time)
	logrus.Info(job.StrCon)

	c := cron.New(cron.WithSeconds())

	c.AddJob(cfg.Time, job)

	c.Start()

	defer c.Stop()

	select {}

}
