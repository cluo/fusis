package bgp

import (
	"fmt"
	"strings"
	"sync"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/deckarep/golang-set"
	"github.com/luizbafilho/fusis/config"
	"github.com/luizbafilho/fusis/state"
	bgp_config "github.com/osrg/gobgp/config"
	"github.com/osrg/gobgp/packet/bgp"
	gobgp "github.com/osrg/gobgp/server"
	"github.com/osrg/gobgp/table"
)

type BgpService struct {
	sync.Mutex
	bgp    *gobgp.BgpServer
	config *config.BalancerConfig
}

type Syncer interface {
	Sync(state state.State) error
}

func (bs *BgpService) Sync(s state.State) error {
	bs.Lock()
	defer bs.Unlock()
	log.Debug("[bgp] Syncing")

	stateSet := bs.getStateBgpPaths(s)
	currentSet, err := bs.getCurrentBgpPaths(s)
	if err != nil {
		return err
	}

	pathsToAdd := stateSet.Difference(currentSet)
	pathsToRemove := currentSet.Difference(stateSet)

	for p := range pathsToAdd.Iter() {
		err := bs.AddPath(p.(string))
		if err != nil {
			return err
		}
		log.Debug("[bgp] Added path: ", p.(string))
	}

	for p := range pathsToRemove.Iter() {
		err := bs.DelPath(p.(string))
		if err != nil {
			return err
		}
		log.Debug("[bgp] Removed path: ", p.(string))
	}

	return nil
}

func (bs *BgpService) getStateBgpPaths(s state.State) mapset.Set {
	paths := mapset.NewSet()

	for _, v := range s.GetServices() {
		paths.Add(v.Address)
	}

	return paths
}

func (bs *BgpService) getCurrentBgpPaths(s state.State) (mapset.Set, error) {
	paths, err := bs.GetPaths()
	if err != nil {
		return nil, err
	}

	currPaths := mapset.NewSet()
	for _, v := range paths {
		currPaths.Add(v)
	}

	return currPaths, nil
}

func NewBgpService(conf *config.BalancerConfig) (*BgpService, error) {
	return &BgpService{
		bgp:    gobgp.NewBgpServer(),
		config: conf,
	}, nil
}

func (bs *BgpService) Serve() {
	go bs.bgp.Serve()

	// global configuration
	global := &bgp_config.Global{
		Config: bgp_config.GlobalConfig{
			As:       bs.config.Bgp.As,
			RouterId: bs.config.Bgp.RouterId,
		},
	}

	if err := bs.bgp.Start(global); err != nil {
		log.Fatal("[bgp] Failed starting service.", err)
	}

	for _, n := range bs.config.Bgp.Neighbors {
		bs.addNeighbor(n)
	}
}

func (bs *BgpService) addNeighbor(nb config.Neighbor) {
	// neighbor configuration
	n := &bgp_config.Neighbor{
		Config: bgp_config.NeighborConfig{
			NeighborAddress: nb.Address,
			PeerAs:          nb.PeerAs,
		},
	}

	if err := bs.bgp.AddNeighbor(n); err != nil {
		log.Fatal("[bgp] Adding BGP Neighbor failed", err)
	}
}

func (bs *BgpService) AddPath(route string) error {
	attrs := []bgp.PathAttributeInterface{
		bgp.NewPathAttributeOrigin(0),
		bgp.NewPathAttributeNextHop("0.0.0.0"),
	}

	if _, err := bs.bgp.AddPath("", []*table.Path{table.NewPath(nil, bgp.NewIPAddrPrefix(32, route), false, attrs, time.Now(), false)}); err != nil {
		return fmt.Errorf("[bgp] Error adding bgp path. %v", err)
	}

	return nil
}

func (bs *BgpService) GetPaths() ([]string, error) {
	paths := []string{}

	var lookupPrefix []*gobgp.LookupPrefix
	_, dsts, err := bs.bgp.GetRib("", bgp.RF_IPv4_UC, lookupPrefix)
	if err != nil {
		return paths, fmt.Errorf("[bgp] Error getting bgp paths. %v", err)
	}

	for k := range dsts {
		paths = append(paths, strings.TrimSuffix(k, "/32"))
	}

	return paths, nil
}

func (bs *BgpService) DelPath(route string) error {
	attrs := []bgp.PathAttributeInterface{
		bgp.NewPathAttributeNextHop("0.0.0.0"),
	}

	if err := bs.bgp.DeletePath([]byte{}, bgp.RF_IPv4_UC, "", []*table.Path{table.NewPath(nil, bgp.NewIPAddrPrefix(32, route), true, attrs, time.Now(), false)}); err != nil {
		return fmt.Errorf("[bgp] Error deleting bgp path. %v", err)
	}

	return nil
}
