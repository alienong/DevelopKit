/**
 * @Author: alienongwlx@gmail.com
 * @Description:
 * @Version: 1.0.0
 * @Date: 2020/4/17 15:51
 */

package dataserver

import (
	"fmt"
	"github.com/unknwon/goconfig"
	"strconv"
	"strings"
	"time"
)

/**
@description: DataService Struct
@attribute configFile: config ini path
@attribute configParam: config section's setting
@attribute preFunc: Before TimerEvent's Function
@attribute timerEvent: Every Interval Minute Function
*/
type DataService struct {
	configFile  string
	configParam []string
	preFunc     func(*map[string]string) error
	timerEvent  func(*map[string]string)
}

/**
@description: NewDataService
@param configFile: config ini path
@param configParam: config section's setting
@param preFunc: Before TimerEvent's Function
@param timerEvent: Every Interval Minute Function
@return: DataService
*/
func NewDataService(configFile string, configParam []string, preFunc func(*map[string]string) error, timerEvent func(*map[string]string)) *DataService {
	return &DataService{configFile, configParam, preFunc, timerEvent}
}

/**
@description: Read Config From ConfigFile
@param : nil
@return: config map
*/
func (ds *DataService) ReadConfig() *map[string]string {
	configs := make(map[string]string, 0)
	c, err := goconfig.LoadConfigFile(ds.configFile)
	if err != nil {
		panic("load config file failed")
	}
	for _, para := range ds.configParam {
		temp, err := c.GetValue("config", para)
		if err != nil {
			panic("load config's " + para + " failed")
		}
		configs[para] = strings.TrimSpace(temp)
	}
	return &configs
}

/**
@description: DataTrans Start
@param : nil
@return: nil
*/
func (ds *DataService) DataTrans() {
	configs := ds.ReadConfig()
	if ds.preFunc != nil {
		err := ds.preFunc(configs)
		if err != nil {
			fmt.Println(fmt.Sprintf("PreFunc error %s", err.Error()))
			return
		}
	}
	interval, err := strconv.Atoi((*configs)["Interval"])
	if err != nil {
		fmt.Println(fmt.Sprintf("Config File Interval Error %s", err.Error()))
		return
	}
	c := make(chan int, 1)
	ticker := time.NewTicker(time.Minute * time.Duration(interval))
	ds.timerEvent(configs)
	go func() {
		for {
			select {
			case <-ticker.C:
				ds.timerEvent(configs)
			}
		}
	}()
	<-c
}
