package main

import "github.com/Sirupsen/logrus"

func report(result *Result) error {
	logrus.Infof("%#v", result)
	return nil
}
