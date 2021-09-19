module github.com/ray-project/kuberay/cli

go 1.16

require (
	github.com/fatih/color v1.12.0
	github.com/kris-nova/logger v0.2.2
	github.com/kris-nova/lolgopher v0.0.0-20210112022122-73f0047e8b65
	github.com/mattn/go-runewidth v0.0.13
	github.com/olekukonko/tablewriter v0.0.5
	github.com/ray-project/kuberay/api v0.0.0
	github.com/rivo/uniseg v0.2.0
	github.com/spf13/cobra v1.2.1
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.8.1
	google.golang.org/grpc v1.40.0
)

replace github.com/ray-project/kuberay/api => ../api
