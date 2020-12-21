package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/genus-machina/phage"
	"github.com/genus-machina/phage/static"
	"github.com/genus-machina/phage/trackerradar"
)

var blacklist = flag.String("blacklist", "", "A list of domains to explictly exclude.")
var whitelist = flag.String("whitelist", "", "A list of domains to explictly allow.")

func main() {
	logger := log.New(os.Stderr, "[phage] ", log.LstdFlags)
	flag.Parse()

	var list phage.DomainList
	list = phage.SetUnion{
		static.NewHttpHostFile(logger, "https://raw.githubusercontent.com/StevenBlack/hosts/master/hosts"),
		static.NewHttpHostFile(logger, "https://someonewhocares.org/hosts/zero/hosts"),
		trackerradar.NewDomainList(
			logger,
			trackerradar.AllOf(
				trackerradar.AnyOf(
					trackerradar.FingerprintingScoreAtLeast(2),
					trackerradar.HasCategoryIn(
						"Action Pixels",
						"Ad Motivated Tracking",
						"Advertising",
						"Analytics",
						"Audience Measurement",
						"Malware",
						"Obscure Ownership",
						"Session Replay",
						"Third-Party Analytics Marketing",
						"Unknown High Risk Behavior",
					),
				),
				trackerradar.Not(
					trackerradar.HasCategoryIn(
						"Badge",
						"CDN",
						"Embedded Content",
						"Federated Login",
						"Non-tracking",
						"Online Payment",
						"SSO",
					),
				),
			),
		),
	}

	if *blacklist != "" {
		list = phage.Union(list, static.NewLocalHostFile(logger, *blacklist))
	}

	if *whitelist != "" {
		list = phage.Difference(list, static.NewLocalHostFile(logger, *whitelist))
	}

	domains, err := list.Domains()
	if err != nil {
		logger.Fatal(err.Error())
	}

	logger.Printf("Found a total of %d unique domains.\n", len(domains))
	for _, domain := range domains {
		fmt.Println(domain)
	}
}
