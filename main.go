package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"

	envLoader "github.com/joho/godotenv"
	"gitlab.com/kulyklev/autoria-parser/cmds"
	"gitlab.com/kulyklev/autoria-parser/common/logger"
)

//var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
//var memprofile = flag.String("memprofile", "", "write memory profile to this file")

func main() {
	cpuprofile := "cpu.prof"
	if cpuprofile != "" {
		f, err := os.Create(cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer f.Close()
		defer pprof.StopCPUProfile()
	}

	/*if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.WriteHeapProfile(f)
		f.Close()
		return
	}*/

	log, err := logger.New("PARSER")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer log.Sync()

	err = envLoader.Load()
	if err != nil {
		log.Errorw("loading env", "ERROR", err)
		log.Sync()
		os.Exit(1)
	}

	if err = cmds.Exec(os.Args, log); err != nil {
		log.Errorw("running command", "ERROR", err)
		log.Sync()
		os.Exit(1)
	}
}
