package net

import (
	"fmt"
	"io/ioutil"
	"net"

	log "github.com/Sirupsen/logrus"
	"github.com/vishvananda/netlink"
)

//AddIp it receives a CIDR Address and add it to the given interface
func AddIp(ip, iface string) error {
	link, err := netlink.LinkByName(iface)
	if err != nil {
		return err
	}

	addr, err := netlink.ParseAddr(ip)
	if err != nil {
		return err
	}

	return netlink.AddrAdd(link, addr)
}

func DelIp(ip, iface string) error {
	link, err := netlink.LinkByName(iface)
	if err != nil {
		return err
	}

	addr, err := netlink.ParseAddr(ip)
	if err != nil {
		return err
	}

	return netlink.AddrDel(link, addr)
}

func DelVips(iface string) error {
	link, err := netlink.LinkByName(iface)
	if err != nil {
		return err
	}

	addrs, err := netlink.AddrList(link, netlink.FAMILY_V4)
	if err != nil {
		return err
	}

	for _, a := range addrs[1:] {
		if err := netlink.AddrDel(link, &a); err != nil {
			return err
		}
	}

	return nil
}

func SetLinkUp(iface string) error {
	link, err := netlink.LinkByName(iface)
	if err != nil {
		return err
	}

	return netlink.LinkSetUp(link)
}

func GetVips(iface string) ([]netlink.Addr, error) {
	link, err := netlink.LinkByName(iface)
	if err != nil {
		return []netlink.Addr{}, err
	}

	return netlink.AddrList(link, netlink.FAMILY_V4)
}

func GetFusisVipsIps(iface string) ([]string, error) {
	addrs, err := GetVips(iface)
	if err != nil {
		return nil, err
	}
	addrs = addrs[1:]
	ips := make([]string, len(addrs))
	for i, addr := range addrs {
		ips[i] = addr.IP.String()
	}
	return ips, nil
}

func GetIpByInterface(iface string) (string, error) {
	link, err := netlink.LinkByName(iface)
	if err != nil {
		return "", err
	}

	addrs, err := netlink.AddrList(link, netlink.FAMILY_V4)
	if err != nil {
		return "", err
	}

	return addrs[0].IP.String(), nil
}

func SetIpForwarding() error {
	return ioutil.WriteFile("/proc/sys/net/ipv4/ip_forward", []byte("1"), 0644)
}

func AddDefaultGateway(ip string) error {
	err := netlink.RouteAdd(&netlink.Route{
		Scope: netlink.SCOPE_UNIVERSE,
		Gw:    net.ParseIP(ip),
	})
	if err != nil {
		log.Errorf("Adding Default Gateway: %s", ip)
		return err
	}
	return nil
}

func GetDefaultGateway() (*netlink.Route, error) {
	routes, err := netlink.RouteList(nil, netlink.FAMILY_V4)
	if err != nil {
		return nil, err
	}

	var route netlink.Route
	for _, v := range routes {
		if v.Gw != nil {
			route = v
		}
	}

	if route.Gw == nil {
		return nil, fmt.Errorf("Default gateway not found")
	}

	return &route, nil
}

func DeleteDefaultGateway(route *netlink.Route) error {
	return netlink.RouteDel(route)
}
