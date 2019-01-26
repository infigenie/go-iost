package run

import (
	"github.com/iost-official/go-iost/itest"
	"github.com/urfave/cli"
	"strconv"
	"github.com/iost-official/go-iost/ilog"
	"time"
)

// AccountCaseCommand is the command of account test case
var AccountCaseCommand = cli.Command{
	Name:      "account_case",
	ShortName: "a_case",
	Usage:     "run account test case",
	Action:    AccountCaseAction,
}

// AccountCaseAction is the action of account test case
var AccountCaseAction = func(c *cli.Context) error {
	itest.Interval = 2 * time.Millisecond
	itest.InitAmount = "1000"
	itest.InitPledge = "1000"
	itest.InitRAM = "3000"
	logger := ilog.New()
	fileWriter := ilog.NewFileWriter(c.GlobalString("log"))
	fileWriter.SetLevel(ilog.LevelInfo)
	logger.AddWriter(fileWriter)
	ilog.InitLogger(logger)
	//anum := c.GlobalInt("anum")
	//output := c.GlobalString("account")
	keysfile := c.GlobalString("keys")
	configfile := c.GlobalString("config")

	it, err := itest.Load(keysfile, configfile)
	if err != nil {
		return err
	}

	for i := 1; i < 101; i++ {
		accounts, err := it.CreateAccountN(10000, false, true, i)
		if err != nil {
			return err
		}

		outputFile := "output_acc" + strconv.FormatInt(int64(i), 10) + ".json"
		//outputFile := output
		ilog.Infof("before dump account %v\n", outputFile)
		if err := itest.DumpAccounts(accounts, outputFile); err != nil {
			return err
		}
	}

	return nil
}
