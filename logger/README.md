# easy_zap_log.go
## This is an encapsulation of "go.uber.org/zap"
- Make Log Configuration Simplest  
The purpose of this configuration is to generate local log files for various projects.   
It has a default configuration. It only needs one line of code for config.   
The default log files will be generated in the "logs" folder in the current directory.   
There will be error and info types. You can try them.  

````golang

    gu.InitEasyZapDefault("golang-util")
    //gu.InitEasyZap("golang-util","logpath", 1 * time.Minute, 1 * time.Minute)

    //  Example 1:

    log := zap.S()
    log.Debug("1111")
    log.Info("2222")
    log.Warn("3333")
    log.Error("4444")

    //  Example 2 (output log faster than example 1 and only use in some special place):
    log := zap.L()
    log.Debug("1111", zap.Int("id",1), zap.String("name", "jack"))
    
````