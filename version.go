// generated by: generate_version_info.sh "1.2.0" "y.nasretdinov@corp.badoo.com" "" "/Applications/Xcode.app/Contents/Developer/usr/bin/make - " 
// DO NOT EDIT!

package main

import (
	"badoo/_packages/service"
	"badoo/_packages/service/stats"
	"github.com/gogo/protobuf/proto"
)

func init() {
	service.VersionInfo = badoo_service.ResponseVersion{
		Version:        proto.String("1.2.0"),
		Maintainer:     proto.String("y.nasretdinov@corp.badoo.com"),
		AutoBuildTag:   proto.String(""),
		BuildDate:      proto.String("Wed Jun 15 10:47:13 UTC 2016"),
		BuildHost:      proto.String("nasretdinov.local"),
		BuildGoVersion: proto.String("go version go1.6 darwin/amd64"),
		BuildCommand:   proto.String("/Applications/Xcode.app/Contents/Developer/usr/bin/make - "),
		VcsType:        proto.String("git"),
		VcsBasename:    proto.String("lsd"),
		VcsNum:         proto.String("65"),
		VcsDate:        proto.String("2016-06-15T13:19:52+0300"),
		VcsBranch:      proto.String("CT-2452_reduce_flood"),
		VcsTag:         proto.String("in-build-1.0.3-100"),
		VcsTick:        proto.String("29"),
		VcsFullHash:    proto.String("fe4022f8d87956d7133f738ebc054646b88c9a99"),
		VcsShortHash:   proto.String("fe4022f"),
		VcsWcModified:  proto.String("1"),
	}
}