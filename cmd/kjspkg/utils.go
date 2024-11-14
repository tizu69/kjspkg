package main

import (
	"fmt"

	"g.tizu.dev/colr"
	"github.com/Modern-Modpacks/kjspkg/pkg/kjspkg"
)

func info(format string, a ...any) { fmt.Printf(colr.Blue(":: ")+format+"\n", a...) }
func warn(format string, a ...any) { fmt.Printf(colr.Yellow(":: ")+format+"\n", a...) }

func LoadLocators() (map[string]kjspkg.PackageLocator, error) {
	var packages map[string]kjspkg.PackageLocator
	packages, err := kjspkg.GetPackageList()
	info("Parsed package list")
	return packages, err
}

func LoadPackage(ref kjspkg.PackageLocator, withStats bool) (kjspkg.Package, error) {
	var pkg kjspkg.Package
	pkg, err := kjspkg.GetPackage(ref, withStats)
	info("Obtained package metadata")
	return pkg, err
}

func LoadPackageById(id string, withStats bool) (kjspkg.Package, kjspkg.PackageLocator, error) {
	var pkg kjspkg.Package
	var loc kjspkg.PackageLocator

	locs, err := LoadLocators()
	if err != nil {
		return pkg, loc, err
	}

	loc, ok := locs[id]
	if !ok {
		return pkg, loc, fmt.Errorf("package does not exist: %s", id)
	}

	pkg, err = LoadPackage(loc, withStats)
	return pkg, loc, err
}