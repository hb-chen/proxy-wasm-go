package main

import (
	"math/rand"
	"time"

	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
)

const tickMilliseconds uint32 = 1000

func main() {
	proxywasm.SetNewRootContext(newHelloWorld)
}

type helloWorld struct {
	// you must embed the default context so that you need not to reimplement all the methods by yourself
	proxywasm.DefaultRootContext
	contextID uint32
}

func newHelloWorld(contextID uint32) proxywasm.RootContext {
	return &helloWorld{contextID: contextID}
}

// override
func (ctx *helloWorld) OnVMStart(vmConfigurationSize int) types.OnVMStartStatus {
	rand.Seed(time.Now().UnixNano())

	proxywasm.LogInfo("proxy_on_vm_start from Go!")
	if err := proxywasm.SetTickPeriodMilliSeconds(tickMilliseconds); err != nil {
		proxywasm.LogCriticalf("failed to set tick period: %v", err)
	}

	return types.OnVMStartStatusOK
}

// override
func (ctx *helloWorld) OnTick() {
	t := time.Now().UnixNano()
	proxywasm.LogInfof("It's %d: random value: %d", t, rand.Uint64())
}
