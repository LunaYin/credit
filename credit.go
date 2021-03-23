package main

import (
	"log"

	"github.com/cloudstateio/go-support/cloudstate"
	"github.com/cloudstateio/go-support/cloudstate/eventsourced"
	"github.com/cloudstateio/go-support/cloudstate/protocol"
	"github.wdf.sap.corp/ICN-Nanjing-Projects/Collaboration-board/credit/entity"
)

func main() {
	server, err := cloudstate.New(protocol.Config{
		ServiceName:    "entity.CreditService",
		ServiceVersion: "0.2.0",
	})
	if err != nil {
		log.Fatalf("cloudstate.New failed: %s", err)
	}
	err = server.RegisterEventSourced(&eventsourced.Entity{
		ServiceName:   "entity.CreditService",
		PersistenceID: "credit",
		EntityFunc:    entity.NewCredit,
	}, protocol.DescriptorConfig{
		Service: "credit.proto",
	}.AddDomainDescriptor("domain.proto"))
	if err != nil {
		log.Fatalf("CloudState failed to register entity: %s", err)
	}
	err = server.Run()
	if err != nil {
		log.Fatalf("Cloudstate failed to run: %v", err)
	}
}
